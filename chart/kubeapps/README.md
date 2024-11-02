# KubeApps

![](https://i.imgur.com/waxVImv.png)
### [View all Roadmaps](https://github.com/kubeapps/all-roadmaps) &nbsp;&middot;&nbsp; [Best Practices](https://github.com/kubeapps/all-roadmaps/blob/main/public/best-practices/) &nbsp;&middot;&nbsp; [Questions](https://www.linkedin.com/in/nholuong/)
<br/>

## TL;DR;

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install --name kubeapps --namespace kubeapps bitnami/kubeapps
```

## Introduction

This chart bootstraps a [kubeapps](https://kubeapps.com) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

It also packages the [Bitnami MongoDB chart](https://github.com/helm/charts/tree/master/stable/mongodb) which is required for bootstrapping a MongoDB deployment for the database requirements of the kubeapps application.

## Prerequisites

- Kubernetes 1.8+ (tested with Azure Kubernetes Service, Google Kubernetes Engine, minikube and Docker for Desktop Kubernetes)
- Helm 2.10.0+
- Administrative access to the cluster to create Custom Resource Definitions (CRDs)

## Installing the Chart

To install the chart with the release name `kubeapps`:

```console
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm install --name kubeapps --namespace kubeapps bitnami/kubeapps
```

> **IMPORTANT** This assumes an insecure Helm installation, which is not recommended in production. See [the documentation to learn how to secure Helm and kubeapps in production](https://github.com/kubeapps/master/docs/user/securing-kubeapps.md).

The command deploys kubeapps on the Kubernetes cluster in the `kubeapps` namespace. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Caveat**: Only one kubeapps installation is supported per namespace

> **Tip**: List all releases using `helm list`

Once you have installed kubeapps follow the [Getting Started Guide](https://github.com/kubeapps/master/docs/user/getting-started.md) for additional information on how to access and use kubeapps.

## Upgrading kubeapps

You can upgrade kubeapps from the kubeapps web interface. Select the namespace in which kubeapps is installed (`kubeapps` if you followed the instructions in this guide) and click on the "Upgrade" button. Select the new version and confirm.

> NOTE: If the chart values were modified when deploying kubeapps the first time, those values need to be set again when upgrading.

You can also use the Helm CLI to upgrade kubeapps, first ensure you have updated your local chart repository cache:

```console
$ helm repo update
```

Now upgrade kubeapps:

```console
$ export RELEASE_NAME=kubeapps
$ helm upgrade $RELEASE_NAME bitnami/kubeapps
```

If you find issues upgrading kubeapps, check the [troubleshooting](#error-while-upgrading-the-chart) section.

## Uninstalling kubeapps

To uninstall/delete the `kubeapps` deployment:

```console
$ helm delete --purge kubeapps
$ # Optional: Only if there are no more instances of kubeapps
$ kubectl delete crd apprepositories.kubeapps.com
```

The first command removes most of the Kubernetes components associated with the chart and deletes the release. After that, if there are no more instances of kubeapps in the cluster you can manually delete the `apprepositories.kubeapps.com` CRD used by kubeapps that is shared for the entire cluster.

> **NOTE**: If you delete the CRD for `apprepositories.kubeapps.com` it will delete the repositories for **all** the installed instances of `kubeapps`. This will break existing installations of `kubeapps` if they exist.

If you have dedicated a namespace only for kubeapps you can completely clean remaining completed/failed jobs or any stale resources by deleting the namespace

```console
$ kubectl delete namespace kubeapps
```

## Configuration

For a full list of configuration parameters of the kubeapps chart, see the [values.yaml](values.yaml) file.

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install --name kubeapps --namespace kubeapps \
  --set chartsvc.service.port=9090 \
    bitnami/kubeapps
```

The above command sets the port for the chartsvc Service to 9090.

Alternatively, a YAML file that specifies the values for parameters can be provided while installing the chart. For example,

```console
$ helm install --name kubeapps --namespace kubeapps -f custom-values.yaml bitnami/kubeapps
```

### Configuring Initial Repositories

By default, kubeapps will track the [community Helm charts](https://github.com/helm/charts) and the [Kubernetes Service Catalog charts](nholuongut/kubeapps/service-catalog). To change these defaults, override the `apprepository.initialRepos` object:

```console
$ cat > custom-values.yaml <<EOF
apprepository:
  initialRepos:
  - name: example
    url: https://charts.example.com
EOF
$ helm install --name kubeapps --namespace kubeapps bitnami/kubeapps -f custom-values.yaml
```

### Configuring connection to a custom namespace Tiller instance

By default, kubeapps connects to the Tiller Service in the `kube-system` namespace, the default install location for Helm.

If your instance of Tiller is running in a different namespace or you want to have different instances of kubeapps connected to different Tiller instances, you can achieve it by setting `tillerProxy.host`:

```console
helm install \
  --set tillerProxy.host=tiller-deploy.my-custom-namespace:44134 \
  bitnami/kubeapps
```

### Configuring connection to a secure Tiller instance

In production, we strongly recommend setting up a [secure installation of Tiller](https://docs.helm.sh/using_helm/#using-ssl-between-helm-and-tiller), the Helm server side component.

In this case, set the following values to configure TLS:

```console
helm install \
  --tls --tls-ca-cert ca.cert.pem --tls-cert helm.cert.pem --tls-key helm.key.pem \
  --set tillerProxy.tls.verify=true \
  --set tillerProxy.tls.ca="$(cat ca.cert.pem)" \
  --set tillerProxy.tls.key="$(cat helm.key.pem)" \
  --set tillerProxy.tls.cert="$(cat helm.cert.pem)" \
  bitnami/kubeapps
```

Learn more about how to secure your kubeapps installation [here](https://github.com/kubeapps/master/docs/user/securing-kubeapps.md).

### Exposing Externally

> **Note**: The kubeapps frontend sets up a proxy to the Kubernetes API service, so when when exposing the kubeapps service to a network external to the Kubernetes cluster (perhaps on an internal or public network), the Kubernetes API will also be exposed on that network. See [#1111](https://github.com/kubeappsissues/1111) for more details.

#### LoadBalancer Service

The simplest way to expose the kubeapps Dashboard is to assign a LoadBalancer type to the kubeapps frontend Service. For example:

```console
$ helm install --name kubeapps --namespace kubeapps bitnami/kubeapps --set frontend.service.type=LoadBalancer
```

Wait for your cluster to assign a LoadBalancer IP or Hostname to the `kubeapps` Service and access it on that address:

```console
$ kubectl get services --namespace kubeapps --watch
```

#### Ingress

This chart provides support for ingress resources. If you have an ingress controller installed on your cluster, such as [nginx-ingress](https://github.com/kubeapps/kubeapps/charts/stable/nginx-ingress) or [traefik](https://github.com/kubeapps/kubeapps/charts/stable/traefik) you can utilize the ingress controller to expose kubeapps.

To enable ingress integration, please set `ingress.enabled` to `true`

##### Hosts

Most likely you will only want to have one hostname that maps to this kubeapps installation, however, it is possible to have more than one host. To facilitate this, the `ingress.hosts` object is an array.

##### Annotations

For annotations, please see [this document](https://github.com/kubernetes/ingress-nginx/blob/master/docs/annotations.md). Not all annotations are supported by all ingress controllers, but this document does a good job of indicating which annotation is supported by many popular ingress controllers. Annotations can be set using `ingress.annotations`.

##### TLS

TLS can be configured using setting the `ingress.hosts[].tls` boolean of the corresponding hostname to true, then you can choose the TLS secret name setting `ingress.hosts[].tlsSecret`. Please see [this example](https://github.com/kubernetes/contrib/tree/master/ingress/controllers/nginx/examples/tls) for more information.

You can provide your own certificates using the `ingress.secrets` object. If your cluster has a [cert-manager](https://github.com/jetstack/cert-manager) add-on to automate the management and issuance of TLS certificates, set `ingress.hosts[].certManager` boolean to true to enable the corresponding annotations for cert-manager as shown in the example below:

```console
helm install --name kubeapps --namespace kubeapps bitnami/kubeapps \
  --set ingress.enabled=true \
  --set ingress.certManager=true \
  --set ingress.hosts[0].name=kubeapps.custom.domain \
  --set ingress.hosts[0].tls=true \
  --set ingress.hosts[0].tlsSecret=kubeapps-tls
```

## Troubleshooting

### Forbidden error while installing the Chart

If during installation you run into an error similar to:

```
Error: release kubeapps failed: clusterroles.rbac.authorization.k8s.io "kubeapps-apprepository-controller" is forbidden: attempt to grant extra privileges: [{[get] [batch] [cronjobs] [] []}]...
```

Or:

```
Error: namespaces "kubeapps" is forbidden: User "system:serviceaccount:kube-system:default" cannot get namespaces in the namespace "kubeapps"
```

This usually is an indication that Tiller was not installed with enough permissions to create the resources by kubeapps. In order to install kubeapps, you will need to install Tiller with elevated permissions (e.g. as a cluster-admin). For example:

```
kubectl -n kube-system create sa tiller
kubectl create clusterrolebinding tiller --clusterrole cluster-admin --serviceaccount=kube-system:tiller
helm init --service-account tiller
```

It is also possible, though less common, that your cluster does not have Role Based Access Control (RBAC) enabled. To check if your cluster has RBAC you can execute:

```console
$ kubectl api-versions
```

If the above command does not include entries for `rbac.authorization.k8s.io` you should perform the chart installation by setting `rbac.create=false`:

```console
$ helm install --name kubeapps --namespace kubeapps bitnami/kubeapps --set rbac.create=false
```

### Error while upgrading the Chart

It is possible that when upgrading kubeapps an error appears. That can be caused by a breaking change in the new chart or because the current chart installation is in an inconsistent state. If you find issues upgrading kubeapps you can follow these steps:

> Note: This steps assume that you have installed kubeapps in the namespace `kubeapps` using the name `kubeapps`. If that is not the case replace the command with your namespace and/or name.

1.  (Optional) Backup your personal repositories (if you have any):

```console
kubectl get apprepository --namespace kubeapps -o yaml <repo name> > <repo name>.yaml
```

2.  Delete kubeapps:

```console
helm del --purge kubeapps
```

3.  (Optional) Delete the App Repositories CRD:

> **Warning**: Don't execute this step if you have more than one kubeapps installation in your cluster.

```console
kubectl delete crd apprepositories.kubeapps.com
```

4.  (Optional) Clean the kubeapps namespace:

> **Warning**: Don't execute this step if you have workloads other than kubeapps in the `kubeapps` namespace.

```console
kubectl delete namespace kubeapps
```

5.  Install the latest version of kubeapps (using any custom modifications you need):

```console
helm repo update
helm install --name kubeapps --namespace kubeapps bitnami/kubeapps
```

6.  (Optional) Restore any repositories you backed up in the first step:

```console
kubectl apply -f <repo name>.yaml
```

After that you should be able to access the new version of kubeapps. If the above doesn't work for you or you run into any other issues please open an [issue](https://github.com/kubeappsissues/new).


# 🚀 I'm are always open to your feedback.  Please contact as bellow information:
### [Contact ]
* [Name: Nho Luong]
* [Skype](luongutnho_skype)
* [Github](https://github.com/kubeapps/)
* [Linkedin](https://www.linkedin.com/in/nholuong/)
* [Email Address](luongutnho@hotmail.com)
* [PayPal.me](https://www.paypal.com/paypalme/kubeapps)

![](https://i.imgur.com/waxVImv.png)
![](Donate.png)
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/nholuong)

# License
* Nho Luong (c). All Rights Reserved.🌟 