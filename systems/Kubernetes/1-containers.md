### Containers
---
##### Images vs Containers 
```
Image     : snapshot of an application along with necessary dependencies
Container : instance of an image running as a process 
```

##### Containers vs VM
```
Virtual Machines
   a) Hypervisors are used for emulating VMs and allocating resources
   b) environmental consistency
   c) ability to run different apps with different OS on the same physical server

Containers
   a) whereas VMs are packaged with their own OS, containers share the host OS and run as a process
   b) removes overhead of hypervisor
   c) lightweight, portable, faster deployment, availability, environmental consistency 
   d) reduces management of multiple OS to a single OS
```

##### Layers
```
Image Layers
   a) Images are built as a series of layers, users only need to copy over layers they don't have 
   b) If changes are made, a new layer is built on top of existing layers 
   c) If two different changes are made on top of the same existing layers, 
      two side by side layers are built on top of existing layers, 
      instead of duplicating the existing layers 

Container Layers 
   a) Containers continue to build layers on top of base images 
   b) Containers are just single read/write layers on top of base images 
   
Tags
   a) Pointer to a particular image commit / version of an image 
   b) Multiple tags can refer to the same commit, so they have the same image ID 
```

##### Separation of Concerns
```
Separation of Concerns / Best Practices
   a) changes should be made to source app, and then containers should be redeployed
   b) changes should not be made directly to containers as the results will not be reproducible
   c) containers should not contain unique data as they might have to be redeployed
```
