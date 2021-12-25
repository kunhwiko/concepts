### Getting Started
---
##### Getting Started
```
Great source to read: https://www.ibm.com/cloud/learn/kubernetes
```

##### Distributed System Design Patterns
```
Sidecar Pattern
   1. In the same Pod, create a separate container from the main application container
   2. This separate container provides supplemental features such as logging aggregation or monitoring
   3. Places less burden on application container

Ambassador Pattern
   1. Like the Sidecar Pattern, create a separate container from the main app container
   2. This container acts as a proxy to the main app container that filters requests
   3. Often used with legacy apps that are risky to modify to extend networking/security configurations
   4. Able to update configurations of ambassador while keeping legacy code

Adapter Pattern
   1. Assume main application has been updated but generates output in a different format
   2. Consumers of the output have not been upgraded to read in the new format
   3. Adapter standardizes output until all consumers have been upgraded 
```

##### Terms
```
Cluster : collection of hosts that provide compute, memory, storage, and networking resources
Node    : a single host that could be a physical or virtual machine 
```

### Control Plane (Master)
---
##### Components
```
etcd: 
   - distributed key-value store to back cluster data (e.g. configuration, state, metadata)
   - a means to restore K8s cluster by recording past snapshots of the cluster 

API Server: 
   - exposes Kubernetes REST API
   - means for Control Plane and Worker Nodes to communicate with one another 

Scheduler
   - assigns Pods to Nodes 

Controller Manager
   - refer below 

coreDNS 
   - a Pod that functions as DNS server in cluster
   - Every service has a DNS name and Pods can receive DNS name too 
```

##### Controller Manager
```
Controller Manager: control loop that oversees the cluster through the apiserver and moves current state to desired state

Deployments: 
   - defines a desired state for Pods and ReplicaSets 
   - enables users to scale number of replicas, control rollout of updates, rollback to previous deployments 
   - enables users to check or update status of Pods  

ReplicaSets:
   - ensures that a specified number of Pod replicas are running at a given time 
   - allows for rollback to previous deployments

StatefulSets:
   - see 4-storage.md
```

##### Raft Protocol
```
Raft protocol: odd number of Master Nodes exist for consensus to be possible 
```

##### High Availability
```
To make Kubernetes cluster highly available:
   - master components should be redundant
   - etcd across nodes should be able to communicate and update cluster data
   - API Server is stateless so there is no need for one to communicate with another
   - multiple Schedulers and Controller Managers means chaos, so these should implement leader election
```

### Worker Node
---
##### Pods 
```
Pods
   - encapsulates one or more containers to be assigned to a Node
   - all containers in a Pod have the same IP address and port
   - contain shared resources such as volumes, network configs, info on how to run containers
```
##### Components
```
kubelet
   - agents on each Node that registers Nodes to the API Server so they can talk with the Control Plane
   - makes sure containers run in Pods in a healthy state 
   - receives Pod specs, downloads secrets from API Server, runs liveness probe, mounts volumes
   - interact Container Runtime Interface (CRI)

kube-proxy:
   - implements networking rules that allow for network communication to Pods
   - finds cluster IPs via environment variables or DNS
```

##### CRI
```
CRI
   - enables Kubernetes to support not only Docker containers but other ontainer runtimes as well
   - uses gRPC framework to enable kubelet to interact with the CRI
```

### Labels and Selectors
---
##### Labels / Annotations / Selectors
```
Labels: 
   - used to identify, select, and group Pods (or other objects) together based on some criteria
   - not used for attaching arbitrary metadata to objects 

Annotations
   - used to attach arbitrary metadata that Kubernetes does not care about 

Selectors
   - chooses objects based on some criteria (two or more selectors imply selector1 AND selector2 instead of OR) 
```

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
