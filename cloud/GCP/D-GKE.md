### Base Concepts
---
##### Deployment Modes
```
GKE clusters can be "zonal" or "regional".
  * Zonal clusters deploy a single master node.
  * Regional clusters deploy three master nodes in different zones that can be upgraded individually to not risk downtime.
```

##### Node Pools
```
a) Node pools are used to put worker nodes into groups with the same configuration.
b) All nodes are put into the default node pool by default.
```

##### Container Registry
```
a) Google Container Registry can be used to store private images, control access to images, and perform vulnerability scans.
b) GKE clusters are able to access registries in the same project by default.
```

##### Quotas
```
a) GKE has limitations per the documentation: https://cloud.google.com/kubernetes-engine/quotas.
b) GKE forces resource quotas that cannot be removed to protect the stability of the Kubernetes cluster.
```

##### Pricing
```
Users are billed for Google Compute Engine VMs that run as worker nodes, but are not billed for master nodes.
```

### Control Plane
---
##### etcd
```
Recommendation is to keep data stored in etcd lower than 6Gi.
This information can be monitored from Google Cloud console.
```

### Scalability
---
##### Cluster Autoscaler
```
a) GKE will consider the relative cost of instance types among various node pools and attempt to expand the cheapest one.
b) GKE will look for under utilized nodes, reallocate existing resources, and shut them down.
   This behavior is prevented for nodes that have the following:
     * PodDisruptionBudget
     * Pods without controllers
     * Pods with local storage attached
     * Pods with particular taints, tolerations, and affinity rules
     * Pods or nodes have labels explicitly specifying not to evict resources 
```