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
ensuring that TLS certificates to talk to the API server have been mounted properly.
```

##### OPA Instances
```
a) 3 OPA instances are typically deployed and backed by the OPA service for load balancing purposes. The instances
   cache policies and data in memory (i.e. do not persist) to make performant policy decisions.
b) OPA instances need to be configured to securely receive and respond to policy queries from Kubernetes 
   admission controllers through TLS.
c) API server sends a webhook request to OPA containing an AdmissionReview object. Refer to the following link for more:
   https://www.openpolicyagent.org/docs/latest/kubernetes-primer/#detailed-admission-control-flow. This object becomes
   bound to the `input` global variable in Rego code and is processed by OPA during decision making.
```

##### Kube Mgmt
```
a) Kube-mgmt runs as a sidecar container that discovers policies stored in ConfigMaps and loads them into memory on OPA. 
   It also has the capability to load ConfigMap JSON data into OPA under the `data` global variable in Rego.
   Kube-mgmt can determine which ConfigMaps to load based on annotations:
     * `openpolicyagent.org/policy=rego` to load Rego policies
     * `openpolicyagent.org/data=opa` to load JSON data
b) Kube-mgmt is able to cache replications of Kubernetes resources through the `--replicate` flag and load them as 
   objects to be used in Rego code. Reference the following for more: https://github.com/open-policy-agent/kube-mgmt. 
```

### OPA Gatekeeper
---
##### OPA Gatekeeper
```
Compared to OPA with kube-mgmt, Gatekeeper provides the following functionality:
  * Extensive policy library for validation and mutation logic that are usable outside the box.
  * Native Kubernetes CRD for defining policies (i.e. ConstraintTemplates instead of ConfigMaps) with control over parameter schema.
  * Native Kubernetes CRD (i.e. Constraints) to provide fine-grained control over policies. This includes controlling target 
    namespaces, resources, enforcement level) through YAML rather than Rego code.
  * Cluster audit functionality.
```

##### Reference Docs
```
a) Gatekeeper docs: https://open-policy-agent.github.io/gatekeeper/website/docs/ 
b) Gatekeeper policy library: https://open-policy-agent.github.io/gatekeeper-library/website/
```

##### OPA Gatekeeper Instances
```
a) Gatekeeper instances (a.k.a. controller managers) are typically deployed as 3 instances for load balancing purposes. 
b) Gatekeeper instances respond to webhook requests from Kubernetes admission controllers and need to be mounted with TLS certs. 
c) Gatekeeper instances internally uses OPA to make decisions. These instances can be configured to sync cluster state into 
   OPA before making decisions.
```

##### Audit Controller
```
Audit controller is responsible for periodically monitoring the state of the cluster for any violations. The audit controller 
itself does not require TLS certificates to communicate with the API server.
```

##### ConstraintTemplate
```
a) ConstraintTemplate is a resource that has Rego code to define the criteria for policy violations. 
b) ConstraintTemplate acts as a template to instantiate Constraint custom resources. 
```

##### ConstraintTemplate Example
```yaml
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: k8srequiredlabels
spec:
  crd:
    spec:
      names:
        kind: K8sRequiredLabels
      validation:
        # Allows for control over the schema for the Constraints custom resource `parameters` field
        openAPIV3Schema:
          properties:
            labels:
              type: array
              items: string
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8srequiredlabels
 
        violation[{"msg": msg, "details": {"missing_labels": missing}}] {
          provided := {label | input.review.object.metadata.labels[label]}
          required := {label | label := input.parameters.labels[_]}
          missing := required - provided
          count(missing) > 0
          msg := sprintf("you must provide labels: %v", [missing])
        }
```

##### Constraint
```
a) Constraint is an instantiation of the ConstraintTemplate resource that describe the enforcement level of policies (e.g. warn, 
   deny), the target resource, target namespace, and parameters. 
b) Constraint holds a violations field that describes violations for the constraint in the cluster as reported by the audit controller.
```

##### Constraint Example
```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredLabels
metadata:
  name: required-labels-pods
spec:
  # specifies action to take when a violation is met
  enforcementAction: warn
  # specifies when the constraint is to be applied
  match:
    kinds:
      - apiGroups: [""]
        kinds: ["Pod"]
    excludedNamespaces:
      - kube-*
  # specifies required parameters (e.g. specified label is required) - used as `input.parameters.labels[_]`
  parameters:                                                           
    labels: ["hello-world"]
status:
  # below is periodically populated by the audit operation
  totalViolations: 2
  violations:
    - enforcementAction: warn
      kind: Pod
      message: 'you must provide labels: {"hello-world"}'
      name: some-pod
      namespace: testns1
    - enforcementAction: warn
      kind: Pod
      message: 'you must provide labels: {"hello-world"}'
      name: some-pod2
      namespace: testns2
```

### Rego
---
##### Rego Definition
```
Rego is a language that helps easily build rules and policies in OPA. Use the following to refer to the 
Rego documentation: https://www.openpolicyagent.org/docs/latest/policy-language/#what-is-rego.
```

##### Rego Rule Basics
```
# Variable code  will be assigned to 200 if all the statements in the {} bracket are true.
# Statements in the {}  brackets represent conditions that need be evaluated.
# If input.request.method == "GET"  is true, code will be assigned to 200. If false, it will be assigned to undefined.
code = 200 {
  input.request.method == "GET"
}

# Boolean expressions here work as an AND statement - v1 will result to true
v1 {                   # this implicitly is `v1 = true`
    x := 42            # OPA searches for expressions that evaluate to booleans
    y := 41
    x > y              # evaluates to true
    true               # evaluates to true
}

# Boolean expressions here work as an AND statement - v2 will result to undefined
v2 {
    "hello" == "hello" # evaluates to true
    "hello" == "world" # evaluates to false
}

# Boolean expressions here work as an OR statement - v3 will result to 200
v3 = 200 {
    true      # evaluates to true
}
v3 = 404 {
    false     # evaluates to false
}

# Rules can be used to compute values
port = res {
    values := split(input.request.host, ":")
    res := to_number(values[1])
}
```

##### Input Global Variable
```yaml
apiVersion: admission.k8s.io/v1
kind: AdmissionReview             # whenever a validating webhook request is sent to OPA, it will receive information here as `input`
request:
  kind:
    group:
    kind: Pod                     # as an example, this value can be queried in Rego via `input.request.kind.group.kind`
    version: v1
  object:
    metadata:
      name: myapp
    spec:
      containers:
        - image: nginx
          name: nginx-frontend
        - image: mysql
          name: mysql-backend
```

##### Data Global Variable
```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  name: policytestdata
  namespace: opa
  labels:
    openpolicyagent.org/data: opa           # this tells kube-mgmt to import as external data
data:
  policytestdata: |-                        # this value can be queried via `data.<opa-namespace>.<configmap-name>.<configmap-key>`
    {
      "opa-test-pod": {"testpolicy" : ["shouldBeAllowed"]}  
    }
```

##### OPA Policies and Packages
```yaml
# OPA "rules" can be organized into "policies", which are a set of rules with a hierarchical name. 
# Policies allow for code reusability through the use of "packages".

# Policies can be retrieved via /v1/policies/<opa-namespace>/<configmap-name>/<configmap-key> endpoint
# Data can be retrieved via /v1/data/<opa-namespace>/<configmap-name>/<configmap-key> endpoint

kind: ConfigMap
apiVersion: v1
metadata:
  name: opa-policy-cm
  namespace: opa
  labels:
    # this tells kube-mgmt to import as OPA policy, which is essentially a package of rules
    openpolicyagent.org/policy: rego
data:
  opa-policy: |
    # specifies package for rules below
    package kubernetes.admission
 
    # `deny` is not a keyword, it is a variable that will be consumed by the main policy. 
    # If the rule evaluates to true, deny will be assigned to an array with [msg].
    deny[msg] {
        some namespace, name
        # checks if the input resource is an ingress.
        input.request.kind.kind == "Ingress"
        # "_" is an iterator to look up the host  field over all rules in the input ingress.
        newhost := input.request.object.spec.rules[_].host
        # OPA has a record of current resources in Kubernetes. This is used to look up the host field by iterating over all rules for all ingresses in all namespaces.
        oldhost := data.kubernetes.ingresses[namespace][name].spec.rules[_].host      # Step 5
        newhost == oldhost
        input.request.object.metadata.namespace != namespace
        input.request.object.metadata.name != name
        # computes msg value 
        msg := sprintf("ingress conflicts with ingress %v/%v", [namespace, name])
    }
    
---

# There should be a system.main policy that acts a main decision point for when OPA receives webhook requests.
# The assigned deny value is used in the system.main policy.

kind: ConfigMap
apiVersion: v1
metadata:
  name: opa-default-system-main
  namespace: opa
  labels:
    openpolicyagent.org/policy: rego 
data:
  main: |
    package system
 
    import data.kubernetes.admission
 
    # this payload will be sent back as a response to control plane
    main = {                                      
      "apiVersion": "admission.k8s.io/v1",
      "kind": "AdmissionReview",
      "response": response,
    }
 
    # `default` keyword sets a value to a variable when none of the rules below succeed instead of defaulting to "undefined"
    default response = {"allowed": true}          
 
    # assigns "response" to "{uid: ..., allowed: false, status: {reason: ...}}" if "reason" is not "" 
    response = {
        "uid": input.request.uid,
        "allowed": false,
        "status": {
            "reason": reason,
        },
    } {
        reason = concat(", ", admission.deny)
        reason != ""
    }
 
    response = {
        "uid": input.request.uid,
        "allowed": true,
        "status": {
            "warning": warning,
        },
    } {
        count(admission.deny) == 0
        warning = concat(", ", admission.warning)
        warning != ""
    }
```
