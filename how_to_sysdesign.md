### Networking & Security 
---
##### Networking Fundamentals 
```
Protocol : set of rules and structures for how computers communicate 
  1) IP : address of where packets come from and where they should be sent  
    * IPv4, IPv6
  2) TCP : responsible for breaking data into packets, delivering/reassembling the packets, checks for corruption
  3) DNS Server : phonebook for finding the IP address of sites 
  4) HTTP : set of rules for how request-response works in the web 

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

###### Security Fundamentals 
```
Man-in-the-Middle(MITM) Attack : malicious activity to intercept or alter IP packets in an HTTP connection

Symmetric Encryption
  1) uses a single key to encrypt/decrypt data and is faster 
  2) makes use of Advanced Encryption Standard (AES) algorithm 

Asymmetric Encryption
  1) uses a public and private key and is slower  
  2) anyone can encrypt with public key, only private key can decrypt messages  
  
HTTPS : HTTP over TLS 
Transport Layer Security (TLS) : security protocol for secure communication 

TLS Handshake : process to establish a secure connection between clients and server 
  Step 1) 
    1) Client sends "client hello"
    2) Server responds with "server hello" + SSL certificate containing the public key
  Step 2) 
    1) Client verifies SSL certificate and sends a "premaster secret" encrypted with the public key
    2) Server decrypts the premaster secret using the private key
  Step 3) 
    1) client hello, server hello, premaster secret are used to create temporary symmetric keys for the session
    2) Symmetric keys are used to figure out whether a connection was established or has failed 

Purpose of SSL Certificates 
  1) MITM attacks can intercept server hellos and public keys, send their own public key, and establish a connection with client 
  2) SSL certificates guarantee where public keys come from
```

### Systems Design
---
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
SQL : relational, structured/predefined, table-based, less scalability, better for ranged queries, strong consistency  
NoSQL : non-relational, unstructured/flexible, key-value paired (JSON objects), better scalability, eventual consistency  

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

Storage Types :
  1) Key-Value Storage : Specializes in storing as a key/value pair (MongoDB)
  2) Blob Store : Specializes in storing massive amounts of unstructured data (Google Cloud Storage, S3, Azure)
  3) Time Series Database : Specializes in time series data / monitoring (InfluxDB)
  4) Graph Database : Stores in a graph form rather than a tabular form, specializes in relations between data (Neo4j)
  5) Spatial Database : Stores data that represents some space (Quad-tree) 
```


##### Database Scalability
```
read problems : as tables grow, it becomes harder to read information that reader needs 
Indexing : uses additional memory to maintain a lookup for faster querying (imagine glossary page) 
  1) Tree Indexing : Allows us to do fast range queries 
  2) Hash Indexing : Allows us to do fast exact queries 

load problems : what if database has too much requests or failures result in inaccessible databases?
Replication : makes copies of the database for backup purposes
Master-Slave Model : slaves are replicas that are read-only to lessen the load on the master server

write problems : 
  1) what if there are too many write requests to master server (replicas are read-only)? 
  2) what if the database has tons of data, is it necessary to replicate all this data?
  3) after writing to the master, how can we solve latency issues of replicating all the data to the slaves?
Sharding : splitting the data across multiple machines 
  1) Vertical Sharding : partitioning master server by feature (profiles, messages, customer support) --> one table might become large
  2) Hash Based Sharding : partitioning through hashing some value (ID) --> same problems with hashing
  3) Directory Based Sharding : a lookup table maintains where data can be found --> lookup table can fail or overload 

Normally good to have a reverse proxy (load balancer) to process client requests and match to databases/shards 
```

##### Leader Election
```
What if we introduce redundancy to servers but we would like to ensure that only one server does a particular request (payments)? 

Leader Election : specify one server to be responsible for a request (replicas will take over if the leader fails)
Consensus Algorithms : complicated algorithms that help select a leader among servers (Zookeeper) 
```

##### Peer-to-Peer Networks
```
What if a server is sending data to thousands of machines? By the time the server is done sending, it might have to send again..or even be late!
If we have multiple servers, then we need to replicate the same data to all the servers..
If we shard the data, then we end up with a server sending data to thousands of machines again...

Peer-to-Peer Networks : 
  1) data to send is broken into many pieces, and instead of sending 1000GB to "1" machine at a time, we send 1GB to "1000" machines
  2) peers communicate and send missing data to other peers AS the server is sending over data
Gossip Protocol : protocol for peers to communicate to each other and spread information
Distributed Hash Table (DHT) : hash table that holds information on what peers hold what data 
```

##### Polling & Streaming/Configurations/Logging & Montoring
```
Polling : sending a request for updated data (packets) in regular intervals (cycle of requests/responses)
Streaming : client opens a channel using "sockets" for servers to send data (client listens on server)


Configurations : "settings" for codebases written normally in JSON or YAML
Static Configuration : Configurations are packaged with codebase, must deploy entire code to test, can create test cases during deployment
Dynamic Configuration : Configurations outside the codebase, takke immediate change, but harder to test 


Usages for data
  1) Logging : the collection of data to use for recording/finding errors/analytics
  2) Monitoring : having transparent ways to analyze data for insights
  3) Alerting : alert of significant changes in data 
```

##### Rate Limiting
```
How do we prevent malicious requests (DoS attacks?)

Rate Limiting : limit number of operations, possibly in different tiers 
  1) Criteria : Request frequency, User/IP Address, Region, Server (10,000 requests at a time)
  2) Typically use another database/cache to check for rate limiting (Redis) because in-memory cache at the server level becomes 
     less useful when clients get redirected to another server
```     
     
##### Pub/Sub Model
```
What if a client was streaming for data and the server goes down. What if really important data doesn't get sent due to interferance?

Pub/Sub Model : publisher sends messages (info) to a topic (channel), and subscribers to the topic can consume the information
  1) servers becomes independent of communicating with clients
  2) clients are guaranteed "at least once delivery" of messages, if a topic loses connection to a subscriber, it will attempt to resend messages 
  3) messages are sent using a queue (guarantees ordering)
  4) subscribers are able to replay messages or filter their subscriptions 
  
Idempotency : a characteristic where the outcome is always the same no matter how many times an operation is performed
  1) Did the subscriber receive the message? (Idempotent)   Increment Youtube View Counter (Non-idempotent)
```

##### MapReduce 
```
MapReduce : Framework for processing large datasets that are split up in distributed systems 

Map : function that takes the data that is spread out and transforms it to key-value pairs
Shuffle : takes key-value pairs and routes them to relevant machines 
Reduce : functions that takes shuffled pairs and transforms them into meaningful data 

Assumptions : 
  1) Distributed file system where large data is spread out across different machines 
  2) Central system that knows where the data resides, how to communicate to make map/reduce work, and knows where the output will reside
  3) MapReduce will run on different databases locally, and will not move the data to somewhere else 
  4) MapReduce is idempotent and fault-tolerant, if a server breaks, MapReduce will re-run on the failed chunk 
  
What do we have to know?
  1) We have to specify what our map function will be 
  2) Understand what sort of data is being mapped
  3) Understand what our key-value pairs will look like
  4) We have to specify what our reduce function will be 
  5) Understand what our output will look like
  
Example cases:
  1) Get the count of views on YouTube per channel 
  2) Count number of payments made per month using logs  
  
Distributed File System : cluster of machines that allow them to work as one large file system (Google File System, Hadoop)
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
##### What is an API
```
1) defines interactions between software such as types of calls/requests, how they are made, data formats, conventions 
2) way of communicating between applications 
3) allows for servers/systems to communicate with one another such that automation is possible 
4) allows information to be shared as a service 
```

##### API Design
```
CRUD operations : Create, Read, Update, Delete

Examples of Entity Definitions 
1) Payment
  - id : uuid 
  - customer_id : uuid 
  - restaurant_id : uuid 
  - amount : int
  - status : enum ["success", "pending", "failed"]

2) Restaurant 
  - id : uuid 
  - name : string 
  - address : string 
  - account : Account 
 
Example of Payment.json and Payment object
{id : "abac1123-bfsdg", customer_id : "bdfsx-123cvxc", restaurant-id : "bac1123-bfsdg", amount : 2000, status : pending}

Example of Restaurant.json and Restaurant object 
{id : "bac1123-bfsddg", name : "Papa Johns", address : "4005 Chestnut Street", account : {Bank : ___, Account No. ___}}


Example of Endpoint Definitions 
1) Payment
  - Payment createPayment(payment: Payment)
    path : POST /v1/payments
  - Payment getPayment(id: uuid)
    path : GET /v1/payments/id 
  - Payment updatePayment(id: uuid, updatedPayment: Payment) 
    path : UPDATE /v1/payments/id
  - Payment[] listPayments(offset: int, limit: int) --> Pagination 
    path : GET /v1/payments

2) Restaurant
  - Restaurant createRestaurant(restaurant: Restaurant)
  - Restaurant getRestaurant(id: uuid)
  - Restaurant deleteRestaurant(id: uuid)
  
Pagination : limit the response of a potentially larger response, usually when retrieving huge lists  
```

##### REST Principles
```
REST principles 
  1) verbs : GET (read), POST (create), PUT/PATCH (update entire/partial), DELETE (delete)
  
REST parameters
broken into endpoint(site) + query parameter(condition)  
ex) api.com/cars?type=SUV&year=2019

REST practices
1) Use "nouns" and not "verbs"
2) Use "plural" for list of items 
3) Use "camel case"
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
