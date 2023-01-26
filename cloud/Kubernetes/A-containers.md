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
Container runtime adds a new writable layer on top of base layers when a new container is launched.
Changes made to the running container (e.g. modification of files) will be written to this layer.
The contents of the container layer are lost when the container is deleted.
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
c) Provides ability to run different apps with a different OS on the same host machine.
```

##### Containers vs Virtual Machines
```
a) VMs virtualize an entire machine down to the hardware layers, and are packaged with their own OS.
   Containers virtualize software layers above the operating system level and share the host OS as isolated user spaces.
b) VMs are relatively slow to start up as they need to boot up an OS.
   Containers are lightweight as they run as OS processes while providing environmental consistency.
c) VMs may incur the overhead of having to do kernel updates if thousands of VMs are active.
   Containers reduce management of multiple OS to a single OS.
d) Containers use Linux cgroups to control CPU, memory, I/O consumption.
e) Containers use Linux namespaces to isolate applications.
```

### Best Practices in Containerization
---
##### Separation of Concerns
```
a) All changes should be made to the source app, and containers should then be redeployed.
b) Changes should not be made to a container directly as the results will not be reproducible.
c) Containers should not contain unique data as they might have to be redeployed.
```
