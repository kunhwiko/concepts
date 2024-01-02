### EKS Basics
---
##### Elastic Kubernetes Service (EKS)
```
EKS is a managed Kubernetes service and AWS's container-as-a-service offering. EKS leverages EC2 for hosting cluster 
nodes and integrates AWS infrastructure (e.g. VPC networking, IAM, availability).
```

### Security
---
##### IAM Roles for Service Accounts (IRSA)
```
IRSA is an EKS native way to allow pods to assume IAM roles through service account annotations. It moves away from 
having to define IAM roles at the node level (i.e. EC2 instance).
```

##### IRSA Workflow
```
Step 1) An OIDC provider for the cluster must be created with a public endpoint. This OIDC provider authenticates and 
        issues new JWT tokens that are projected into pods via service accounts.
Step 2) An IAM role must be created with a trust relationship to the OIDC provider:
          * The following is used to indicate that any service authenticated by the OIDC provider can be trusted:
            "Principal": {
              "Federated": "arn:aws:iam::$account_id:oidc-provider/$oidc_provider"
            }
          * The following is used to indicate that the caller must pass a web identity token that indicates it has been
            authenticated by the OIDC provider:
            "Action": "sts:AssumeRoleWithWebIdentity"
Step 3) Pods should be annotated with the ARN of the IAM role to assume. A mutating webhook from the Pod Identity Webhook 
        will inject environment variables representing the IAM Role ARN and path to the pod's service account token.
Step 4) When a pod makes an AWS API call, the request is sent to STS to evaluate the trust policy of the IAM role. STS 
        verifies with the public OIDC endpoint to verify that the pod's JWT token is issued by the same OIDC provider.
Step 5) If verified, STS issues temporary credentials which effectively assigns the IAM role to the pod.
```