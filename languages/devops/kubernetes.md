### Kubernetes Basics
---
##### Orchestration 
```
Orchestration deals with certain problems:
    1. How do we automate container life cycles?
    2. How can we easily make our service scalable?
    3. How do we make our system fault tolerant and replace containers without downtime?
    4. How can we ensure containers run on trusted servers?
    5. How can we deal with security issues? 
```

##### Kubernetes Terms
```
K8s: abbreviated version of the Kubernetes orchestration engine
Kubetcl (Kube Control): CLI to configure K8s

Cluster: set of nodes that run containerized apps 
Control Plane: set of master nodes that manage K8s cluster

Node: individual workers in a K8s cluster
Kubelet: agents on each node that allows nodes to talk with control plane, makes sure containers run in pods 
Pods: 
    - encapsulates one or more containers to be assigned to a node 
    - contain shared resources such as volumes, network configs, info on how to run containers 

Raft protocol
    - odd number of master nodes exist for consensus to be possible 
```

##### Master Container Components
```
1) some container such as Docker 
2) etcd: distributed key-value store to back cluster data 
3) kube-apiserver: means for control plane and nodes to communicate with one another 
4) kube-scheduler: assigns pods to nodes 
5) kube-controller-manager: control loop that oversees the cluster through the apiserver and moves current state to desired state
6) coreDNS: functions as DNS in cluster 
```

##### Node Container Components
```
1) some container such as Docker
2) kubelet
3) kube-proxy: implements networking rules 
```