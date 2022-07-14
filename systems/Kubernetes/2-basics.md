### Getting Started
---
##### Basic Infrastructure
```
Cluster
   a) Collection of hosts that provide compute, memory, storage, and networking resources.

Node
   a) A single host that could be a physical or virtual machine.
```

### Control Plane / Master
---
##### etcd 
```
etcd
   a) Distributed key-value store to back cluster-level data (e.g. configuration, state, metadata).
   b) A means to restore the Kubernetes cluster by recording past snapshots of the cluster.

Improvements from etcd2 to etcd3
   a) Uses gRPC (HTTP/2 protocol) over REST which enables single TCP connections for multiple streams of requests and responses.
   b) Uses leases over TTL to reduce keep-alive traffic.
```

##### API Server
```
Kubernetes API Server
   a) Exposes Kubernetes REST API.
   b) Means for control plane and worker nodes to communicate with one another and understand the status of the cluster.

API Server cache
   a) Kubernetes components operate on snapshots of the state rather than real-time updates.
      These snapshots are saved in etcd, and API Server has an in-memory read cache where the snapshot can be read from.
   b) Caching states in the API server decreases read latency and reduces load on etcd  

Protocol Buffers
   a) REST APIs typically use JSON as serialization format, and typically requires marshaling / unmarshaling JSON.
      JSON parsing becomes expensive because so many REST APIs are called to the API Server.
      Internal components in Kubernetes instead communicate via a protocol buffer serialization format to reduce this expense.
```

##### Controller Manager
```
Controller Manager
   a) Control loop that oversees the cluster through the API server and moves current state to desired state.

Deployments
   a) Defines a desired state for pods and replica sets.
   b) Enables users to scale number of replicas, control rollout of updates, rollback to previous deployments.
   c) Enables users to check or update status of pods.  

Replica Sets
   a) Ensures that a specified number of pod replicas are running at a given time.
   b) Allows for rollback to previous deployments.
```

##### Scheduler
```
Scheduler
   a) Assigns pods to nodes. 
```

##### CoreDNS
```
CoreDNS 
   a) A pod that functions as a DNS server in the cluster.
   b) Every service has a DNS name and pods can receive DNS names too. 
```

##### Raft Protocol
```
Raft Protocol
   a) Odd number of master nodes exist for consensus to be possible. 
```

### Worker Node
---
##### Pods 
```
Pods
   a) Encapsulates one or more containers to be assigned to a node.
   b) All containers in a pod share the same IP address and port.
   c) Contains shared resources such as volumes, network configs, info on how to run containers.
```

##### Components
```
Kubelet
   a) Agents on each node that registers nodes to the API server.
   b) Makes sure containers run in pods in a healthy state. 
   c) Receives pod specs, downloads secrets from API server, runs liveness probes, mounts volumes.
   d) Communicates to API server typically via TLS.
   e) Interacts with Container Runtime Interface (CRI).

Kube Proxy
   a) Implements networking rules that allow for network communication to pods.
   b) Typically discovers resources (e.g. services) either through DNS or environment variables.
```

##### Container Runtime Interface (CRI)
```
CRI
   a) Enables Kubernetes to support a general interface for not only Docker containers but other container runtimes as well.
   b) Kubelet interacts with CRI via gRPC to determine what container runtime to use.
```

##### Pod Lifecycle Event Generator (PLEG)
```
Problems
   a) Kubelet used to poll container runtimes from each pod for info.
      This consistently puts a lot of pressure on the container runtime and to each pod.

PLEG
   a) Lists the state of pods and containers and compares with the previous state.
      Kubelet then knows which pods need to sync again and only polls those pods.
```

### Labels and Selectors
---
##### Labels / Annotations / Selectors
```
Labels
   a) Used to identify, select, and group pods (or other objects) together based on some criteria.
   b) Not used for attaching arbitrary metadata to objects.

Annotations
   a) Used to attach arbitrary metadata that Kubernetes does not care about.

Selectors
   a) Chooses objects based on some criteria (two or more selectors imply selector1 AND selector2 instead of OR).
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
