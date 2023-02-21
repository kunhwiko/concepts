### Cluster Federation Basics 
---
##### Federation History
```
The original proposal for cluster federation was to reuse existing Kubernetes APIs to federate clusters.
However, there were many underlying complications to this approach.
Therefore, a new proposal was published with dedicated APIs for federation.
```

##### Cluster Federation
```
a) Cluster Federation aggregates multiple Kubernetes clusters that are treated as a single logical cluster.
b) Cluster Federation control plane presents a single unified view of the system.
```

##### Cluster Federation Components
```
KubeFed
  a) KubeFed is the control plane for the Kubernetes cluster federation.
   
Host Cluster
  a) Kubernetes cluster that exposes KubeFed APIs and runs as the KubeFed control plane.

Member Cluster
  a) Kubernetes cluster that is registered with the KubeFed API.
     Member clusters do not need to know they are part of the federation.
  b) KubeFed controllers have authentication credentials to contact this cluster.
```

##### Federation Control Plane
```
Consists of federation API server and federation controller manager.

Federation API Server
  a) Forwards requests to all member clusters in the federation.
     Member functions do not need to know they are part of the federation.
  b) Manages the state of clusters registered in the federation in etcd.
     The state of each cluster is stored in etcd of that cluster.

Federation Controller Manager
  a) Performs the duties of the controller manager (i.e. ensure desired state matches actual state) for the entire federation.
     This is achieved by observing the controller manager of each individual member cluster.
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

##### Advantages
```
Cloud Bursting
  a) Allows users to primarily run load on on-prem infra but use managed clusters when peak capacity is reached.

Private Data Protection
  a) Sensitive data and workloads are subject to external auditing and may need to run on on-prem.
  b) Certain data by law may need to stay in a particular geographical area.

Availability
  a) Provides availability and redundancy in case a cloud provider shuts down.
  b) Denies vendor lock ins and allows for the ability to negotiate prices.
  c) Able to distribute across availability zones, regions, and cloud providers.
```

##### Challenges
```
Location Affinity
  a) At a cluster federation layer, pod affinities and anti-affinities across clusters becomes a harder topic.

Federated Data Access
  a) Accessing data remotely could cause high latency for certain clusters in different geographical zones.
  b) Replicating data in each cluster could be expensive and would require smart syncing mechanisms.
  c) Using a hybrid solution to cache frequently accessed data is complicated and could result in caching stale data.

Federated Auto Scaling
  a) It is difficult to dynamically pick certain clusters to scale resources under load.
     Normally, scaling would require scaling resources on all member clusters.
  b) It is difficult to autoscale new member clusters due to setup time and public cloud quotas.
     Hybrid approach would involve increasing capacity of existing clusters and preparing new ones when need be.
```

### Networking
---
##### Domain / ServiceDNSRecord
```
Problem Statement
  a) If a federated service creates load balancers in multiple clusters, there needs to be a single endpoint that users can access.
  b) DNS allows for a manual mapping of all IP addresses as an A record to a single domain, but this can be a tedious process.

Domain
  a) Object that specifies a domain/subdomain to be used for FederatedService LoadBalancer objects.

ServiceDNSRecord
  a) Kubernetes CRD that links a FederatedService object to a domain object.
  b) For every service to be registered with DNS, both Domain and ServiceDNSRecord CRD object must be created. 
```

##### IngressDNSRecord
```
Kubernetes CRD that identifies FederatedIngress objects to understand and aggregate all ingress targets throughout the federation.
Multi-cluster ingress controller helps to read all ingress objects of the federation.
```

##### DNSEndpoint
```
Kubernetes CRD that represents endpoints for a FederatedService and are created when DNSRecord objects are created. 

ServiceDNSEndpoint A Records
  a) For services, the following endpoints will be generated per zone and region:
       * <service>.<namespace>.<federation>.svc.<domain>
       * <service>.<namespace>.<federation>.svc.<region>.<domain>
       * <service>.<namespace>.<federation>.svc.<availability-zone>.<region>.<domain>
  b) Region and zone are determined from node labels.
     If these labels are not present, nodes must manually be labeled.
```

##### ExternalDNS
```
a) ExternalDNS synchronizes exposed services and ingresses with DNS providers.
b) External DNS will watch for LoadBalancer services, ExternalType services, and ingress hostnames.
   For newly scanned resources, ExternalDNS will upsert DNS records to external DNS providers.
```

##### Federated Networking
```
Step 1) User creates a Domain + ServiceDNSRecord or IngressDNSRecord for every service to be registered with a DNS provider.
Step 2) KubeFed controller watches for new records and generates a DNSEndpoint which External DNS will read from.
Step 3) External DNS upserts DNS records.

More here: https://faun.pub/multi-cloud-multi-region-kubernetes-federation-part-2-e8d403150d4f
```