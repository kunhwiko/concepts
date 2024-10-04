### Observability Practices
---
##### Observability Considerations
```
Logging
  * What should be the format of logs? Should it be plain text, binary logs, structured logs?
  * How should logs be stored? Should it be file system, database, cloud storage?
  * How should logs be aggregated?
  * For remote logging, how do we ensure that sensitive data is not sent to third parties?

Metrics
  * What metrics are valuable and of concern?
  * How should data be visualized and what should be visualized?
  * When should we trigger alerts?

Distributed Tracing
  * How can requests that are bounced across multiple microservices be traced effectively?
```

##### Log Collection Strategies
```
Direct Logging
  * The application container is responsible for sending logs to a remote logging service.
  * Applications need to be aware of the remote logging service and there is no separation of concerns in this approach.

Node Agent
  * Application containers will dump stdout/stderr to nodes and agents will deliver them to a remote logging service.
  * Node agents are typically deployed as DaemonSets.
  * Efficient but requires control of worker nodes.

Sidecar Container
  * Sidecar containers will collect logs from the application container and deliver them to a remote logging service.
  * This is not as efficient as having a node agent, but is a good approach when there is no control over worker nodes. 
```

### Prometheus
---
##### Prometheus
```
Prometheus is a CNCF graduated project and a time series database that focuses on metrics collection and alert 
management. It has the following features:
  * Pulls metrics from components over HTTP by default. A push model is supported through a push gateway.
  * Built in time series database that can be queried through PromQL.
  * Built in alert manager and the ability to define alerting rules.
```

##### Prometheus Operator
```
Custom controller that manages Prometheus, ServiceMonitor, PodMonitor, PrometheusRule, Alertmanager custom resources.  
```

##### ServiceMonitor / PodMonitor
```
Defines a set of targets to be monitored by Prometheus. These monitors can select services/pods to scrape metrics from 
based on label selectors, namespace selectors, and ports. The Prometheus object specifies which service monitors and
pod monitors it will pull metrics from.
```

##### Kube State Metrics
```
Monitoring tool that listens to the Kubernetes API server to provide metrics that focus primarily on the state and 
health of objects in the cluster. These metrics include pod status, available replicas, resource limits, annotations etc.
The full list is given here: https://github.com/kubernetes/kube-state-metrics/tree/main/docs. It is not to be confused
with the metrics server which primarily is used to measure resource utilization and for scaling.
```

##### Prometheus Node Exporter
```
Daemonset that collects low level information on Kubernetes nodes, including network connections, CPU load, memory 
usage, thread count, number of scrapes etc. These metrics are more detailed than the ones provided by kube state metrics.
```

##### PrometheusRule
```
Defines alerting and recording rules for a Prometheus instance. The Prometheus object specifies which rules it needs to 
look for.
```

##### Alertmanager
```
Component that can be configured to send alerts to various communication stacks (e.g. Slack). Prometheus can be 
configured to know which alertmanager to forward alerts to.
```