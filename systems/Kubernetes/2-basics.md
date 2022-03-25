### Getting Started
---
##### Getting Started
```
Great source to read : https://www.ibm.com/cloud/learn/kubernetes
```

##### Terms
```
Cluster : collection of hosts that provide compute, memory, storage, and networking resources
Node    : a single host that could be a physical or virtual machine 
```

### Control Plane / Master
---
##### Components
```
etcd 
   a) distributed key-value store to back cluster data (e.g. configuration, state, metadata)
   b) a means to restore K8s cluster by recording past snapshots of the cluster 

API server
   a) exposes Kubernetes REST API
   b) means for control plane and worker nodes to communicate with one another 

Scheduler : assigns pods to nodes 

CoreDNS 
   a) a pod that functions as a DNS server in cluster
   b) every service has a DNS name and pods can receive DNS names too 

Controller Manager : refer below 
```

##### Controller Manager
```
Controller Manager : control loop that oversees the cluster through the API server and moves current state to desired state

Deployments
   a) defines a desired state for pods and replica sets 
   b) enables users to scale number of replicas, control rollout of updates, rollback to previous deployments 
   c) enables users to check or update status of Pods  

Replica Sets:
   a) ensures that a specified number of pod replicas are running at a given time 
   b) allows for rollback to previous deployments

Stateful Sets : see 6-storage.md
```

##### Raft Protocol
```
Raft Protocol : odd number of master nodes exist for consensus to be possible 
```

### Worker Node
---
##### Pods 
```
Pods
   a) encapsulates one or more containers to be assigned to a node
   b) all containers in a pod have the same IP address and port
   c) contain shared resources such as volumes, network configs, info on how to run containers
```

##### Components
```
kubelet
   a) agents on each node that registers nodes to the API server so they can talk with the control plane
   b) makes sure containers run in pods in a healthy state 
   c) receives pod specs, downloads secrets from API server, runs liveness probe, mounts volumes
   d) interact with Container Runtime Interface (CRI)

kube-proxy
   a) implements networking rules that allow for network communication to pods
   b) finds cluster IPs via environment variables or DNS
```

##### CRI
```
CRI
   a) enables Kubernetes to support not only Docker containers but other container runtimes as well
   b) uses gRPC framework to enable kubelet to interact with the CRI
```

### Labels and Selectors
---
##### Labels / Annotations / Selectors
```
Labels
   a) used to identify, select, and group Pods (or other objects) together based on some criteria
   b) not used for attaching arbitrary metadata to objects 

Annotations : used to attach arbitrary metadata that Kubernetes does not care about 

Selectors : chooses objects based on some criteria (two or more selectors imply selector1 AND selector2 instead of OR) 
```

##### Example
```yaml
kind: Deployment
  spec:
    # Deployments use the template field to create Pods with labels "app:test"
    template:
      metadata:
        labels:
          app: test 
          
    selector:
      matchLabel:
        # why do Deployments require Selectors?
        
        # Deployments use Selectors to know which Pods it needs to manage
        # This field must be predefined (expect for version v1beta1) to 
        # prevent mutation of what Pods the Deployments should manage  
        app: test 
```
