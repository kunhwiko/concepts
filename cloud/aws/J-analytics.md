### Athena
---
##### Athena Basics
```
a) Athena is a serverless query service that allows users to query data in S3 using SQL. Billing is based on amount of 
   data scans performed, so it is recommended to use columnar formats (e.g. Apache Parquet), data compression, and 
   dataset partitions to increase Athena's performance and lower costs.
b) Data source connectors that run on Lambda can be used to query data from other AWS services such as DynamoDB, 
   ElastiCache, CloudWatch logs, and AuroraDB. 
```

##### Glue
```
a) Glue is a serverless ETL service that is used to transform data for analytical purposes. Glue can be used to
   transform data from various sources (e.g. S3) into a parquet format that can be queried by Athena.
b) Glue Data Catalog uses data crawlers on connected sources (e.g. S3, RDS, DynamoDB) to collect metadata on datasets.
   The catalog can be used by Glue jobs, Athena, Redshift Spectrum, and EMR.
c) Glue Job Bookmarks help AWS maintain state of ETL jobs and prevent the reprocessing of old data.
```

### Redshift
---
##### Redshift Basics
```
a) Redshift is a managed data warehouse service based on PostgreSQL that is optimized for analytical processing. Data
   is stored as columnar storage (not row-based) and has an SQL interface for querying. Redshift is great for join
   and aggregation heavy queries due to indexing.
b) Redshift can run in provisioned or serverless mode. Leader nodes are used to coordinate queries and compute nodes
   are used to perform queries in parallel and send results to the leader nodes.
c) Redshift supports snapshots and can be configured to automatically be copied into other AWS regions.
```

##### Redshift Spectrum
```
Spectrum allows users to query data in S3 without loading it. When a query is started by a Redshift cluster, it is 
submitted to the Spectrum layer that employs massive number of nodes for parallelism.  
```

### EMR
---
##### EMR Basics
```
Elasic MapReduce makes it simple to run distributed processing frameworks lik Hadoop or Spark. Master nodes manage the
cluster and coordinate jobs, core nodes store data and run tasks, and task nodes are optional nodes that solely run tasks. 
```

### OpenSearch
---
##### OpenSearch Basics
```
OpenSearch is a distributed search and analytics engine that is used to search, analyze, and visualize data in real-time.
The service can run in provisioned or serverless mode.
```

##### OpenSearch Integrations
```
a) Lambda functions can capture DynamoDB Streams and write to OpenSearch. OpenSearch can then be used for search purposes
   (i.e. since DynamoDB can only query on primary keys and indexes) to know what items to retrieve from DynamoDB.
b) CloudWatch logs or Kinesis Data Streams can forward data to OpenSearch via Lambda or Kinesis Data Firehose.
```