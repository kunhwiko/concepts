### Master Node Components
---
##### API Server
```
a) API server exposes Kubernetes REST API.
b) API server is a means for the control plane and worker nodes to communicate with one another and understand the overall status of the cluster.
c) API server is designed to be stateless as data is meant to be stored in etcd. 
```

##### API Server Optimizations
```
API Server cache
  a) Kubernetes components operate on snapshots of the cluster state rather than on real-time updates.
     Snapshots are saved in etcd, and API Server maintains an in-memory read cache where the snapshot can be read from.
  b) Caching states in the API server decreases read latency and reduces load on etcd.  

Protocol Buffers
  a) REST APIs typically use JSON as a serialization format, which typically requires marshaling / unmarshaling JSON.
     JSON parsing can become expensive because so many REST APIs are called to the API Server.
     Instead, internal components in Kubernetes communicate via a protocol buffer serialization format to reduce this expense.
```

##### etcd 
```
a) etcd is a highly reliable distributed key-value store to back cluster level data (e.g. configuration, state, metadata).
b) etcd can act as a means to restore the Kubernetes cluster by recording past snapshots of the cluster.
```

##### etcd Optimizations
```
Improvements from etcd2 to etcd3
  a) Uses gRPC (HTTP/2 protocol) over REST which enables single TCP connections for multiple streams of requests and responses.
  b) Uses leases over TTL to reduce keep-alive traffic.
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

### Worker Node Components
---
##### Kube Proxy
```
a) Kube-proxy discovers cluster IPs through DNS or environment variables.
b) Kube-proxy maintains networking rules (e.g. IP tables) on each node.
c) Kube-proxy is able to forward traffic via TCP and UDP forwarding.
d) Kube-proxy is able to load balance requests if there are multiple pod backends.
```

##### Kubelet
```
a) Agents on each node that registers nodes to the API server.
   Communicates to API server typically via TLS.
b) Makes sure containers run in pods via the configured runtime in a healthy state.
   Also reports the status of the node and pods. 
c) Downloads secrets from API server, runs liveness probes, and mounts volumes.
d) Able to manually configure network MTU, otherwise network plugins will attempt to deduce optimal MTU.
```

### Container Runtime Interface (CRI)
---
##### CRI
```
a) Enables Kubernetes to support a general interface for various container runtimes.
b) Kubelet interacts with custom implementations of CRI via gRPC to determine what the container runtime should do.
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
