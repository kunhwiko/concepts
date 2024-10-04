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
b) Liveness probes let users define what it means for a container to be running.
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
        # HTTP status must return 200 ~ 399 for container to be considered alive
        httpGet:
          path: /healthz
          port: 8080
        # give time for the pod to initialize
        # kubelet will not check liveness during this period
        initialDelaySeconds: 30
```

##### Readiness Probe
```
a) Readiness probes let users define what it means for a container application to be able to serve requests.
b) Container may be initialized, but necessary dependencies might still be or have become unavailable.
c) When readiness probe fails, the container's pod is removed from registered service endpoints but containers are not 
   restarted. This ensures requests do not flood a pod with requests it cannot process.
```

##### Init Containers
```
a) Containers that run to completion before all other containers start. These containers can provide setup scripts and 
   take care of non-deterministic initializations. 
b) Init containers can be configured such that probes will start only after init containers are started.
c) If multiple init containers exist, they will run one at a time in sequential order. If any fail, the pod will restart 
   until the containers succeed.
```

### Availability
---
##### Pod Disruption Budget
```
a) Limits the number of replicas that can be down at a given time due to voluntary disruptions (e.g. node draining).
b) Involuntary disruptions cannot be prevented by PDBs but do count towards the budget.
c) Unavailability due to rolling upgrades cannot be prevented by PDBs but do count towards the budget. Availability 
   during rolling upgrades should be configured in the workload's spec.
```
