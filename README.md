# Kubeapps is a web-based UI for deploying and managing applications in Kubernetes clusters. kubeapps allows you to:

![](https://i.imgur.com/waxVImv.png)
### [View all Roadmaps](https://github.com/nholuongut/all-roadmaps) &nbsp;&middot;&nbsp; [Best Practices](https://github.com/nholuongut/all-roadmaps/blob/main/public/best-practices/) &nbsp;&middot;&nbsp; [Questions](https://www.linkedin.com/in/nholuong/)
<br/>

- Browse and deploy [Helm](https://github.com/helm/helm) charts from chart repositories
- Inspect, upgrade and delete Helm-based applications installed in the cluster
- Add custom and [private chart repositories](docs/user/private-app-repository.md) (supports [ChartMuseum](https://github.com/helm/chartmuseum) and [JFrog Artifactory](https://www.jfrog.com/confluence/display/RTF/Helm+Chart+Repositories))
- Browse and provision external services from the [Service Catalog](nholuongut/kubeapps/service-catalog) and available Service Brokers
- Connect Helm-based applications to external services with Service Catalog Bindings
- Secure authentication and authorization based on Kubernetes [Role-Based Access Control](docs/user/access-control.md)

## Quick Start

Use the Helm chart to install the latest version of kubeapps: 

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install --name kubeapps --namespace kubeapps bitnami/kubeapps
```

For detailed instructions on how to install and use kubeapps follow the [Getting Started Guide](docs/user/getting-started.md).

## Developer Documentation

Please refer to:

- The [kubeapps Build Guide](docs/developer/build.md) for instructions on setting up the build environment and building kubeapps from source.
- The [kubeapps Developer Documentation](docs/developer/README.md) for instructions on setting up the developer environment for developing on kubeapps and its components.

## Next Steps

If you have followed the instructions for [installing kubeapps](docs/user/getting-started.md) check how to [use kubeapps](docs/user/dashboard.md) to easily manage your applications running in your cluster, or [look under the hood to see what's included in kubeapps](docs/architecture/overview.md).

## Useful Resources

- [Walkthrough for first-time users](docs/user/getting-started.md)
- [Detailed installation instructions](chart/kubeapps/README.md)
- [kubeapps Dashboard documentation](docs/user/dashboard.md)
- [kubeapps components](docs/architecture/overview.md)
- [Roadmap](https://github.com/kubeappswiki/Roadmap)

## Differences from Monocular

The [Monocular](https://github.com/helm/monocular) project was designed to run a public search and discovery website for Helm repositories (e.g. https://github.com/kubeapps/kubeapps). Following its 1.0 release, Monocular is focused on delivering the experience for the Helm Hub.

Versions of Monocular 0.7 and older include a basic ability to install, view and delete Helm releases in a Kubernetes cluster. To focus on the Helm Hub experience, these [features have been removed](https://github.com/helm/monocular#looking-for-an-in-cluster-application-management-ui) since the 1.0 release and are no longer supported. We recommend users of Monocular's in-cluster features try kubeapps as it provides a more featured, robust and secure solution to managing Helm applications in your cluster.

# 🚀 I'm are always open to your feedback.  Please contact as bellow information:
### [Contact ]
* [Name: Nho Luong]
* [Skype](luongutnho_skype)
* [Github](https://github.com/nholuongut/)
* [Linkedin](https://www.linkedin.com/in/nholuong/)
* [Email Address](luongutnho@hotmail.com)
* [PayPal.me](https://www.paypal.com/paypalme/nholuongut)

![](https://i.imgur.com/waxVImv.png)
![](Donate.png)
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/nholuong)

# License
* Nho Luong (c). All Rights Reserved.🌟

