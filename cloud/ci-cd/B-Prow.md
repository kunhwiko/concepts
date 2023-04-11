### Prow
---
##### Prow Definition
```
Prow is a Kubernetes native CI/CD system that performs tasks based on various events (e.g. GitHub events).
Prow can define a pre-commit step (e.g. before PR is merged) and a post-commit step (e.g. after PR is merged).
Prow will by default run ProwJobs but can also run Tekton jobs if specified.
```

##### Prow Architecture
```
The overall Prow architecture is described here: https://docs.prow.k8s.io/docs/overview/architecture/.
```