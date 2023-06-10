### Namespace and Secret
---
##### Namespace
```
Provides a partition among users, and each namespace can be sealed with credentials.
```

##### Secrets
```
a) Secrets contain sensitive info such as credentials and tokens that are namespace scoped. To limit access to secrets, 
   they should be put in namespaces accessible only to a limited set of users or services.
b) Secrets are stored by default as plaintext in etcd and are base64 encoded. Kubelets can download and access these
   secrets by invoking the API server via TLS. Kubelets will then store secrets in node memory (i.e. never to disk)
   and are deleted when not needed.
c) Secrets can be mounted as files via volumes, specified in service accounts, or be picked up as container environment 
   variables. The same secret can be mounted to multiple pods.
```

##### ImagePullSecrets
```
Keys used to pull images from a private registry.
```

### Service Account
---
##### User Accounts
```
When a human tries to access a Kubernetes cluster, they are typically authenticated via a user account. These accounts 
typically give "admin" access, giving global access to all namespaces. Best practice is to give users privilege to a 
limited set of namespaces.
```

##### Service Accounts
```
a) When pods are instantiated, they are assigned a service account. A default service account is used if one is not
   assigned, which limits access to resources in the current namespace.
b) Service accounts come with tokens (i.e. credentials to talk to the API server) that are mounted as secrets. Pod 
   processes that need to contact the API server are authenticated with this token.
```

##### Service Account Admission Controller
```
a) Assigns at pod creation time a custom or default service account. Also ensures a default service account exists in 
   every namespace.
b) Adds a volume to the pod with a token that is used to authenticate to the API server.   
c) If specified, ensures pod has ImagePullSecrets if images need to be pulled from a remote registry.
   If not specified in the pod spec, uses the service account's ImagePullSecrets instead.
```

##### Token Controller
```
Whenever a service account is created, creates and adds a token as a secret that is used to authenticate to the API server.
```

### Security Context and RBAC
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

##### RBAC Authorization
```
ClusterRoles
  * Represent a set of permissions to resources in all namespaces.
  * ClusterRoles are available for use in all namespaces.

ClusterRoleBindings
  * Grants permissions defined by ClusterRoles.
  * ClusterRoleBindings are available for use in all namespaces.

Roles
  * Represent a set of permissions to resources in a given namespace.    

RoleBindings
  * Grants permissions defined by a Role in a given namespace.
  * RoleBindings can also reference a ClusterRole and bind it to the given namespace.
```

### Authentication and Authorization to API Server
---
##### kubeconfig
```
a) Kubernetes uses a YAML file called kubeconfig to store cluster auth information for kubectl.
b) kubeconfig contains a list of contexts to which kubectl will refer to.
c) kubeconfig contains API server endpoints that kubectl will contact on port 443.
d) Path to the kubeconfig file can be specified via the environment variable KUBECONFIG.
```

##### Step 1: Authentication
```
a) Users use keys and certificates to authenticate against the API server over TLS.
b) Cluster admins can choose what authentication strategy to use.
c) If at least one authentication step succeeds, authentication is granted.

Impersonation
  * Possible for users to impersonate different users (e.g. troubleshoot some issue for a different user).
  * Requires passing impersonation headers to API request (e.g. kubectl --as / --as-group parameters).
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
a) Admission controller is code that intercepts requests to the API server after authorization but prior to storing objects into etcd.
   All admission control steps must pass for the request to succeed.
b) Admission controllers can validate (e.g. denies requests if exceeding resource limits) or mutate (e.g. create namespace if it does not exist) requests.
   Mutating admission controllers are invoked before validating admission controllers and can modify requests and objects sent to the API server.
c) Admission plugins can be compiled or dynamic (i.e. run as webhooks and do not require restart of API server).
```

### Security Tools
---
##### AppArmor
```
Linux kernel security module that allows you to create profiles to do the following:
  * Restrict network access of processes in container.
  * Restrict Linux capabilities of container.
  * Restrict file permissions of container.
  * Provide improved auditing through logs.
```

### Malicious Attacks
---
##### Node Attacks
```
Attacks
  * External user can replace kubelet to communicate normally with API server by sending false data.
    Meanwhile, they can use the node to run their own workloads.
  * External user can gain access to shared resources and secrets.
  * External user can send malicious messages and disrupt the cluster or cause resource drain.

Mitigation
  * Place a limit on resource capacities.
  * Carefully consider interaction between containers and limit privileges as much as possible.
```

##### Network Attacks
```
Considerations between ease of discovery vs security
  * Which endpoints should be publicly accessible, and if public, how do we authenticate users?
  * What privileges can we take away? What containers do not need to talk with each other?
  * How can we control access among internal services in case someone obtains internal access?
  * What data is considered sensitive enough that we should be encrypting it?
```

##### Image Attacks
```
Attacks
  * Malicious images can be designed to hack into systems.
  * Vulnerable images are good targets for malicious attacks.

Mitigation
  * Integrate static image analyzers.
  * Limit resource access of containers.
  * Patch known vulnerabilities as soon as possible.
```
