### Namespace & Cgroup
---
##### Namespace
```
Linux kernel feature that partitions kernel resources such that one set of processes are isolated from the host system
and other set of processes. Namespaces allow for the isolation of UID/GID assignment, process runtimes (e.g. PID), 
network stack, file system, and IPC resources.
```

##### User Namespace
```
User namespace has its own set of UIDs/GIDs for assignment to processes and Linux capabilities. This means that a process 
can have root privilege within its user namespace without having it in other user namespaces.
```

##### PID Namespace
```
PID namespace assigns a set of PIDs to processes that are independent from the set of PIDs in other namespaces.
```

##### Network Namespace
```
Network namespace has its own network stack, including its own network devices (i.e. IP addresses), IP routing tables,
sockets (i.e. ports), firewall rules. Virtual ethernet pairs can be used to create tunnels between network namespaces
and can be used to create a bridge to a physical network device in another namespace. 
```

##### Mount Namespace
```
Mount namespace has its own independent list of mount points seen by processes in a given namespace. This allows for 
mounting file systems without affecting the file system of other namespaces.
```

##### Interprocess Communication (IPC) Namespace
```
IPC namespace has its own IPC resources (e.g. POSIX message queue).
```

##### Control Group (cgroup)
```
Linux kernel feature that limits, isolates, and monitors resource usage (CPU, memory, disk I/O, network) of a collection 
of processes. It also allows for the prioritization of resource allocation under resource limits.  
```