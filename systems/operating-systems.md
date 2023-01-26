### Operating Systems 
---
##### Basic Terminology 
```
What does an OS do?
1) process management : scheduling, allow concurrency
2) memory management : memory allocation, allow different files to share storage  
3) I/O management 
4) security : network, security, and access control 

Kernel vs Shell
1) kernel : core of operating system that controls tasks 
2) shell : interface to communicate with kernel 
```

##### Processes / Threads 
```
Process vs Thread 
1) process : program that is being executed (heavy, isolated memory, takes time to switch)
2) threads : segments of a process (light, shares memory allocated to the process, fast switch times)  

Dispatching 
1) process control block : data structure to store info of each process
  a) save execution state of each thread/process (running, blocked, waiting)
  b) track scheduling information 
  c) track memory usage and open files 
2) dispatcher : performs context switching 
3) context switching : saves state of the old thread/process, then loads the state of new thread/process

Trap vs Interrupt 
Makes the dispatcher run 
1) trap : events within current thread that causes state switches 
  - illegal instructions, zero division error, page fault 
2) interrupt : events outside current thread that causes state switches 
  - completion of some operation 
  - keyboard input 
  
Process Creation 
Step 1) Initialize a process control block
Step 2) Load program into memory 
Step 3) Create thread structures and initial value for the state of threads 
Step 4) Make thread known to the dispatcher 

When a program is loaded into memory and becomes a process, memory is divided into 4 sections 
1) Text : current activity as represented by contents of the registers 
2) Data : section for global and static variables 
3) Stack : temporary data such as function parameters, addresses, local variables 
4) Heap : dynamically allocated memory during runtime 
```

###### Concurrency / Scheduling
```
Locks 
1) synchronization : forces threads to run one at a time to prevent race conditions 
2) mutual exclusion : using methods (e.g. locks) to achieve synchronization 
3) locks : prevents race conditions between threads 
4) deadlocks : processes are blocked because each process holds the resource needed for the other to go further

Scheduling 
1) scheduler : process that chooses what process to run and for how long 
2) time slicing : divide CPU time between available threads 
3) FIFO scheduling : run thread in FIFO order 
4) round robin : run thread for one time slice, then move it back to queue 
5) SRPT scheduling : run thread that will finish most quickly
6) priority scheduling : supports a priority order 
```

##### Dynamic Linking / Memory Management  
```
Executing
1) compiler : generate one object file for each source code file (may reference other source files)
2) linker : combines all object files into a single executable file (re-organizes storage allocation)
3) OS : loads executable files into memory (allows several different processes to share memory at once)
4) Runtime library : works with OS to execute dynamic allocation (malloc, free)

Components of Object Files
1) sections : distinct info such as text/data, starting addresses, and initial contents  
2) symbol table : name and current location of each variable 
3) relocation record : addresses in the object file that linker will adjust to a final memory allocation

Linking Object Files (Static Linking)
1) Read section and compute memory layout 
2) Read symbols and construct a complete symbol table 
3) Read relocation record, update addresses, write executable file 

Dynamic Linking
Instead of linking programs and resolving references beforehand, shared references are resolved during runtime
Redundancy of loading the same shared libraries is replaced with shared memory 
1) copy shared libraries from persistent storage into RAM 
2) read symbol table from libraries 
3) fill a jump table with addresses for each function to "jump" to a function in the shared libraries 

Shared Memory 
1) Multiple processes should be allowed inside memory at once
2) No process should be aware that memory is being shared
3) Processes must not corrupt one another 
4) Efficiency should not be affected by sharing 

Memory Management 
1) Stack Allocation
  a) used when allocation and freeing is predictable 
  b) all free spaces are stacked on top of one another and is very efficient for memory management 
  
2) Heap Allocation
  a) used when allocation and freeing is unpredictable
  b) spaces might not be contigously allocated
  c) must keep track of storage that needs to be freed (potential for memory leak)

3) Reference Counting / Generational Garbage Collection
(reference the Python md file for more details)
```

##### Virtual Memory 
```
Intuition
1) The amount of addressable/needed memory could be larger than the actual physical memory
2) We don't need all processes to have data in physical memory, just the ones being executed
3) For processes being executed, we just need parts that are currently being used 

Virtual Memory
1) Memory-Management Unit (MMU) gives each program its own virtual address space
2) Maps executing programs into memory and idle programs into disk, but give illusion that everything is in memory

Paging
  - virtual (addressable) memory is broken into pages 
  - pages in use reside in physical memory and are called "frames"
  - other pages reside in a backing store/hard disk
  - page table : maps virtual pages into page frames in physical memory 
  - demand paging : brings programs from disk into physical memory when needed  

Page Fault : causes trap to OS when accessing memory not in physical memory (RAM) but in the backing store 
Resolving page faults 
Step 1) save current instruction states and register content 
Step 2) locate the virtual page needed and allocates to a free page frame 

Page Replacement : determine which pages to throw out of memory (FIFO, Random, LRU, Clock Algorithm)

Page Fetching : determine when to bring pages into memory 
  - start process with no pages loaded, and use demand paging to load pages into memory when it gets referenced 
  
Inter-Process Communication
Allow processes to communicate with one another and synchronize actions
1) Shared Memory Method : one process produces an item into shared memory, where another process will consume the item
2) Message Parsing Method : Establish a communication link and exchange messages 
```

##### File Systems and I/O
```
Goals 
1) manage efficient use of disk space, fast access, and sharing between users
2) information stored must be durable to survive failures 
3) must guarantee isolation between users 

Inode 
- OS data structure that carries info about a particular file (file size, occupied space, access times, id)
- stored on disk along with the file, and kept in memory when file is opened 
- various representations of how disk can represent bytes of a file (linked list of 4096 byte "blocks")

Directories : special structures used to map text names to inode id numbers (directories are organized as a tree structure) 
Working Directory : OS remembers the inode number of the current directory 

I/O or Disk Scheduling : order of executing I/O
1) FIFO
2) shortest positioning time first : chooses the request that is closest to the previous one 
3) elevator algorithm / scan :  have head move in one direction serving all requests along the way, then move in opposite direction

File System Crash Recovery 
1) consistency check on reboot : restores consistency, but does not prevent information loss 
2) logging : use a log entry to ensure consistency (refer to systems-design.md)
```

##### Distributed File Systems 
```
Distributed File Systems 
- client/server application where clients can access files from the server as if it were from the local machine 
- server sends a copy of the file that gets cached into the client machine
- typically has file/database replication across multiple servers to protect against failures 
- must have a mechanism in place to organize updates across all replicas 
```

### Virtualization
---
##### Hypervisor
```
Software layer that breaks dependencies of an operating system on the host machine.
It then allows virtual machines to share hardware of that host machine. 
```