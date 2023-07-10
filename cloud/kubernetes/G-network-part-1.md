### Services
---
##### Problem Statement
```
Pods are assigned internal IP addresses that are exposed across the entire cluster. This IP address can be used to 
directly route requests to certain pods. However, these IP addresses are generally instable as pod restarts will assign 
new IP addresses.
```

##### Service
```
a) Services provide a stable access point to send requests to pods and identifies the pods it needs to send requests to
   via labels.
b) Services provide load balancing features done by kube-proxy and operate at layer 3 (TCP/UDP) networking.
```

##### Discoverability
```
a) Services can be discovered via environment variables. When a pod runs on a node, the kubelet will inject env variables 
   representing the host and port of each active service. However, env variables will not be updated for services created 
   after the pod's creation. 
b) Services can be discovered via DNS (i.e. <service-name>.<namespace>.svc.cluster.local). CoreDNS will register the 
   service name internally in the cluster.  
```

##### Ports
```
Port vs TargetPort: https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ports-targetport-nodeport-service.html
```

##### Endpoint
```
Endpoint objects are used by services to keep track of all pod endpoints that the service routes to. These objects are 
represented as a DNS mapping of hosts to IP addresses/ports.
```

##### Headless Services
```
Problem Statement
  * Connections to services are load balanced and forwarded randomly to a backing pod.
  * It is difficult to get A records of all backing pods through services.

Headless Services
  * DNS entry of headless service returns all A records of pods backed by headless service.
  * All backing pods are able to connect with each other.
  * Headless service does not have a clusterIP.
```

### Types of Services
---
##### ClusterIP
```
Service that exposes an internal IP that is recognizable only from within the cluster.
```

##### NodePort
```
a) Kubernetes nodes have publicly accessible IPs. Utilizing NATs, nodeports can be exposed on the same port number on 
   all nodes. This makes the service externally accessible via <node-ip>:<node-port>.
b) Requests to nodeports via <node-ip>:<node-port> will be routed to clusterIPs on <clusterIP>:<port> by the kube-proxy.
   Nodeports are therefore a superset of clusterIPs.
```

##### Load Balancer
```
a) Load balancers are assigned a fixed external IP that is accessible from outside the cluster. Note that Multiple ports 
   and protocols can be defined on a single load balancer.
b) Load balancers are a superset of nodeports, meaning both nodeports and clusterIPs will be created. Requests are 
   typically forwarded from load balancers to nodeports.
c) Load balancers are mostly used with managed cloud services and might require spinning up resources (e.g. NLBs) for a 
   cost. Cloud providers can also decide how requests will be load balanced.
```

##### Load Balancer vs NodePort
```
Limitations of NodePorts  
  * Users must know the IP of the node they are looking for, which can be difficult when many nodes exist or crash.
  * Nodeports expose at most a single service per port.

Advantages over NodePorts
  * Users only need to know the IP address of the load balancer. Requests to <loadbalancer-ip>:<port> are sent to the
    appropriate <node-ip>:<node-port>.
  * While load balancers typically create and forward requests to nodeports, this can optionally be disabled. This is 
    only possible if the given cloud provider has implemented load balancers this way.
```

##### Ingress
```
An ingress is technically not a service, and comes in two components, a load balancer and an ingress controller. Ingress 
controllers enable controlled routing based on a set of predefined rules. Ingress load balancers performs routing.
```

##### Ingress vs Load Balancer
```
Limitations of Load Balancers 
  * A single service is exposed per load balancer. With multiple services, this can lead to a large overhead.

Advantages over Load Balancers
  * Enable rule based routing to various services with a single load balancer and operates at the HTTP layer.
  * Supports request limits, URL rewrites, TCP/UDP load balancing, SSL termination, authentication etc.
```

##### External Name
```
a) Maps a service to a user specified DNS name as a means to get traffic out to an external source.
b) Adds a CNAME DNS record to CoreDNS such that looking up the service will route to the user specified DNS.
```

### Network Policies
---
##### Network Policies
```
a) Network policies act as a set of firewall rules that manage network traffic for selected pods and namespaces. These
   policies aim to implement the least privilege principle.
b) Network policies use labels to whitelist applicable pods and namespaces. Usage of labels to define virtual network 
   segments is much more flexible than defining CIDR / subnet masking.
c) If a network policy is created for ingress requests, then a response is allowed back regardless of egress network 
   policies. From a Kubernetes perspective, we just need to be worried about the originating request.  
d) Network policies are part of the standard Kubernetes API but differ in implementation per networking solution / CNI 
   plugin. Kubernetes provides the ability to define network policies through APIs, but it is up to the networking 
   solution whether to enforce those policies.
e) Network policies are cluster scoped and rules are unified if multiple network policies exist.
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
Kubernetes resources that represent policies used to control and deny outbound traffic.
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
  # what policies we are allowing
  policyTypes:
  - Ingress
  - Egress
  # specifies what requests can come into the above pods
  ingress:
  - from:
    # rules inside the same element work as an AND condition
    # separate rules work as an OR condition
    - ipBlock:
        cidr: 172.17.0.0/16
        except:
        - 172.17.1.0/24
    # refer to https://kubernetes.io/docs/concepts/services-networking/network-policies/#behavior-of-to-and-from-selectors
    - namespaceSelector:
        matchLabels:
          project: test
      podSelector:
        matchLabels:
          role: app-frontend
    # network protocol and ports that are allowed
    ports:
    - protocol: TCP
      port: 8888
  # specifies what request can leave the above pods
  egress:
  - to:
    - ipBlock:
        cidr: 192.168.5.10/32
    ports:
    - protocol: TCP
      port: 80
        
```
