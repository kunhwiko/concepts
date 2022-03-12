### Secrets
---
##### Secrets
```
Secrets
   - stores credentials and tokens that by default are stored as plaintext in etcd
   - Pods can mount Secrets as files and same Secret can be mounted by multiple Pods
   - Secrets in a Pod are mounted in-memory (ephemeral) for security purposes 

ImagePullSecrets
   - keys to pull image from a private registry
```

### Roles & Security Context
---
##### RBAC Authorization
```
ClusterRole 
   - allows users to access namespaced/cluster-wide resources 
    
Role 
   - allows users to access namespaced resources   
    
ClusterRoleBinding
   - grants permissions granted by Roles/ClusterRoles cluster-wide 
    
RoleBinding
   - grants permissions granted by Roles/ClusterRoles within a specific namespace 
```

##### Groups 
```
supplementalGroup: 
   - ability to supply additional GIDs 

fsGroup:
   - provides a supplementalGroup to change ownership of Volumes to be owned by the Pod 
   - Kubernetes will change the permission of all files in the volumes to the GID 
   - could harm other processes that were accessing the volumes with a different GID 
   - could cause slow startup for large volumes as permissions need to be modified 
```

##### Service Accounts
```
Namespace:
   - provides a partition among users, and each namespace can be sealed with credentials
   
User Accounts:
   - When a human tries to access a Kubernetes cluster, they are typically authenticated as a particular User Account
   - Typically gives "admin" access, giving global access to all namespaces 
   
Service Accounts:
   - When processes in containers inside Pods contact the API Server, they are authenticated via a Service Account
   - When Pods are instantiated, they are assigned a Service Account
   - Default Service Account is used if one is not assigned, which limits access to resources only in current namespace
   - Service Accounts carry credentials used to talk to the API server via a Secret volume mount 
   
Service Account Admission Controller
   - Assigns at Pod creation time a custom or default Service Account
   - Ensures Pod has ImagePullSecrets when pulling from remote registry (if not specified in Pod spec, uses Service Account's ImagePullSecrets)
   - Adds a volume with an API token for API access
   - Ensures default Service Account exists in every namespace

Token Controller
   - Whenever a Service Account is created, creates and adds an API token to the secret volume
```

### Malicious Attacks
---
##### Node Attacks
```
Problems
   1. Replace kubelet to run other workloads while communicating normally with API Server
   2. Access to shared resources and secrets
   3. Send malicious messages or cause malicious resource drain
   4. Malicious containers within (or even outside) a Pod share networks / resources with other containers, and can escalate attacks

Mitigation
   1. Limit resource
   2. Carefully consider interaction between containers 
```

##### Network Attacks
```
Considerations between ease of discovery vs security
   1. Which endpoints should be publicly accessible and how to authenticate users?
   2. Need to control access within internal services in case someone has internal access
   3. Sensitive data must be encrypted 
```

#### Image Attacks
```
Problems
   1. Malicious images designed to hack into systems
   2. Vulnerable images are good targets for malicious attacks

Mitigation
   1. Integrate static image analyzers
   2. Limit resource access of containers
   3. Patch known vulnerabilities as soon as possible
```
