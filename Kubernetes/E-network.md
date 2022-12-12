### Definition of Service
---
##### Problem Statement
```
Pods are assigned internal IP addresses that are exposed across the entire cluster.
This IP address can be used to directly route requests to certain pods.
However, these IP address are generally instable as pod restarts will assign new IP addresses.
```

##### Service
```
a) Provides a stable access point to send requests to pods.
b) Services provide load balancing features done by kube-proxy. 
c) Services identify pods it needs to send requests to via labels. 
d) Services operate at layer 3 (TCP/UDP) networking.
```

###### Discoverability
```
a) Services can be discovered via environment variables. 
   When a pod runs on a node, the kubelet will add env variables for the host and port of each active service.
   However, env variables will not be updated for services created after the pod's creation. 
b) Services can be discovered via DNS name (i.e. <service-name>.<namespace>.svc.cluster.local).
   Kubernetes CoreDNS will register the service name internally in the cluster.  
```

##### Ports
```
Port vs TargetPort: https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ports-targetport-nodeport-service.html
```

##### Endpoint
```
Object that shows a current/valid DNS record of host to IP addresses.
```

##### Headless Services
```
Problem Statement
  a) Connections to services are load balanced and forwarded randomly to a backing pod.
  b) It is difficult to get A records of all backing pods through services.

Headless Services
  a) DNS entry of headless service returns all A records of pods backed by headless service.
  b) All backing pods are able to connect with each other.
  c) Headless service does not have a clusterIP.
```

### Types of Services
---
##### ClusterIP
```
Exposes service on an internal IP in the cluster reachable only from within the cluster.
```

##### NodePort
```
a) Kubernetes nodes have publicly accessible IPs. Utilizing NATS, nodeports can be exposed on the same port number on all nodes. 
   This makes the service externally accessible via <node-ip>:<node-port>.
b) Requests to nodeports via <node-ip>:<node-port> will be routed to clusterIPs on <clusterIP>:<port> by the kube-proxy.
   Nodeports are therefore a superset of clusterIPs.
```

##### Load Balancer
```
a) Load balancers are assigned a fixed external IP that is accessible from outside the cluster.
b) Load balancers are a superset of nodeports, meaning both nodeports and clusterIPs will be created.
   Requests are typically forwarded from load balancers to nodeports.
c) Load balancers are mostly used with managed cloud services and might require spinning up resources (e.g. NLBs) for a cost.
   Cloud providers also decide how requests will be be load balanced.
d) Multiple ports and protocols can be defined on a single load balancer.
```

##### Load Balancer vs NodePort
```
Limitations of NodePorts  
  a) Users must know the IP of the node they are looking for, which can be difficult when many nodes exist or crash.
  b) Nodeports expose at most a single service per port.

Advantages over NodePorts
  a) Users only need to know the IP address of the load balancer.
     Requests to <loadbalancer-ip>:<port> are sent to appropriate <node-ip>:<node-port>.
  b) While load balancers typically create and forward requests to nodeports, this can optionally be disabled.
     This is only possible if the given cloud provider has implemented load balancers this way.
```

##### Ingress
```
a) Ingress is technically not a service.
b) Ingress comes in two components, a load balancer and ingress controller.
   Ingress controller enables controlled routing based on a set of predefined rules.
   Ingress load balancer performs actual routing.
```

##### Ingress vs Load Balancer
```
Limitations of Load Balancers 
  a) One service is exposed per load balancer, and with multiple services, this can lead to a large overhead.

Advantages over Load Balancers
  a) Enable routing to multiple services with a single load balancer.
  b) Supports request limits, URL rewrites, TCP/UDP load balancing, SSL termination, authentication etc.
  c) Ingress operates at HTTP layer.
```

##### External Name
```
a) Maps a service to a user specified DNS name as a means to get traffic out to an external source.
b) Adds a CNAME DNS record to coreDNS such that looking up the service will route to the user specified DNS.
```

### Kubernetes Internal Networking
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
Step 2) When two pods communicate from within the node, requests are resolved via ARP.
        The request will jump from the current namespace's veth to the root namespace's veth, and then to the target namespace's veth.
Step 3) For pod communications across different nodes, subnet masking first determines if endpoint is on the same network.
        If not, ARP will check for the MAC address of the Kubernetes default gateway.
        The request will jump from the current namespace's veth to the root namespace's veth, and then to the default gateway to be routed to the right node. 
```

##### Service Networking
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

##### Asynchronous Networking
```
Queues
  a) Queues are a great way to achieve asynchronous communication and decouple various Kubernetes components.
     Queues abstract away the need to know about IP addresses, ports, and Kubernetes services.
  b) Containers can listen or respond to messages, perform actions, and post progress to queues.
     It is also easy to add or remove listeners and keep track of progress by monitoring the queue.
  c) The queue can be used alongside databases and must be highly available.
```

### Kubernetes External Networking
---
##### Cloud Provider Networking
```
External access from the cluster usually involves a public load balancer that directs requests to any Kubernetes node.
Kube proxy on the node will then redirect the request to the correct pod on the correct node.
This is typically the process for cloud providers.
```

##### Client IP Preservation
```
service.spec.externalTrafficPolicy
   a) Specifies whether to route external traffic to node-local endpoint or cluster-wide endpoint.

Cluster Wide Endpoint
   a) Default behavior.
   b) Spreads load well.
   c) Does not preserve client IP address.
   d) Could make network hops.

Node Local Endpoint
   a) Does not spread load well.
   b) Preserves client IP address.
   c) does not perform network hops.
```

##### External Load Balancing
```
External Load Balancing
   a) External load balancers operate at a node level rather than a pod level.
      If 3 pods exist in node A and 1 pod exists in node B for the service, load will likely be distributed equally to both nodes.
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

### Definition of Network Policies
---
##### Network Policies
```
a) Network policies act as a set of firewall rules that manages network traffic for selected pods and namespaces.
b) Network policies aim to minimize security breaches and maximize parts of systems that don't need to talk to each other.
c) Network policies use labels to select applicable pods/namespaces and to define whitelist rules.
   Usage of labels to define virtual network segments is much more flexible than defining CIDR / subnet masking.
d) Network policies are cluster scoped and rules are unified if multiple network policies exist. 
```

##### Whitelisting
```
a) By default, all access is forbidden to a certain pod if targeted by at least one network policy.
b) By default, all access is granted to a certain pod if targeted by no network policy.    
```

##### Examples
```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
spec: 
  # what pods this policy applies to
  podSelector:
    matchLabels:
      role: app-backend
  # which namespace and pod can access above pods 
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
Policies used to control and deny outbound traffic.
```

##### CNI Plugin Implementations
```
Relationship between CNI plugins and network policies
   a) Implementation of network policy differs between CNI plugins.
      Some CNI plugins both implement network connectivity and enable network policies, while others only do one or the other.
   b) CNI plugins are able to collaborate with one another.
      As an example, Calico can implement networking connectivity + Flannel can enable network policies.
```

##### Execution of Network Policies
```
Enforcement of Network Policies
   Step 1) Policy is posted and sent to Kubernetes master nodes.
   Step 2) Kubernetes master nodes forwards the policy to a policy controller.
   Step 3) Policy controller pushes the policy to gatekeepers on each worker node.
   Step 4) Each gatekeeper intercepts traffic, verifies against policies, and forward/rejects requests.
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

##### Keepalived Virtual IP
```
Problem
   a) Clients need a stable endpoint, but pods and sometimes load balancers move around in Kubernetes.
   b) DNS resolution for load balancers and services are not good enough due to performance issues.

Solution
   a) Keepalived provides high performance virtual IP address that serves address of load balancers / ingress controllers.

Linux functionalities
   a) IPVS (IP virtual server).
   b) High availability via Virtual Redundancy Router Protocol (VRRP).
   c) Operates at networking layer 4 level.
```
