# kubeapps apprepository-controller Developer Guide

The `apprepository-controller` is a Kubernetes controller for managing Helm chart repositories added to kubeapps.

## Prerequisites

- [Git](https://git-scm.com/)
- [Make](https://www.gnu.org/software/make/)
- [Go programming language](https://golang.org/dl/)
- [Docker CE](https://www.docker.com/community-edition)
- [Kubernetes cluster (v1.8+)](https://kubernetes.io/docs/setup/pick-right-solution/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Telepresence](https://telepresence.io)

*Telepresence is not a hard requirement, but is recommended for a better developer experience*

## Environment

```bash
export GOPATH=~/gopath/
export PATH=$GOPATH/bin:$PATH
export kubeapps_DIR=$GOPATH/src/github.com/kubeapps/kubeapps
```

## Download the kubeapps source code

```bash
git clone --recurse-submodules https://github.com/kubeapps/kubeapps $kubeapps_DIR
```

The `apprepository-controller` sources are located under the `cmd/apprepository-controller/` directory of the repository.

```bash
cd $kubeapps_DIR/cmd/apprepository-controller
```

### Install kubeapps in your cluster

kubeapps is a Kubernetes-native application. To develop and test kubeapps components we need a Kubernetes cluster with kubeapps already installed. Follow the [kubeapps installation guide](../../chart/kubeapps/README.md) to install kubeapps in your cluster.

### Building `apprepository-controller` binary

```bash
go build
```

This builds the `apprepository-controller` binary in the working directory.

### Running in development

Before running the `apprepository-controller` binary on the development host we should stop the existing controller that is running in the development cluster. The best way to do this is to scale the number of replicas of the `apprepository-controller` deployment to `0`.

```bash
kubectl -n kubeapps scale deployment apprepository-controller --replicas=0
```

> **NOTE** Remember to scale the deployment back to `1` replica when you are done

You can now execute the `apprepository-controller` binary on the developer host with:

```bash
./apprepository-controller --logtostderr --repo-sync-image=quay.io/helmpack/chart-repo:myver --kubeconfig ~/.kube/config
```

Performing application repository actions in the kubeapps dashboard will now trigger operations in the `apprepository-controller` binary running locally on your development host.

### Running tests

To start the tests on the `apprepository-controller` execute the following command:

```bash
go test
```

## Building the kubeapps/apprepository-controller Docker image

To build the `kubeapps/apprepository-controller` docker image with the docker image tag `myver`:

```bash
cd $kubeapps_DIR
make IMAGE_TAG=myver kubeapps/apprepository-controller
```
