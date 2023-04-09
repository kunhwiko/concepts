### Tekton
---
##### Tekton Definition
```
Tekton is a Kubernetes native project for CI/CD that leverages Kubernetes to run a series of user defined tasks that verify user provided behaviors. 
Tekton API reference documentations are found here: https://tekton.dev/docs/pipelines/.
```

##### Steps
```
Steps are containers executed in a pod (i.e. inside of Tekton tasks).
These containers perform work as specified by the user.
```

##### Tasks
```
a) Tasks are resources that define a sequence of steps that run in order and are executed by TaskRun resources.
   Tasks will start a pod that executes all the steps mentioned in the task definition.
b) Steps and tasks can share a common Tekton workspace to share data by mounting the same persistent volume.
   EmptyDirs can also be used to share data across steps of the same task, but not among different tasks.
```

##### Pipelines
```
Pipelines are resources that define a chain of Tasks that could run sequentially or in parallel and are executed by PipelineRun resources. 
PipelineRuns will trigger TaskRuns as it executes the pipeline resource.
```

##### EventListeners
```
EventListener is a deployment with a service that listens for HTTP requests.
When a request is received, the resource defines what TriggerBindings and TriggerTemplates to run.
```

##### TriggerBindings
```
TriggerBindings take events and parses information to use as parameters.
```

##### TriggerTemplates
```
TriggerTemplates invoke PipelineRuns when an event occurs based on parameters from TriggerBindings.
```