### Volumes
---
##### emptyDir
```
a) Ephemeral volume mounted to a pod that starts with empty contents. Containers in the same pod can access the same 
   content. Note that containers can have different mount paths even to the same emptyDir.
b) Contents are erased upon pod being deleted, but not erased when containers crash. Contents are not erased upon node 
   reboot as contents are still stored in disk.
```

##### RAM backed emptyDir
```
a) Faster read performance.
b) Contents are lost upon node restart as memory is reset.
c) By default, size of allocatable memory is half of node's RAM.
```

##### hostPath
```
a) Mounts directory from host node's file system into a pod. All pods on the same node that have this volume mounted can 
   read the same content.
b) Containers accessing host directories must be privileged to do so (e.g. root user) to allow writing.
```

### Persistent Volume Claims
---
##### Local Persistent Volume
```
a) Local persistent volume mounts a local disk directly to a node.
b) Local persistent volumes can specify node affinities to define which nodes to be attached to.
c) Unlike hostpath volumes, scheduler will ensure pods requiring persistent volumes will be scheduled on the same node.
```

##### Local vs Remote Persistent Volumes
```
a) Local storages provide better IOPS and lower latency as networking is not involved.
b) Local storages do not support dynamic volume provisioning and must be statically pre-provisioned and be tied to a
   specific node.
c) Most remote storages implement synchronous replication while most local storages do not. Loss of the disk or node 
   could result in loss of data on the persistent storage.
```

##### Provisioning
```
a) Static Provisioning  : Creates storage ahead of time and can be claimed later by containers.
b) Dynamic Provisioning : Provision storage on the fly if not available.
```

##### Persistent Volume Configurations
```
Capacity
  * Storage claims are satisfied by persistent volumes that have at least that amount of storage. Even if 10 pvs with 
    50G capacity are provisioned, a container claiming 100G will not be satisfied. In the same example, a PVC claiming
    10G will be bound to a volume with 50G capacity.

Volume Mode
  * Specifies whether the volume should use a file system or be raw block storage for efficient data transport.

Access Mode
  * ReadOnlyMany  (ROX) : Can be mounted as read only by many nodes.
  * ReadWriteOnce (RWO) : Can be mounted as read/write by a single node.
  * ReadWriteMany (RWX) : Can be mounted as read/write by several nodes.
  * Storages are mounted to nodes, so multiple pods on the same node can still read/write to RWO volumes. If this causes 
    an issue, it can be solved with a mode called ReadWriteOncePod.

Reclaim Policy
  * Reclaim policies determine what happens when a claim to the volume is deleted.
  * Retain : Volume is not deleted, it is not available for reuse by other claims, and must be reclaimed manually.
  * Delete : Volume is deleted.
  * Recycle: Data in the volume is deleted before making reuse.

Storage Class
  * Only volume claims that specify the same storage class name can claim the volume.

Volume Types
  * e.g. NFS, AWS EBS
```

##### Persistent Volume Examples
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-pv
  labels: 
    release: production
spec:
  capacity:
    storage: 100Gi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
  # can only be bound via PVCs specifying this storage class
  # PVs with no storageClassName can only be bound to PVCs that request no particular class
  storageClassName: example-class
  nfs:
    path: /tmp
    server: 192.19.1.22
```

##### Persistent Volume Claim
```
a) PVCs are a mechanism to claim a persistent volume based on matching specs (e.g. capacity, label selector access mode, 
   storage class, labels). Once a persistent volume is claimed by a PVC, it cannot be claimed by other PVCs.  
b) Kubernetes will attempt to try to bind PVCs to the smallest capacity volume available.
```

##### Persistent Volume Claim Example
```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi
  # must match storage class name of persistent volume claim
  # if storageClassName: "", matches any volume with no storage class name
  # if storageClassName is not specified, default storage class will be used
  storageClassName: example-class
  selector:
    # allows for filtering volumes by labels
    matchLabels:
      release: "production"
    # allows for filtering volumes by certain specs
    matchExpressions:
      - {key: capacity, operator: In, values: [80Gi, 100Gi]}
```

##### Storage Class
```
Describes the classes of storages that are offered. These classes might map to different quality of service levels or 
varying backup policies.  
```

##### Storage Class Examples
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: example-class
provisioner: kubernetes.io/aws-ebs
parameters:
  type: io1
  iopsPerGB: "10"
  fsType: ext4
# WaitForFirstCustomer indicates the PV is dynamically created only after a pod consumes a PVC
# Immediate indicates the PV is dynamically created after the PVC is created
volumeBindingMode: WaitForFirstCustomer
```

### CSI
---
##### Problem Statement
```
In the past, storage vendors would need to develop a volume plugin and add it to the source code to integrate their 
storage system. Development of volume plugins would heavily depend on Kubernetes release version and would also require 
the source code to be publicly available. Source code changes could also cause bugs that are difficult to test.
```

##### Container Storage Interface (CSI)
```
Specification for writing custom storage drivers via plugins that are integratable with various container orchestrators. 
With CSI, users are able to select storage solutions and the container orchestration system of their choice according to 
user needs. Vendors do not need to worry about Kubernetes source code or be locked down to Kubernetes.
```

### ConfigMaps
---
##### ConfigMaps
```
a) ConfigMaps are a means to keep configuration separate away from container images.
b) ConfigMap configurations can be consumed as environment variables, volumes, or secrets.
```

### Projections, Snapshots, and Cloning
---
##### Projections
```
Projection projects multiple volumes as a single volume on a single volume mount. This is supported for secrets, 
downward API, and configmaps.
```

##### Snapshots
```
a) Kubernetes allows for the snapshotting of a volume at a certain point of time.
b) Volumes can be provisioned from a snapshot.
c) Utilizes a sidecar container to create and delete snapshots.
d) Only available for CSI drivers.
```

##### Cloning
```
a) Volumes clones are new volumes populated with the content of an existing volume.
b) Only available for dynamic provisioning. 
c) Uses the storage class of the source volume.
d) Only available for CSI drivers.
```
