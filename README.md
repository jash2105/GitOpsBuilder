# GitOps-Builder

GitOps-Builder is an innovative Kubernetes operator that shifts the paradigm of continuous integration processes from being event-driven to a GitOps-centric approach. It leverages the GitOps Toolkit from Flux to craft CI pipelines based on GitOps principles.

This tool primarily focuses on `DockerBuild` operations, enabling users to execute `docker build` and `docker push` commands within a repository hosting a `Dockerfile`.

## Introduction

To get started with GitOps-Builder, ensure you have the [GitOps Toolkit](https://fluxcd.io/docs/components/) installed on your Kubernetes cluster. For those new to Flux, begin your journey at [https://fluxcd.io/docs/get-started/](https://fluxcd.io/docs/get-started/) for a comprehensive guide.

### Setting Up

**IMPORTANT:** Ensure Docker is running on your system to execute the controller locally.

Follow these steps to install Flux on your Kubernetes cluster:

```bash
flux install
```

Next, enable port forwarding to the source-controller artifacts server:

```bash
kubectl -n flux-system port-forward svc/source-controller 8080:80
```

Set the `SOURCE_CONTROLLER_LOCALHOST` environment variable:

```bash
export SOURCE_CONTROLLER_LOCALHOST=localhost:8080
```

To run GitOps-Builder locally, execute:

```
make install
make run
```

### Quick Guide
#### Establish a DockerHub Repository

Initiate by creating a `gitops-builder-example` repository on DockerHub.

#### Configure a Git Repository Source

Define a source object linking to a Git repository that houses your application and a `Dockerfile`:

```yaml
apiVersion: source.toolkit.fluxcd.io/v1beta1
kind: GitRepository
metadata:
  name: example
  namespace: default
spec:
  interval: 5m
  url: https://github.com/jash2105/GitOpsBuilder
  ref:
    branch: main
```

#### Docker Hub Secret Creation

Generate a Docker Hub secret with your DockerHub credentials and server address.

```bash
kubectl create secret generic dockerhub-credentials \
--from-literal=username=yourusername \
--from-literal=password=yourpassword \
--from-literal=server="https://index.docker.io/v1/"
```

Resulting Secret:
```yaml
apiVersion: v1
data:
  password: cGFzc3dvcmQ=
  server: aHR0cHM6Ly9pbmRleC5kb2NrZXIuaW8vdjEv
  username: dXNlcm5hbWU=
kind: Secret
metadata:
  name: dockerhub-credentials
  namespace: default
```

#### Setting Up a DockerBuild

Establish a `DockerBuild` resource that refers to the previously configured `GitRepository`.

```yaml
apiVersion: build.contrib.flux.io/v1alpha1
kind: DockerBuild
metadata:
  name: example
  namespace: default
spec:
  interval: 5m
  buildMode: buildPush
  sourceRef:
    kind: GitRepository
    name: example
  containerRegistry:
    repository: <yourusername>/gitops-builder-example
    tagStrategy: commitSHA
    authConfigRef:
      name: dockerhub-credentials
      namespace: default
```

### Development Notes

To run `make test`, you must set `DOCKERBUILD_USERNAME`, `DOCKERBUILD_PASSWORD`, and `DOCKERBUILD_SERVER` as environment variables to perform `docker push` operations to an existing registry.

```
make DOCKERBUILD_USERNAME="yourusername" DOCKERBUILD_PASSWORD="yourtoken" DOCKERBUILD_SERVER="docker-server-address" test
```
=======
```

