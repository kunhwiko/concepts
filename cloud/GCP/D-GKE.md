### Base Concepts
---
##### Deployment Modes
```
GKE clusters can be "zonal" or "regional".
  * Zonal clusters deploy a single master node in a single zone.
    If this option is selected, it is possible to deploy identical number of worker nodes to additional zones. 
  * Regional clusters deploy three master nodes in different zones that can be upgraded individually to not risk downtime.
    If this option is selected, it is possible to manually select zones to deploy to.
```

##### Container Registry
```
a) Google Container Registry can be used to store private images, control access to images, and perform vulnerability scans.
b) GKE clusters are able to access registries in the same project by default.
```

##### IAM
```
Kubernetes Engine Admin                   : Provides full access to management of clusters and to Kubernetes resources inside clusters.
Kubernetes Engine Cluster Admin           : Provides full access to management of GKE clusters.
Kubernetes Engine Cluster Viewer          : Provides read only access to get and list GKE clusters.
Kubernetes Engine Developer               : Provides full access to Kubernetes resources inside clusters.
Kubernetes Engine Viewer                  : Provides read only access to Kubernetes resources inside clusters.
Kubernetes Engine Host Service Agent User : Provides GKE service account in the host project to configure shared network resources for cluster management.
                                            Also provides access to inspect firewall rules in the host project.
Kubernetes Engine Node Service Account    : Provides the minimum privileged role to use as a service account for GKE nodes.
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

### Node Pools
---
##### Node Pools
```
a) Node pools are used to put worker nodes into groups with the same configuration.
b) All nodes are put into the default node pool by default.
```

##### Node Pool Upgrades
```
a) Master and worker nodes are upgraded separately.
b) GKE will drain nodes to be upgraded and attempt to reschedule pods before the upgrade.
c) Master nodes can only work with nodes up to two minor versions older than their own version.
d) GKE node pools are configured to auto-update by default.
   Only one node is upgraded at a given time even if multiple node pools exists.
```

##### Node Pool Repairs
```
GKE node pools are configured to auto-repair by default to keep nodes healthy.
Recreation of nodes are triggered when nodes are reported as unhealthy, does not report back, or is out of boot disk space.
Before recreation, GKE will attempt to drain nodes if nodes are responsive or otherwise perform a shut down.
```

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

### Control Plane
---
##### etcd
```
a) The maximum size of the etcd database is 6GB.
b) etcd usage can be monitored from Google Cloud Console.
```