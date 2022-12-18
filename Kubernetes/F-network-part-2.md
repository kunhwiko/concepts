### Cross Component Networking
---
##### Container to Container Networking
```
a) Containers in the same pod share IP addresses and network namespaces.
   This means containers in the same pod can communicate with one another via localhost. 
b) CRI is responsible for creating new Linux namespaces on the Kubernetes node.
   Each pod is assigned to a Linux namespace and gets its own IP address.
```

##### Pod to Pod Networking
```
Step 1) CNI sets up a virtual ethernet in the pod's Linux namespace.
        This veth is connected to the veth of the node's root namespace via a network bridge.
Step 2) Subnet masking determines if endpoint is on the same network.
Step 3) If two pods are communicating from within the same node, requests are resolved via ARP.
        The request will jump from the current namespace's veth --> root namespace's veth --> target namespace's veth.
Step 4) For pod communications across different nodes, ARP will check for the MAC address of the Kubernetes default gateway.
        The request will jump from the current namespace's veth --> root namespace's veth --> default gateway --> route to correct node. 
```

##### Service Networking
```
How Services Work
  a) Services are pieces of data stored in etcd and are configurations on Linux Netfilter and IP Tables.
     Kube proxy will update IP Tables for each node based on info stored in etcd.

Pod to Pod Networking via Services
  Step 1) ARP will check for the MAC address of the Kubernetes default gateway.
  Step 2) Netfilter hooks are triggered and IP Table chains are applied.
          DNAT will rewrite the packet's destination address to the backend Pod of the service.
  Step 3) Conntrack will keep track of the origin so the target pod can send back a response to the requesting pod.
```

##### Asynchronous Networking
```
Queues
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
This is typically the process for cloud providers.
```

##### Cluster Wide Endpoint, Node Local Endpoint, and Client IP Preservation 
```
a) Kubernetes has the ability to preserve information about the originating client IP address. 
b) Ability to preserve client IP address differs based on how service.spec.externalTrafficPolicy is set.
   This field specifies whether to route external traffic to node-local endpoint or cluster-wide endpoint.

Cluster Wide Endpoint
  a) Default Kubernetes behavior that does not preserve client IP address.
  b) Might make network hops to other nodes but spreads load efficiently.

Node Local Endpoint
  a) Does not preserve client IP address.
  b) Does not make network hops but does not spread load well.
```

##### External Load Balancing
```
External load balancers operate at a node level rather than a pod level.
If 3 pods exist in node A and 1 pod exists in node B, load will still be distributed equally to both nodes.
```

### Container Network Interface (CNI)
---
##### CNI
```
Initiative and specification to write various networking solutions via plugins that are integratable with various container orchestrators.
Users are able to adopt networking solutions and the container orchestration system of their choice according to different needs.
Vendors do not need to worry about Kubernetes source code or be locked down to Kubernetes this way.
```

##### CNI Plugin
```
a) CNI Plugin must do:
    * Add network interface to container network namespace and bridge the container to the host via a veth pair.
    * Assign unique IP addresses to CNI containers (in Kubernetes, these are pods) via an IP Address Management (IPAM) plugin.
    * Take care of routing logic.
b) Plugin must implement and support:
    * ADD / DEL container to network
    * CHECK container's network status
    * VERSION reporting
c) Container runtimes invoke CNI plugins as an executable (e.g. invoke ADD verb) and passes a JSON configuration payload to the CNI. 
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
Step 3) Supplementary information can be further provided via environment variables.
        Examples of environment variables:
          * Desired operations (e.g. ADD)
          * Container ID
          * Path to network namespace file
          * Name of the network interface that will be set up
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