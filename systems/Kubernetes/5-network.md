### Service Types 
---
##### Service
```
Service
   a) means for a way to connect to Pods
   b) provides a DNS entry to a group of Pods that persists even if some Pods update their IP 
   c) allows for load balancing between Pods  
   d) Service operates at layer 3 (UDP/TCP), Ingress operates at HTTP layer

Services normally get published as either
   a) environment variables that are picked up by Pods (e.g. SOME_NAME_SERVICE_HOST, SOME_NAME_SERVICE_PORT)
   b) discovery through CoreDNS
   c) more here: https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/
```

##### Service Types
```
Types
   a) ClusterIP
      * exposes Service on an internal IP in the cluster 
      * reachable only from within the cluster 
   b) NodePort 
      * using NATs, exposes Service on the same port of each selected Node
      * makes a Service accessible outside the cluster using <NodeIP>:<NodePort>
      * requests to NodePorts get routed to ClusterIP services  
      * superset of ClusterIP
   c) LoadBalancer
      * mostly used with cloud services 
      * sets up clusterIPs / NodePorts and a great means to get external traffic to come into Service 
      * creates an external load balancer and assigns a fixed, external IP to the Service 
      * superset of NodePort
   d) ExternalName
      * means to get traffic out to an external source
      * adds CNAME DNS record to CoreDNS 

Also reference : https://www.ibm.com/cloud/blog/kubernetes-ingress
```

##### Ports 
```
NodePort vs Port vs TargetPort : https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ports-targetport-nodeport-service.html
```

##### NodePort vs Load Balancer
```
Things to know
   a) Nodes have publicly accessible IPs
   b) NodePorts will open the same port number on all Nodes
   c) <NodeIP>:<NodePort> will convert requests to <ClusterIP>:<Port>  
    
Limitations of NodePorts  
   a) only exposes a single service per port 
   b) must maintain and know the NodeIP of the Node you're looking for, which is difficult when many Nodes exist / crash

Pros of Load Balancers 
   - only need to know the IP address of the Load Balancer 
   - transfers request of <LB IP>:<Port> to appropriate <NodeIP>:<NodePort>
   - ability to open multiple ports and protocols per service 
```

##### Load Balancer vs Ingress 
```
Limitations of Load Balancers 
   a) one service is exposed per LoadBalancer, and with multiple services, this costs a lot of overhead

Ingress Components
   a) Load Balancer: performs the actual routing
   b) Ingress Controller: enables controlled routing based on a set of predefined rules

Pros of Ingress
   a) enable routing to multiple services with a single load balancer 
   b) supports multiple protocols and authentication rules  
```

### Network Policies
---
##### Network Policies
```
Network Policy
   a) set of firewall rules
   b) labels are used to define virtual network segments (more flexible than IP address range / subnet masking)

Purposes
   a) network segmentation allows for multi-tenancy and minimizes security breaches
   b) maximize parts of system that don't need to talk to each other

Scope
  a) network policies are cluster-wide
  b) rules are unified if multiple network policies exist
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

##### Whitelist
```
Whitelist
   a) by default, all access is forbidden to a certain pod if targeted by at least one network policy
   b) by default, all access is granted to a certain pod if targeted by no network policy
```

##### Egress
```
Egress : ability to control outbound traffic
```
