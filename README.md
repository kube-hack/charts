# Overview

This is an educational resource designed to demonstrate various application and platform vulnerabilities within Kubernetes, as well as provide an opportunity to practice different hacking techniques.

## \*\*\*\*\**DISCLAIMER*\*\*\*\*\*

This contains applications with a built-in security vulnerabilities. Please don't deploy the Helm charts into a production environment. There are also instructions showing how to exploit different application and platform vulnerabilities, so please don't use this to break any laws :grin:.

## How to use this repository

### Requirements

- Latest version of [Helm](https://helm.sh/docs/intro/install/)
- Latest version of [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- A fully-compliant Kubernetes distribution (i.e. microk8s, k3s, k3d) that is running on Linux/amd64, and is using containerd or Docker as the runtime.

### Using the Helm repository

You can add the Helm repository locally by running the below command:

```sh
helm repo add kube-hack https://kube-hack.github.io/charts
helm repo update
```

If you would like to see the available charts in the `kube-hack` repository, run the below command:

```sh
helm search repo kube-hack
```

### Navigating the linked repositories

Linked below are repositories containing source code and walkthroughs for each vulnerability. Each will include an introductory `README.md` file and a `solutions` directory. If you would prefer to practice hacking an application without knowing too many details, follow the installation instructions in the root `README.md` file, and don't read the files in the solutions directory.

# Linked Repositories

## Application Vulnerabilities

- [SQL Injection](https://github.com/kube-hack/sql-injection)
- [Command Injection](https://github.com/kube-hack/command-injection)

## Platform Vulnerabilities

More coming soon!

# Contributing

## Requirements

- Latest version of [Go](https://go.dev/dl/)
- A computer running Linux or MacOS (this might work on Windows, but we haven't tested :grimacing:)

## Steps to add a chart to the repository

1. Create a public Github repository with a directory named `chart` containing the Helm chart files (see the linked repositories for examples)

2. Fork the `kube-hack/charts` repository and clone it to your computer. If you want to test the newly-added files on your fork, you'll need to configure your Github repository to answer GET requests for `YAML` and `tar` files. See Helm's [Chart Repository Guide](https://helm.sh/docs/topics/chart_repository/) for more details.

3. In your terminal, navigate to the root directory of the cloned charts repo and run the command below with the URLs of the repositories you wish to add as arguments.  Running this command will clone the repositories provided, package their respective Helm charts, and re-index the Helm repository:
    ```sh
    go run main.go https://github.com/example-owner/example-repo-1 https://github.com/example-owner/example-repo-2
    ```
4. Add, commit, and push the changes to the charts repository to your fork, then make a pull request to the `kube-hack/charts` repo. We will review the code and determine if the chart is a meaningful addition to the project.