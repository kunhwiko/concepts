### Prometheus
---
##### Features
```
Features
   a) Collects and stores metrics as time series data that are organized as key-value labels.
   b) Pull model where Prometheus server periodically scrapes targets rather than applications needing to push data.
   c) Time series collection and pulling is done via HTTP transport.
   d) Supports PromQL query language that makes it easy to fetch metrics.
   e) Designed for reliability and fault tolerance, but not for accuracy and completeness. 
```

##### Components
```
Prometheus Server
   a) Scrapes from targets and stores time series data.
   
Push Gateway
   a) Short lived jobs can push metrics to the Gateway where Prometheus can then pull data from.
   
Alertmanager
   a) Prometheus servers can push metric alerts to Alertmanager, which can then send alert notifications.

Grafana
   a) Visualization tool that can query and pull data from the Prometheus server.
```
