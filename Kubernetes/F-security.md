### Definition of Secret
---
##### Secrets
```
a) Secrets contain sensitive info such as credentials and tokens and are namespace-scoped.
b) Secrets are stored by default as plaintext in etcd and are base64 encoded.
c) Kubelet can download and access secrets from the API server via TLS.
d) Secrets can be mounted as files via volumes or be picked up as container environment variables.
   The same secret can be mounted in multiple pods.
e) Kubelet stores secrets in node memory (never to disk) and are deleted when not needed.
```

##### ImagePullSecrets
```
Keys used to pull images from a private registry.
```

### Roles & Security Context
---
##### RBAC Authorization
```
Cluster Role
   a) Allows users to access namespaced/cluster-wide resources.     

Role
   a) Allows users to access namespaced resources.   
    
Cluster Role Binding
   a) Grants permissions granted by roles/clusterroles cluster-wide.

Role Binding
   a) Grants permissions granted by roles/clusterroles within a specific namespace.
```

##### Security Context
```
Security Context
   a) Applied at container level.
   b) Overrides podSecurityContext.
   c) Does not apply to volumes.

Pod Security Context
   a) Applies to all containers.
   b) Applies to volumes.

Supplemental Group
   a) Ability to supply additional GIDs. 

FsGroup
   a) Provides a supplementalGroup to change ownership of volumes to be owned by the pod.
   b) Kubernetes will change the permission of all files in the volumes to the GID. 
   c) Could harm other processes that were accessing the volumes with a different GID. 
   d) Could cause slow startup for large volumes as permissions need to be modified. 
```

##### Service Accounts
```
Namespace
   a) Provides a partition among users, and each namespace can be sealed with credentials.
   
User Accounts
   a) When a human tries to access a Kubernetes cluster, they are typically authenticated as a particular user account.
   b) Typically gives "admin" access, giving global access to all namespaces.
   
Service Accounts
   a) When processes in containers inside pods contact the API server, they are authenticated via a service account.
   b) When pods are instantiated, they are assigned a service account.
   c) Default service account is used if one is not assigned, which limits access to resources only in current namespace.
   d) Service accounts carry credentials used to talk to the API server via a secret volume mount.
   
Service Account Admission Controller
   a) Assigns at pod creation time a custom or default service account.
   b) Ensures pod has ImagePullSecrets when pulling from remote registry.
      If not specified in Pod spec, uses service account's ImagePullSecrets.
   c) Adds a volume with an API token for API access.
   d) Ensures default service account exists in every namespace.

Token Controller
   a) Whenever a service account is created, creates and adds an API token to the secret volume.
```

### Authentication and Authorization to API Server
---
##### Authentication
```
Step 1
   a) Users use keys and certificates to authenticate against the cluster over TLS.
   b) Cluster admins can choose what authentication strategy to use.
   c) If at least one authentication step succeeds, authentication is granted.

Impersonation
   a) Possible for users to impersonate different users (e.g. troubleshoot some issue for a different user).
   b) Requires passing impersonation headers to API request (e.g. kubectl --as / --as-group parameters).
```

##### Authorization
```
Step 2
   a) Authorization requests include info such as authenticated username and request verb.
   b) Cluster admins can choose what authorization strategy to use.
   c) All authorization steps must succeed for authorization to be granted.
   d) kubectl auth can-i ... verifies whether user can perform certain actions.
```

##### Admission Control Plugins
```
Step 3
   a) Address global cluster concerns after all authorization steps pass.
   b) Either validates (e.g. check violation/quota limits) or mutates (e.g. create namespace if it does not exist).
   c) All admission control steps must succeed for request to succeed.
   d) Cluster admins can choose what admission control plugins to use.
   e) Possible to define a custom dynamic admission control.
      https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/
```

### Security Tools
---
##### AppArmor
```
AppArmor : Linux kernel security module that allows you to create profiles to do the following
   a) Restrict network access of processes in container.
   b) Restrict Linux capabilities of container.
   c) Restrict file permissions of container.
   d) Provide improved auditing through logs.
```

### Malicious Attacks
---
##### Node Attacks
```
Problems
   a) Replace kubelet to run other workloads while communicating normally with API server.
   b) Access to shared resources and secrets.
   c) Send malicious messages or cause malicious resource drain.
   d) Malicious containers within (or even outside) a pod share networks / resources with other containers, and can escalate attacks.

Mitigation
   a) Limit resource.
   b) Carefully consider interaction between containers.
```

##### Network Attacks
```
Considerations between ease of discovery vs security
   a) Which endpoints should be publicly accessible and how to authenticate users?
   b) Need to control access within internal services in case someone obtains internal access.
   c) Sensitive data must be encrypted.
```

#### Image Attacks
```
Problems
   a) Malicious images designed to hack into systems.
   b) Vulnerable images are good targets for malicious attacks.

Mitigation
   a) Integrate static image analyzers.
   b) Limit resource access of containers.
   c) Patch known vulnerabilities as soon as possible.
```