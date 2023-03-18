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