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

##### Container Optimized OS
```
GCE nodes in GKE by default use Google's Container-Optimized OS, which is based on Chromium OS.
This OS is optimized for running containerized applications and is stripped of unnecessary features to reduce attack surface.
```

##### Google Container Registry
```
a) Google Container Registry can be used to store private images, control access to images, and perform vulnerability scans.
b) GKE clusters are able to access registries in the same project by default.
```

##### Workload Identities
```
Workload identities allow Kubernetes service accounts to act as IAM service accounts.
This is done by adding an IAM policy binding to the Google service account.
Afterwards an annotation is placed on the Kubernetes service account.
```

##### GKE IAM Roles
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

### Networking
---
##### Control Plane Networking
```
a) Master nodes and worker nodes in GKE run in different VPCs and are connected through VPC peering.
b) Master node IP addresses are static and do not change, but it is possible to initiate an IP rotation.
```

##### Public and Private Clusters
```
GKE supports both public and private clusters.
  * Public clusters expose the control plane and nodes with external IP addresses.
  * Private clusters do not expose nodes with external IP addresses.
    In this mode, public endpoints to the control plane are still exposed by default but can be disabled.
    In this mode, user must specify a /28 CIDR range for the private endpoint to the control plane.
```

##### Routes Based Networking
```
Routes-based networking use Google Cloud Routes for routing, but this method is no longer recommended.
In this mode, pod IP addresses are not assigned to the VPC and are not natively routed.
Rather, nodes reserve a unique /24 CIDR range for pods that are allocated to itself.
GKE automatically creates routes using this CIDR range as a destination IP range and the node as the next hop.
```

##### VPC Native Networking
```
VPC-native clusters use alias IP addresses and is the recommended network mode.
In this mode, primary CIDR ranges are reserved for nodes, but secondary CIDR ranges can be reserved for pods and services.
The NIC for nodes will continue to be assigned an IP address from the primary CIDR range.
Using alias IP features, the NIC is additionally assigned a /24 IP address range from the secondary CIDR range.

This feature creates the following advantages over routes-based networking:
  a) Pod IP address ranges do not depend on static routes and pod IP addresses are natively routable.
  b) Firewall rules can be applied to pod IP address ranges.
  c) Pod IP address ranges are accessible from on-prem networks connected with Cloud VPN or Interconnect.
```

##### IP Masquerade Agent
```
IP Masquerade Agent is a feature that can SNAT the source IP of packets sent from a pod to be the node's IP address.
```