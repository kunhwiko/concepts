### Storages  
---
##### Storage
```
Kubernetes is designed to:
   - keep containers ephemeral, immutable, replaceable 
   - we want a cloud based database to store data 
   - we do not want our cluster to be stateful 

Sometimes stateful workloads are inevitable, so we use StatefulSets + Persistent Volumes. 

Persistent Volumes are used to: 
   - outlive the life of Pods 

StatefulSets are used to:
   - controller will guarantee the Pods are ordered and have a unique identifier 
   - Pods will individually use their own PVCs instead of sharing
   - each unique identifier will fetch data from its corresponding PVC  
   - cannot roll back to previous versions
   - more here: https://medium.com/stakater/k8s-deployments-vs-statefulsets-vs-daemonsets-60582f0c62d4    
``` 

##### Ephemeral Volumes 
```
Ephemeral Volumes: volumes with the same lifetime of a Pod but persists beyond containers  
    
emptyDir: 
   - initially empty volume created when a Pod is assigned to a Node 

configMap: 
   - mounts a configMap (configuration data) as a volume to inject into Pods  
```

##### Persistent Volumes
```
persistentVolumeClaim
   - mechanism to claim persistent volumes without knowing details of the particular cloud environment 

hostPath
   - mounts directory from host Node's filesystem into a Pod 
   - the above means the directory and Pod must be in the same Node
   - not recommended as a means for persistent storage, but suitable for development/testing purposes 
```