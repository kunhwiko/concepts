### Operating Systems 
---
##### Basic Terminology 
```
What does an OS do?
1) process management, scheduling, and allow concurrency
2) memory management and allocation 
3) I/O management 
4) allow different files to share storage 
5) network, security, and access control 

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
When a program is loaded into memory and becomes a process, memory is divided into 4 sections 
1) Text : current activity as represented by contents of the registers 
2) Data : section for global and static variables 
3) Stack : temporary data such as function parameters, addresses, local variables 
4) Heap : dynamically allocated memory during runtime 

Memory Allocation
1) compiler : generate one object file for each source code file (may reference other source files)
2) linker : combines all object files into a single executable file (re-organizes storage allocation)
3) OS : loads executable files into memory and allows several processes to share memory at once 
4) Runtime library : works with OS to execute dynamic allocation (malloc, free)

Components of Object Files
1) sections : distinct info such as text/data, starting addresses, and initial contents  
2) symbol table : name and current location of each procedure or variable 
3) relocation record : addresses in the object file that linker will adjust to a final memory allocation

Linking Object Files (Static Linking)
1) Read section and compute memory layout 
2) Read symbols and construct a complete symbol table 
3) Read relocation record, update addresses, write executable file 

Dynamic Linking
Instead of linking programs and resolving references beforehand, shared references are resolved during runtime
1) copy shared libraries from persistent storage into RAM 
2) read symbol table from libraries 
3) fill a jump table with addresses for each function to "jump" to a function in the shared libraries 

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
Shared Memory 
1) Multiple processes should be allowed inside memory at once
2) No process should be aware that memory is being shared
3) Processes must not corrupt one another 
4) Efficiency should not be affected by sharing 

Intuition 
1) The amount of addressable/needed memory could be larger than the actual physical memory
2) We don't need all processes to have data in physical memory, just the ones being executed
3) For processes being executed, we just need parts that are currently being used 

Paging / Dynamic Memory Allocation
Virtual (addressable) memory is broken into pages 
  - pages in use reside in physical memory and are called "frames"
  - other pages reside in a backing store/hard disk 
  - page table : maps virtual pages into page frames in physical memory 
  
Page Fault : causes trap to OS when accessing memory not in physical memory (RAM) but in the backing store 
Resolving page faults 
Step 1) save current instruction states and register content 
Step 2) locate the virtual page needed and allocates to a free page frame 

Page Replacement : determine which pages to throw out of memory (FIFO, Random, LRU, Clock Algorithm)
Page Fetching : determine when to bring pages into memory 
  - demand fetching : start process with no pages loaded, and load pages into memory when it gets referenced 
  
Inter-Process Communication
Allow processes to communicate with one another and synchronize actions
1) Shared Memory Method : one process produces an item into shared memory, where another process will consume the item
2) Message Parsing Method : Establish a communication link and exchange messages 
```


<br />

### Object Oriented Programming
---
##### Abstraction 
```
hide all but important features to reduce complexity and think about things at a higher level

ex) create an object TV to focus on a higher level, rather than the specifics/internals of a TV
``` 

##### Encapsulation 
```
bundling data fields as private to prevent open access and change  
```

##### Inheritance
```
parent class : Dog 
child class : Retriever 

child class can inherit all data fields and methods of parent class, while defining its own fields or methods 

Advantages : 
  1) Reusability : Do not have to write already existing lines of code 
```

##### Polymorphism
```
Triangle shape = new Triangle() 
Shape shape = new Triangle()           class triangle is a part of class shape
Shape shape = new Circle()             class circle is a part of class shape 
Triangle shape = new Shape()           a triangle is a shape, but a shape is not a triangle (invalid)

public void run(Shape shape) {}        classes shape, triangle, and circle can all be passed to the parameter

Advantages :
  1) Flexibility 
```

##### Composition 
```
defines a has-a relationship
public class Point3D {
  public Point3D(int x, int y, int z) {
    p = new Point2D(x,y);
  }
}
```

##### Dependency Injection 
```
Suppose you have a class that depends on another object 

public Car() {
  Wheel wheel = new CanadianWheel();
  Windows window = new CanadianWindow();
}

Wheel and windows are 'dependencies' of car because they are needed to construct the car 

What if we want to create multiple objects of car, but with different wheels, companies, and materials?

Instead of instantiating and setting dependencies within the class, one can create a dependency from the outside 

This exterior dependency (e.g. wheel) can then be injected into the class car  

public Car(Wheel wh, Windows wi) {
  this.wheel = wh;
  this.windows = wi;
}

Advantages :
  1) Makes testing easier as dependency classes do not have to be instantiated 
  2) Allows for independent development of classes 
  3) Can easily switch out different dependencies to the class 
  4) Can use the dependencies in other classes 
```

##### Object Relational Mapping
```
Converting data between relational databases and OOP languages such that they become compatible with each other
Query database using an object-oriented paradigm (graph of objects) instead of SQL (tabular format) 
```

<br />

### Terms in Software Engineering 
---
##### Languages 
```
compiled language : compiler translates program to machine code before execution, time needed to compile every time changes are made, fast during runtime 
interpreted language : interpreter reads and executes program without compilation, dynamic typing, smaller memory size, slow during runtime   

strong typing : interpreter does not allow operations with incompatible types ("1" + 1)
weak typing : interpreter does allow operations with incompatible types ("1" + 1 = 2)

static typing : type is checked during compile time 
dynamic typing : type is checked during runtime
```

##### Programming Paradigms 
```
declarative programming : tell program specifically what you want it to do (this is A, this is B made from A, return C made from B)
imperative programming : tell program how you would like it to do something (this is A, we will go through this loop, we check this, if B, return A)

functional programming :
  1) paradigm where programs are composed of functions that are independent of outside states 
    ex) x = 5 
        func double():  
          x *= 2        // changes the state of an outside variable 
          return x  

        instead, 
        func double(num):
          return num * 2  // does not change outside variables 
  2) does not share states 
  3) paradigm where functions are treated as "first-class citizens" (can save as variable, pass as argument, be returned)
    ex) func double(func2):
          return func3   
```

##### Terms 
```
Web Server : return content of file following HTTP protocols (Apache HTTP)
Application Server : execute and display results of file following various protocols (Oracle WebLogic, Apache Tomcat)

Regression Testing : testing to confirm that recent program change does not impact existing functionality 
Automated Testing : reduce time, cost, and errors by automating certain test cases that are repetitive, tedious, or difficult to test manually
  ex) QTP, Rational Robot, Selenium
```

##### Methodology
```
Agile
  1) approach to break development into stages and constantly collaborate with end users
  2) advocates adaptive planning, evolutionary development, early delivery, and continual improvement 
 
DevOps 
  1) set of practices to reduce the time between committing a change, and the change being placed into production 
  2) integrate development and operations teams to improve collaboration by automating testing, infrastructure, workflows 
 
CI/CD
  1) continous integration, continous delivery 
  2) bridges gap between development and operations through automation to allow DevOps procedures
  
ETL 
  1) extract, transform, and load one database to another
```
