### RDS
---
##### Relational Database Service (RDS)
```
Managed DB service using SQL as a query language. RDS supports databases such as PostgreSQL, MySQL, Aurora etc. AWS 
handles provisioning, autoscaling, monitoring, OS upgrades, automated backups, and disaster recovery. Users can also 
create and migrate snapshots or do point in time restoration.
```

##### RDS Read Replicas
```
a) RDS supports read-only replicas within AZ, cross AZ, and cross region. Replications within the same region do not
   incur data transfer costs, but do occur for cross region replications.
b) RDS replications to read replicas are asynchronous and eventually consistent.
c) Read replicas can be promoted to a standalone instance with its own lifecycle, which typically means the replica can
   now be written to. Note that this means the replica also leaves the previous DB cluster.
```

##### RDS Standby Instances
```
a) RDS supports standby instances in multi-AZ for disaster recovery. When a new standby instance is created, a snapshot
   of the master database is taken to create the new instance. 
b) Replications to standby instances are synchronous. 
c) Standby instances are used for failover purposes and cannot be used as read replicas.
```

##### RDS Encryption
```
a) For at-rest encryption, data in RDS can only be encrypted during DB creation time using KMS. If the master is not 
   encrypted, read replicas will not be encrypted. To encrypt an unencrypted master, a new database must be created 
   from a DB snapshot.
b) For in-flight encryption, TLS is ready to use by default and clients need to simply use AWS root certificates.
c) IAM roles can be used to connect to the database and retrieve data.
d) RDS supports security groups to control inbound and outbound traffic to the database.
```

##### RDS Proxy
```
a) Proxy for RDS can be setup to handle connection pooling, which means multiple applications can share established DB 
   connections instead of having multiple individual connections. This reduces load and the number of open connections 
   on the RDS instance.
b) Proxies are serverless, can autoscale, and can handle failovers by reconnecting to a new DB instance.
c) RDS proxies enforces IAM authentication to the DB and can only be accessed from within a VPC. 
```

### AuroraDB
---
##### AuroraDB Basics
```
a) AuroraDB is a high-performant database optimized for AWS that can automatically scale with use and also grants both 
   compatability and performance improvements over MySQL and PostgreSQL.
b) AuroraDB maintains multiple copies of data across availiability zones and has self healing capabilities with 
   peer-to-peer replication for high availability. In a zone, data is sharded across multiple volumes.   
c) AuroraDB read instances share the same underlying storage, avoiding the need to perform writes to replicas and 
   vastly reduces replication latency. Any of the replicas can also become the master instance for immediate failover.  
d) All write requests are sent to the master instance, and connections to read requests are load balanced across
   available read replicas. Users can also create custom endpoints to route read requests to specific read replicas
   (e.g. run heavy analytic workloads on larger replica instances).
```

##### AuroraDB Serverless
```
AuroraDB supports serverless for infrequent and unpredictable workloads. Clients interact with AuroraDB through a proxy 
fleet that spins up DB instances on demand.
```

##### AuroraDB Global
```
AuroraDB can be setup to have a primary region and multiple secondary read-only regions. Replication across regions 
and promoting another region to primary during failover is highly optimized for low latency. 
```

### DynamoDB
---
##### DynamoDB Basics
```
a) DynamoDB is a highly available and performant NoSQL database made up of tables. Each table is made up of items (rows), 
   which are made up of various attributes and queries can only be executed on primary keys or indexes. TTLs can be set 
   to delete items after a certain time. 
b) DynamoDB supports a provisioned mode where users specify expected read/write capacity units, or an on-demand mode 
   where DynamoDB automatically scales based on traffic.
c) DynamoDB Accelerator (DAX) can be used as a caching layer to reduce read congestion on DynamoDB by caching frequently 
   accessed data.
d) Changes to DynamoDB tables can be stream processed to DynamoDB Streams or Kinesis Data Streams for further processing.
   When DynamoDB Stream is enabled, tables can be replicated across regions. 
```

### ElastiCache
---
##### ElastiCache Basics
```
Managed cache service that supports Redis or Memcached. AWS handles provisioning, monitoring, OS upgrades, automated
backups, and disaster recovery.
```

##### ElastiCache Patterns
```
a) ElastiCache supports lazy loading (i.e. loads data into cache only after there was a cache miss) and write-through 
   (i.e. update cache when data is written to the database). 
b) Using session affinity with ELBs can store session data but the data is lost if the bound EC2 instance goes down.
   User cookies can also be used but they can only hold limited data and can make HTTP requests heavier. Instead, user
   cookies can send small session IDs where servers can then look up corresponding session data in ElastiCache.
```
