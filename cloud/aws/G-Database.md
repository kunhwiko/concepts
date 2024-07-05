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
c) Read replicas can be set up as multi-AZ and can be promoted to its own master database for disaster recovery purposes.
```

##### RDS Standby Instances
```
a) RDS supports standby instances in multi-AZ for disaster recovery. When a new standby instance is created, a snapshot
   of the master database is taken to create the new instance. 
b) Replications to these standby instances are synchronous. 
c) Standby instances are not used as read replicas and therefore are not used for scalability purposes.
```
