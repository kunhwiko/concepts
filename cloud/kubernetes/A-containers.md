### Definition of Container and Image
---
##### Images
```
a) Images are an executable package that includes code, libraries, and system tools.
b) Images represent a snapshot of an application with necessary dependencies.
```

##### Image Layers
```
Images are built as a series of layers, and only new layers need to be copied over. 
If changes are made, a new layer is built on top of existing layers. 
If two different changes are made on top of the same existing layer, two layers will exist side by side. 
```

##### Containers
```
a) Containers represent instances of images and run as processes on a host operating system.
b) Containers can have their own operating system but share the host's kernel through syscall requests.
c) Container processes are executed in an isolated Linux namespace and controlled by cgroups.
```

##### Containers vs Virtual Machines
```
a) VMs are packed with their own OS and virtualize an entire machine down to the hardware layers.
   Containers share the kernel of the host OS and run as processes in isolated user spaces.
b) VMs are relatively slow to start up as they need to boot up an OS.
   Containers are lightweight as there is no need to bootstrap a kernel but still provide environmental consistency.
c) VMs may incur the overhead of having to do kernel updates if thousands of VMs are active.
   Containers reduce management of multiple OS to a single OS.
```

##### Container Layers
```
Container runtime adds a new writable layer on top of base layers when a new container is launched.
Changes made to the running container (e.g. modification of files) will be written to this layer.
The contents of the container layer are lost when the container is deleted.
```

##### Container Registry
```
a) A container registry is a service for storing container images.
b) If a container registry is not specified, Kubernetes nodes will pull from docker.io. 
```

##### Tags
```
a) Pointer to a particular image commit / version of an image.
b) Multiple tags can refer to the same commit, so they have the same image ID. 
```

##### Separation of Concerns
```
a) All changes should be made to the source app, and containers should then be redeployed.
b) Changes should not be made to a container directly as the results will not be reproducible.
c) Containers should not contain unique data as they might have to be redeployed.
```

### Docker
---
##### Docker Daemon
```
Docker daemon (dockerd) listens for Docker API requests and manages Docker objects (e.g. images, containers, networks, volumes).
Docker daemon may send responses and stream output back to the Docker client (e.g. output to terminal after running a docker command).
The Docker infrastructure is better explained here: https://docs.docker.com/get-started/overview/#docker-architecture
```
