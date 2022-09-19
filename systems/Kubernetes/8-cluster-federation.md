### Cluster Federation  
---
##### History
```
History
   Step 1) Original proposal was to reuse existing Kubernetes APIs to federate clusters.
           However, there were many underlying problems to this approach.
   Step 2) New proposal published dedicated APIs for federation.
```

##### Basic Concepts
```
Cluster Federation
   a) Aggregate multiple Kubernetes clusters that are treated as a single logical cluster.
   b) Federation control plane presents a single unified view of the system.

Federation Control Plane
   a) Consists of federation API server and federation controller manager
```

##### Advantages
```
Cloud Bursting
   a) Primarily runs on-premise infra but uses cloud computing when peak capacity is reached.

Private Data Protection
   a) Sensitive data and workloads are subject to external auditing and may need to run on-premise.
   b) Workloads can dynamically change from non-sensitive to sensitive data, although policies should prevent this.
      If it cannot be prevented, sensitive data needs to be migrated to non-public clusters.
   c) Certain data by law needs to stay in a particular geographical area.

Availability
   a) Provides availability and redundancy in case a cloud provider shuts down.
   b) Denies vendor lock inds and allows for the ability to negotiate prices.
   c) Able to distribute across availability zones, regions, and cloud providers.
```