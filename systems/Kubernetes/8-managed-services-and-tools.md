### Cloud Providers
---
##### Managed Kubernetes Services
```
Kubernetes Flavors
   a) AKS: K8s on Azure
   b) EKS: K8s on AWS
   c) GKE: K8s on GCP

Container Management
   a) examples: Azure Container Instance (ACI), AWS Fargate, Google Cloud Run
   b) compatible and are used as part of AKS, EKS, GKE
   b) offer means to deploy containers without having to provision instances or install container runtimes
   b) offer means to automatically find a machine to run containers for you

Examples with ACI
   a) utilizes the concept of provisioning virtual nodes, which are not backed by actual VMs
   b) virtual nodes do not rely on cluster autoscaler and so do not need to wait for number of VM nodes to scale up 
   b) deploys containers only when necessary in a quick fashion
```

### Security Tools
---
##### AppArmor
```
AppArmor : Linux kernel security module that allows you to create profiles to do the following
   a) restrict network access of processes in container
   b) restrict Linux capabilities of container
   c) restrict file permissions of container
   d) provide improved auditing through logs
```

### Monitoring Tools
---
##### Prometheus / Grafana
```
Prometheus : time series database that scrapes metrics such as CPU and memory usage via HTTP requests

Grafana : provides means to visualize metrics and provide alerts
```

### Testing Tools
---
##### Kubemark
```
Kubemark
   a) does not test actual real life behavior for the sake of cost
   b) runs mock hollow nodes, hollow kubelets, hollow proxies that fake functionalities in a lightweight manner
   c) still very good at testing improvements and regressions to the cluster
```
