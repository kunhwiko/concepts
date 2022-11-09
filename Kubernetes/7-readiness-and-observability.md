### Probing & Readiness
---
##### Probes
```
Startup Probes
   a) Defines what it means for a container to have started.
   b) Disables liveness and readiness checks until startup succeeds.
   c) Provides time for the container application to initialize.

Liveness Probes
   a) Defines what it means for a container to be alive.
   b) Kubelets have basic restart policies if a container process crashes, but it might not be sufficient.

Readiness Probes
   a) Defines what it means for a container application to be able to serve requests.
   b) Container may be initialized, but necessary dependencies might still be or have become unavailable.
   c) When readiness probe fails, containers are not killed but requests will not be received.
   d) Ensures requests don't flood a pod with requests it cannot process.
```

##### Readiness Gates
```
Problems
   a) Readiness probes help address pod level readiness, but not an infrastructure level readiness.
   b) Services, network policies forwarding traffic might not be ready yet.

Readiness Gates
   a) Provide an extra podSpec to specify set of conditions for when things are ready.
```

##### Probe Examples
```yaml
apiVersion: v1
kind: Pod
spec:
  containers:
    - name: app
      image: test-app
      args:
      - /bin/sh
      - -c
      - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
      livenessProbe:
        exec:
          command: 
          - cat 
          - /tmp/healthy
        # give time for the pod to initialize
        # kubelet will not check liveness for this period
        initialDelaySeconds: 30
```

##### Init Containers
```
Init Containers
   a) Containers that run to completion before all other containers start.
   b) Can configure such that probes will start only after init containers are started.
   c) Init containers can provide setup scripts.
```

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
