### Open Policy Agent (OPA)
---
##### OPA Definition
```
OPA is a CNCF project that provides a policy as code framework to help decouple policy decision making 
away from applications. These policies are structured based on hierarchical structured data (e.g. YAML and JSON).
OPA can be used to make admission control decisions in Kubernetes after authentication and authorization.
```

##### Deployment Strategies
```
OPA is typically deployed through one of the following methods.
  * OPA with kube-mgmt
  * OPA Gatekeeper

The architecture is explained in details here: https://www.openpolicyagent.org/docs/latest/kubernetes-introduction/
```


### OPA with kube-mgmt
---
##### Init Containers
```
Validating or mutating webhooks should first be configured to call the OPA service running in some namespace.
Init containers can then be used to ensure OPA instances have the proper configurations necessary, such as 
ensuring TLS certificates to talk to the Kubernetes API have been mounted properly.
```

##### OPA Instances
```
a) 3 OPA instances are typically deployed and backed by the OPA service for load balancing purposes. The instances
   cache policies and data in memory (i.e. do not persist) to make performant policy decisions.
b) OPA instances need to be configured to securely receive and respond to policy queries from Kubernetes 
   admission controllers through TLS.
c) API server sends a webhook request to OPA containing an AdmissionReview object. Refer to the following link for more:
   https://www.openpolicyagent.org/docs/latest/kubernetes-primer/#detailed-admission-control-flow. This object becomes
   bound to the `input` global variable in Rego.
```

##### Kube Mgmt
```
a) Kube-mgmt runs as a sidecar container that discovers policies stored in configmaps and loads them into memory on OPA. 
   It also has the capability to load configmap JSON data into OPA under the `data` global variable in Rego.
   Kube-mgmt can determine which configmaps to load based on annotations:
     * `openpolicyagent.org/policy=rego` to load rego policies
     * `openpolicyagent.org/data=opa` to load JSON data
b) Kube-mgmt is able to cache replications of Kubernetes resources through the `--replicate` flag and load them as 
   objects to be used in rego code. Reference the following for more: https://github.com/open-policy-agent/kube-mgmt. 
```

### OPA Gatekeeper
---
##### Components
```
OPA Gatekeeper for Kubernetes admission control typically consists of the following:
  * 3 Gatekeeper instances known as controller-managers that are load balanced. These instances respond to
    webhook requests from Kubernetes admission controllers. Each Gatekeeper uses OPA to make admission 
    decisions based on defined policies. Gatekeepers have the option to sync information of cluster state
    into OPA before decisions are made. 
  * 1 audit controller responsible for periodically monitoring violations in the cluster.
```

##### Constraints
```
For OPA Gatekeeper, it takes two separate Kubernetes objects for a policy to take effect.
  * Constraint Template has Rego code that defines the criteria for policy violations. This resource acts as
    a template for actual constraints to instantiate.
  * Constraint is an instantiation of a constraint template. It describes the enforcement level of policies 
    (e.g. warn, deny) and where to enforce those policies. It also holds a violations field that describes 
    all violations for the constraint in the cluster.
```

##### OPA Gatekeeper Architecture
```
Step 1) Admission controllers are configured to call the OPA gatekeeper service deployed on a given namespace.
Step 2) Gatekeeper instances are configured to handle requests from admission controllers. Gatekeeper instances
        should mount certificates to securely communicate with webhook controllers.
Step 3) Audit controller instances should be deployed to monitor the Kubernetes cluster for policy violations.
        These instances are not required to listen on the admission controllers.
```

### Rego
---
##### Rego Definition
```
Rego is a language that helps easily build rules in OPA.
More information can be found here: https://www.openpolicyagent.org/docs/latest/policy-language/#what-is-rego.
It is also worth reviewing the course here: https://academy.styra.com/.
```

##### Packages
```
Rules make different policy decisions and typically act as if statement. Rules are organized into policies, 
which are a set of rules with a hierarchical name. Policies allow for code reusability.
```

##### Examples of Rules
```
# input
request:
  method: GET
  path: "api/v1/products"
  token:
    user: alice
    roles:
      - engineering

# 'code' evaluates to 200 if statements inside the rule evaluate to true
# statements in a rule works as an AND statement
# rules could be invoked via localhost:8181/v1/data/<path-to-policy>/code
code = 200 {
  is_read        # true
  is_alice_user  # true
}

# is_read rule here works as an OR statement
# is_read is true if the OR statement here is true
is_read {
  input.request.method == "GET"
}
is_read {
  input.request.method == "PUT"
}

is_alice_user {
  input.request.token.user == "alice"
}
``` 

##### Examples of Packages
```
# user.rego
package util.user

is_authenticated { ... }
is_admin { ... }

# compute.rego
package authz.compute

import data.util.user as u

allow {
  u.is_authenticated
  ...
}

# network.rego
package authz.network

import data.util.user as u

allow {
  u.is_authenticated
  ...
}

# main.rego
# packages could be invoked via localhost:8181/v1/data/authz/main
package authz.main

allow {
  input_is_compute
  data.authz.compute.allow
}

allow {
  input_is_network
  data.authz.network.allow
}
```
