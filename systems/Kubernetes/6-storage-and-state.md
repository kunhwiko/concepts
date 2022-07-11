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

##### Persistent Volumes and Claims
```
storageClass
   a) describes the class of storage 
   b) maps quality of service levels, backup policies

Provisioning
   a) static  : creates storage ahead of time and can be claimed later by containers
   b) dynamic : provision storage on the fly if not available
   
Persistent Volume Configurations
   a) capacity
      * storage claims are satisfied by persistent volumes that have at least that amount of storage
      * even if 10 pvs with 50G capacity are provisioned, a container claiming 100G will not be satisfied
   b) volume mode
      * file system vs raw block storage
      * block storage : direct access to block device without filesystem abstraction for efficient data transport
   c) access mode
      * ROX (read only by many nodes) vs RWO (read-write by one node) vs RWX (read-write by many nodes)
      * storages are mounted to nodes, so multiple containers/pods in a node can still write to a RWO storage
      * for RWO storages, if one claim is satisfied no other claim with RWO can be satisfied but ROX can still be satisfied
   d) reclaim policy
      * retain (volume must be reclaimed manually) vs delete vs recycle (retain but contents are deleted)
      * dynamically provisioned volumes always have delete policy
   e) storage class
      * only PVCs that specify the same storage class name can claim the volume
      * empty storage class name on PVC : match PV with no storage class name
      * no storage class line on PVC    : use default storage class
   f) volume type
      * e.g. nfs

PVC 
   a) mechanism to claim persistent volumes without knowing details of the particular cloud environment 
   b) finds a matching persistent volume based on specs (e.g. capacity, access mode, storage class, labels)
   c) Kubernetes will attempt to try to match to the smallest capacity volume available
   d) pods can choose which pvcs to mount by pvc name
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

##### Container Storage Interface (CSI)
```
Problems
   a) Storage vendors relied on Kubernetes in-tree (source code) volume plugins to support storage connectivity
   b) Vendors would need to develop a volume plugin and add it to the Kubernetes source code to integrate their storage system
   c) Source code changes could cause bugs and are difficult to test
   d) Development of volume plugin depends on Kubernetes release version and requires the source code to be publicly available

CSI
   a) Kubernetes has a CSI-compliant plugin that acts as a standard adapter between containerized workloads and storages
   b) Vendors need to develop a CSI-compliant driver in which the plugin will interact with through an API  
   c) Vendors do not need to be worried about the Kubernetes source code
   d) Plugin can be used for all container orchestrators that have a CSI-compliant plugin
```

##### Projections, Snapshots, Cloning
```
Projections
   a) project multiple volumes into a single volume on a single volume mount
   b) supported for secrets, downward API, configmaps

Snapshots 
  a) Kubernetes allows for the snapshotting of a volume at a certain point of time
  b) Only available for CSI drivers

Cloning
   a) New volumes populated with the content of the existing volume
   b) Works for dynamic provisioning and uses the storage class of the source volume
   c) Only available for CSI drivers
```

### Stateful Applications
---
##### Config Maps
```
Config Maps
   a) means to keep configuration separate from container image
   b) configurations can be consumed as environment variables, volumes, or secrets
```

##### Stateful Sets
```
Stateful Sets
   a) controller will guarantee pods are ordered and have unique identifiers 
   b) pods are associated with their own dynamic PVCs instead of sharing
   c) each unique identifier will fetch data from its corresponding PVC
   d) headless service is used to manage network identity of pods
   e) cannot roll back to previous versions

More here: https://medium.com/stakater/k8s-deployments-vs-statefulsets-vs-daemonsets-60582f0c62d4 
```
