### Secrets
---
##### Secrets
```
Secrets
   - stores credentials and tokens that by default are stored as plaintext in etcd
   - Pods can mount Secrets as files and same Secret can be mounted by multiple Pods
   - Secrets in a Pod are mounted in-memory (ephemeral) for security purposes 
```

### Roles
---
##### RBAC Authorization
```
ClusterRole 
   - allows users to access namespaced/cluster-wide resources 
    
Role 
   - allows users to access namespaced resources   
    
ClusterRoleBinding
   - grants permissions granted Roles/ClusterRoles cluster-wide 
    
RoleBinding
   - grants permissions granted by Roles/ClusterRoles within a specific namespace 
```

### Security Context
---
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