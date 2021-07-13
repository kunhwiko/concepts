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
K8s: abbreviation for Kubernetes 
Kubectl (Kube Control): CLI for K8s

Cluster: set of nodes that run containerized apps 
Control Plane: set of master nodes that manage K8s cluster
Node: individual workers in a K8s cluster
Kubelet: agents on each node that allows nodes to talk with control plane, makes sure containers run in Pods 

Pod
    - encapsulates one or more containers to be assigned to a node 
    - contain shared resources such as volumes, network configs, info on how to run containers 
    - basic unit of deployment 

Controller Manager
    - control loop that oversees the cluster through the apiserver and moves current state to desired state
    - Deployments: 
        * defines a desired state for Pods and ReplicaSets 
        * enables users to scale number of replicas, control rollout of updates, rollback to previous deployments 
        * enables users to check or update status of Pods  
    - ReplicaSets:
        * ensures that a specified number of Pod replicas are running at a given time 

Service
    - provides a DNS entry to easily locate a group of Pods even if changes exist
    - allows for load balancing between Pods   

Raft protocol
    - odd number of master nodes exist for consensus to be possible 
```

##### Master Container Components
```
1) some container such as Docker 
2) etcd: distributed key-value store to back cluster data 
3) kube-apiserver: means for control plane and nodes to communicate with one another 
4) kube-scheduler: assigns Pods to nodes 
5) kube-controller-manager
6) coreDNS: functions as DNS in cluster 
```

##### Node Container Components
```
1) some container such as Docker
2) kubelet
3) kube-proxy: implements networking rules 
```

### Kubernetes Commands 
---
##### Basic Commands
```
kubectl version 
    --> check version

kubectl run 
    --> used for pod creation

kubectl create
    --> used to create resources via CLI / YAML

kubectl apply 
    --> create or update via YAML
```

##### Run / Create / Get / Delete 
```
kubectl run kunko --image=nginx 
    --> start an nginx pod named kunko

kubectl create deployment kunko --image=nginx
    --> creates an nginx deployment named kunko

kubectl get pods 
    --> show pods 

kubectl delete deployment <deployment name>
    --> deletes deployment
```

##### Scale / Describe
```
kubectl scale deployment kunko --replicas 3
    --> create 3 Pods in the kunko deployment 

kubectl describe pod <pod name>
    --> get details on Pod 

kubectl describe deployment <deployment name>
    --> get details on Deployment
```