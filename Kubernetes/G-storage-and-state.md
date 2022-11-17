### Volumes  
---
##### emptyDir
```
EmptyDir
   a) Ephemeral volume mounted on a particular pod and starts with empty contents.
   b) Contents are erased upon pod being deleted, but not erased when containers crash.
   c) Each container can have different mount paths to the same emptyDir.
   c) Contents are not erased upon node reboot as contents are stored in disk.

RAM backed EmptyDir
   a) Faster reads but more volatile.
   b) Contents are lost upon node restart.
   c) By default, size of memory is half of node's RAM.
```

##### hostPath
```
HostPath
   a) Mounts directory from host node's filesystem into a pod.
   b) Containers accessing host directories must be privileged to do so (e.g. root user).
```

##### Persistent Volumes and Claims
```
StorageClass
   a) Describes the class of storage.
   b) Maps quality of service levels, backup policies.

Provisioning
   a) Static  : Creates storage ahead of time and can be claimed later by containers.
   b) Dynamic : Provision storage on the fly if not available.
   
Persistent Volume Configurations
   a) Capacity
      * Storage claims are satisfied by persistent volumes that have at least that amount of storage.
      * Even if 10 pvs with 50G capacity are provisioned, a container claiming 100G will not be satisfied.
   b) Volume Mode
      * File System vs Raw Block Storage
      * Block storage : Direct access to block device without filesystem abstraction for efficient data transport.
   c) Access Mode
      * ROX (read only by many nodes) vs RWO (read-write by one node) vs RWX (read-write by many nodes)
      * Storages are mounted to nodes, so multiple containers/pods in a node can still write to a RWO storage.
      * For RWO storages, if one claim is satisfied no other claim with RWO can be satisfied but ROX can still be satisfied.
   d) Reclaim Policy
      * Retain (volume must be reclaimed manually) vs Delete vs Recycle (retain but contents are deleted)
      * Dynamically provisioned volumes always have delete policy.
   e) Storage Class
      * Only PVCs that specify the same storage class name can claim the volume.
      * Empty storage class name on PVC : Match PV with no storage class name.
      * No storage class line on PVC    : Use default storage class.
   f) Volume Type
      * e.g. nfs

PVC 
   a) Mechanism to claim persistent volumes without knowing details of the particular cloud environment.
   b) Finds a matching persistent volume based on specs (e.g. capacity, access mode, storage class, labels).
   c) Kubernetes will attempt to try to match to the smallest capacity volume available.
   d) Pods can choose which pvcs to mount by pvc name.
```

##### Local Persistent Volume
```
Local Persistent Volume
   a) Local disk physically attached to nodes.
   b) Can use node affinity to bind the volume to particular nodes.

Local Persistent Volume vs HostPath
   a) Kubernetes scheduler understands which node a local persistent volume belongs to.
   b) Ensures pod using a local persistent volume is always scheduled to the same node.
```

##### Container Storage Interface (CSI)
```
Problems
   a) Storage vendors relied on Kubernetes in-tree (source code) volume plugins to support storage connectivity.
   b) Vendors would need to develop a volume plugin and add it to the Kubernetes source code to integrate their storage system.
   c) Source code changes could cause bugs and are difficult to test.
   d) Development of volume plugin depends on Kubernetes release version and requires the source code to be publicly available.

CSI
   a) Kubernetes has a CSI-compliant plugin that acts as a standard adapter between containerized workloads and storages.
   b) Vendors need to develop a CSI-compliant driver in which the plugin will interact with through an API.
   c) Vendors do not need to be worried about the Kubernetes source code.
   d) Plugin can be used for all container orchestrators that have a CSI-compliant plugin.
```

##### Projections, Snapshots, Cloning
```
Projections
   a) Project multiple volumes into a single volume on a single volume mount.
   b) Supported for secrets, downward API, configmaps.

Snapshots 
  a) Kubernetes allows for the snapshotting of a volume at a certain point of time.
  b) Only available for CSI drivers.

Cloning
   a) New volumes populated with the content of the existing volume.
   b) Works for dynamic provisioning and uses the storage class of the source volume.
   c) Only available for CSI drivers.
```

### Stateful Applications
---
##### ConfigMaps
```
a) Means to keep configuration separate from container image.
b) Configurations can be consumed as environment variables, volumes, or secrets.
```