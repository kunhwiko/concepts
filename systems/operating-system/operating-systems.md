### Base Concepts
---
##### Operating System (OS)
```
a) OS is a software system that helps to manage and facilitate hardware resources and execution of user programs.
b) OS may also provide a high level interface to help users to interact with the computer system. 
c) OS typically provides the following functionality:
     * process management : scheduling tasks, concurrency handling, executing applications
     * memory management  : memory allocation, file system management
     * device management  : I/O management, hardware management
     * security           : network, security, and access control
```

##### Kernel
```
a) Kernel is a core component of the OS that provides low level services to manage and communicate with computer hardware.
b) Kernel provides functionality such as memory management, process scheduling, device driver management, and system call handling.
```

##### Shell
```
User interface to interact with kernel.
```

##### Driver
```
Drivers represent code that runs in the kernel to communicate with hardware devices.
```

##### Module
```
Modules represent code that can extend the behavior of the kernel and be removed on demand. 
```

### Kernel and User Space
---
##### Kernel Space
```
Kernel space is system memory reserved for running kernel, kernel extensions, and drivers.
```

##### User Space
```
a) User space is memory area where user processes from user applications run.
b) User space is isolated from kernel space through permission layers known as a "ring".
```

##### Protection Ring
```
Protection rings form layers of varying privileges that provides different levels of access to resources.
These privileges are assigned based on the CPU architecture that the OS is running on.
```

##### System Call (syscall)
```
System calls provide a standardized interface for programs to interact with the OS and leverage its capabilities.
These calls allow for user space applications to request for kernel level functionality.
If user space application has proper permissions, a context switch to the kernel space is made and the user space application awaits a response. 
```

### Process and Thread
---
### Process
```
a) Process is an instance of a computer program that is being executed and a unit of work that the OS executes.
   Each process has a unique ID that is assigned by the OS to manage and track the process.
b) Process has its own isolated memory space and system resources.
c) Process can communicate with other processes through inter-process communication (pipes, sockets, and shared memory).
d) Process takes time for context switching.
```

##### Thread
```
a) Thread is a segment of a process that can run concurrently with other components of the same process.
b) Thread shares resources with other threads in the same process.
c) Thread can communicate with other threads through shared variables and synchronization mechanisms.
d) Thread takes less time for context switching.
```

##### Scheduler
```
Scheduler is responsible for selecting the next process to run through a scheduling algorithm.
As part of this, it also chooses how long the process should run.
```

##### Dispatcher
```
Dispatcher is responsible for performing context switching of processes that the scheduler selects to run.
Context switching involves saving the state of the current process and loading state of the new process.
It ensures to take the selected process from the ready queue to be executed by CPUs.
```

### Virtualization
---
##### Hypervisor
```
Hypervisor is a software layer that breaks dependencies from the host OS and allows VMs to share hardware of the host.
It accomplishes this by assigning virtualized hardware resources to VMs that are isolated from other VMs.
```