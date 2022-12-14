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

##### Others
```
Other Tools
   a) Fission
   b) Kubeless
   c) Riff
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