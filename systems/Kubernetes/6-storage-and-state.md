### Volumes  
---
##### emptyDir
```
emptyDir
   a) ephemeral volume mounted on a particular pod and starts with empty contents
   b) contents are erased upon pod being deleted, but not erased when containers crash
   c) each container can have different mount paths to the same emptyDir
   c) contents are not erased upon node reboot as contents are stored in disk

RAM backed emptyDir
   a) faster reads but more volatile
   b) contents are lost upon node restart
   c) by default, size of memory is half of node's RAM  
```

##### hostPath
```
hostPath
   a) mounts directory from host node's filesystem into a pod 
   b) containers accessing host directories must be privileged to do so (e.g. root user)
```

##### Persistent Volume Claims
```
storageClass : use a provisioner to allocate storage to pods

Provisioning
   a) static  : creates storage ahead of time and can be claimed later by containers
   b) dynamic : provision storage on the fly if not available

PVC
   a) mechanism to claim persistent volumes without knowing details of the particular cloud environment 
   b) can either be an external storage or disk physically attached to node
```

##### Local Persistent Volume
```
Local Persistent Volume
   a) local disk physically attached to nodes
   b) can use node affinity to bind the volume to particular nodes

Local Persistent Volume vs hostPath
   a) Kubernetes scheduler understands which node a local persistent volume belongs to
   b) ensures pod using a local persistent volume is always scheduled to the same node
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
