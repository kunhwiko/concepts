### Availability Designs
----
##### High Availability
```
a) Master components should be redundant.
b) etcd across nodes should be able to communicate with one another and update cluster data.
c) API server is stateless so there is no need for one to communicate with another.
d) Multiple schedulers and controller managers means chaos, so these should implement leader election.
```

##### Systems Availability
```
a) Provide redundancy for all systems.
b) Automate hot swapping for when components fail.
c) Ensure to have robust logging, monitoring, alerting, testing.
d) Persist raw data in case processed data is corrupted (raw or old data can be kept in cheap storage).
```

### Autoscaling
---
##### Horizontal Pod Autoscaler (HPA)
```
a) HPA is able to specify the min and max pod count that a user needs.
b) HPA interacts with replicasets or deployments instead of with pods directly as a source of truth.
   In general, it is recommended to bind HPA to deployments and not replicasets.
   This prevents HPA from being bound to an old replicaset during rolling updates.
c) Scaling does not happen immediately to reduce thrashing issues where average load is around scaling thresholds.
d) HPA respects and evaluates all existing metrics and autoscales based on largest number of replicas required.
```

##### Cluster Autoscaler
```
a) Provisions a new node when not enough nodes necessary are provisioned.
b) Downscaling times usually have longer delays to prevent thrashing.
```

##### Vertical Pod Autoscaler
```
a) Provides additional resources (e.g. CPU, memory) to pods.
b) Cannot update running pods (i.e. must bring down existing pods).
c) Cannot be executed together with HPA.
```

##### Vertical Pod Autoscaler Components
```
a) Recommender      : Watches resource usage and recommends new values.
b) Updater          : Kills managed pods where resource requests don't match recommend values of Recommender.
c) Admission Plugin : Sets CPU/memory requests for new pods based on recommended values.
```

### Live Updating
---
##### Rolling Update
```
Gradually updates components from the current version to the next.

Example
  Step 1) Deployment v1 --> Pod v1, Pod v1
  Step 2) Update to Deployment v2 
  Step 3) Deployment v2 --> Pod v1, Pod v1, Pod v2
  Step 4) Deployment v2 --> Pod v1, Pod v2, Pod v2
  Step 5) Deployment v2 --> Pod v2, Pod v2
```

##### Adapter Service
```
Translates requests/responses during an update.

Example
  Step 1) Pod A v1 depends on pod B v1.
  Step 2) Pod B v2 is introduced and is now incompatible with pod A v1.
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
More subtle process of blue-green deployments that changes gradually over time.

Example
  Step 1) Replace 10% of production pods to canary pods (pods hosting new feature).
  Step 2) Gradually increase number of canary pods to production.
```

### Scheduling
---
##### Node Selector
```
A pod spec that specifies which nodes it should be scheduled to.
```

##### Taints
```
a) Taint a Kubernetes node to prevent pods from being scheduled onto that node.
b) A given node can have multiple taints.
c) NoSchedule taints prevent pods from being scheduled.
d) NoExecute taints prevent pods from being scheduled and also evict existing pods without proper tolerations.
```

##### Tolerations
```
Specify that a pod can tolerate a specific taint.
```

##### Node Affinity
```
a) Sophisticated selection criteria of assigning pods to particular nodes.
b) Ability to specify "preferred" requirements rather than hard requirements.
c) Ability to specify anti-affinity requirements.
d) If both node selectors and node affinities exist, pods will be scheduled to nodes that match both requirements.
```

##### Pod Affinity
```
a) Able to assign pods based on labels of existing pods.
b) Ability to assign a pod only if some other pod co-exists.
c) Ability to not co-locate a pod with some other pod (i.e. anti-affinity).
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
