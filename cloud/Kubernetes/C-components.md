### Master Node Components
---
##### API Server
```
a) API server exposes Kubernetes REST API.
b) API server is a means for the control plane and worker nodes to communicate with one another and understand the overall status of the cluster.
c) API server is designed to be stateless as data is meant to be stored in etcd. 
```

##### etcd 
```
a) etcd is a highly reliable distributed key-value store to back cluster level data (e.g. configuration, state, metadata).
b) etcd can act as a means to restore the Kubernetes cluster by recording past snapshots of the cluster.
```

##### Controller Managers
```
Kube Controller Manager
  a) Control loop that oversees the cluster through the API server and moves current state to desired state (e.g. ReplicaSets, Service Controller).

Cloud Controller Manager
  a) Provides the ability for cloud providers to replace certain Kube controller manager functionality.
     This allows cloud providers to integrate their own platform logic for managing routes, services, volumes, nodes.
```

##### Scheduler
```
Assigns pods to nodes. 
```

##### CoreDNS
```
a) CoreDNS is a pod that functions as a DNS server in the cluster.
b) Every service, except for headless services, receive a DNS name and pods can receive DNS names as well. 
```

### Master Node Component Optimizations
---
##### API Server Cache
```
a) Various Kubernetes components operate on snapshots of the cluster state rather than on real-time updates.
b) Current state of the cluster is maintained by the API server through an in-memory read cache.
c) Snapshots are saved in etcd, and the in-memory read cache will be updated by etcd watches.  
d) Caching states in the API server decreases read latency and reduces load on etcd.
```

##### Pod Lifecycle Event Generator (PLEG)
```
Problems
  a) Kubelet used to poll container runtimes from each pod for info. 
     This consistently puts pressure on CPU usage for the container runtimes.

PLEG
  a) Kubelet lists the state of pods and containers and compares with the previous state.
     Kubelet then knows which pods need to sync again and only polls those pods.
```

### Protocol Buffers
```
REST APIs typically use JSON as a serialization format, which typically requires marshaling / unmarshaling JSON to native data structures.
JSON parsing can become expensive because so many REST APIs are called to the API Server.
Instead, internal components in Kubernetes communicate via a protocol buffers serialization format to reduce this expense.
```

##### etcd3
```
a) Uses gRPC over REST (etcd2), utilizing HTTP/2 to enable a single TCP connection for multiple streams of requests and responses.
b) etcd2 uses Time to Live (TTL) per key to expire keys, while etcd3 uses leases with TTLs such that multiple keys can share the same key.
c) Stores state as protocol buffers to reduce JSON serialization overhead.
```

### Worker Node Components
---
##### Kube Proxy
```
a) Kube-proxy discovers cluster IPs through DNS or environment variables.
b) Kube-proxy watches for new endpoints and helps maintain/update networking rules (e.g. IP tables) on each node.
c) Kube-proxy is able to forward traffic via TCP and UDP forwarding.
d) Kube-proxy is able to round robin requests if there are multiple pod backends.
```

##### Kubelet
```
a) Kubelets are agents on each node that registers nodes to the API server.
b) Kubelet makes sure containers run in pods via their configured container runtime. 
c) Kubelet makes sure that containers run in a healthy state.
   Kubelet reports the status of the node and pods to the control plane.
   Communication to API server is typically done through TLS.
d) Kubelet makes sure to inject necessary environment variables (e.g. service names) to pods before they start.
   Kubelet also downloads secrets from API server, runs liveness probes, and mounts volumes.
e) Kubelet allows you to manually configure network MTU.
   Otherwise by default, network plugins will attempt to deduce optimal MTU.
```

### Container Runtime Interface (CRI)
---
##### CRI
```
a) Enables Kubernetes to support a general interface for various container runtimes.
b) Kubelet interacts with custom implementations of CRI via gRPC to determine what the container runtime should do.
```
