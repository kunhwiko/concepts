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

##### Systems Availability
```
Best Practices
   - Provide redundancy for all systems
   - Automate hot swapping for when components fail
   - Robust logging, monitoring, alerting
   - Testing
   - Persist raw data in case processed data is corrupted (raw or old data can be kept in cheap storage)
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

### Live Updating
---
##### Rolling Update
```
Updates that gradually update components from the current version to the next

Example
   1. Deployment v1 --> Pod v1, Pod v1
   2. Update to Deployment v2 
   3. Deployment v2 --> Pod v1, Pod v1, Pod v2
   4. Deployment v2 --> Pod v1, Pod v2, Pod v2
   5. Deployment v2 --> Pod v2, Pod v2
```

##### Adapter Service
```
Translates requests/responses during an update

Example
   1. Pod A v1 depends on Pod B v1
   2. Pod B v2 is introduced and is incompatible with Pod A
   3. Introduce adapter that translates requests/responses between Pod A and B
```

##### Blue-green Deployments
```
Process
   1. Prepare a copy of a production environment green with the new version
   2. Use green to test active requests on existing environment blue
   3. Assuming stateless components only, switch active environment to green
   4. Roll back to blue if there are problems
```

##### Canary Deployments
```
More subtle process of Blue-green Deployments that changes little by little at a time

Example
   1. Replace 10% of production Pods to canary Pods (Pods hosting new feature)
   2. Gradually increase number of canary Pods to production
```

### Scheduling
---
##### Node Selector
```
Node Selector
   - Pod spec that specifies which nodes to schedule to
```

##### Taint and Tolerations
```
Taint
   - Taint a Kubernetes Node to prevent Pods from being scheduled onto that Node 
```
