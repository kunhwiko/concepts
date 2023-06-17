### Cluster and Node
---
##### Cluster
```
Collection of node instances that provide compute, memory, storage, and networking resources.
```

##### Nodes
```
Nodes are physical or virtual machines.
  * Master nodes act as a control plane for Kubernetes and manage worker nodes.
    These nodes follow raft protocol, meaning an odd number of master nodes must exist for consensus to be possible. 
  * Worker nodes are instances that carry out actual computation work.
```

### Kubernetes Resource Model (KRM)
---
##### KRM
```
a) KRM is a declarative format used to talk to the Kubernetes API to express the desired state of the cluster.
   The declarative format makes it particularly helpful for GitOps.
b) KRM is typically expressed in YAML format and follows OpenAPI standards.
   This allows developers to extend KRM with tools that consume OpenAPIs (e.g. SDK generator, API doc generator). 
c) KRM allows environments to be configured and deployed in a repeatable and reliable way.
```

##### GitOps
```
GitOps is a means to automate deployment of infrastructure changes much like how software development lifecycle is automated through CI/CD.
Once declarative configuration files (i.e. infrastructure as code) are committed, reproducible infrastructure updates are made through pipelines.
```

### Master Node Components
---
##### etcd 
```
a) etcd is a highly reliable distributed key-value store to back cluster level configuration data.
   This includes the state of the cluster, various configuration data (network, authorization), and API metadata.
b) etcd provides watch functionality, allowing API server to be notified of any key-value changes.
c) Periodically taking snapshots of etcd is useful in the event of a cluster failure.
```

##### API Server
```
a) API server exposes Kubernetes REST API and is a means for components in the control plane and worker nodes to communicate with one another.
b) API server is the only component connected to etcd, meaning all other components must pass through the API server to work with cluster state.
c) API server is responsible for all authentication and authorization (e.g. kubectl) mechanisms.
d) API server provide watch functionality, allowing watchers (e.g. scheduler, controllers) to be notified of changes.
```

##### Controller Managers
```
a) Controllers act as non-terminating control loops that continues to poll from the API server to oversee the state of the cluster.
   This information is used to trigger events to move from the current state to a desired state if necessary (e.g. ReplicaSet).
b) Cloud controller managers provide the ability for cloud providers to replace certain controller manager functionality.
   This allows cloud providers to integrate their own platform logic for managing routes, services, volumes, nodes.
c) Customer controller managers can be written to manage custom resources through KRM.
   This could include resources outside the cluster (e.g. Cloud SQL, IAM) using tools like Google Config Connector which follow Kubernetes style YAML.
```

##### Scheduler
```
Scheduler acts as a non-terminating loop that continues to poll from the API server to oversee the state of the cluster.
This information is used to trigger events to assign pods to nodes based on availability (e.g. node availability, affinities, constraints). 
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
b) Kube-proxy is able to forward traffic via TCP and UDP forwarding.
c) Kube-proxy watches the API server for new services and endpoints.
   It will then help maintain networking rules (e.g. IP tables) on its node.
d) Kube-proxy is able to load balance requests if there are multiple pod backends.
   If IP tables are in use, load balancing is typically done through round robin.
```

##### Kubelet
```
a) Kubelets are agents on each node that registers nodes and reports the status of nodes/pods to the API server.
   Communication to the API server is typically done via TLS.
b) Kubelets receive information about pods that need to be newly created from the API server.
   Kubelets will then interact with CRIs to start pods and containers via their configured container runtime.
c) Kubelets ensure to inject necessary environment variables (e.g. service names) to pods before they start.
d) Kubelets monitor and ensure that containers run in a healthy state.
e) Kubelets help download secrets from API server, run liveness probes, and mount volumes.
f) Kubelet allows you to manually configure network MTU.
   Otherwise by default, network plugins will attempt to deduce optimal MTU.
```

### CRI
---
##### Container Runtime Interface (CRI)
```
CRI is a standardized interface for various container runtimes to implement. The interface to implement is as follows: 
https://github.com/kubernetes/kubernetes/blob/release-1.5/pkg/kubelet/api/v1alpha1/runtime/api.proto. CRI allows 
Kubernetes to support various container runtimes based on user needs without the need to recompile.
```

##### CRI Architecture
```
Kubelets will interact with CRI compatible container runtimes or CRI shims (i.e. a bridge that implements CRI and 
translates to something container runtimes can understand) over Unix sockets using the gRPC framework. 
```
