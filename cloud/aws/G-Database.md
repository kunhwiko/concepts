### RDS
---
##### Relational Database Service (RDS)
```
Managed DB service using SQL as a query language. RDS supports databases such as PostgreSQL, MySQL, Aurora etc. As a
managed service, AWS handles provisioning, autoscaling, OS upgrades, and disaster recovery. Users can also create and
migrate snapshots or do point in time restoration.
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
