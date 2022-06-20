### Designs
----
##### Availability Designs
```
High Availability
   a) master components should be redundant
   b) etcd across nodes should be able to communicate and update cluster data
   c) API server is stateless so there is no need for one to communicate with another
   d) multiple schedulers and controller managers means chaos, so these should implement leader election

Systems Availability
   a) provide redundancy for all systems
   b) automate hot swapping for when components fail
   c) robust logging, monitoring, alerting, testing
   e) persist raw data in case processed data is corrupted (raw or old data can be kept in cheap storage)
```

##### Distributed System Design Patterns
```
Sidecar Pattern
   Step 1) in the same pod, create a separate container from the main application container
   Step 2) this separate container provides supplemental features such as logging
   Step 3) places less burden on application container

Ambassador Pattern
   Step 1) like the Sidecar Pattern, create a separate container from the main app container
   Step 2) this container acts as a proxy to the main app container that filters requests
   Step 3) often used with legacy apps that are risky to modify to extend networking/security configurations
   Step 4) able to update configurations of ambassador while keeping legacy code

Adapter Pattern
   Step 1) assume main application has been updated but generates output in a different format
   Step 2) consumers of the output have not been upgraded to read in the new format
   Step 3) adapter standardizes output until all consumers have been upgraded 
```

### Live Updating
---
##### Rolling Update
```
Rolling Update : gradually update components from the current version to the next

Example
   Step 1) Deployment v1 --> Pod v1, Pod v1
   Step 2) update to Deployment v2 
   Step 3) Deployment v2 --> Pod v1, Pod v1, Pod v2
   Step 4) Deployment v2 --> Pod v1, Pod v2, Pod v2
   Step 5) Deployment v2 --> Pod v2, Pod v2
```

##### Adapter Service
```
Adapter Service : translates requests/responses during an update

Example
   Step 1) pod A v1 depends on pod B v1
   Step 2) pod B v2 is introduced and is incompatible with pod A
   Step 3) introduce adapter that translates requests/responses between pod A and B
```

##### Blue-green Deployments
```
Process
   Step 1) prepare a copy of a production environment green with the new version
   Step 2) use green to test active requests on existing environment blue
   Step 3) assuming stateless components only, switch active environment to green
   Step 4) Roll back to blue if there are problems
```

##### Canary Deployments
```
Canary Deployments : more subtle process of Blue-green Deployments that changes gradually over time

Example
   Step 1) replace 10% of production pods to canary pods (pods hosting new feature)
   Step 2) gradually increase number of canary pods to production
```

### Resource Quotas
---
##### Quota Types
```
Resource Quota
   a) there cannot be conflicts for ResourceQuota object per namespace
      * kubectl get quota
   b) can specify CPU, memory, GPU quotas
   c) can specify total amount of storage and number of PVCs / ephemeral storage per cluster or per storage class
   d) can specify a limit for the number of each Kubernetes resource
   e) can specify quota scopes
      * quota can be specified for only non-terminating pods
      * even if quota is exceeded, new pods can be scheduled if existing pods are terminating
```

##### Priority Classes
```
Priority Class : Prioritize scheduling of pods when resources are scarce
```

### Scaling
---
##### Horizontal Pod / Cluster Autoscaling
```
Horizontal Pod Autoscaling
   a) increases number of pods to handle more requests for a particular service
   b) synchronizes effort and interaction with replica controller instead of with pods directly
   c) scaling is not done immediately to reduce thrashing issues where average load is around scaling thresholds

Cluster Autoscaling : provisions a new node when there are not enough resources in cluster
```

##### Autoscaling Behaviors
```
Metrics
   a) respects and evaluates all existing metrics and autoscales based on largest number of replicas required
   b) for custom metrics, must enable API aggregation layer and then register resource / custom metrics API

Rolling Updates
   a) during rolling updates, HPA is bound to the old replication controller
   b) recommended to bind autoscaling to the deployment and not the replication controller
```

##### Vertical Pod Autoscaling
```
Vertical Pod Autoscaling : Provide additional resources (e.g. CPU, memory) to pods

Components
   a) Recommender      : watches resource usage and recommends new values
   b) Updater          : kills managed pods who resource requests don't match recommend values of Recommender
   c) Admission Plugin : sets CPU/memory requests for new pods based on recommended values

Downfalls
   a) cannot update running pods
   b) cannot be executed together with HPA
```

### Scheduling
---
##### Node Selector & Affinity
```
Node Selector : pod spec that specifies which nodes to schedule to

Node Affinity
   a) more sophisticated selection criteria to assign pods to particular nodes
   b) ability to specify "preferred" requirements rather than hard requirements
   c) assigned pods will not be evicted from nodes even if labels are changed and no longer satisfy node affinity

Pod Affinity
   a) able to assign pod based on labels of existing pods
   b) can make sure to assign a pod only if some other pod co-exists
   c) can make sure that a pod does not co-locate with some other pod (anti-affinity)
```

##### Taint and Tolerations
```
Taint
   a) taint a Kubernetes node to prevent pods from being scheduled onto that node
   b) nodes can have multiple taints

Toleration : specify that a pod can tolerate a specific taint
```

##### Daemon Sets
```
Daemon Set
   a) ensures that a pod runs on all or a designated subset of nodes 
   b) useful for monitoring or aggregating multiple small requests into a single network request
```
