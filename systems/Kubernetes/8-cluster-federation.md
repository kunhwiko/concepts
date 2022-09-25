### Cluster Federation Basics 
---
##### Federation History
```
History
   Step 1) Original proposal was to reuse existing Kubernetes APIs to federate clusters.
           However, there were many underlying problems to this approach.
   Step 2) New proposal published dedicated APIs for federation.

Cluster Federation
   a) Aggregate multiple Kubernetes clusters that are treated as a single logical cluster.
   b) Federation control plane presents a single unified view of the system.

Federation Control Plane
   a) Consists of federation API server and federation controller manager
```

##### Basic Terminology
```
KubeFed
   a) Control plane of the federation.
   
Host Cluster
   a) Cluster that exposes KubeFed APIs and runs as the KubeFed control plane.

Member Cluster
   a) Cluster that is registered with the KubeFed API.
   b) KubeFed controllers have authn credentials to contact this cluster.
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

### Networking
---
##### Getting Started
```
Good examples of Cluster Federation:
   a) [CockroachDB Networking](https://faun.pub/multi-cloud-multi-region-kubernetes-federation-part-2-e8d403150d4f)
   b) [CockroachDB Resources](https://faun.pub/multi-cloud-multi-region-kubernetes-federation-part-3-182fe2eecc85)
```

##### Domain
```
Domain
   a) Describes a domain that is to be set by external DNS provider.

ServiceDNSRecord
   a) Resource that specifies how a service will be registered to the federation.

IngressDNSRecord
   a) Resource that specifies how an ingress will be registered to the federation.
```
