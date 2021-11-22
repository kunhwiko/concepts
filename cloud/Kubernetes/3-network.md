### Networks 
---
##### Service
```
Service
  - means to connect Pods to external services 
  - provides a DNS entry to a group of Pods that persists even if some Pods update their IP 
  - allows for load balancing between Pods  
  - Service operates at layer 3 (UDP/TCP), Ingress operates at HTTP layer

Services normally get published as either:
  - environment variables that are picked up by Pods (e.g. SOME_NAME_SERVICE_HOST, SOME_NAME_SERVICE_PORT)
  - discovery through CoreDNS
  - more here: https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/
```

##### Service Types
```
Types
  - ClusterIP:
      * exposes Service on an internal IP in the cluster 
      * reachable only from within the cluster 
  - NodePort 
      * using NATs, exposes Service on the same port of each selected Node
      * makes a Service accessible outside the cluster using <NodeIP>:<NodePort>
      * requests to NodePorts get routed to ClusterIP services  
      * superset of ClusterIP
  - LoadBalancer
      * mostly used with cloud services 
      * sets up clusterIPs / NodePorts and a great means to get external traffic to come into Service 
      * creates an external load balancer and assigns a fixed, external IP to the Service 
      * superset of NodePort
  - ExternalName
      * means to get traffic out to an external source
      * adds CNAME DNS record to CoreDNS 

Also reference: https://www.ibm.com/cloud/blog/kubernetes-ingress
```

##### Ports 
```
NodePort vs Port vs TargetPort 
  - reference: https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ports-targetport-nodeport-service.html
```

##### NodePort vs Load Balancer
```
Things to know:
  - Nodes have publicly accessible IPs
  - NodePorts will open the same port number on all Nodes
  - <NodeIP>:<NodePort> will convert requests to <ClusterIP>:<Port>  
    
Limitations of NodePorts  
  - Only exposes a single service per port 
  - Must maintain and know the NodeIP of the Node you're looking for, which is difficult when many Nodes exist / Nodes crash or update 

Pros of Load Balancers 
  - Only need to know the IP address of the Load Balancer 
  - Transfers request of <LB IP>:<Port> to appropriate <NodeIP>:<NodePort>
  - Ability to open multiple ports and protocols per service 
```

##### Load Balancer vs Ingress 
```
Limitations of Load Balancers 
  - One service is exposed per LoadBalancer, and with multiple services, this costs a lot of overhead

Ingress Components
  - Load Balancer: performs the actual routing
  - Ingress Controller: enables controlled routing based on a set of predefined rules

Ingress Pros
  - Enable routing to multiple services with a single load balancer 
  - Supports multiple protocols and authentication rules  
```