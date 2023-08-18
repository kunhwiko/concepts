### Clusters and Nodes
---
##### Cluster
```
Collection of node instances that provide compute, memory, storage, and networking resources.
```

##### Nodes
```
Nodes are physical or virtual machines.
  * Master nodes act as a control plane for Kubernetes and manage worker nodes. These nodes follow raft protocol, 
    meaning an odd number of master nodes must exist for consensus to be possible. 
  * Worker nodes are instances that carry out actual computation work.
```

### KRM
---
##### KRM (Kubernetes Resource Model)
```
a) KRM is a declarative format used to talk to the Kubernetes API to express the desired state of the cluster. The 
   declarative format makes it particularly helpful for GitOps.
b) KRM is typically expressed in YAML format and follows OpenAPI standards. This allows developers to extend KRM with 
   tools that consume OpenAPIs (e.g. SDK generator, API doc generator). 
c) KRM allows environments to be configured and deployed in a repeatable and reliable way.
```

### Master Node Components
---
##### etcd 
```
a) etcd is a highly reliable distributed key-value store to back cluster level configuration data. This includes the 
   state of the cluster, various configuration data (network, authorization), and API metadata.
b) etcd provides watch functionality, allowing API server to be notified of any key-value changes.
c) Users can take snapshots of data stored in etcd and restore the cluster in the event of a cluster failure.
```

##### API Server
```
a) API server exposes Kubernetes REST API and is a means for components in the control plane and worker nodes to 
   communicate with one another.
b) API server is the only component connected to etcd, meaning all other components must pass through the API server to 
   work with cluster state.
c) API server is responsible for all authentication and authorization (e.g. kubectl) mechanisms.
d) API server provide watch functionality, allowing watchers (e.g. scheduler, controllers) to be notified of changes.
```

##### Controller Managers
```
a) Controllers act as non-terminating control loops that continue to poll from the API server to oversee the state of 
   various components. Controllers communicates with the API server to trigger events that move from the current state 
   to a desired state when necessary. Basic controllers such as node, service account, job controllers are compiled into
   a single binary and run as a single process known as the "kube controller manager".
b) Cloud controller managers allow the cluster to interact with the cloud provider's API. These controllers allow cloud 
   providers to implement platform logic for managing nodes, routes, volumes, load balancer services etc. 
c) Customer controller managers can be written to manage user defined custom resources following the KRM approach.
```

##### Scheduler
```
Scheduler acts as a non-terminating loop that continues to poll from the API server to oversee the state of scheduling.
If the scheduler identifies that a pod is not assigned to a node, it will determine the best node to schedule the pod to 
considering node availability, affinities, constraints etc. It sends this information back to the API server, which will
persist the info in etcd and communicate with the kubelet of the node for where the pod will be allocated.  
```

##### CoreDNS
```
a) CoreDNS is a pod that functions as a DNS server in the cluster. It has become the default over KubeDNS.
     * Every service receives an A record: <service-name>.<namespace>.svc.cluster.local
     * Every pod receives an A record: <pod-ip-address>.<namespace>.pod.cluster.local
b) When a pod is scheduled, kubelets will write the IP address of the DNS server to /etc/resolv.conf. 
```

### Master Node Component Optimizations
---
##### High Availability
```
a) Master nodes and their components should be redundant in case of an outage.
b) API server is stateless, so multiple API servers can be active at a time. A load balancer can be placed in front of
   master nodes to split traffic between API servers.
c) Schedulers and controller managers perform leader election by periodically attempting to obtain a lock. This prevents
   schedulers or controller managers from taking duplicate actions at the same time.
d) etcd can either run on the master node or run on an isolated server. Leader election for etcd is done through raft 
   protocol. All write requests are forwarded to the leader, and the leader will ensure a copy of the data is then sent 
   to all etcd instances. When a new etcd instance comes up, data is copied to it.
```

##### API Server Cache
```
a) Various Kubernetes components operate on snapshots of the cluster state rather than on real-time updates.
b) Current state of the cluster is maintained by the API server through an in-memory read cache.
c) Snapshots are saved in etcd, and the in-memory read cache will be updated by etcd watches.  
d) Caching states in the API server decreases read latency and reduces load on etcd.
```

##### Pod Lifecycle Event Generator (PLEG)
```
Kubelets used to poll container runtimes from each pod for info, which consistently puts pressure on the container 
runtimes' CPU usage. To overcome this, PLEG was implemented such that the Kubelet lists the state of pods and containers 
and compares with the previous state. The Kubelet then knows which pods need to sync again and only polls those pods.
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
Kube-proxies are agents on each node that watches the API server for new services and endpoints. It will then take this 
information and create forwarding rules (e.g. iptables) on its respective node.
```

##### Kubelet
```
a) Kubelets are agents on each node that registers nodes and reports the status of nodes/pods to the API server.
   Communication to the API server is typically done via TLS.
b) Kubelets receive information about pods that need to be newly created from the API server. Kubelets will interact 
   with CRIs to start pods and containers via their configured container runtime. Kubelets will then monitor and ensure
   containers run in a healthy state.
c) Kubelets ensure to inject necessary environment variables (e.g. service names) to pods before they start.
d) Kubelets help download secrets from API server, run liveness probes, and mount volumes.
e) Kubelet allows you to manually configure network MTU. Otherwise by default, network plugins will attempt to deduce 
   optimal MTU.
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
Step 1) Kubelets will interact with CRI compatible container runtimes or CRI shims (i.e. a bridge that implements CRI and 
        translates to something container runtimes can understand) over Unix sockets using the gRPC framework. 
Step 2) The CRI will start containers based on its logic. This includes setting up process and network namespaces. This 
        could also include invoking other plugins such as CNI as well.
Step 3) Once the containers are created and the pod is up, kubelets will reach out to the API server to persist this 
        information into etcd. 
```
