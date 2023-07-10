### Live Updating
---
##### Recreate Strategy
```
Specifies the deployment to kill all existing pods and create new ones.
```

##### Rolling Update Strategy
```
Specifies the deployment to gradually update components from the current version to the next. During a rolling update, 
it is possible to specify the following:
  * Maximum number of pods that can be unavailable.
  * Maximum number of pods that can exist over the desired number of pods (i.e. max surge).
  * Deadline for when the deployment should have finished upgrading.
  * Number of revisions (i.e. replicasets) to retain to allow rollback. 
```

##### Adapter Service
```
Translates requests/responses during an update.

Step 1) Pod A v1 depends on pod B v1.
Step 2) Pod B v2 is introduced and is now incompatible with pod A v1.
Step 3) Introduce adapter that translates requests/responses between pod A and B.
```

##### Blue Green Deployments
```
Step 1) Prepare a copy of a production environment green with the new version.
Step 2) Use green to test active requests on existing environment blue.
Step 3) Assuming stateless components only, switch active environment to green.
Step 4) Rollback to blue if there are problems.
```

##### Canary Deployments
```
More subtle process of blue-green deployments that changes gradually over time.

Step 1) Replace 10% of production pods to canary pods (pods hosting new feature).
Step 2) Gradually increase number of canary pods to production.
```

### Node Maintenance
---
##### Pod Eviction Timeout
```
Time that it takes for pods to be evicted once a node is reported to be unhealthy. The default time for pod eviction is 
5 minutes, and any pods not backed by a controller will be destroyed afterwards. Refer to the following document for 
more: https://dbafromthecold.com/2020/04/08/adjusting-pod-eviction-time-in-kubernetes/.
```

##### Node Drain
```
A node can be drained to mark it as unschedulable and gracefully move existing workloads to other nodes. Once drained,
the node must be uncordoned to be considered ready.
```

##### Node Cordon
```
A node can be cordoned to mark it as unschedulable. This does not migrate existing workloads away from the node. Once
cordoned, the node must be uncordoned to be considered ready.
```

##### Node Upgrades
```
For non-managed clusters using kubeadm, refer to: https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/
```

##### etcdctl
```
For non-managed clusters, refer to: https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/ for 
ETCD related commands.
```