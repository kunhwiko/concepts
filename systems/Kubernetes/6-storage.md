### Storages  
---
##### Storage
```
Best Practices
   a) keep containers ephemeral, immutable, replaceable 
   b) cloud based database should be used to store data 
   c) cluster should ideally not be stateful 

Sometimes stateful workloads are inevitable, so we use stateful sets & persistent volumes. 

Persistent Volumes are used to
   a) outlive the life of pods 

Stateful Sets are used to
   a) controller will guarantee the pods are ordered and have a unique identifier 
   b) pods will individually use their own PVCs instead of sharing
   c) each unique identifier will fetch data from its corresponding PVC  
   d) cannot roll back to previous versions
   e) more here: https://medium.com/stakater/k8s-deployments-vs-statefulsets-vs-daemonsets-60582f0c62d4    
``` 

##### Ephemeral Volumes 
```
Ephemeral Volume : volumes with the same lifetime of a pod but persists beyond containers  
    
emptyDir : initially empty volume created when a pod is assigned to a node 

configMap : mounts a config map (configuration data) as a volume to inject into pods  
```

##### Persistent Volumes
```
persistentVolumeClaim : mechanism to claim persistent volumes without knowing details of the particular cloud environment 

hostPath
   a) mounts directory from host node's filesystem into a pod 
   b) the above means the directory and pod must be in the same node
   c) not recommended as a means for persistent storage, but suitable for development/testing purposes 
```