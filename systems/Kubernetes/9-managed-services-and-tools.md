### Azure AKS
---
##### Azure Container Instance
```
ACI
   a) Non-K8s specific solution that allows for on-demand containers on Azure.
   b) Deploys containers onto a virtual node, removing need to provision instances or install container runtimes.
      Kubernetes recognizes the virtual node as a node with infinite resources.
   c) Behind the scenes, offers a means to automatically find a machine to run containers for you.
   b) For known workloads, nodes should be provisioned in a traditional way.
      Extra loads will then be bursted dynamically to ACI.
```

### Amazon EKS
---
##### Fargate
```
Fargate
   a) Non-K8s specific solution that allows for on-demand containers on AWS.
   b) While a single infinite virtual node represents ACI, Fargate puts each pod onto its own virtual node.
      This still removes the need to provision instances or install container runtimes.
```

### Google GKE
---
##### Google Cloud Run
```
Cloud Run
   a) Non-K8s specific solution that allows for on-demand containers on GCP.
   b) Cloud Run is based off of Knative.
```

### Serverless Computing Tools
---
##### Knative
```
Knative
   a) Supports use cases for scaling to zero for non-frequent operations.
   b) Project that aims to build serverless and event driven applications.
   c) Supported by two components, Knative Serving and Knative Eventing.

Knative Serving
   a) Comprised of 4 CRDs: Service, Route, Configuration, Revision.
   b) Service
      * Manages the lifecycle of workloads and acts as both a Kubernetes deployment and service. 
      * Responsible for creating Routes and Configurations.
      * Manages new Revisions whenever the Service is updated.
   c) Route
      * Able to customize logic to specify how to route traffic to different Revisions.
   d) Configuration
      * Contains records of the latest Service, number of generations, and Revisions.
      
Knative Eventing
   a) Instead of hitting API endpoints, Knative uses event publishing.
      Events can be modelled as an an EventType CRD.
      A broker will identify these events and match them to consumers (e.g. Knative Service) via triggers.
```

### Testing Tools
---
##### Kubemark
```
Kubemark
   a) Does not test actual real life behavior for the sake of cost.
   b) Runs mock hollow nodes, hollow kubelets, hollow proxies that fake functionalities in a lightweight manner.
   c) Still very good at testing improvements and regressions to the cluster.
```


# Below is Deprecated

### Network Tools
---
##### Keepalived Virtual IP
```
Problem
   a) Clients need a stable endpoint, but pods and sometimes load balancers move around in Kubernetes.
   b) DNS resolution for load balancers and services are not good enough due to performance issues.

Solution
   a) Keepalived provides high performance virtual IP address that serves address of load balancers / ingress controllers.

Linux functionalities
   a) IPVS (IP virtual server).
   b) High availability via Virtual Redundancy Router Protocol (VRRP).
   c) Operates at networking layer 4 level.
```

### Security Tools
---
##### AppArmor
```
AppArmor : Linux kernel security module that allows you to create profiles to do the following
   a) Restrict network access of processes in container.
   b) Restrict Linux capabilities of container.
   c) Restrict file permissions of container.
   d) Provide improved auditing through logs.
```
