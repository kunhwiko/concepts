### Tekton
---
##### Tekton Definition
```
Tekton is a CI/CD framework that leverages Kubernetes to run user specified pipelines. Tekton API reference docs can be 
found here: https://tekton.dev/docs/pipelines/.
```

##### Steps
```
Steps are containers executed in a Kubernetes pod (i.e. inside of Tekton tasks). These containers perform work (e.g. 
shell scripts) as specified by the user.
```

##### Tasks
```
a) Tasks are executed by TaskRun resources and run as a pod to complete a sequence of steps in sequentially order.
b) Steps and tasks can use a common Tekton workspace to share data by mounting the same persistent volume. EmptyDirs can 
   also be mounted to share data across steps of the same task, but not among different tasks.
```

##### Pipelines
```
Pipelines are executed by PipelineRun resources and define a chain of Tasks that could run sequentially or in parallel.
Note that PipelineRuns will trigger TaskRuns on a given Tekton pipeline.
```

##### EventListeners
```
EventListener is a deployment with a service that listens for HTTP requests. When a request is received, the resource 
defines what TriggerBindings and TriggerTemplates to run.
```

##### TriggerBindings
```
TriggerBindings take events and parses information to use as parameters for TriggerTemplates.
```

##### TriggerTemplates
```
TriggerTemplates invoke PipelineRuns when an event occurs based on parameters from TriggerBindings.
```
