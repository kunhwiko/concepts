### Secrets
---
##### Secrets
```
Secrets
   a) stores credentials and tokens that by default are stored as plaintext in etcd
   b) pods can mount secrets as files and the same secret can be mounted by multiple pods
   c) secrets in a pod are mounted in-memory (ephemeral) for security purposes 

ImagePullSecrets : keys to pull image from a private registry
```

### Roles & Security Context
---
##### RBAC Authorization
```
ClusterRole : allows users to access namespaced/cluster-wide resources     
Role        : allows users to access namespaced resources   
    
ClusterRoleBinding : grants permissions granted by roles/clusterroles cluster-wide 
RoleBinding        : grants permissions granted by roles/clusterroles within a specific namespace 
```

##### Security Context
```
securityContext
   a) applied at container level
   b) overrides podSecurityContext
   c) does not apply to volumes

podSecurityContext
   a) applies to all containers
   b) applies to volumes

supplementalGroup : ability to supply additional GIDs 

fsGroup
   a) provides a supplementalGroup to change ownership of volumes to be owned by the pod 
   b) Kubernetes will change the permission of all files in the volumes to the GID 
   c) could harm other processes that were accessing the volumes with a different GID 
   d) could cause slow startup for large volumes as permissions need to be modified 
```

##### Service Accounts
```
Namespace : provides a partition among users, and each namespace can be sealed with credentials
   
User Accounts
   a) when a human tries to access a Kubernetes cluster, they are typically authenticated as a particular user account
   b) typically gives "admin" access, giving global access to all namespaces 
   
Service Accounts
   a) when processes in containers inside pods contact the API server, they are authenticated via a service account
   b) when pods are instantiated, they are assigned a service account
   c) default service account is used if one is not assigned, which limits access to resources only in current namespace
   d) service accounts carry credentials used to talk to the API server via a secret volume mount 
   
Service Account Admission Controller
   a) assigns at pod creation time a custom or default service account
   b) ensures pod has ImagePullSecrets when pulling from remote registry (if not specified in Pod spec, uses service account's ImagePullSecrets)
   c) adds a volume with an API token for API access
   d) ensures default service account exists in every namespace

Token Controller : whenever a service account is created, creates and adds an API token to the secret volume
```

##### AppArmor
```
AppArmor : Linux kernel security module that allows you to create profiles to do the following
   a) restrict network access of processes in container
   b) restrict Linux capabilities of container
   c) restrict file permissions of container
   d) provide improved auditing through logs
```

### Authentication and Authorization to API Server
---
##### Authentication
```
Step 1
   a) users use keys and certificates to authenticate against the cluster over TLS
   b) cluster admins can choose what authentication strategy to use
   c) if at least one authentication step succeeds, authentication is granted

Impersonation
   a) possible for users to impersonate different users (e.g. troubleshoot some issue for a different user)
   b) requires passing impersonation headers to API request (e.g. kubectl --as / --as-group parameters)
```

##### Authorization
```
Step 2
   a) authorization requests include info such as authenticated username and request verb
   b) cluster admins can choose what authorization strategy to use
   c) all authorization steps must succeed for authorization to be granted
   d) kubectl auth can-i ... verifies whether user can perform certain actions
```

##### Admission Control Plugins
```
Step 3
   a) address global cluster concerns after all authorization steps pass
   b) either validates (e.g. check violation/quota limits) or mutates (e.g. create namespace if it does not exist)
   c) all admission control steps must succeed for request to succeed
   d) cluster admins can choose what admission control plugins to use
   e) possible to define a custom dynamic admission control (https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/)
```

### Malicious Attacks
---
##### Node Attacks
```
Problems
   a) replace kubelet to run other workloads while communicating normally with API server
   b) access to shared resources and secrets
   c) send malicious messages or cause malicious resource drain
   d) malicious containers within (or even outside) a pod share networks / resources with other containers, and can escalate attacks

Mitigation
   a) limit resource
   b) carefully consider interaction between containers 
```

##### Network Attacks
```
Considerations between ease of discovery vs security
   a) which endpoints should be publicly accessible and how to authenticate users?
   b) need to control access within internal services in case someone obtains internal access
   c) sensitive data must be encrypted 
```

#### Image Attacks
```
Problems
   a) malicious images designed to hack into systems
   b) vulnerable images are good targets for malicious attacks

Mitigation
   a) integrate static image analyzers
   b) limit resource access of containers
   c) patch known vulnerabilities as soon as possible
```
