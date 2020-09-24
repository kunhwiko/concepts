### Systems Design
---
##### Networking
```
Protocol : set of rules and structures for how computers communicate 
  1) IP : address of where packets come from and where they should be sent  
    * IPv4, IPv6
  2) TCP : responsible for breaking data into packets, delivering/reassembling the packets, checks for corruption
  3) DNS Server : phonebook for finding the IP address of sites 
  4) HTTP : set of rules for how request-response works in the web 
  5) HTTPS : encrypted HTTP 


Networking : 
  Packets : small segments of data of a larger message 
    * IP Packet Header : holds the source and destination address
    * TCP Packet Header : order of how packets should be reassembled
    * IP Packet Data : holds the data of the packet 
  
  Handshake : TCP sends requests by sending packets to destination server asking for a connection
    
    
Ports : 
  1) docking point for where information is received or sent
  2) how multiple programs listen for new network connections on the same machine without collision  
  3) IP address is like a mailbox to an apartment complex, and ports are the specific apt number


Client-Server Model 
Step 1)
  Client makes a DNS query, retrieves the IP address of some domain, and contacts the server  
  IP addresses can be granted by cloud providers where you can run your server  
  
Step 2)
  Sends HTTP request to Server along with source address (address of sender)
  
Step 3)
  Server listens to requests on ports, and sends HTML/CSS/JS files to source address 
```

##### Fundamentals
```
Storage
  1) Disk Storage : permanent / persistent storage with high latency (hard disk)
  2) Memory Storage : temporary / transient storage with low latency (RAM)
 
 
Capacity
  1) Latency : time between stimulation and response 
  2) Throughput : actual output of a system or machine in a given time / how many requests are handled  
  3) Bandwidth : theoretical output of a system or machine in a given time 
  4) Bottleneck : constraint of a system 


Availability : uptime in a given amount of time (usually per year)
  1) SLA (Service Level Agreement) : an assurance for the uptime of a service 
  2) Redundancy : having alternatives when a failure happens 
    * Passive Redundancy : Uses excess capacity to reduce failures 
    * Active Redundancy : Monitors and reconfigures capacity in downtimes 


Proxy : a server that acts as a middleman between a client and another server
  1) Forward Proxy : acts on the behalf of the client, could mask the identity of client (VPNs) 
  2) Reverse Proxy : acts on the behalf of the server (load balancer, cache, filter, logging) 
```

##### Caching
```
Caching : save certain data/files/results in a caching layer to retrieve them faster 


Read Examples
  1) Servers cache results retrieved from databases 
  2) Clients cache computationally heavy operation 
  3) Browsers cache HTML / JS / image files 
  4) DNS servers cache DNS records
  5) CDNs cache website contents and static files such as images or HTML files


Content Delivery Network (CDN)
  1) Pull : when a file is accessed for the first time, load it to the CDN, and will be cached thereafter (initially slow)
  2) Push : put files into the CDN to be cached (initially fast but some files may never be used)
  
  
Write Examples
  1) Write Through Cache : posts are saved to both cache and database at the same time 
  2) Write Back Cache : posts are saved to cache, and asynchronously updates database 
  
  
Problems of Caching :
  Caching becomes difficult for mutable data as this might result in displaying stale (outdated) data to certain users
  

Caching Eviction :
  Rules to evict cached data (LRU Cache, LFU Cache, FIFO Cache)
```


##### Load Balancing
```
Load Balancer : balances and allocates requests to DNS/servers/databases to maintain availability and throughput 


Horizontal Scaling : increase number of hardware
Vertical Scaling : increase performance of existing hardware 


Rules for load balancing 
1) Round Robin : start at the first item of a list of servers, sequentially look for available servers 
2) Weighted Round Robin : ability to weigh different servers based on how powerful they are, and distribute work based on weight 
3) Load Based Server Selection : monitor the performance and load for each server and dynamically allocate based on calculations 
4) IP Hash Based Selection : hash IP address to determine where to send request (useful for geographical servers and when servers cache requests)
5) Service Based Selection : different servers handle different services 


Tools to imitate systems:
  1) NGINX : Reverse Proxy / Load Balancer
  2) Express.js : Server, Cache
  3) Redis : Cache
```

##### Hashing
```
Hashing : convert an input into a fixed size value 
Collision : when two values are consistently hashed to the same value 
SHA (Secure Hash Algorithms) : cryptographic hash functions


problems : 
  1) if a server fails, hashing might still allocate requests to the failed server 
  2) when new servers are added and hashing formula changes, previous keys will be remapped, making previous caches become useless


Consistent Hashing : 
uses a hash ring where servers are distributed via a hash function (can be distributed more than once!)
after hashing a request, the request will move clockwise/counter-clockwise to the nearest server 
this does not solve but greatly reduces the problem of previous keys being remapped 


Rendezvous Hashing (Highest Ranking Hashing) : 
clients rank servers by ranking and get distributed to the highest ranking server 
when a server gets deleted, clients who lost their highest ranking server will redistribute to the next highest server
also solves the problem of previous keys being remapped 
```

##### Database Types
```
Why not write scripts to query data?
You must potentially load all the data into memory when running on Python/Java. 


Relational Database : Data stored in table(relations) form and organized in a strict, predefined way (usually supports SQL)  
Non-relational Database : Flexible (non-tabular) form not precisely organized in a predefined way 


ACID principles for SQL 
  1) Atomicity : guarantee that when one operation fails(succeeds), all other following operations fails(succeeds) 
  2) Consistency : each transaction ensures that the database moves from one valid state to another valid state (does not corrupt data)
  3) Isolation : when you run operations concurrently, the result will be as if you ran the operations in sequence
  4) Durability : once the data is stored in the database, it will remain to do so


BASE principles for NoSQL
  1) Basically Available : system guarantees availability
  2) Soft State : state of system and replicas might change over time even without input 
  3) Eventual Consistency 


Consistency : read request for any of the copies should return the same data 
  1) Locks : a method to keep data consistent by allowing only certain users to update the database at a time  
  2) Strong Consistency : must become consistent immediately, offers updated data indefinitely at higher latency
  3) Eventual Consistency : becomes consistent eventually, offers low latency but risks returning non-updated data
  

Considerations : 
  1) Do we want strong vs eventual consistency?
  2) Do we want in-memory (caching) vs disk storage?

```


##### Databases 
```
read problems : as tables grow, it becomes harder to read information that reader needs 
Indexing : allow for short cuts to data by specifying matching values (query by date, age, id)

failure problems : what if database fails (too much load) and you cannot access the database?
Replication : makes copies of the database for backup purposes
Master-Slave Model : slaves are replicas that are read-only to lessen the load on the master server

write problems : 
  1) what if there are too many write requests to master server (replicas are read-only)? 
  2) after writing to the master, how can we solve latency issues of replicating all the data to the slaves?
Sharding : splitting the data across multiple machines 
  1) Vertical Sharding : partitioning master server by feature (profiles, messages, customer support) --> one table might become large
  2) Hash Based Sharding : partitioning through hashing some value (ID) --> same problems with hashing
  3) Directory Based Sharding : a lookup table maintains where data can be found --> lookup table can fail or overload 

Types
SQL : relational, structured/predefined, table-based, less scalability, better for ranged queries, strong consistency  
NoSQL : non-relational, unstructured/flexible, key-value paired (JSON objects), better scalability, eventual consistency  
```

##### Other Concepts
```
Pub/Sub Model : publisher sends info to a topic, and subscribers to the topic can consume the information
  1) saves time for the publisher to have to maintain a roster queue or independently sending messages to subscriber
  2) solves the problem of publisher failing to send messages to everyone due to an interference 

Leader Election : situation in which you want to specify one server to be responsible for a request 

Polling : sending a request for updated data (packets) in regular intervals (cycle of requests/responses)
Streaming : sending a request that opens a channel using "sockets" (single request/response)

Endpoint Protection : protect system from too many operations (strong against DOS)
  1) Rate Limiting : limit number of operations (weak against DDOS)

Usages for data
  1) Logging : the collection of data to use for analytics
  2) Monitoring : analyze data for insights
  3) Alerting : alert of significant changes in data 
```

##### Microservices 
```
architectural design that breaks a monolithic application into smaller pieces, which communicate through HTTP/API

pros 
  1) independent services with ability to use different programming languages
  2) easier to understand codebase and modify 
  3) when failure arises, only particular service goes down
  4) easier to scale specific process / service
  5) less commitment to a tech stack and able to change to new tech faster 

cons 
  1) testing can be complicated and managing the whole product becomes more difficult 
  2) information barriers between different services
  3) must implement means of communicating between services
  4) large upfront investment in automation as manual deployment becomes more difficult 
```

### API Design
---
##### What is API
```
1) interface that defines interactions between software such as types of calls/requests, how they are made, data formats, conventions 
2) way of communicating between applications 
3) allows for servers/systems to communicate such that automation is possible 
4) allows information to be shared as a service 
```

##### REST Principles
```
REST principles 
  1) verbs : GET (read), POST (create), PUT/PATCH (update entire/partial), DELETE (delete)
  
REST parameters
broken into endpoint(site) + path(topic) + query parameter(condition)  
ex) api.com/cars?type=SUV&year=2019
```

##### REST Practices
```
Use "nouns" and not "verbs"
Use "plural" for list of items 
Use "camel case"

GET || List of items || individual item

POST || Create item || Error

PUT || Replace items || Replace certain item

DELETE || Delete items || Delete certain item 
```

##### REST Mappings
```
One-To-Many Mapping
api.com/tickets/145/messages/4 --> find 4th message for 145th ticket 
a single ticket has N unique messages associated with that ticket

Many-To-Many Mapping
api.com/groups/200/users/56 --> find user of id 56 in 200th group
a user might also be in different groups    
```

##### Status Codes
```
1xx : Request received and understood 
2xx : Request by client was received, understood, and accepted 
  1) 201 Resource Created (for POST methods)
  2) 202 Accepted 
  3) 204 No content (for DELETE methods)
3xx : Client must take additional actions 
4xx : Client screwed up (for wrong GET,DELETE requests)
5xx : Server screwed up
  1) 500 Internal Server Error
  2) 504 Gateway Timeout
```


### Computer Science Concepts
---
##### OS 
```
synchronous : statements in sequence
asynchronous : statements executing at different times 

process : program that is being executed (heavy, isolated memory, takes time to switch)
threads : segments of a process (light, shared memory, fast switch times)  
locks : prevents race conditions between threads 
deadlocks : processes are blocked because each process holds the resource needed for the other to go further

kernel : core of operating system that controls tasks 
shell : interface to communicate with kernel 
```

##### Compilers
```
compiled language : compiler translates program to machine code before execution, time needed to compile every time changes are made, fast during runtime 
interpreted language : interpreter reads and executes program without compilation, dynamic typing, smaller memory size, slow during runtime   
dynamic typing : type is checked during runtime
```

##### Servers 
```
Web Server : return content of file following HTTP protocols (Apache HTTP)
Application Server : execute and display results of file following various protocols (Oracle WebLogic, Apache Tomcat)
```

##### Testing
```
1) Regression Testing : testing to confirm that recent program change does not impact existing functionality 
2) Automated Testing : reduce time, cost, and errors by automating certain test cases that are repetitive, tedious, or difficult to test manually
ex) QTP, Rational Robot, Selenium
```

##### Methodology
```
Agile
  1) approach to break development into stages and constantly collaborate with end users
  2) advocates adaptive planning, evolutionary development, early delivery, and continual improvement 
  
CI/CD
  1) continous integration, continous delivery 
  2) bridges gap between development and operations through automation to allow DevOps procedures
  
ETL 
  1) extract, transform, and load one database to another
```



### Object Oriented Programming
---
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
Shape shape = new Triangle()           class triangle is a part of class shape (polymorphism)
Shape shape = new Circle()             class circle is a part of class shape 
Triangle shape = new Shape()           a triangle is a shape, but a shape is not a triangle (invalid)

public void run(Shape shape) {}        classes shape, triangle, and circle and all be passed to the parameter

Advantages :
  1) Flexibility 
```

##### Dependency Injection 
```
Suppose you have a class car and you must also instantiate classes wheel, company, material

class Car {
  Company company = new Company("Audi");
}

Wheel, company, and material are 'dependencies' of car because car uses these classes

What if we want to create multiple objects of car, but with different wheels, companies, and materials?

Instead of instantiating and setting dependencies within the class, one can create a dependency from the outside 

This exterior dependency (e.g. wheel) can then be injected into the class car  

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



### Language Specifics
---
##### HTML/CSS/XML/XSLT
```
HTML
  1) static language used to display data
  2) hypertext language : defines links between web pages 
  3) markup language : defines text within tags that defines structures of web pages 
  
CSS : style sheet for HTML 

XML
  1) dynamic language used to transfer data 
  2) markup language
  3) able to define new tags 

XSL : style sheet for XML

XSLT : transforms XML to other documents like HTML
  
XPath : navigate through an XML tree  
  
XQuery : queries through XML using XPath
```
