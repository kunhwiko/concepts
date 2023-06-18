### Workload Resources
---
##### Deployments
```
a) Controller that defines a desired state for pods and replicasets.
b) Enables users to scale number of replicas, control rollout of updates, and allow rollbacks to a previous state.
c) Enables users to check or update status of pods in a centralized way.  
```

##### ReplicaSets
```
Controller that is typically created by deployments and ensures that a specified number of pod replicas are running at 
a given time. Replicasets make rollbacks and rolling updates an easier process.
```

##### StatefulSets
```
a) Controller that uniquely identifies pods with a stable hostname and ordering.
b) Pods become associated with their own stable storage (i.e. dynamic PVC) that are linked based on hostname and order.
c) Headless services are used to manage network identity of pods.
d) Does not support rollbacks to previous versions.
```

##### DaemonSets
```
Controller that ensures that pods run on all or a designated subset of nodes.
```

##### Pods 
```
a) Pods encapsulate one or more containers and are assigned to nodes. While containers could be separate user/PID 
   Linux namespaces, containers in a pod are always scheduled together on the same node and share the same network and 
   mount namespace. This means containers share the same IP address and port and can communicate via localhost.
b) Pods specify network/storage configs for containers, specifications for how to run containers, and can help setup
   necessary environments via init containers.   
c) Pods are ephemeral by nature but can be distinguished by unique IDs.
```

### Label, Annotation, and Selector
---
##### Labels
```
a) Key-value pairs to identify, select, and group pods or other objects together based on some criteria.
b) Keys are comprised of a prefix and name and must be unique on an object.
c) Key prefixes are optional, must be a valid DNS subdomain, and can be at most 253 characters long.
d) Key names are mandatory, only allow for certain characters, and must be at most 63 characters long.
e) Labels should not be used for attaching arbitrary metadata to objects.
```

##### Label Selectors
```
a) Chooses objects based on some criteria.
   Two or more selectors imply selector1 AND selector2 instead of OR.
b) Selectors can be equality-based or set-based.
```

##### Label Examples
```yaml
kind: Deployment
  spec:
    # This deployment uses the template field to create Pods with labels "app: test"
    template:
      metadata:
        labels:
          app: test 
          
    selector:
      matchLabel:        
        # Deployments use selectors to know which Pods it needs to manage.
        # This field is predefined to prevent mutation of what pods the deployments should manage.  
        app: test 
```

##### Annotations
```
a) Key-value pairs to attach arbitrary metadata that Kubernetes does not care about.
b) Keys are comprised of a prefix and name and follow similar restrictions to labels.
```

### Jobs
---
##### Jobs
```
a) Manages one or more pods to execute some operation until it is successful.
b) If a pod fails, a new pod runs to finish the operation.
c) Jobs and their pods will not be cleared after completion. 
   Deleting the job will clear all related resources.
``` 

##### Job Example
```yaml
# here only 3 pods will run in parallel as not all 5 are required
# when a job is done, pods that complete their task will update their status as "completed"
# run kubectl logs <pod-name> to view output of the job

apiVersion: batch/v1
kind: Job
metadata:
  name: factorial5
spec:
  # by default, this value is 1
  # specifies how many successful completions are required 
  completions: 3

  # determines max number of pods to launch to run job in parallel
  # job will not launch more pods than required
  parallelism: 5

  template:
    spec:
      containers:
      - name: factorial5
        image: some-python-image-1.32
        command: ["python",  "-c", "from math import factorial; print(factorial(5))"]
      # job should not restart after a successful completion
      restartPolicy: Never
```

##### Cron Jobs
```
a) Jobs that run periodically.
b) Each invocation launches a new job object along with corresponding pods.
c) Deleting a cron job will delete existing jobs and pods.
```

### Distributed System Design Patterns
---
##### Sidecar Pattern
```
Step 1) In the same pod, create a separate container from the main application container.
Step 2) This separate container will take care of supplemental features such as logging.
Step 3) This provides advantages such as lessening burden of application container and makes it easier to switch different components when upgrading.
```

##### Ambassador Pattern
```
Step 1) Like the Sidecar Pattern, a separate container from the main app container is created.
Step 2) This ambassador container acts as a proxy to the main app container and can filter requests or enforce certain policies.
Step 3) Often used with legacy apps to extend networking/security capabilities when the legacy apps are risky to modify.
        As an example, the main app container can connect to a Redis cluster by first communicating to the ambassador via localhost.
Step 4) The ambassador will filter the requests and send write requests to the Redis master and send read requests to Redis replicas.
Step 5) This enables configurations to be updated on the ambassador side while not having to worry about legacy code. 
```

##### Adapter Pattern
```
Step 1) Assume main application has been updated but generates output in a different format.
Step 2) Consumers of the output have not been upgraded to read in the new format.
Step 3) Adapter standardizes output until all consumers have been upgraded.
```