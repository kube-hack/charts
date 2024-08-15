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
helm search repo
```

### Navigating the linked repositories

Linked below are repositories containing source code and walkthroughs for each vulnerability. Each will include an introductory `README.md` file and a `solutions` directory. If you would prefer to practice hacking an application without knowing too many details, follow the installation instructions in the root `README.md` file, and don't read the files in the solutions directory.

# Linked Repositories

## Application Vulnerabilities

- [SQL Injection](https://github.com/kube-hack/sql-injection)

## Platform Vulnerabilities

More coming soon!