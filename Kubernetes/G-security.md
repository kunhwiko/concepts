### Definition of Namespace and Secret
---
##### Namespace
```
Provides a partition among users, and each namespace can be sealed with credentials.
```

##### Secrets
```
a) Secrets contain sensitive info such as credentials and tokens and are namespace scoped.
   To limit access to secrets, they should be put in namespaces accessible to a limited set of users or services.
b) Secrets are stored by default as plaintext in etcd and are base64 encoded.
c) Kubelet can download and access secrets from the API server via TLS.
d) Secrets can be mounted as files via volumes, specified in service accounts, or be picked up as container environment variables.
   The same secret can be mounted in multiple pods.
e) Kubelet stores secrets in node memory (never to disk) and are deleted when not needed.
```

##### ImagePullSecrets
```
Keys used to pull images from a private registry.
```

### Definition of Service Account
---
##### User Accounts
```
a) When a human tries to access a Kubernetes cluster, they are typically authenticated via a user account.
b) Typically gives "admin" access, giving global access to all namespaces.
c) Best practices are to give users privileges to a limited set of namespaces.
```

##### Service Accounts
```
a) When pods are instantiated, they are assigned a service account.
b) Pod processes that need to contact the API server are authenticated via a service account.
c) Default service account is used if one is not assigned, which limits access to resources in the current namespace.
d) Service accounts carry credentials (e.g. token to communicate with API server) that are mounted in a secret volume.
```

##### Service Account Admission Controller
```
a) Assigns at pod creation time a custom or default service account.
b) If specified, ensures pod has ImagePullSecrets if images need to be pulled from a remote registry.
   If not specified in the pod spec, uses the service account's ImagePullSecrets instead.
c) Adds a volume to the pod with a token that is used to authenticate to the API server.
d) Ensures default service account exists in every namespace.
```

##### Token Controller
```
Whenever a service account is created, creates and adds a token that is used to authenticate to the API server to the secret volume.
```

### Definition of Security Context
---
##### Container Security Context
```
a) OS level security settings such as UID that are applied at the container level.
b) Container security contexts override pod security contexts.
c) Cannot be applied to volumes.
```

##### Pod Security Context
```
a) OS level security settings such as UID that are applied at the pod level.
   The security settings are applied to all containers in the pod.
b) Applied to volumes.
```

##### SupplementalGroup
```
Ability to supply additional GIDs. 
```

##### fsGroup
```
a) Provides a SupplementalGroup to change ownership of volumes to be owned by a particular pod.
b) Kubernetes will change the permission of all files in the volumes to the GID. 
c) Could harm other processes that were accessing the volumes with a different GID. 
d) Could cause slow startup for large volumes as permissions need to be modified. 
```

### Definition of RBAC
---
##### RBAC Authorization
```
ClusterRoles
  a) Represent a set of permissions to resources in all namespaces.
  b) ClusterRoles are available for use in all namespaces.

ClusterRoleBindings
  a) Grants permissions defined by ClusterRoles.
  b) ClusterRoleBindings are available for use in all namespaces.

Roles
  a) Represent a set of permissions to resources in a given namespaces.    

RoleBindings
  a) Grants permissions defined by a Role in a given namespace.
  b) RoleBindings can also reference a ClusterRole and bind it to the given namespace.
```

### Authentication and Authorization to API Server
---
##### Step 1: Authentication
```
a) Users use keys and certificates to authenticate against the API server over TLS.
b) Cluster admins can choose what authentication strategy to use.
c) If at least one authentication step succeeds, authentication is granted.

Impersonation
  a) Possible for users to impersonate different users (e.g. troubleshoot some issue for a different user).
  b) Requires passing impersonation headers to API request (e.g. kubectl --as / --as-group parameters).
```

##### Step 2: Authorization
```
a) Authorization requests include info such as authenticated username and request verb.
b) Cluster admins can choose what authorization strategy to use.
c) All authorization steps must succeed for authorization to be granted.
d) kubectl auth can-i ... verifies whether user can perform certain actions.
```

##### Step 3: Admission Control Plugins
```
a) Address global cluster concerns after all authorization steps pass.
   All admission control steps must succeed for request to succeed.
b) Cluster admins can choose what admission control plugins to use.
c) Admission controllers can validate requests (e.g. denies requests if exceeding resource limits) or mutate requests (e.g. create namespace if it does not exist).
d) Alongside compiled admission plugins, it is possible to define custom dynamic admission controls.
```

##### Dynamic Admission Control
```
a) Admission plugins that run as webhooks during runtime and do not need to compiled in Kubernetes.
b) Validating admission webhooks deny requests if the validation criteria is not met.
c) Mutating admission webhooks are invoked first over validating admission webhooks.
   These webhooks can modify requests and objects sent to the API server.

More info: https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/
```

### Security Tools
---
##### AppArmor
```
Linux kernel security module that allows you to create profiles to do the following:
  a) Restrict network access of processes in container.
  b) Restrict Linux capabilities of container.
  c) Restrict file permissions of container.
  d) Provide improved auditing through logs.
```

### Malicious Attacks
---
##### Node Attacks
```
Attacks
  a) External user can replace kubelet to communicate normally with API server by sending false data.
     Meanwhile, they can use the node to run their own workloads.
  b) External user can gain access to shared resources and secrets.
  c) External user can send malicious messages and disrupt the cluster or cause resource drain.

Mitigation
  a) Place a limit on resource capacities.
  b) Carefully consider interaction between containers and limit privileges as much as possible.
```

##### Network Attacks
```
Considerations between ease of discovery vs security
  a) Which endpoints should be publicly accessible, and if public, how do we authenticate users?
  b) What privileges can we take away? What containers do not need to talk with each other?
  c) How can we control access among internal services in case someone obtains internal access?
  d) What data is considered sensitive enough that we should be encrypting it?
```

#### Image Attacks
```
Attacks
  a) Malicious images can be designed to hack into systems.
  b) Vulnerable images are good targets for malicious attacks.

Mitigation
  a) Integrate static image analyzers.
  b) Limit resource access of containers.
  c) Patch known vulnerabilities as soon as possible.
```