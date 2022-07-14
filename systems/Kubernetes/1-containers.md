### Containers
---
##### Images vs Containers 
```
Image
   a) Snapshot of an application along with necessary dependencies.

Container
   a) Instance of an image running as a process. 
```

##### Containers vs VM
```
Virtual Machines
   a) Hypervisors are used for emulating VMs and allocating resources.
   b) Provide environmental consistency.
   c) Provide ability to run different apps with different OS on the same physical server.

Containers
   a) Whereas VMs are packaged with their own OS, containers share the host OS and run as a process.
   b) Removes overhead of needing a hypervisor.
   c) Are lightweight, portable, provide faster deployment, more availability, and environmental consistency. 
   d) Reduces management of multiple OS to a single OS.
```

##### Layers
```
Image Layers
   a) Images are built as a series of layers, users only need to copy over layers they don't have. 
      If changes are made, a new layer is built on top of existing layers. 
      If two different changes are made on the same existing layer, two side by side layers will be built on the existing layer, 
      instead of duplicating and splitting all of the previously existing layers. 

Container Layers 
   a) Containers continue to build layers on top of base images.
   b) Containers are just single read/write layers on top of base images. 
   
Tags
   a) Pointer to a particular image commit / version of an image.
   b) Multiple tags can refer to the same commit, so they have the same image ID. 
```

##### Separation of Concerns
```
Separation of Concerns / Best Practices
   a) Changes should be made to source app, and then containers should be redeployed.
   b) Changes should not be made directly to containers as the results will not be reproducible.
   c) Containers should not contain unique data as they might have to be redeployed.
```
