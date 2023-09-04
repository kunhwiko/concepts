### Namespace and Secret
---
##### Namespace
```
Provides a partition among users, and each namespace can be sealed with credentials.
```

##### Secrets
```
a) Secrets are namespaced resources that can contain sensitive info such as credentials and tokens. To limit access to
   secrets, they should be put in namespaces accessible only to a limited set of users or services.
b) Secrets are stored as unencrypted plain text in etcd and are typically base64 encoded when kubelets download them to
   worker nodes. Kubelets will typically download these secrets through the API server via TLS and store them in node
   memory (i.e. never to disk), and are deleted when not needed.
c) Secrets can be mounted as files via volumes, specified in service accounts, or be picked up as container environment 
   variables. The same secret can be mounted to multiple pods.
```

##### EncryptionConfiguration
```
Configuration on the API server to control how API data is encrypted in etcd (e.g. encryption algorithms, what resources
to encrypt). As an example, this can be used to encrypt secret objects before they are persisted in etcd.
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
a) Service accounts represent an identity for a pod that allows for authenticating to the API server or represent what
   privileges the pod has. When pods are instantiated, they are assigned a service account. A default service account is 
   used if one is not assigned.
b) Before Kubernetes 1.22, service accounts used to be attached with non-expiring token secrets that applications can 
   use to authenticate against the Kubernetes API. Nowadays, tokens with bounded lifetimes are obtained via the 
   TokenRequest API and are mounted inside a projected volume.
```

##### Service Account Admission Controller
```
a) Assigns at pod creation time a custom or default service account. Also ensures a default service account exists in 
   every namespace.
b) Mutates the pod during pod creation time to add a projected volume to the pod containing a token to authenticate
   against the API server. If a token does not exist, the kubelet will refresh the token via the TokenRequest API.   
c) If specified, ensures pod has ImagePullSecrets if images need to be pulled from a remote registry. If not specified 
   in the pod spec, uses the service account's ImagePullSecrets instead.
```

##### Token Controller
```
Whenever a service account is created, creates and adds a token as a secret that is used to authenticate to the API 
server. Note that TokenRequest based tokens are now in favor of secret based tokens as of Kubernetes 1.22.
```

### Security Context and RBAC
---
##### Pod Security Context
```
a) OS level security settings such as UID that are applied at the pod level. The security settings are applied to all 
   containers in the pod.
b) Applied to volumes.
```

##### Container Security Context
```
a) OS level security settings such as UID that are applied at the container level.
b) Container security contexts override pod security contexts.
c) Cannot be applied to volumes.
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
ClusterRoles: 
  * Represents a set of permissions to resources in all namespaces. ClusterRoles need to be used to grant permissions to 
    cluster-scoped resources (e.g. namespace, node, pv, csr). 
  * ClusterRoles are available for use in all namespaces.

ClusterRoleBindings
  * Grants permissions defined by ClusterRoles.
  * ClusterRoleBindings are available for use in all namespaces.

Roles
  * Represents a set of permissions to resources in a given namespace.    

RoleBindings
  * Grants permissions defined by a Role in a given namespace.
  * RoleBindings can also reference a ClusterRole and bind it to the given namespace.
  
Refer to the following for system component RBAC: 
  * https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
```

### Component Authentication and Authorization
---
##### TLS
```
Various components (e.g. etcd, API server, scheduler, kubelets) in Kubernetes will communicate with one another via TLS.
Note the following:
  * Root ceritifcates (e.g. ca.key, ca.crt) need to first be generated for the cluster. These certificates should be 
    stored in a safe environment.
  * Certificates for individual components need to be signed with the CA key pair. While a single certificate can be 
    used as both the client and server certificate, best practice is to separate them in case the certificate becomes 
    compromised. For example, the API server can have a server certificate, a client certificate for kubelets, and a 
    client certificate for etcd.
  * Client certificates can specify a group that they belong to (e.g. system:masters) depending on the permissions that
    are needed. Refer to: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-subjects.
```

##### Certificate Signing Requests (CSR)
```
A CertificateSigningRequest object is used to request that a certificate be signed by a specified signer, after which 
the request may be approved or denied by administrators before finally being signed. Refer to the following for signers: 
https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers.
```

##### Kubeconfig
```
Kubernetes uses a YAML file called kubeconfig to store cluster auth information (e.g. certs, keys). The path to this 
file can be specified via the environment variable KUBECONFIG. This file consists of the following:
  * Clusters: HTTP endpoint to the Kubernetes API server along with the certificate of the root CA.
  * Contexts: Links a user to a cluster.
  * Users: Contains the client certificate and the client private key.
```

##### Step 1: Authentication
```
a) Users use keys and certificates to authenticate against the API server over TLS.
b) Cluster admins can choose what authentication strategy to use. If at least one authentication step succeeds, 
   authentication is granted.
c) Administrators can impersonate different users (e.g. troubleshoot some issue for a different user) through the --as 
   and --as-group parameters.
```

##### Step 2: Authorization
```
a) Authorization requests include info such as authenticated username and request verb.
b) Cluster admins can choose what authorization strategy to use. When multiple authorization modules are configured, each 
   is checked in sequence. If any authorizer approves or denies a request, that decision is immediately returned and no 
   other authorizer is consulted.
c) kubectl auth can-i ... verifies whether user can perform certain actions.
```

##### Node Authorization
```
The node authorizer is a special purpose authorization mode that authorizes API requests made by kubelets. It recognizes
requests from kubelets as their certificates should be configured with a "system:nodes" group with a username of 
"system:node:<node-name>". Refer to: https://kubernetes.io/docs/reference/access-authn-authz/node/.
```

##### Step 3: Admission Control Plugins
```
a) Admission controller is code that intercepts requests to the API server after authorization but prior to storing 
   objects into etcd. All admission control steps must pass for the request to succeed.
b) Admission controllers can validate (e.g. deny requests that exceed resource limits) or mutate (e.g. create namespace 
   if it does not exist) requests. Mutating admission controllers are invoked before validating admission controllers.
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
  * External user can replace kubelet to communicate normally with API server by sending false data. Meanwhile, they can 
    use the node to run their own workloads.
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
