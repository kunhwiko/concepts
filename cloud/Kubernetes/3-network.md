### Networks 
---
##### Service
```
1) means to connect Pods to external services 
2) provides a DNS entry to a group of Pods that persists even if some Pods update their IP 
3) allows for load balancing between Pods   

Types
    1) ClusterIP:
        * exposes Service on an internal IP in the cluster 
        * reachable only from within the cluster 
    2) NodePort 
        * using NATs, exposes Service on the same port of each selected Node
        * makes a Service accessible outside the cluster using <NodeIP>:<NodePort>
        * requests to NodePorts get routed to ClusterIP services  
        * superset of ClusterIP
    3) LoadBalancer
        * mostly used with cloud services 
        * sets up clusterIPs / NodePorts and a great means to get external traffic to come into Service 
        * creates an external load balancer and assigns a fixed, external IP to the Service 
        * superset of NodePort
    4) ExternalName
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
    1. Nodes have publicly accessible IPs
    2. NodePorts will open the same port number on all Nodes
    3. <NodeIP>:<NodePort> will convert requests to <ClusterIP>:<Port>  
    
Limitations of NodePorts  
    1. Only exposes a single service per port 
    2. will have to maintain and know the NodeIP of the Node you're looking for, which is difficult when many Nodes exist / Nodes crash or update 

Pros of Load Balancers 
    1. Only need to know the IP address of the Load Balancer 
    2. Transfers request of <LB IP>:<Port> to appropriate <NodeIP>:<NodePort>
    3. Ability to open multiple ports and protocols per service 
```

##### Load Balancer vs Ingress 
```
Limitations of Load Balancers 
    1. One service is exposed per LoadBalancer, and with multiple services, this costs a lot of overhead

Ingress Components
    1. Load Balancer: performs the actual routing
    2. Ingress Controller: enables controlled routing based on a set of predefined rules

Ingress Pros
    1. Enable routing to multiple services with a single load balancer 
    2. Supports multiple protocols and authentication rules  
```