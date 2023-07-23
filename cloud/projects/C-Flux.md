### Flux
---
##### Flux Definition
```
Flux is a CNCF graduated GitOps tool for keeping Kubernetes clusters in sync with sources of configuration (e.g. Git 
repo). Flux is built with Kubernetes in mind, and integrates well with other tools such as Prometheus, Helm, Istio etc.
```

##### Source Controller
```
Controller that polls commits from a Git repo and Helm charts from artifactories given that proper credentials have been 
provided. This controller will then call Git APIs and use the commit history to apply any new changes (e.g. Kubernetes 
manifests).
```

##### Helm Controller
```
Controller responsible for managing Helm artifacts and applying changes based on a desired state (i.e. described through 
HelmRelease custom resources). 
```

##### Flux Commands
```
fluxctl install  : used to install Flux onto a Kubernetes cluster, typically with Git configurations
fluxctl identity : gives back the RSA key that can be used to read from a given Git repo
fluxctl sync     : used to sync Flux with the Git repo
```
