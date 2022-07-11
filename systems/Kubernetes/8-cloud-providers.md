### Cloud Providers
---
##### Kubernetes Managed Services
```
Kubernetes Flavors
   a) AKS: K8s on Azure
   b) EKS: K8s on AWS
   c) GKE: K8s on GCP

Container Management
   a) examples: Azure Container Instance (ACI), AWS Fargate, Google Cloud Run
   b) compatible and are used as part of AKS, EKS, GKE
   b) offer means to deploy containers without having to provision instances or install container runtimes
   b) offer means to find a machine to run containers

Examples with ACI
   a) utilizes the concept of provisioning virtual nodes, which are not backed by actual VMs
   b) virtual nodes do not rely on cluster autoscaler and so do not need to wait for number of VM nodes to scale up 
   b) deploys containers only when necessary in a quick fashion
```