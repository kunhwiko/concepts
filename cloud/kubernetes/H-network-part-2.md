### Cross Component Networking
---
##### Container to Container Networking
```
Containers in the same pod share network/IPC namespaces and have a common IP address. This means containers in the 
same pod can communicate with one another via localhost or IPC mechanisms. Container runtimes are responsible for setting 
up new Linux namespaces for containers and CNI is responsible for assigning an IP address to the pod.
```

##### Pod to Pod Networking
```
Step 1) CNI sets up a veth pair, one to the pod's network namespace and the other end to the root namespace's network 
        bridge (this depends on CNI implementation - Linux bridges, tunnels, L3 routers). The next steps assume this 
        networking model: https://sookocheff.com/post/kubernetes/understanding-kubernetes-networking-model/.
Step 2) A pod will send a packet to the veth device on its network namespace, which will route to the other end of the 
        veth pair on the root namespace. 
Step 3) Once the request is forwarded to the network bridge, a MAC address table is used to forward requests. This table
        is populated via the ARP protocol.
Step 4) If the packet needs to be forwarded to a pod in the same node, it will forward to the veth pair of the target 
        pod's network namespace. 
Step 5) If the packet needs to be forwarded to a different node, the bridge will send the request to the default gateway. 
        The default gateway will route the packet to the gateway of the node where the target pod resides (logic to know 
        which node to forward to based on the pod's IP will be cloud provider specific). From here, the request will be 
        sent to the node's network bridge and the CNI specific networking solution will apply.
```

##### Pod to Service Networking
```
Step 1) Services are pieces of data stored in etcd that use endpoint objects as a lookup table to fetch target pod IP 
        addresses. When new endpoints are added, all kube proxies will update iptables on the root namespace of their 
        node based on info stored in etcd.
Step 2) Before packets leave the source node, prerouting chains as part of iptable rules will be triggered. Rules 
        specific to the service will be searched in sequential order until a match is found.
Step 3) Once a match is found, rules for what pod the packet should be sent to must be chosen. If there is more than one 
        pod that is backed by the service, iptables use a Linux module called "statistic" to choose which pod's IP to
        DNAT (i.e. rewrite packet's destination IP) into via round robin. DNAT can also be used to change port numbers.
Step 4) The packet will leave the source node to the correct target node with a modified destination IP address. 
```

##### Service to Pod Networking
```
After a packet reaches the target pod, a response will be sent back. When the response comes back into the source node, 
iptables will leverage conntrack to rewrite the response's source IP address to be the service's IP address through SNAT. 
Conntrack is a Linux kernel feature that iptables leverage to remember previous routing choices that were made. This is 
necessary as the source expects a response back from the target service's IP address and not the target pod's IP address.
```

##### IPTable vs IPVS Networking
```
Problem Statement
  * IPtable is not meant to provide load balancing functionality and is purely meant for firewall rules.
  * IPtable only allows for simple round robin load balancing.
  * Sequential search is done across all iptable rules to find which rules are applicable to the packet. This results in 
    O(n) lookup time where n is the number of services and pods in the entire cluster.

IPVS Networking
  * IPVS creates a dummy interface on each node and binds service IP addresses to the dummy interface. IPVS virtual 
    servers are created for each service IP address.
  * IPVS supports multiple load balancing algorithms.
  * When using IPVS, ipsets are used in cases where iptables are required for an efficient O(1) lookup.
  * More information here: https://www.tkng.io/services/clusterip/dataplane/ipvs/
```

##### Asynchronous Networking
```
a) Queues are a great way to achieve asynchronous communication and decouple various Kubernetes components. Queues 
   abstract away the need to know about IP addresses, ports, and Kubernetes services.
b) Containers can listen or respond to messages, perform actions, and post progress to queues. It is also easy to add or 
   remove listeners and keep track of progress by monitoring the queue.
c) The queue can be used alongside databases and must be highly available.
```

### Load Balancing
---
##### Cloud Provider Networking
```
External access from the cluster usually involves a public load balancer that directs requests to any Kubernetes node.
Kube proxy on the node will then redirect the request to the correct pod on the correct node. This is the typical 
process for cloud providers.
```

##### External Load Balancing
```
External load balancers distributes load at a node level rather than a pod level. If 3 pods exist in node A and 1 pod 
exists in node B, load will be distributed equally to both nodes.
```

##### Cluster Wide Endpoint, Node Local Endpoint, and Client IP Preservation 
```
a) Kubernetes has the ability to preserve information about the originating client IP address. 
b) Ability to preserve client IP address differs based on how 'service.spec.externalTrafficPolicy' is set. This field 
   specifies whether to route external traffic to a node-local endpoint or cluster-wide endpoint.

Cluster Wide Endpoint
  * Default Kubernetes behavior that does not preserve client IP address.
  * Might make network hops to other nodes but spreads load efficiently.

Node Local Endpoint
  * Preserves client IP address.
  * Does not make network hops but does not spread load well.
```

### CNI
---
##### Container Network Interface (CNI)
```
Initiative and specification to write various networking solutions via plugins to configure network interfaces for 
containers. Users are then able to adopt networking solutions and the container orchestration system of their choice 
according to networking needs. Vendors do not need to worry about Kubernetes source code or be locked down to Kubernetes.
```

##### CNI Plugin
```
CNI plugin is responsible for:
  * Add network interface to network namespaces and bridge the container to the host via a VETH pair.
  * Assign unique IP addresses to CNI containers (e.g. pods for Kubernetes) via an IP Address Management (IPAM) plugin.
  * Set up routes and take care of routing logic.

CNI plugin must implement and support:
  * ADD / DEL container to network
  * CHECK container's network status
  * VERSION reporting

CNI plugin execution:
  * Container runtimes pass network configurations in JSON format to CNI plugins and invoke CNI plugins as executables 
    (e.g. invoke ADD verb).
```

##### CNI Workings
```
Step 1) Container runtime specifies actions (e.g. add container) it wants to execute on CNI plugin.
Step 2) Input network configurations stored as JSON files are picked up and streamed to the plugin via STDIN. Runtimes 
        typically specify what path to look for when picking up JSON files. Examples of input configurations:
          * CNI version
          * Type of plugin to use (e.g. bridge plugin)
          * IPAM (e.g. subnet CIDR range, gateway IP address) and DNS (e.g. nameserver) configurations 
Step 3) Supplementary information are further provided via environment variables. Examples of environment variables:
          * Desired operations (e.g. ADD)
          * Container ID
          * Path to network namespace file
          * Name of the network interface to set up
          * Path to CNI plugin executable
          * Additional environment variable arguments
Step 4) CNI plugin takes STDIN and environment variables to perform operations.
Step 5) CNI plugin outputs the result (i.e. generated network interfaces) as STDOUT in JSON format.
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
