### Probes
---
##### Startup Probe
```
a) Defines what it means for a container to have started.
b) Disables liveness and readiness checks until startup succeeds.
c) Provides time for the container application to initialize.
```

##### Liveness Probe
```
a) Kubelets have basic restart policies if a container process crashes, but it might not be sufficient.
b) Liveness probes let users define what it means for a container to be alive.
```

##### Liveness Probe Example
```yaml
apiVersion: v1
kind: Pod
spec:
  containers:
    - name: app
      image: test-app
      livenessProbe:
        # HTTP status must return 200 ~ 399 for container to be alive
        httpGet:
          path: /healthz
          port: 8080
        # give time for the pod to initialize
        # kubelet will not check liveness for this period
        initialDelaySeconds: 30
```

##### Readiness Probe
```
a) Readiness probes let users define what it means for a container application to be able to serve requests.
b) Container may be initialized, but necessary dependencies might still be or have become unavailable.
c) When readiness probe fails, container's pod are removed from registered service endpoints but containers are not restarted.
   This ensures requests do not flood a pod with requests it cannot process.
```

##### Init Containers
```
a) Containers that run to completion before all other containers start. These containers can provide setup scripts and 
   take care of non-deterministic initializations. 
b) Init containers can be configured such that probes will start only after init containers are started.
c) If multiple init containers exist, they will run one at a time in sequential order. If any fail, the pod will restart 
   until the containers succeed.
```

##### Readiness Gates
```
a) Readiness probes help address pod level readiness, but not an overall infrastructure level readiness.
   Readiness gates address can address infrastructure level readiness.
b) A pod spec called readinessGates can specify set of conditions for when things are ready.
```

### Availability
---
##### Pod Disruption Budget
```
a) Limits the number of replicas that can be down at a given time due to voluntary disruptions. When using PDBs, it is
   recommended to drain nodes rather than directly deleting pods or deployments.
b) Pods deleted by involuntary disruptions and rolling upgrades both count towards the budget. Workload resources are however
   not limited by PDBs when conducting rolling upgrades, and availability should be configured via each workload's spec.
```
