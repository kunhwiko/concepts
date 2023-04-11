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

##### Discoverability
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
a) Endpoint object shows a DNS mapping of hosts to IP addresses.
b) Endpoint objects are used by services to keep track of all IP addresses of Pods corresponding to the service.
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
Technically not a service, ingress comes in two components, a load balancer and ingress controller.
Ingress controller enables controlled routing based on a set of predefined rules.
Ingress load balancer performs actual routing.
```

##### Ingress vs Load Balancer
```
Limitations of Load Balancers 
  a) Like nodeports, a single service is exposed per load balancer.
     With multiple services, this can lead to a large overhead.

Advantages over Load Balancers
  a) Enable routing to multiple services with a single load balancer.
  b) Supports request limits, URL rewrites, TCP/UDP load balancing, SSL termination, authentication etc.
  c) Ingress operates at HTTP layer.
```

##### External Name
```
a) Maps a service to a user specified DNS name as a means to get traffic out to an external source.
b) Adds a CNAME DNS record to CoreDNS such that looking up the service will route to the user specified DNS.
```

### Definition of Network Policies
---
##### Network Policies
```
a) Network policies act as a set of firewall rules that manages network traffic for selected pods and namespaces.
b) Network policies aim to minimize security breaches and maximize parts of systems that don't need to talk to each other.
c) Network policies use labels to whitelist applicable pods and namespaces.
   Usage of labels to define virtual network segments is much more flexible than defining CIDR / subnet masking.
d) Network policies are cluster scoped and rules are unified if multiple network policies exist. 
e) Network policies are part of the standard Kubernetes API but differ in implementation per networking solution / CNI plugin.
   In other words, Kubernetes provides the ability to define and store network policies through APIs. 
   However, enforcing that network policy is left to the networking solution.
```

##### Execution of Network Policies
```
Step 1) Policy is posted and sent to Kubernetes master nodes.
Step 2) Kubernetes master nodes forwards the policy to a policy controller.
Step 3) Policy controller pushes the policy to agents on each worker node.
Step 4) Each agent intercepts traffic, verifies against policies, and forward/rejects requests.
```

##### Whitelisting of Network Policies
```
a) By default, all access is forbidden to a certain pod if targeted by at least one network policy.
b) By default, all access is granted to a certain pod if targeted by no network policy.    
```

##### Egress
```
Policies used to control and deny outbound traffic.
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