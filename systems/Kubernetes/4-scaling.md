### Scaling
---
##### Control Plane Availability
```
To make Kubernetes cluster highly available:
   - master components should be redundant
   - etcd across nodes should be able to communicate and update cluster data
   - API Server is stateless so there is no need for one to communicate with another
   - multiple Schedulers and Controller Managers means chaos, so these should implement leader election
```

##### Horizontal Pod Autoscaling
```
Increases number of Pods to handle more requests for a particular service
```

##### Vertical Pod Autoscaling
```
Provide additional resources (e.g. CPU, memory) to Pods

Components
   - Recommender      : watches resource usage and recommends new values
   - Updater          : kills managed Pods who resource requests don't match recommend values of Recommender
   - Admission Plugin : sets CPU/memory requests for new Pods based on recommended values

Downfalls
   - cannot update running Pods
   - cannot be executed together with HPA
```

##### Cluster Autoscaling
```
Provisions a new Node when there are not enough resources in cluster
```