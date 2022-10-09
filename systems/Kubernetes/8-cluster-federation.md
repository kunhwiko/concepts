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

##### Resource Federation
```
FederatedTemplate
   a) Holds the base specifications of a resource.
   b) FederatedDeployment holds the base specs for Deployment.
      This pushes Deployment object to all member clusters by default.

FederatedPlacement
   a) Holds the specifications of the clusters that resources should be distributed to.
   b) FederatedReplicaSetPlacement holds specs about which clusters FederatedReplicaSets go to.
   
FederatedOverrides
   a) Holds specifications of how Template resources should vary for certain clusters.
   b) FederatedReplicaSetOverrides shows how FederatedReplicaSet should vary for some clusters.
   
ReplicaSchedulingPreference
   a) Allows user to specify total number of replicas for a particular Template resource.
   b) Allows user to specify weighted distribution and min/max replica to member clusters.
   c) Allows for an option to dynamically rebalance resources if a resource cannot be scheduled.
```

##### Federation Control Plane
```
Federation Control Plane
   a) Consists of federation API server and federation controller manager.

Federation API Server
   a) Keeps the state of clusters registered in the federation in etcd.
   b) Interacts with federation controller manager.
   c) Routes requests to member clusters.
      Member clusters do not need to know they are part of the federation.

Federation Controller Manager
   a) Ensures the federation's desired state matches current state.
   b) Binary contains multiple controllers for different federated resources.
```

##### Advantages
```
Cloud Bursting
   a) Primarily runs on-premise infra but uses managed clusters when peak capacity is reached.

Private Data Protection
   a) Sensitive data and workloads are subject to external auditing and may need to run on-premise.
   b) Workloads can dynamically change from non-sensitive to sensitive data, although policies should prevent this.
      If it cannot be prevented, sensitive data needs to be migrated to non-public clusters.
   c) Certain data by law needs to stay in a particular geographical area.

Availability
   a) Provides availability and redundancy in case a cloud provider shuts down.
   b) Denies vendor lock ins and allows for the ability to negotiate prices.
   c) Able to distribute across availability zones, regions, and cloud providers.
```

##### Challenges
```
Pod Location Affinity
   a) If two pods needs to be strictly coupled, this could be done by putting them in all member clusters.
   b) Best architecture is for pods to be loosely coupled.
   c) Difficult to make pods preferentially coupled or strictly decoupled.

Data Access
   a) Access data remotely if need be but risk high latency.
   b) Replica data in each cluster but will be expensive to store and will require syncing mechanisms.
   c) Use a hybrid solution by caching most used data but is complicated and could result in caching stale data.

Federated Scaling
   a) Difficult to dynamically pick just a few clusters to scale resources under load.
      Normally will require scaling resources on all member clusters.
   b) Difficult to start new member clusters due to public cloud quota issues.
      Hybrid approach is to increase capacity of existing clusters and preparing new ones when need be.
```

### Networking
---
##### Resources
```
Reference: https://faun.pub/multi-cloud-multi-region-kubernetes-federation-part-2-e8d403150d4f
```

##### Domain / DNSRecord
```
Problems
   a) If a service is created in multiple clusters, they are bound to have different IP addresses.
      These IP addresses need to turn into a single endpoint that users can use.
   b) DNS lets us map all IP addresses as an A record to a single domain, but IP addresses can change.

Domain
   a) Configuration of domain object is necessary for FederatedService LoadBalancer objects.
   b) Specifies a domain/subdomain that can be used for all linked FederatedServices.
   c) Name metadata for this object specifies the <federation> record of DNSEndpoints.

ServiceDNSRecord
   a) CRD that links service object to domain object.
   b) For every service to be registered with DNS, a DNSRecord CRD object must be created. 
```

##### DNSEndpoint
```
DNSEndpoint
   a) Once a DNSRecord object is created, KubeFed controller will create a DNSEndpoint CRD object.
   b) This object will be read by ExternalDNS to create DNS records.

DNSEndpoint A Records
   a) For services, will create three A records:
      * <service>.<namespace>.<federation>.svc.<domain>
      * <service>.<namespace>.<federation>.svc.<region>.<domain>
      * <service>.<namespace>.<federation>.svc.<availability-zone>.<region>.<domain>
   b) Region and zone are determined from node labels.
      If they are not there, manual processes will need to be involved.
```

##### ExternalDNS
```
ExternalDNS
   a) Synchronizes exposed services and ingresses with DNS providers.
      Watches and lists LoadBalancer services, ExternalType services, and ingress hostnames.
   b) For federations where services are replicated across clusters with varying endpoints, DNSEndpoint should be used.
   c) For newly scanned resources, upserts DNS records in external DNS providers.
```

### Tools
---
##### Gardener
```
Garden Cluster
   a) Responsible for managing seed clusters.

Seed clusters
   a) Responsible for managing shoot clusters.
   b) Control planes of all shoot clusters run as pods and services in the seed cluster.

Shoot clusters
   a) Runs actual workloads and contains only worker nodes.
```
