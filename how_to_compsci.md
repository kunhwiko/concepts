### Systems Concepts
---
##### Protocols 
```
protocol : set of rules and structures for how computers communicate 
1) IP : obtains the address of where packets come from and where they should be sent 
2) TCP : responsible for breaking data into packets and delivering/reassembling the packets
3) HTTP : set of rules for how request-response works in the web 
```

##### Systems
```
disk storage : permanent/ persistent storage with high latency (hard disk)
memory storage : temporary / transient storage with low latency (RAM)

latency : time between stimulation and response 
throughput : how much a machine or system can output 
bottleneck : constraint of a system (system is only as fast as the server with minimum throughput) 

availability : uptime in a given ammount of time 
SLA : an assurance for the uptime of a service 
redundancy : having an alternative when a failure happens 

caching : save certain data/results to retrieve faster 
1) CDNs cache website contents
2) Browsers cache HTML/JS/image files
3) DNS servers cache DNS records 

proxy : a server that acts as a middleman between a client and another server
1) forward proxy : acts on the behalf of the client, could mask the identity of client (VPNs) 
2) reverse proxy : acts on the behalf of the server (load balancer) 
```

##### Load Balancing
```
load balancer : balances and allocates request load to servers to maintain availability and throughput 

horizontal scaling : increase number of hardware
vertical scaling : increase performance of existing hardware 

Methods of load balancing 
1) round robin : start at the first item of a list of servers, sequentially look for available servers 
2) weighted round robin : ability to weigh different servers based on how powerful they are, and distribute work based on weight 
3) load based server selection : monitor the performance and load for each server and dynamically allocate based on calculations 
4) IP hash based selection : hash IP address to determine where to send request (useful for geographical servers or when servers cache requests)
5) service based selection : different servers handle different services 
```

##### Hashing
```
hashing : convert an input into a fixed size value 
collision : when two values are consistently hashed to the same value 

problems : 
1) if a server fails, hashing might still allocate requests to the failed server 
2) when new servers are added and hashing formula changes, previous keys will be remapped, making previous caches become useless

consistent hashing : 
uses a hash ring where servers can be distributed more than once throughout the ring 
after hashing a request, the request will go to the nearest server on the hash ring 
this does not solve but greatly reduces the problem of previous keys being remapped 
```

##### Databases 
```
SQL : relational, structured/predefined, table-based, less scalability
NoSQL : non-relational, unstructured/flexible, key-value paired (JSON objects), better scalability (database can be scattered into distributed systems)

indexing : short cutting to the record with matching values (query by age)
replication : makes copies of the database (backup purposes)

problems : 
1) replicating large amounts of data may be harmful for latency issues

sharding : breaks a database into smaller chunks, allowing for portions of the database to be replicated

consistency : read request for any of the copies should return the same data 
strong consistency : must become consistent immediately, offers updated data indefinitely at higher latency
eventual consistency : becomes consistent eventually, offers low latency but risks returning non-updated data 
```

##### Servers
```
Leader Election : situation in which you want to specify one server to be responsible for a request 

strong consistency : 
```


##### Concepts 
```
synchronous : statements in sequence
asynchronous : statements executing at different times 

kernel : core of operating system that controls tasks 
shell : interface to communicate with kernel 

compiled language : compiler translates program to machine code before execution, time needed to compile every time changes are made, fast during runtime 
interpreted language : interpreter reads and executes program without compilation, dynamic typing, smaller memory size, slow during runtime   
dynamic typing : type is checked during runtime
```

##### Methodology
```
Agile
  1) approach to break development into stages and constantly collaborate with end users
  2) advocates adaptive planning, evolutionary development, early delivery, and continual improvement 
  
CI / CD
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



### Servers and Testing
---
##### Servers
```
client <--> web server <--> database
client <--> web server <--> application server <--> database

Web server : return content of file following HTTP protocols (Apache HTTP)

Application server : execute and display results of file following various protocols (Oracle WebLogic, Apache Tomcat)

DNS server : phonebook of the Internet responsible for finding the IP address of sites

API
  1) interface that defines interactions between software such as types of calls/requests, how they are made, data formats, conventions 
  2) way of communicating between applications 
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

##### Environments
```
Development Environment --> User Acceptance Testing (UAT) Environment --> Production Environment 

UAT allows for beta testing before production, and to use actual production data to test results 
```

##### Regression Testing
```
testing to confirm that recent program change does not impact existing functionality 
```

##### Automated Testing
```
reduce time, cost, and errors by automating certain test cases that are repetitive, tedious, or difficult to test manually
ex) QTP, Rational Robot, Selenium
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
