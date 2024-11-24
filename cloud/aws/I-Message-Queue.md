### SQS
---
##### Simple Queue Service (SQS)
```
Managed queue service where producers send messages to a queue and consumers poll and delete messages from the queue.
```

##### Standard Queue
```
a) Standard queue has unlimited throughput (i.e. send as many messages as required), supports unlimited number of 
   messages in the queue, and low latency on sending/receiving.
b) Each message has a short retention period (i.e. max 14 days) and a limit of 256KB. After a message is polled by a 
   consumer, it is hidden from other consumers for a short period and must be processed within that time frame. A consumer 
   can use the ChangeMessageVisibility API to extend the visibility timeout for a message.
c) Standard queue supports best-effort ordering (i.e. messages are not guaranteed to be ordered) and at-least-once 
   delivery model (i.e. duplicate messages can exist). Applications should take this into account.
d) Queue length from CloudWatch metrics can be used with ASGs to scale instances based on the number of messages.
e) SQS supports long polling, which allows consumers to wait for messages to arrive for a certain period if the queue is 
   empty. This decreases the number of API calls made to SQS and improves latency.
```

##### FIFO Queue
```
At the cost of limited throughput (i.e. restricted messages per second), FIFO queues guarantee messages are received in 
the order they were sent and support an exactly-once send model through message deduplication in the event of retries.
```

##### SQS Security
```
a) Supports in-flight encryption through HTTPS endpoints and at-rest encryption through AWS managed SQS keys and KMS keys.
   Client-side encryption is also supported if the client wants to perform encryption themselves.
b) IAM policies can be used to regulate access to SQS APIs (e.g. ProcessMessage, DeleteMessage).
c) SQS access policies, much like S3 bucket policies, can be used to regulate access to queues. These policies can be
   used to allow for cross-account access or allowing other services to access the queue.
```

### SNS
---
##### Simple Notification Service (SNS)
```
Mananged pub/sub service where an event producer sends messages to an SNS topic and multiple event receivers can listen
and filter for notifications. Like SQS, SNS supports FIFO topics for ordered delivery.
```

##### SNS Security
```
a) Supports in-flight encryption through HTTPS endpoints and at-rest encryption through AWS managed SNS keys and KMS keys.
   Client-side encryption is also supported if the client wants to perform encryption themselves.
b) IAM policies can be used to regulate access to SNS APIs.
c) SNS access policies, much like S3 bucket policies, can be used to regulate access to SNS topics. These policies can be
   used to allow for cross-account access or allowing other services to access the queue.
```

##### SNS + SQS Integration
```
SNS can be used to fan out messages to multiple SQS queues. Each SQS queue can subscribe to relevant messages which are
then consumed by respective consumers.
```

### Kinesis
---
##### Kinesis Data Streams
```
a) Data streams help capture, process, and store streams of data in real-time. A stream is divided into shards, which 
   are the base throughput unit. The number of shards can be provisioned manually or can be adjusted on-demand.
b) Producers send data to streams in the form of records, which holds a partition key and a data blob. The partition key
   is used to determine which shard the record belongs to through a hash. Data is ordered at the shard level.
c) Consumers read data from shards and process records. Each record is uniquely identified by a partition key, sequence 
   number and a data blob. Multiple consumers can consume the same stream.
d) Data streams have a retention period and data can be reprocessed within that time frame. Once data is added, it can't
   be deleted (i.e. immutable) and data that shares the same partition key will go to the same shard (i.e. ordering).   
```

##### Kinesis Data Streams Security
```
a) Supports in-flight encryption through HTTPS endpoints and at-rest encryption through AWS managed Kinesis keys and KMS 
   keys. Client-side encryption is also supported if the client wants to perform encryption themselves.
b) IAM policies can be used to regulate access to Kinesis APIs.
```

##### Kinesis Data Firehose
```
a) Data Firehose loads data streams from producers (e.g. apps, Kinesis Data Streams, CloudWatch) as records, can optionally 
   convert/compress data or perform custom transform logic through Lambda, and batch writes to destinations (e.g. S3, 
   3rd party data stores, or custom endpoints).
b) Data Firehose is serverless and provides automatic scaling. Unlike Kinesis Data Streams, Firehose itself is not a 
   data store and is solely responsible for ingesting data.
c) Buffer sizes and intervals can be configured to control latency and costs.
```