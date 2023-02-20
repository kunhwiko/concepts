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
##### Resources
```
Reference: https://faun.pub/multi-cloud-multi-region-kubernetes-federation-part-2-e8d403150d4f
```

##### Domain / DNSRecord
```
Problems
   a) If a service is created in multiple clusters, they are bound to have different IP addresses.
      These IP addresses need to turn into a single endpoint that users can use.
   b) DNS lets us manually map all IP addresses as an A record to a single domain, but IP addresses can change.

Domain
   a) Configuration of domain object is necessary for FederatedService LoadBalancer objects.
   b) Specifies a domain/subdomain that can be used for all linked FederatedServices.
   c) Name metadata for this object specifies the <federation> record of DNSEndpoints.

ServiceDNSRecord
   a) CRD that links service object to domain object.
   b) For every service to be registered with DNS, a ServiceDNSRecord CRD object must be created. 
```

##### IngressDNSRecord
```
IngressDNSRecord
   a) Cluster may need to send requests from the receiving cluster to a different cluster.
      IngressDNSRecord CRD objects can be created for KubeFed to register a DNSEndpoint CRD object.
```

##### DNSEndpoint
```
DNSEndpoint
   a) Once a DNSRecord object is created, KubeFed controller will create a DNSEndpoint CRD object.
   b) This object will be read by ExternalDNS to create DNS records.

ServiceDNSEndpoint A Records
   a) For services, three A records will be generated:
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
   b) For newly scanned resources, upserts DNS records in external DNS providers.
```