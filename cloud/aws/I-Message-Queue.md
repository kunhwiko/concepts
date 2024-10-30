### SQS
---
##### Simple Queue Service (SQS)
```
Managed queue service where producers send messages to a queue and consumers poll and delete messages from the queue.
```

##### Standard Queue
```
a) Standard queue has unlimited throughput (i.e. send as many messages as required), supports unlimited number of 
   messages in the queue, and low latency on pub/sub.
b) Each message has a short retention period (i.e. max 14 days) and a limit of 256KB. After a message is polled by a 
   consumer, it is hidden from other consumers for a short period and must be processed within that time frame. A consumer 
   can use the ChangeMessageVisibility API to extend the visibility timeout for a message.
c) Standard queue supports best effort ordering (i.e. messages are not guaranteed to be ordered) and at least once 
   delivery model (i.e. duplicate messages can exist). Applications should take this into account.
d) Queue length from CloudWatch metrics can be used with ASGs to scale instances based on the number of messages.
e) SQS supports long polling, which allows consumers to wait for messages to arrive for a certain period if the queue is 
   empty. This decreases the number of API calls made to SQS and improves latency.
```

##### FIFO Queue
```
FIFO queue guarantees ordering, an exactly once send model, and message deduplication at the cost of limited throughput 
(i.e. messages per second is restricted).
```

##### SQS Security
```
a) Supports in-flight encryption through HTTPS endpoints and at-rest encryption through AWS managed SQS keys and KMS keys.
a) IAM policies can be used to regulate access to SQS APIs (e.g. ProcessMessage, DeleteMessage).
b) SQS access policies, much like S3 bucket policies, can be used to regulate access to queues. These policies can be
   used to allow for cross-account access or allowing other services to access the queue.
```