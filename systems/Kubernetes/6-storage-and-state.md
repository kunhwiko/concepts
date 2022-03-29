### Storages  
---
##### emptyDir
```
emptyDir : initially empty volume created when a pod is assigned to a node 
```

##### Other Ephemeral Volumes 
```
configMap : mounts a config map (configuration data) as a volume to inject into pods
secrets   : see 5-security.md
```

##### hostPath
```
hostPath
   a) mounts directory from host node's filesystem into a pod 
   b) the above means the directory and pod must be in the same node
   c) not recommended as a means for persistent storage, but suitable for development/testing purposes 
```

##### persistentVolumeClaims
```
persistentVolumeClaim : mechanism to claim persistent volumes without knowing details of the particular cloud environment 
```

### Stateful Applications
---
##### Stateful Sets
```
Stateful Sets
   a) controller will guarantee the pods are ordered and have a unique identifier 
   b) pods will individually use their own PVCs instead of sharing
   c) each unique identifier will fetch data from its corresponding PVC  
   d) cannot roll back to previous versions
   e) more here: https://medium.com/stakater/k8s-deployments-vs-statefulsets-vs-daemonsets-60582f0c62d4 
```
