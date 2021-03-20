### Databases 
---
##### Getting Started with Databases 
```
Database Management Systems (DBMS) 
  1) a software package designed to manage databases
  2) allows for reliability, recovery, data access, querying, updates 

Database Benefits vs Programs/Files
  1) users do not have to be concerned with low level details such as indexes, concurrency, disk speed, data structures 
  2) scalable and efficient for large datasets by providing random access (for files, this requires scanning the entire document)
  3) provides permanence, consistency even for concurrent actions 
```

##### Relational Databases / SQL 
```
Relational Database : data stored in tables(relations) organized in a strict, predefined way 
SQL : relational, structured/predefined, table-based, less scalability, better for ranged queries, strong consistency  

ACID principles for SQL 
  1) Atomicity : guarantee that when one operation fails(succeeds), all other following operations fails(succeeds) 
  2) Consistency : each transaction ensures that the database moves from one valid state to another valid state (does not corrupt data)
  3) Isolation : when you run operations concurrently, the result will be as if you ran the operations in sequence
  4) Durability : once the data is stored in the database, it will remain to do so

Atomicity 
  a) read-copy-update : keep a copy of the original database before some query execution 
  b) journaling : logs the updates, and reverses operations if failure arises 
```

##### Isolation Levels of DBMS 
```
Serialization 
  1) Serial Execution vs Interleave Execution 
      - Serial Execution: queries executed in sequential order 
      - Interleave Execution: queries executed concurrently (can cause conflicts)
  2) Serializable: good executions are those that are equal to some serial execution
  3) Modern databases implement locks to prevent bad access to data and ensure serialization 
  4) Deadlocks can happen, in which DBMS must detect and abort one of the conflicting transactions 

Undesired Phenomena 
  1) Overwrite Uncommitted Data: T1 updates, T2 updates-->commits, T1 commits
  2) Dirty Read: 
       - reading from data that has not yet been committed in another transaction 
       - if the uncommitted data fails to commit, we have read the wrong information 
  3) Unrepeatable Read: reads the same item twice and gets different values 
  4) Phantom Read: 
       - retrieve collections of rows and see different results, although these collections are not modified
       - this might happen if some rows are inserted/deleted

Isolation Levels
  1) Serializable: ensures total isolation 
  2) Repeatedable Read: allows phantom reads 
  3) Read Committed: allows phantom reads + unrepeatedable reads 
  4) Read Uncommitted: allows phantom reads + unrepeatedable reads + dirty reads (read-only application)

Snapshot Isolation 
  1) copies and reads the last committed version (snapshot) of the database before the operation starts 
  2) will only commit if the operation does not conflict with concurrent commits  
  3) low latency, but guarantees less consistency and uses more memory resources   
```

##### Database Costs 
```
Page: unit of storage in main memory ("block" in disk storage)

e.g.)
Student(id int, name char(40), major char(4), hobby char(30))
  - size of row: 4+40+4+30 = 78(bytes)
  - page size: 8,000 
  - number of rows: 10,000
floor(8000/78) = 102 (student rows per page)
ceiling(10000/102) = 98 (pages to store all student records)

File Storages 
  1) Heap: unordered data
       - find: O(n), insert: O(1), delete: O(n)
  2) Sorted
       - find: O(logn), insert: O(1), delete: O(n)
       - great for ranged queries 
  3) Hashed
       - find: O(1), insert: O(1), delete: O(1)
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


##### Non-relational Database / NoSQL 
```
Non-relational Database : data stored in a flexible form
NoSQL : non-relational, unstructured/flexible, key-value paired (JSON objects), better scalability, eventual consistency 

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
  1) Key-Value Store : Specializes in storing as a key/value pair (Apache HBase)
    ex) id -> [name age experience]
  2) Wide Column Store : Organizes related facts into "column families", 2-D Key-Value (BigTable, Cassandra)
    ex) id -> [personal]  [professional] 
               - name      - experience
               - age 
  3) Document Oriented Store : organized as documents, usually JSON format (MongoDB)
    ex) {id : 1, name : __, experience : __}
  4) Blob Store : Specializes in storing massive amounts of unstructured data (S3)
  5) Time Series Store : Specializes in time series data / monitoring (InfluxDB)
  6) Graph Store : Stores in a graph form rather than a tabular form, specializes in relations between data (Neo4j)
```
