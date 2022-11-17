### Designs
----
##### Availability Designs
```
High Availability
   a) Master components should be redundant.
   b) etcd across nodes should be able to communicate and update cluster data.
   c) API server is stateless so there is no need for one to communicate with another.
   d) Multiple schedulers and controller managers means chaos, so these should implement leader election.

Systems Availability
   a) Provide redundancy for all systems.
   b) Automate hot swapping for when components fail.
   c) Robust logging, monitoring, alerting, testing.
   e) Persist raw data in case processed data is corrupted (raw or old data can be kept in cheap storage).
```

### Live Updating
---
##### Rolling Update
```
Rolling Update
   a) Gradually update components from the current version to the next.

Example
   Step 1) Deployment v1 --> Pod v1, Pod v1
   Step 2) Update to Deployment v2 
   Step 3) Deployment v2 --> Pod v1, Pod v1, Pod v2
   Step 4) Deployment v2 --> Pod v1, Pod v2, Pod v2
   Step 5) Deployment v2 --> Pod v2, Pod v2
```

##### Adapter Service
```
Adapter Service
   a) Translates requests/responses during an update.

Example
   Step 1) Pod A v1 depends on pod B v1.
   Step 2) Pod B v2 is introduced and is incompatible with pod A.
   Step 3) Introduce adapter that translates requests/responses between pod A and B.
```

##### Blue-green Deployments
```
Example
   Step 1) Prepare a copy of a production environment green with the new version.
   Step 2) Use green to test active requests on existing environment blue.
   Step 3) Assuming stateless components only, switch active environment to green.
   Step 4) Roll back to blue if there are problems.
```

##### Canary Deployments
```
Canary Deployments
   a) More subtle process of blue-green Deployments that changes gradually over time.

Example
   Step 1) Replace 10% of production pods to canary pods (pods hosting new feature).
   Step 2) Gradually increase number of canary pods to production.
```

### Quotas
---
##### Quota Types
```
Quota
   a) Observable through `kubectl get quota`.

Compute Quotas
   a) Can specify CPU, GPU, memory quotas.
      If quota is specified, resource request and limit must be set at the container level as well.

Storage Quotas
   a) Can specify total amount of storage and number of PVCs / ephemeral storage per cluster.
      Can also be specified per storage class.

Object Count Quotas
   a) Can specify a limit for the number of each Kubernetes resource.
      Replica sets with zero replicas can still overwhelm the API server because validation must still happen.

Quota Scopes
   a) Quota can be customized to target certain resources at a certain state (e.g. target only non-terminating pods).
      If specified for non-terminating pods and quota is exceeded, new pods can still be scheduled if existing pods are terminating.
```

##### Limit Ranges
```
Problems
   a) If a compute quota is set, users must specify the resource request and limit for each container.

Limit Ranges
   a) Provides a means to set default resource request and limit values for containers, if not specified.
   b) Observable through `kubectl get limit`.
```

##### Priority Classes
```
Priority Class
   a) Prioritize scheduling of pods when resources are scarce.
```

### Scaling
---
##### Horizontal Pod / Cluster Autoscaling
```
Horizontal Pod Autoscaling
   a) Increases number of pods to handle more requests for a particular service.
   b) Synchronizes effort and interaction with replica controller instead of with pods directly.
   c) Scaling is not done immediately to reduce thrashing issues where average load is around scaling thresholds.

Cluster Autoscaling
   a) Provisions a new node when there are not enough resources in cluster.
```

##### Autoscaling Behaviors
```
Metrics
   a) Respects and evaluates all existing metrics and autoscales based on largest number of replicas required.
   b) For custom metrics, must enable API aggregation layer and then register resource / custom metrics API.

Rolling Updates
   a) During rolling updates, HPA is bound to the old replication controller.
      In general, it is recommended to bind autoscaling to the deployment and not the replication controller.
```

##### Vertical Pod Autoscaling
```
Vertical Pod Autoscaling
   a) Provide additional resources (e.g. CPU, memory) to pods.

Components
   a) Recommender      : Watches resource usage and recommends new values.
   b) Updater          : Kills managed pods who resource requests don't match recommend values of Recommender.
   c) Admission Plugin : Sets CPU/memory requests for new pods based on recommended values.

Downfalls
   a) Cannot update running pods.
   b) Cannot be executed together with HPA.
```

### Scheduling
---
##### Node Selector & Affinity
```
Node Selector
   a) Pod spec that specifies which nodes to schedule to.

Node Affinity
   a) More sophisticated selection criteria to assign pods to particular nodes.
   b) Ability to specify "preferred" requirements rather than hard requirements.
   c) Assigned pods will not be evicted from nodes even if labels are changed and no longer satisfy node affinity.

Pod Affinity
   a) Able to assign pod based on labels of existing pods.
   b) Can make sure to assign a pod only if some other pod co-exists.
   c) Can make sure that a pod does not co-locate with some other pod (anti-affinity).
```

##### Taint and Tolerations
```
Taint
   a) Taint a Kubernetes node to prevent pods from being scheduled onto that node.
   b) Nodes can have multiple taints.

Toleration
   a) Specify that a pod can tolerate a specific taint.
```

##### Daemon Sets
```
Daemon Set
   a) Ensures that a pod runs on all or a designated subset of nodes.
   b) Useful for monitoring or aggregating multiple small requests into a single network request.
```
