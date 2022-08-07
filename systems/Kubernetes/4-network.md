### Service Types 
---
##### Service
```
Problems
   a) Pods have internal IP addresses that can be used to directly access that pod.
      These IP address are instable as pod restarts will assign new IP addresses.

Service
   a) Provides a stable access point to pods via a DNS entry to pods that persists even if pods update their IP.
   b) Allows for load balancing between pods.
   c) Services operates at layer 3 (UDP/TCP), ingress operates at HTTP layer.

Discoverability
   a) Environment variables that are picked up by pods (e.g. SOME_NAME_SERVICE_HOST, SOME_NAME_SERVICE_PORT).
   b) DNS name that pods can use. 
   c) more here: https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/
```

##### Ports
```
NodePort vs Port vs TargetPort : https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ports-targetport-nodeport-service.html
```

##### ClusterIP
```
ClusterIP
   a) Exposes service on an internal IP in the cluster reachable only from within the cluster.
```

##### NodePort
```
NodePort
   a) Nodes have publicly accessible IPs. Using NATS, node ports will expose the same port number on all nodes. 
      Makes the service externally accessible via <node-ip>:<node-port>, which will convert requests to <clusterIP>:<port>.
   b) Requests to node ports will get routed to clusterIPs (node ports are supersets of clusterIPs).
```

##### Load Balancer
```
LoadBalancer
   a) Mostly used with managed cloud services and might spin up resources (e.g. Network Load Balancers) for a cost.
   b) Sets up clusterIPs / node ports and are a great means to get external traffic inbound (load balancers are superset of node ports).
   c) Assigns a fixed external IP to the service
   d) Additional references: https://www.ibm.com/cloud/blog/kubernetes-ingress

Limitations of NodePorts  
   a) Only exposes a single service per port.
   b) Must maintain and know the IP of the node you're looking for, which is difficult when many nodes exist / crash.

Advantages over NodePorts
   a) Only need to know the IP address of the load balancer.
   b) Transfers request of <loadbalancer-ip>:<port> to appropriate <node-ip>:<node-port>.
   c) Ability to open multiple ports and protocols per service. 
```

#### Ingress
```
Ingress Components
   a) Load Balancer      : Performs the actual routing.
   b) Ingress Controller : Enables controlled routing based on a set of predefined rules.

Limitations of Load Balancers 
   a) One service is exposed per load balancer, and with multiple services, this costs a lot of overhead.

Advantages over Load Balancers
   a) Enable routing to multiple services with a single load balancer.
   b) Supports multiple protocols and authentication rules.   
```

##### Headless Services
```
Limitations of Services
   a) Connections to services are load balanced and forwarded randomly to a backing pod.
      It is difficult to get A records of all backing pods through services.

Headless Services
   a) DNS entry of headless service returns all A records of pods backed by headless service.
   b) All backing pods are able to connect with each other.
   c) Headless service does not have a clusterIP.
```

##### External Name
```
ExternalName
   a) Means to get traffic out to an external source
   b) Adds CNAME DNS recourd to coreDNS
```

### Networking Methods
---
##### Container to Container Networking
```
Container to Container Networking
   a) CRI is responsible for creating new Linux namespaces.
      Each pod is assigned to a Linux namespace on a Kubernetes node and gets their own IP address.
   b) Containers in a pod share IP addresses and network namespaces, and therefore can communicate via localhost. 
```

##### Pod to Pod Networking
```
Pod to Pod Networking
   Step 1) CNI sets up a virtual ethernet in the pod's Linux namespace.
           This veth is connected to the veth of the node's root namespace via a network bridge.
   Step 2) When two pods communicate from within the node, requests are resolved via ARP.
           The request will jump from the current namespace's veth to the root namespace's veth, and then to the target namespace's veth.
   Step 3) For pod communications across different nodes, subnet masking first determines if endpoint is on the same network.
           If not, ARP will check for the MAC address of the Kubernetes default gateway.
           The request will jump from the current namespace's veth to the root namespace's veth, and then to the default gateway to be routed to the right node. 

IP Address Uniformity
   a) The same pod IP address is used for within the node and across the entire cluster.  
      This IP address is exposed across the entire cluster.
```

##### Pod to Pod Networking via Services
```
How Services Work
   a) Services are pieces of data stored in etcd and built on top of Linux Netfilter and IP Tables.
      Kube proxy will update IP Tables for each node based on info stored in etcd.

Pod to Pod Networking via Services
   Step 1) ARP will check for the MAC address of the Kubernetes default gateway.
   Step 2) Netfilter hooks are triggered and IP Table chains are applied.
           DNAT will rewrite the packet's destination address to the backend Pod of the service.
   Step 3) Conntrack will keep track of the origin so the target pod can send back a response to the requesting pod.
```

##### Pod to Pod Networking via Queues
```
Pod to Pod Networking via Queues
   a) Containers can listen or respond to messages, perform actions, post progress status via queues
   b) Queues decouple the need to know about IP addresses.
   c) Easy to keep track of progress by monitoring the queue, and great for large scale systems.
   d) Easy to add or remove listeners.
   e) The queue must be highly available and may requires some time to set up.
   f) Queues can be used alongside databases.
      After a container stores job results or data into the database, the container can ping the container.
      Other containers in the system can then pick up this data.
```

##### Cloud Provider Networking
```
Cloud Provider Networking
   a) Cloud provider load balancers are not Kubernetes aware. 
      Typically these load balancers will direct to any Kubernetes node.
      Kube proxy will then redirect to the correct pod in the correct node. 
```

### Container Network Interface (CNI)
---
##### CNI
```
CNI
   a) Specification and a set of libraries for writing CNI plugins.
   b) Initiative to make different networking solutions integratable with various container orchestration systems.
   c) Users are able to adopt networking solutions and the container orchestration system of their choice according to different needs.
```

##### CNI Plugin
```
CNI Plugin
   a) Plugin must do the following:
      * Add network interface (e.g. bridge) to container network namespace, bridge the container to the host via veth pairs.
      * Assign unique IP addresses to CNI containers (in Kubernetes, these are pods) via an IP Address Management (IPAM) plugin.
      * Take care of routing logic.
   b) Plugin must implement and support the following interface:
      * ADD / DEL container to network
      * CHECK container's network status
      * VERSION reporting
   c) Container runtimes invoke CNI plugins as an executable (e.g. invoke ADD verb) and passes a JSON configuration payload to the CNI. 
```

##### CNI Workings
```
How CNI Works
   Step 1) Container runtime specifies actions (e.g. add container) it wants to execute on CNI plugin.
   Step 2) Input network configurations stored as JSON files are picked up and streamed to the plugin via STDIN.
           Runtimes typically specify what path to look for when picking up JSON files. 
           Examples of configurations:
              * CNI version
              * Type of plugin to use (e.g. bridge plugin)
              * IPAM and DNS configurations 
   Step 3) Container runtime can pass other additional environment variables.
           Examples of environment variables:
              * Desired operations (e.g. ADD)
              * Path to network namespace file
              * Name of the network interface that will be set up
              * Path to CNI plugin executable
   Step 4) CNI plugin performs operations based on the configurations.
   Step 5) CNI plugin outputs the result (generated network interfaces) as STDOUT in JSON format.
```

##### Flat Networking vs Overlay Networking Model
```
CNI with Flat Networking
   a) All pods in the cluster are assigned IP addresses from the cluster's IP pool.
   b) Easy to set up and monitor network traffic but could easily exhaust all of the available IP addresses.

CNI with Overlay Networking
   a) Encapsulates packets coming from the underlaying network at the secondary network level when going to another node.
   b) Typically uses VXLAN (tunneling L2 domains over L3 networks).
```

### Network Policies
---
##### Network Policies
```
Network Policy
   a) Set of firewall rules.
   b) Labels are used to define virtual network segments (more flexible than IP address range / subnet masking).

Purposes
   a) Network segmentation allows for multi-tenancy and minimizes security breaches.
   b) Maximize parts of system that don't need to talk to each other.

Scope
  a) Network policies are cluster-wide.
  b) Rules are unified if multiple network policies exist.

Whitelist
   a) By default, all access is forbidden to a certain pod if targeted by at least one network policy.
   b) By default, all access is granted to a certain pod if targeted by no network policy.  
```

##### Examples
```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
spec: 
  # what pods are targeted by this policy?
  podSelector:
    matchLabels:
      role: app-backend
  # namespace and pods with these labels can access target pods 
  # current namespace is exempt and does not need to match labels here
  ingress:
    - from:
      - namespaceSelector:
          matchLabels:
            project: test
      - podSelector:
          matchLabels:
            role: app-frontend
  # network protocol and ports that are allowed
  ports:
    - protocol: tcp
      port: 8888
```

##### Egress
```
Egress
   a) Ability to control outbound traffic.
```
