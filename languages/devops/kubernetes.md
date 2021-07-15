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

Cluster: set of Nodes that run containerized apps 
Control Plane: set of master Nodes that manage K8s cluster
Node: individual workers in a K8s cluster
Kubelet: agents on each Node that allows Nodes to talk with control plane, makes sure containers run in Pods 

Raft protocol
    - odd number of master Nodes exist for consensus to be possible 
```

##### Pods 
```
1) encapsulates one or more containers to be assigned to a Node 
2) contain shared resources such as volumes, network configs, info on how to run containers 
3) basic unit of deployment 
```

##### Controller Manager
```
1) control loop that oversees the cluster through the apiserver and moves current state to desired state

Types
    1) Deployments: 
        * defines a desired state for Pods and ReplicaSets 
        * enables users to scale number of replicas, control rollout of updates, rollback to previous deployments 
        * enables users to check or update status of Pods  
    2) ReplicaSets:
        * ensures that a specified number of Pod replicas are running at a given time 
```

##### Service
```
1) means to connect Pods to external services 
2) provides a DNS entry to a group of Pods that persists even if some Pods update 
3) allows for load balancing between Pods   

Types
    1) ClusterIP:
        * default Service type 
        * exposes Service on an internal IP in the cluster 
        * reachable only from within the cluster 
    2) NodePort 
        * exposes Service on the same port of each selected Node using NAT 
        * makes a Service accessible outside the cluster using <NodeIP>:<NodePort>
    3) LoadBalancer
        * mostly used with cloud services 
        * sets up clusterIPs / NodePorts and a great means to get external traffic to come into Service 
        * creates an external load balancer and assigns a fixed, external IP to the Service 
    4) ExternalName
        * means to get traffic out to an external source
        * adds CNAME DNS record to CoreDNS 
```

##### Master Container Components
```
1) some container such as Docker 
2) etcd: distributed key-value store to back cluster data 
3) kube-apiserver: means for control plane and Nodes to communicate with one another 
4) kube-scheduler: assigns Pods to Nodes 
5) kube-controller-manager
6) coreDNS: functions as DNS server in cluster 
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

##### Expose
```
kubectl expose deployment <deployment name> --port=8888
    --> start ClusterIP
    --> Deployment listens on port 8888

kubectl expose deployment <deployment name> --port=8888 --name=test --type=NodePort 
    --> start NodePort named test 
    --> Deployment listens on port 8888, Node listens to external sources on port 30000 - 32767
    --> you can curl into the high ports (30000 - 32767) to access the K8s cluster 
```
