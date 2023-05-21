### Containers and Images
---
##### Images
```
Images are a read-only template that represent snapshots of an application with necessary dependencies
at a given point in time. The snapshot holds all necessary info to instantiate and run a container.
```

##### Image Layers
```
Images are built as a series of layers, or intermediate images that each represent changes from previous layers
(e.g. ENV, COPY, RUN commands in Dockerfiles each represent a layer). These layers are read only and if changes 
are made, only the new layers need to be pulled or pushed on top of existing layers residing in an image cache. 
```

##### Containers
```
Containers represent instances of images that run as processes in isolated user spaces on the host OS.
Container processes use Linux features such as namespace and cgroups to provide this isolation.
Containers can have a different OS from the host and will share the host's kernel through syscall requests.
```

##### Container Layers
```
Container runtime adds a new read-write layer on top of base layers when a new container is launched.
Changes made to the running container (e.g. modification of files) will be written to this layer.
The contents of the container layer are lost when the container is deleted.
```

##### Containers vs Virtual Machines
```
a) VMs are packed with their own OS and virtualize an entire machine down to the hardware layers.
   Containers share the kernel of the host OS and run as processes in isolated user spaces.
b) VMs are relatively slow to start up as they need to boot up an OS. Containers are lightweight as 
   there is no need to bootstrap a kernel but still provide environmental consistency.
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

### Dockerfile
---
##### Dockerfile
```
Dockerfiles list instructions (e.g. FROM, ADD, COPY, LABEL, WORKDIR, RUN, CMD) to build out an image.
Each instruction typically represents a single layer and will be its own intermediary image up to that instruction.
It is recommended to use .dockerignore and avoid copying unnecessary files or packages to keep image size small.
```

##### CMD vs ENTRYPOINT
```
a) If multiple CMD commands exist, all except the last one are ignored (i.e. there is an expectation that it 
   can be overriden). It could be further overrriden through arguments to the `docker run` command. 
b) ENTRYPOINT instructions on the other hand will always be executed and are useful in multi-stage builds
   when a command always needs to be executed.
c) CMD can be used alongside ENTRYPOINT to specify parameters or arguments to the ENTRYPOINT executable.
```

##### ADD vs COPY
```
a) COPY supports copying of local files into a container and has a much more clear purpose of what to port over.
b) ADD has additional functionalities such as automatic tar extraction and fetching of packages from remote URLs.
   While decompression of compressed files is a valid feature, COPY and curl for remote URLs is generally recommended.  
```

##### Multiple Staging
```
Multiple staging allows users to keep image sizes small by using multiple images from a single Dockerfile to build a 
final image. Below is an example:

Step 1) First stage of a Dockerfile will download a repository holding Golang code and necessary dependencies.
Step 2) This stage will compile the Golang code into a binary using necessary dependencies.
Step 3) The second stage specifies a new base image and can copy just the binary. The repository, other 
        dependencies, and the previous base image are discarded and only the final stage is kept.
```

### Daemons
---
##### Open Container Initiative (OCI)
```
Set of vendor-neutral standards for container runtime that defines how images should be structured (e.g. 
file system, metadata) and how containers should be launched, managed, and interact with host system
(e.g. process isolation, resource constraints, lifecycle management). 
```

##### runc
```
Lightweight container runtime that is universal per OCI standards and manages the following via system calls:
  * Sets up container namespaces.
  * Sets up cgroups and capabilities.
  * Sets up file system.
  * Typically invoked via higher level container runtimes (e.g. containerd) and not by end users.
```

##### containerd
```
containerd is a container runtime that serves as a foundation for containers and provides the following:
  * Supports pulling and pushing container images.
  * Handles the creation, execution, and management of containers.
  * Manages the creation and handling of container filesystem snapshots.
  * Integratable with networking and storage plugins to enable container networking and persistence.
```

##### Docker Daemon (dockerd)
```
a) dockerd listens for Docker API requests and manages Docker objects (e.g. images, containers, networks, volumes).
b) dockerd can invoke containerd via a gRPC request to start a container.
c) dockerd may send responses and stream output back to the client (e.g. output to terminal after running docker run).

The Docker infrastructure is better explained here: https://docs.docker.com/get-started/overview/#docker-architecture.
```

##### Kaniko
```
Open source tool that enables a container to build other container images without requiring privileged access or the
need for dockerd. It also supports authentication to push images to private registries. 
```
