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
##### Group ID 
```
supplementalGroup: 
   - ability to add additional GIDs 

fsGroup:
   - used to set the group that owns the pod volumes 
   - Kubernetes will change the permission of all files in the volumes to the GID 
   - could harm other processes that were accessing the volumes with a different GID 
   - could cause slow startup for large volumes as permissions need to be modified 
```