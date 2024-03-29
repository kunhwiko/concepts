### Autoscaling
---
##### Horizontal Pod Autoscaler (HPA)
```
a) HPA is able to specify the min and max pod count that a user requires.
b) HPA interacts with replicasets or deployments instead of with pods directly as a source of truth. In general, it is 
   recommended to bind HPA to deployments and not replicasets. This prevents HPA from being bound to an old replicaset 
   during rolling updates.
c) Scaling does not happen immediately to reduce thrashing issues where average load is around scaling thresholds.
```

##### Metrics Server
```
a) HPA requires a metrics server to scrape resource utiliziation (e.g. CPU, memory usage) to be able to know when it 
   needs to scale resources. 
b) The Metrics API can be accessed via "kubectl top".
c) The metrics server is not meant for non-autoscaling purposes (i.e. monitoring purposes). 
```

##### HPA Scaling Metrics
```
a) HPA usually requires a metric server to measure resource usage. 
b) Custom metrics can be configured and exposed for more complex scaling (e.g. memory based scaling).
c) HPA respects and evaluates all existing metrics and autoscales based on largest number of replicas required.
```

##### Cluster Autoscaler
```
Cluster autoscaler provisions a new node when not enough nodes necessary are provisioned. Downscaling times usually 
will have longer delays to prevent thrashing.
```

##### Vertical Pod Autoscaler (VPA)
```
VPA provides additional resources (e.g. CPU, memory) to pods, but it cannot update running pods (i.e. must bring down 
existing pods). It also cannot be executed together with HPA.
```

##### Vertical Pod Autoscaler Components
```
a) Recommender      : Watches resource usage and recommends new values.
b) Updater          : Kills managed pods where resource requests don't match recommend values of Recommender.
c) Admission Plugin : Sets CPU/memory requests for new pods based on recommended values.
```

### Scheduling
---
##### Scheduling Algorithm
```
Step 1) By default, the scheduler will place pods that need to be scheduled into a priority queue. The priority of
        those pods are determined through the "PriorityClass" API.
Step 2) The scheduler goes through a filtering process to remove nodes that are incapable of supporting the pod to be
        scheduled. The process considers free ports, node selectors, node resources, volume locality, cordons, network
        availability, and taints.
Step 3) The scheduler goes through a scoring process through the remaining filtered nodes. The process will consider
        affinities, user provided policies over node preferences (e.g. favor nodes to pack pods into smallest set of nodes),
        and image locality etc.
Step 4) The scheduler will bind the pod to the chosen node and notify the API server.

The scheduling process is extensible and can be customized via plugins based on user needs. Different schedulers with
custom scheduling algorithms can also be deployed, and pods can then choose which scheduler to use.   
```

##### Manual Scheduling
```
If a pod needs to be forcefully scheduled to a particular node, it is possible to manually specify the "nodeName" field on 
the pod's spec. Otherwise, the scheduler will decide what the value of this field will be.
```

##### Multiple Schedulers
```
Users can create custom scheduler profiles that specifies what scheduling plugins to enable or disable. A scheduler
deployment can use a single KubeSchedulerConfiguration that has multiple scheduling profiles, which will prevent
race conditions (e.g. pod allocation to the same node). Pods can then choose which profile to use for scheduling
based on the "schedulerName" field under the pod's spec. 
```

##### Node Selector
```
A pod spec that specifies which nodes it should be scheduled to.
```

##### Taints
```
a) Taint a Kubernetes node to prevent pods from being scheduled onto that node. Note that a given node can have 
   multiple taints.
b) NoSchedule taints prevent pods from being scheduled.
c) NoExecute taints prevent pods from being scheduled and also evict existing pods without proper tolerations.
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

##### Topology Spread Constraints
```
Controls how pods are spread across a Kubernetes cluster (e.g. regions, zones, nodes, and other user-defined topology 
domains).
```

### Quotas
---
##### Resource Quota
```
a) Provides constraints that limit aggregate resource consumption per namespace.
b) Observable through `kubectl get quota`.
```

##### Resource Quota Types
```
Compute Quotas
  * Quotas can specify CPU, GPU, and memory quotas at a namespace level.
  * If a quota is specified, resource request and limit must be set at the container level as well. Container level 
    resource requests and limits are implemented through Linux cgroups.

Storage Quotas
  * Quotas can specify storage quotas at a namespace level. This can also specified per storage class.
  * Quotas can specify total number of PVCs and ephemeral storage. This can also be specified per storage class.

Object Count Quotas
  * Quotas can specify a limit on the number of each Kubernetes resource at a namespace level. Note that replicasets 
    with zero replicas can still overwhelm the API server because validation must still happen.
```

##### Resource Quota Scope
```
Quotas can be customized to target certain resources of a certain state. For example, if only non-terminating pods are
targeted, new pods can be scheduled if existing pods are terminating.
```

##### Limit Ranges
```
If a compute quota is set, users must specify the resource request and limit for each container. Limit range objects
create a default resource request and limit value for containers if not already specified. This is observable through
'kubectl get limit'.
```
