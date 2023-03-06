### GCE Storage
---
##### Boot Disk
```
a) Default storage for GCE is a standard persistent disk (block storage).
   The default storage can be changed to SSD for more IOPS at a higher price.
b) New instances can be booted with existing snapshots or even existing disks.
```

##### Snapshots
```
a) Snapshots of existing and running persistent disks can be used when a new instance is being created.
b) Snapshots are global resources that can be accessed by any resource in the same project. 
c) Snapshots can be periodically scheduled and are incremental.
   This means the first snapshot will be a full snapshot, and subsequent snapshots will contain diffs of previous snapshots.
   Each snapshot is stored across multiple locations and can be stored across projects with the correct permissions. 
```

### GCE Networking
---
##### Bastion
```
If there is no external IP to access GCE VMs, a bastion host (i.e. terminal server) must be used.
Bastion is configured with external IP addresses and act as a gateway to internal resources.
```

### GCE Management
---
##### Labels
```
Labels allow for organizing and filtering instances for billing and reporting purposes.
```