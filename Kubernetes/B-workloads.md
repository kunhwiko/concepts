### Definition of Cluster and Node
---
##### Cluster
```
Collection of hosts that provide compute, memory, storage, and networking resources.
```

##### Nodes
```
a) A single host that could be a physical or virtual machine.
b) Master nodes act as a control plane for Kubernetes and manage worker nodes.
   Master nodes follow raft protocol, meaning an odd number of master nodes must exist for consensus to be possible. 
c) Worker nodes carry out actual work.
```

### Workload Resources
---
##### ReplicaSets
```
a) Ensures that a specified number of pod replicas are running at a given time.
b) Makes it easy for rollbacks and rolling updates.
```

##### Deployments
```
a) Defines a desired state for pods and replicasets.
b) Enables users to scale number of replicas, control rollout of updates, rollback to previous deployments.
c) Enables users to check or update status of pods.  
```

##### Daemon Sets
```
Ensures that a pod runs on all or a designated subset of nodes.
```

##### StatefulSets
```
a) Controller uniquely identifies pods with a stable hostname and ordering.
b) Pods become associated with their own stable storage (i.e. dynamic PVC) that are linked based on hostname and order.
c) Headless services are used to manage network identity of pods.
d) Cannot roll back to previous versions.

More info: https://medium.com/stakater/k8s-deployments-vs-statefulsets-vs-daemonsets-60582f0c62d4 
```

##### Pods 
```
a) Encapsulates one or more containers and are assigned to nodes. 
   Containers in a pod are always scheduled together to the same machine.
b) All containers in a pod share the same IP address and port, meaning they can communicate via localhost.
c) Pods can specify storage and network configs to be used by containers.
d) Pods are ephemeral by nature, but can be distinguished by unique IDs.
```

### Definition of Label, Annotation, and Selector
---
##### Labels
```
a) Key-value pairs to identify, select, and group pods or other objects together based on some criteria.
b) Keys are comprised of a prefix and name and must be unique on an object.
c) Key prefixes are optional, must be a valid DNS subdomain, and can be at most 253 characters long.
d) Key names are mandatory, only allow for certain characters, and must be at most 63 characters long.
e) Labels should not used for attaching arbitrary metadata to objects.
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

### Definition of Job
---
##### Jobs
```
a) Manages one or more pods to execute some operation until it is successful.
b) If a pod fails, then a new pod runs to finish the operation.
c) Jobs and their pods will not be cleared after completion (delete the job to clear all resources).
``` 

##### Example
```yaml
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
      # job should not restart after completing
      restartPolicy: Never

# here only 3 pods will run in parallel as not all 5 are required
# when a job is done, pods that completed the task will be status "completed"
# run kubectl logs <pod-name> to view output of the job
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