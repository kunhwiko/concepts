### Systems Concepts
---
##### Protocols 
```
Protocol : set of rules and structures for how computers communicate 

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

horizontal scaling : increase number of hardware
vertical scaling : increase performance of existing hardware 

availability : uptime in a given ammount of time 
SLA : an assurance for the uptime of a service 
redundancy : having an alternative when a failure happens 


load balancing : process of distributing tasks over resources efficiently 


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

##### Databases 
```
SQL : relational, structured/predefined, table-based, less scalability
NoSQL : non-relational, unstructured/flexible, key-value paired (JSON objects), better scalability (database can be scattered into distributed systems)
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
