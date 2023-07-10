### Volumes
---
##### emptyDir
```
a) Ephemeral volume mounted on a particular pod that starts with empty contents.
   Containers in the pod can read the same content.
b) Each container can have different mount paths to the same emptyDir.
c) Contents are erased upon pod being deleted, but not erased when containers crash.
   Contents are not erased upon node reboot as contents are stored in disk.
```

##### RAM backed emptyDir
```
a) Faster read performance.
b) Contents are lost upon node restart.
c) By default, size of allocatable memory is half of node's RAM.
```

##### hostPath
```
a) Mounts directory from host node's file system into a pod.
   All pods on the same node that have this volume mounted can read the same content.
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
a) Local storages provides more consistent high performance as networking is not involved.
b) Local storages do not support dynamic volume provisioning.
c) Local storages must be pre-provisioned and be tied to a specific node.
   There must be sufficient space on the node for pods requiring the local storage to be scheduled.
d) Most remote storages implement synchronous replication while most local storages do not.
   Loss of the disk or node could result in loss of all data in the local storage.
```

##### Provisioning
```
a) Static Provisioning  : Creates storage ahead of time and can be claimed later by containers.
b) Dynamic Provisioning : Provision storage on the fly if not available.
```

##### Persistent Volume Configurations
```
Capacity
  a) Storage claims are satisfied by persistent volumes that have at least that amount of storage.
  b) Even if 10 pvs with 50G capacity are provisioned, a container claiming 100G will not be satisfied.

Volume Mode
  a) Specifies whether the volume should use a file system or be raw block storage for efficient data transport.

Access Mode
  a) ReadOnlyMany (ROX)  : Can be mounted as read only by many nodes.
  b) ReadWriteOnce (RWO) : Can be mounted as read write by a single node.
  c) ReadWriteMany(RWX)  : Can be mounted as read write by mane nodes.
  d) Storages are mounted to nodes, so multiple pods on the same node can still read write to RWO volumes.
     If this causes an issue, it can be solved with a mode called ReadWriteOncePod.

Reclaim Policy
  a) Reclaim policies determine what happens when a volume claim is deleted.
  b) Retain : Volume must be reclaimed manually.
  c) Delete : Volume is deleted.

Storage Class
  a) Only volume claims that specify the same storage class name can claim the volume.
  c) Not specifying a storage class means claims that do not specify a storage class can bind to the volume.

Volume Types
  a) e.g. NFS
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
  storageClassName: example-class
  nfs:
    path: /tmp
    server: 192.19.1.22
```

##### Persistent Volume Claim
```
a) Mechanism to claim a persistent volume based on matching specs (e.g. capacity, access mode, storage class, labels).
b) Kubernetes will attempt to try to match to the smallest capacity volume available.
c) Pods mount claims and not volumes, meaning volumes will only be available once the claim is satisfied.  
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
  # if storageClassName is empty (""), matches volumes with no storage class name
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
Describes the classes of storages that are offered.
These classes might map to different quality of service levels or to backup policies. 
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
```

### CSI
---
##### Problem Statement
```
a) Storage vendors relied on Kubernetes in-tree (source code) volume plugins to support storage connectivity.
b) Storage vendors would need to develop a volume plugin and add it to the source code to integrate their storage system.
c) Source code changes could cause bugs that are difficult to test.
d) Development of volume plugin depends on Kubernetes release version and requires the source code to be publicly available.
```

##### Container Storage Interface (CSI)
```
Initiative and specification to write various storage solutions via plugins that are integratable with various container orchestrators.
Users are able to adopt storage solutions and the container orchestration system of their choice according to different needs.
Vendors do not need to worry about Kubernetes source code or be locked down to Kubernetes this way.
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

### ConfigMaps
---
##### ConfigMaps
```
a) ConfigMaps are a means to keep configuration separate away from container images.
b) ConfigMap configurations can be consumed as environment variables, volumes, or secrets.
```
