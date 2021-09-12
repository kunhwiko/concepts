### Basics
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

##### Terms
```
K8s: abbreviation for Kubernetes 
Kubectl (Kube Control): CLI for K8s

Control Plane: set of master Nodes that manage K8s cluster
Node: individual workers in a K8s cluster

Raft protocol: odd number of master Nodes exist for consensus to be possible 
```

<br />

### Components
---
##### Getting Started
```
Great source to read: https://www.ibm.com/cloud/learn/kubernetes
```

##### Control Plane Components
```
1) etcd: 
     - distributed key-value store to back cluster data (e.g. configuration, state, metadata)
     - a means to restore K8s cluster by recording past snapshots of the cluster 
2) kube-apiserver: 
     - means for Control Plane and Nodes to communicate with one another 
     - a front end component that opens access to a K8s cluster (e.g. access via CLI) 
3) kube-scheduler: assigns Pods to Nodes 
4) kube-controller-manager
5) coreDNS: functions as DNS server in cluster 
```

##### Controller Manager
```
Controller Manager: control loop that oversees the cluster through the apiserver and moves current state to desired state

Types
    1) Deployments: 
        * defines a desired state for Pods and ReplicaSets 
        * enables users to scale number of replicas, control rollout of updates, rollback to previous deployments 
        * enables users to check or update status of Pods  
    2) ReplicaSets:
        * ensures that a specified number of Pod replicas are running at a given time 
    3) StatefulSets:
        * used when Pods need a persistent storage volume in the cluster (see "Storage" section below)
        * guarantees ordering / uniqueness of Pods and stable network identifiers 
```

##### Node Components
```
1) Pods 
2) kubelet
    - agents on each Node that registers Nodes to the kube-apiserver allowing Nodes to talk with the Control Plane
    - makes sure containers run in Pods in a healthy state 
3) kube-proxy:
    - implements networking rules that allow network communication to Pods from network sessions inside / outside the cluster
```

##### Pods 
```
1) encapsulates one or more containers to be assigned to a Node 
2) contain shared resources such as volumes, network configs, info on how to run containers 
3) basic unit of deployment 
```

<br />

### Networks 
---
##### Service
```
1) means to connect Pods to external services 
2) provides a DNS entry to a group of Pods that persists even if some Pods update their IP 
3) allows for load balancing between Pods   

Types
    1) ClusterIP:
        * exposes Service on an internal IP in the cluster 
        * reachable only from within the cluster 
    2) NodePort 
        * using NATs, exposes Service on the same port of each selected Node
        * makes a Service accessible outside the cluster using <NodeIP>:<NodePort>
        * requests to NodePorts get routed to ClusterIP services  
        * superset of ClusterIP
    3) LoadBalancer
        * mostly used with cloud services 
        * sets up clusterIPs / NodePorts and a great means to get external traffic to come into Service 
        * creates an external load balancer and assigns a fixed, external IP to the Service 
        * superset of NodePort
    4) ExternalName
        * means to get traffic out to an external source
        * adds CNAME DNS record to CoreDNS 
```

##### Ports 
```
NodePort vs Port vs TargetPort 
    - reference: https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ports-targetport-nodeport-service.html
```

##### NodePort vs Load Balancer
```
Things to know:
    1. Nodes have publicly accessible IPs
    2. NodePorts will open the same port number on all Nodes
    3. <NodeIP>:<NodePort> will convert requests to <ClusterIP>:<Port>  
    
Limitations of NodePorts  
    1. Only exposes a single service per port 
    2. will have to maintain and know the NodeIP of the Node you're looking for, which is difficult when many Nodes exist / Nodes crash or update 

Pros of Load Balancers 
    1. Only need to know the IP address of the Load Balancer 
    2. Transfers request of <LB IP>:<Port> to appropriate <NodeIP>:<NodePort>
    3. Ability to open multiple ports and protocols per service 
```

##### Load Balancer vs Ingress 
```
Limitations of Load Balancers 
    1. Each service exposed with a LoadBalancer costs money, and can be expensive with multiple LoadBalancers
```

<br />

### Storages  
---
##### Storage
```
Kubernetes is designed to:
    - keep containers ephemeral, immutable, replaceable 
    - we want a cloud based database to store data 
    - we do not want our cluster to be stateful 

Sometimes stateful workloads are inevitable, so we use StatefulSets / Persistent Volumes. 

Use Persistent Volumes if
    1. you need volumes that outlive the life of Pods 

Use StatefulSets if 
    1. Pods need access to the same persistent volume when restarted / redeployed 
    2. App needs to communicate with replicas using predefined network identifiers  

To learn more about the different controllers, states, persistent volumes and more:
https://medium.com/stakater/k8s-deployments-vs-statefulsets-vs-daemonsets-60582f0c62d4    
``` 

##### Ephemeral Volumes 
```
Ephemeral Volumes: volumes with the same lifetime of a Pod but persists beyond containers 
    - mountPath: directory on the container to access the volume  
    
1) emptyDir: initially empty volume created when a Pod is assigned to a Node 

2) configMap: mounts a configMap (configuration data) as a volume to inject into Pods  
```

##### Persistent Volumes
```
1) persistentVolumeClaim
    - used to mount PersistentVolumes into a Pod 
    - mechanism to claim persistent storage without knowing details of the particular cloud environment 

2) hostPath
    - mounts directory from host Node's filesystem into a Pod 
    - the above means the directory and Pod must be in the same Node
    - not recommended as a means for persistent storage, but suitable for development/testing purposes 
```

<br />

### Miscellaneous
---
##### Labels / Annotations / Selectors
```
Annotations: comments that provide extra context 

Labels: used to identify, select, and group Pods (or other objects) together based on some criteria  

Selectors: chooses objects based on some criteria (two or more selectors imply selector1 AND selector2 instead of OR) 
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

##### RBAC Authorization
```
ClusterRole 
    - allows users to access namespaced/cluster-wide resources 
    
Role 
    - allows users to access namespaced resources   
    
ClusterRoleBinding
    - grants permissions granted Roles/ClusterRoles cluster-wide 
    
RoleBinding
    - grants permissions granted by Roles/ClusterRoles within a specific namespace 
```

<br />

### Commands  
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
    
kubectl rollout status <deployment name> 
    --> waits until the deployment is ready 

kubectl delete deployment <deployment name>
    --> deletes deployment
```

##### Scale / Describe
```
kubectl scale deployment kunko --replicas 3
    --> create 3 replica Pods in the kunko deployment 

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

<br />

### YAML configurations (also reference Helm Charts)
---
##### Apply
```
kubectl apply -f filename.yml
    --> applies changes based on YAML file configs 

kubectl diff -f filename.yml
    --> finds diff of YAML file and previously applied configuration
```

##### Reference
```
kubectl api-resources
    --> quick lookup for names, versions, and kind for YAML file 

kubectl explain services/deployment/pods --recursive
    --> quick way to remind yourself how to build YAML files for a specific kind 

kubectl explain services.spec.<subfield1>.<subfield2>
    --> spec is a subfield of services
    --> add and remove subfields with dot notations to modify what you want to see 
```
