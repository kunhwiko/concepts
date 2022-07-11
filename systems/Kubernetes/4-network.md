### Service Types 
---
##### Service
```
Service
   a) provides a stable access point to pods
   b) provides a DNS entry to pods that persists even if pods update their IP 
   c) allows for load balancing between pods  
   d) services operates at layer 3 (UDP/TCP), ingress operates at HTTP layer

Services are published/discovered via
   a) environment variables that are picked up by pods (e.g. SOME_NAME_SERVICE_HOST, SOME_NAME_SERVICE_PORT)
   b) DNS name that pods can use 
   c) more here: https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/

Service Uses
   a) services are pieces of data stored in etcd
   b) kube-proxy will update iptables for each node based on info stored in etcd
```

##### Service Types
```
Types
   a) ClusterIP
      * exposes service on an internal IP in the cluster 
      * reachable only from within the cluster 
   b) NodePort 
      * using NATs, exposes service on the same port of each selected node
      * makes a service accessible outside the cluster using <node-ip>:<node-port>
      * requests to nodeports get routed to clusterIP services  
      * superset of clusterIP
   c) LoadBalancer
      * mostly used with cloud services 
      * sets up clusterIPs / nodeports and a great means to get external traffic to come into service 
      * creates an external load balancer and assigns a fixed, external IP to the service 
      * superset of nodeport
   d) ExternalName
      * means to get traffic out to an external source
      * adds CNAME DNS record to coreDNS 

Also reference : https://www.ibm.com/cloud/blog/kubernetes-ingress
```

##### Ports 
```
NodePort vs Port vs TargetPort : https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ports-targetport-nodeport-service.html
```

##### NodePort vs Load Balancer
```
Things to know
   a) nodes have publicly accessible IPs
   b) nodeports will open the same port number on all nodes
   c) <node-ip>:<node-port> will convert requests to <clusterIP>:<port>  
    
Limitations of NodePorts  
   a) only exposes a single service per port 
   b) must maintain and know the IP of the node you're looking for, which is difficult when many nodes exist / crash

Pros of Load Balancers 
   a) only need to know the IP address of the load balancer 
   b) transfers request of <loadbalancer-ip>:<port> to appropriate <node-ip>:<node-port>
   c) ability to open multiple ports and protocols per service 
```

##### Load Balancer vs Ingress 
```
Limitations of Load Balancers 
   a) one service is exposed per load balancer, and with multiple services, this costs a lot of overhead

Ingress Components
   a) Load Balancer      : performs the actual routing
   b) Ingress Controller : enables controlled routing based on a set of predefined rules

Pros of Ingress
   a) enable routing to multiple services with a single load balancer 
   b) supports multiple protocols and authentication rules  
```

##### Headless Services
```
Limitations of Services
   a) connections to services are load balanced and forwarded randomly to a backing pod
   b) difficult to get A records of all backing pods through services

Headless Services
   a) DNS entry of headless service returns all A records of pods backed by headless service
   b) all backing pods are able to connect with each other
   c) headless service does not have a clusterIP
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
