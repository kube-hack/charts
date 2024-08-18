package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	git "github.com/go-git/go-git/v5"
)

type Repo struct {
	URL   string
	Name  string
	Owner string
}

func (r *Repo) set(repoURL string) error {
	urlStruct, err := url.Parse(repoURL)
	if err != nil {
		return fmt.Errorf("unable to parse url: %v", err)
	}

	if urlStruct.Scheme != "https" {
		return fmt.Errorf("the repo url %q needs to have 'https://' as the scheme", repoURL)
	}

	pathArray := strings.Split(urlStruct.Path, "/")

	if len(pathArray) != 3 {
		return fmt.Errorf("the path of the repository URL %q should only have two layers: /repo-owner/repo-name", repoURL)
	}

	r.URL = repoURL
	r.Name = pathArray[2]
	r.Owner = pathArray[1]

	return nil
}

func main() {
	var repos []Repo
	reposToClone := os.Args[1:]
	for _, repoToClone := range reposToClone {
		var repo Repo
		if err := repo.set(repoToClone); err != nil {
			panic(err)
		}

		repos = append(repos, repo)
	}

	// Cloning the linked repositories, packaging the helm charts, then moving them into the kube-hack repo
	repoDir := path.Join("https:", "github.com")
	for _, repo := range repos {
		if err := cloneRepo(repo.URL); err != nil {
			panic(err)
		}

		chartDir := path.Join(repoDir, repo.Owner, repo.Name, "chart")

		// Packaging the chart in the root directory
		if err := packageChart(".", chartDir); err != nil {
			panic(err)
		}
	}

	// Updating the index.yaml file to include the new packages
	if err := indexCharts(); err != nil {
		panic(err)
	}

	// Removing the cloned repos
	if err := os.RemoveAll("https:"); err != nil {
		panic(err)
	}
}

func cloneRepo(url string) error {
	/*
		This will create the following directory structure:
		- https:
			- github.com
				- repository-owner
					- repository-name
		Example directory: https:/github.com/kube-hack/charts
	*/

	if _, err := git.PlainClone(url, false, &git.CloneOptions{URL: url}); err != nil {
		return fmt.Errorf("unable to clone directory %q: %v", url, err)
	}

	return nil
}

func packageChart(packageDir, chartDir string) error {
	cmd := exec.Command("helm", "package", chartDir, "--destination", packageDir)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unable to package Helm chart in %q: %v\n%s", packageDir, err, string(outputBytes))
	}

	return nil
}

func indexCharts() error {
	cmd := exec.Command("helm", "repo", "index", ".")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unable to index repository: %v\n%s", err, string(outputBytes))
	}

	return nil
}
