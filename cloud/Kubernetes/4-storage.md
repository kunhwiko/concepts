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

StatefulSets are used for: 
   - Pods need access to the same persistent volume when restarted / redeployed 
   - App needs to communicate with replicas using predefined network identifiers  

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