### Observability Practices
---
##### Observability Considerations
```
Logging
  a) What should be the format of logs? Should it be plain text, binary logs, structured logs?
  b) How should logs be stored? Should it be file system, database, cloud storage?
  c) How should logs be aggregated?
  d) For remote logging, how do we ensure that sensitive data is not sent to third parties?

Metrics
  a) What metrics are valuable and of concern?
  b) How should data be visualized and what should be visualized?
  c) When should we trigger alerts?

Distributed Tracing
  a) How can requests that are bounced across multiple microservices be traced effectively?
```

##### Log Collection Strategies
```
Direct Logging
  a) The application container is responsible for sending logs to a remote logging service.
  b) Applications need to be aware of the remote logging service and there is no separation of concerns in this approach.

Node Agent
  a) Application containers will dump stdout and stderr in which node agents will intercept and deliver to a remote logging service.
  b) Node agents are typically deployed as DaemonSets.
  c) Efficient but requires control of worker nodes.

Sidecar Container
  a) Sidecar containers will collect logs from the application container and deliver them to a remote logging service.
  b) This method is not as efficient as having a node agent, but is a good approach when there is no control over worker nodes. 
```