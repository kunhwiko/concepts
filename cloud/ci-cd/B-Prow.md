### Prow
---
##### Prow Definition
```
Prow is a Kubernetes native CI/CD system that performs tasks based on various events (e.g. GitHub events).
Prow can define a pre-commit step (e.g. before PR is merged) and a post-commit step (e.g. after PR is merged).
```

##### Prowjobs
```
Prowjobs are the basic unit of work in Prow that performs certain actions based on certain events.
If specified, Prow can also run Tekton pipelines instead of Prowjobs.
```

### Prow Architecture
---
##### Prow Architecture
```
The overall Prow architecture is described here: https://docs.prow.k8s.io/docs/overview/architecture/.
Core components in Prow are better described here: https://docs.prow.k8s.io/docs/components/core/
```

##### Crier
```
Crier reports the status of ProwJobs to a user specified platform (e.g. GitHub, Slack).
OAuth tokens will need to be given for Crier to be able to communicate to GitHub or Slack.
```

##### Deck
```
Deck shows what jobs are running or have been executed on the Prow UI.
```

##### Hook
```
Hooks are webhooks that listen on events (e.g. user types "/retest" on GitHub) to trigger Prowjobs or Tekton PipelineRuns.
```

##### Horologium
```
Horologium is a utility that manages the execution and lifecycle of periodic Prowjobs (i.e. cron-based, interval-based).
```

##### Sinker
```
Sinker cleans up old jobs based on user specified parameters.
```

##### Plank
```
Plank is a Go program that decorates Prowjobs through Pod Utilities.
Decoration is done through init containers as well as sidecar containers.
```

##### Pod Utilities
```
Pod Utilities help to ensure the following:
  * Execution of test code is done against the correct versions of the source code (i.e. clones and targets relevant source code)
  * Prepare test environments
  * Upload job metadata/status/artifacts to a cloud storage
```

##### Tide
```
a) Tide has logic to automatically tests PRs based on a given criteria (i.e. tide-in).
b) Tide logic to merge PRs based on a given criteria (i.e. tide-out). 
```
