### Definition of Container and Image
---
##### Images
```
Snapshot of an application along with necessary dependencies.
```

##### Image Layers
```
a) Images are built as a series of layers, and only new layers need to be copied over. 
b) If changes are made, a new layer is built on top of existing layers. 
c) If two different changes are made on top of the same existing layer, two layers will exist side by side. 
```

##### Containers vs Images
```
Containers are an instance of an image and run as a process. 
```

##### Container Layers
```
a) Containers continue to build layers on top of base images.
b) Container layers are just single read/write layers on top of base images. 
```

##### Tags
```
a) Pointer to a particular image commit / version of an image.
b) Multiple tags can refer to the same commit, so they have the same image ID. 
```

### Definition of Container and Virtual Machine
---
##### Virtual Machines
```
a) Hypervisors are necessary in allocating resources from host machine and emulating VMs.
b) Provides environmental consistency.
c) Provides ability to run different apps with different OS on the same host machine.
```

##### Containers vs Virtual Machines
```
a) VMs are packaged with their own OS, containers share the host OS and run as a process.
b) Containers do not require a hypervisor and therefore are more lightweight.
c) Containers provide environmental consistency but are also portable and faster to deploy. 
d) Containers reduces management of multiple OS to a single OS.
```

### Best Practices in Containerization
---
##### Separation of Concerns
```
a) All changes should be made to the source app, and containers should then be redeployed.
b) Changes should not be made to a container directly as the results will not be reproducible.
c) Containers should not contain unique data as they might have to be redeployed.
```
