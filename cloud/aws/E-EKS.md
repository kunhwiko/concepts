### Security
---
##### IAM Roles for Service Accounts (IRSA)
```
IRSA is an EKS native way to allow pods to assume IAM roles through service account annotations. It moves away from 
having to define IAM roles at the node level (i.e. EC2 instance).

Step 1) IAM role must be created with a trust relationship to the EKS OIDC provider. Each EKS cluster can be configured 
        to have an OIDC provider with a public endpoint to authenticate users to the cluster.
Step 2) Pods should be annotated with the ARN of the IAM role to assume. A mutating webhook called the Pod Identity Webhook
        will inject environment variables representing the IAM Role ARN and path to the pod's service account token.
Step 3) When a pod makes an AWS API call, STS evaluates the trust policy of the IAM role. STS checks with the trusted 
        OIDC endpoint to verify that the JWT token is issued by the cluster's OIDC provider.
Step 4) STS issues temporary credentials which effectively assigns the IAM role to the pod.
```