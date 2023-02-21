### Cross Component Networking
---
##### Container to Container Networking
```
a) Containers in the same pod share IP addresses and network namespaces.
   This means containers in the same pod can communicate with one another via localhost. 
b) CRI is responsible for creating new Linux namespaces on the Kubernetes node.
   Each pod is assigned a Linux namespace and CNI is responsible for giving the pod a unique IP address.
```

##### Pod to Pod Networking
```
Step 1) CNI sets up a virtual ethernet in the pod's Linux namespace.
        This VETH is paired with the VETH of the node's root namespace via a network bridge.
Step 2) Subnet masking determines if endpoint is on the same network.
Step 3) If two pods are communicating from within the same node, requests are resolved via ARP.
        The request will jump from the current namespace's VETH --> root namespace's VETH --> target namespace's VETH.
Step 4) For pod communications across different nodes, ARP will check for the MAC address of the Kubernetes default gateway.
        The request will jump from the current namespace's VETH --> root namespace's VETH --> default gateway --> route to correct node. 
```

##### Pod to Service Networking
```
Step 1) Services are pieces of data stored in etcd that use endpoint objects as a lookup table to fetch target pod IP addresses.
        When new endpoints are added, all kube proxies will update IP tables on the root namespace of their node based on info stored in etcd.
Step 2) Packets leaving a pod will trigger prerouting chains as part of IP table rules on the node.
        Rules specific for the service will be searched in sequential order until a match is found.
Step 3) Once a match is found, rules for what pod the packet should be sent to must be chosen.
        If there are more than one pod that are backed by the service, there will be several applicable rules to choose from.
        IP table uses a Linux module called "statistic" to choose between those rules through round robin. 
Step 4) IP table rewrites the destination IP address via DNAT to a target pod's IP address chosen by round robin.

More here: https://www.tkng.io/services/clusterip/dataplane/iptables/
```

##### Service to Pod Networking
```
Step 1) After a packet reaches a target pod, a response needs to be sent back.
        However, the source expects a response back from the service's IP address and not the pod's IP address.
Step 2) IP tables use Linux connection tracking to remember previous routing choices.
Step 3) IP tables will masquerade the packet's source IP address to be the service's IP address through SNAT.
```

##### IP Table vs IPVS Networking
```
Problem Statement
  a) IP table is not meant to provide load balancing functionality and is purely meant for firewall rules.
  b) IP table only allows for round robin load balancing.
  c) Sequential search is done across all IP table rules to find which rules are applicable to the packet.
     This results in O(n) lookup time where n is the number of services and pods in the entire cluster.

IPVS Networking
  a) IPVS creates a dummy interface on each node and binds service IP addresses to the dummy interface.
     IPVS virtual servers are created for each service IP address.
  b) IPVS optimizes lookup through hash tables managed by the kernel to achieve a O(1) lookup time.
  c) IPVS supports multiple load balancing algorithms.
  d) When using IPVS mode, IP sets are used in cases IP tables are required (e.g. packet filtering, SNAT) for more efficient lookup.

More here: https://www.tkng.io/services/clusterip/dataplane/ipvs/
```

##### Asynchronous Networking
```
a) Queues are a great way to achieve asynchronous communication and decouple various Kubernetes components.
   Queues abstract away the need to know about IP addresses, ports, and Kubernetes services.
b) Containers can listen or respond to messages, perform actions, and post progress to queues.
   It is also easy to add or remove listeners and keep track of progress by monitoring the queue.
c) The queue can be used alongside databases and must be highly available.
```

### Load Balancing
---
##### Cloud Provider Networking
```
External access from the cluster usually involves a public load balancer that directs requests to any Kubernetes node.
Kube proxy on the node will then redirect the request to the correct pod on the correct node.
This is the typical process for cloud providers.
```

##### Cluster Wide Endpoint, Node Local Endpoint, and Client IP Preservation 
```
a) Kubernetes has the ability to preserve information about the originating client IP address. 
b) Ability to preserve client IP address differs based on how 'service.spec.externalTrafficPolicy' is set.
   This field specifies whether to route external traffic to a node-local endpoint or cluster-wide endpoint.

Cluster Wide Endpoint
  a) Default Kubernetes behavior that does not preserve client IP address.
  b) Might make network hops to other nodes but spreads load efficiently.

Node Local Endpoint
  a) Preserves client IP address.
  b) Does not make network hops but does not spread load well.
```

##### External Load Balancing
```
External load balancers distributes load at a node level rather than a pod level.
If 3 pods exist in node A and 1 pod exists in node B, load will be distributed equally to both nodes.
```

### Container Network Interface (CNI)
---
##### CNI
```
Initiative and specification to write various networking solutions via plugins to configure network interfaces for Linux containers.
Users are then able to adopt networking solutions and the container orchestration system of their choice according to different needs.
Vendors do not need to worry about Kubernetes source code or be locked down to Kubernetes.
```

##### CNI Plugin
```
CNI plugin performs:
  a) Add network interface to container network namespace and bridge the container to the host namespace via a VETH pair.
  b) Assign unique IP addresses to CNI containers (e.g. pods in the case of Kubernetes) via an IP Address Management (IPAM) plugin.
  c) Take care of routing logic.

CNI plugin must implement and support:
  a) ADD / DEL container to network
  b) CHECK container's network status
  c) VERSION reporting

CNI plugin execution:
  a) CRIs pass network configurations in JSON format to CNI plugins and invoke CNI plugins as executables (e.g. invoke ADD verb).
```

##### CNI Workings
```
Step 1) Container runtime specifies actions (e.g. add container) it wants to execute on CNI plugin.
Step 2) Input network configurations stored as JSON files are picked up and streamed to the plugin via STDIN.
        Runtimes typically specify what path to look for when picking up JSON files. 
        Examples of input configurations:
          * CNI version
          * Type of plugin to use (e.g. bridge plugin)
          * IPAM and DNS configurations 
Step 3) Supplementary information are further provided via environment variables.
        Examples of environment variables:
          * Desired operations (e.g. ADD)
          * Container ID
          * Path to network namespace file
          * Name of the network interface to set up
          * Path to CNI plugin executable
Step 4) CNI plugin takes STDIN and environment variables to perform operations.
Step 5) CNI plugin outputs the result (generated network interfaces) as STDOUT in JSON format.
```

##### CNI Flat Networking Model
```
a) All pods in the cluster are assigned IP addresses from the cluster's IP pool.
b) Easy to set up and monitor network traffic but could easily exhaust all of the available IP addresses.
```

##### CNI Overlay Networking Model
```
a) Encapsulates packets coming from the underlaying network at the secondary network level when going to another node.
b) Typically uses VXLAN (tunneling L2 domains over L3 networks).
```

### Network Tools
---
##### Networking Solutions
```
Networking Solution Tools
  a) Flannel
  b) Calico
  c) Weave Net
  d) Romana
  e) Contiv
```