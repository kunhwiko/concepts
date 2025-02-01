### KMS
---
##### Key Management Service (KMS)
```
a) KMS is a service that manages encryption keys. Symmetric keys can both encrypt and decrypt data, but users cannot get
   direct access to these keys and must invoke KMS APIs to use. Asymmetric keys are split into public and private key 
   pairs, and only the public key is accessible to use for encryption outside AWS or for users who cannot invoke KMS.   
b) KMS is fully integrated with IAM for authorization and usage can be audited by CloudTrail.
c) KMS is by default region specific. When doing S3 or EBS replications across regions or accounts, objects and data 
   need to be decrypted in the source account and then reencrypted with a new key in the new region. While it is possible
   to copy keys into multi-region, this is not recommended due to higher exposure and possible security risks.
```

##### KMS Key Types
```
a) AWS Owned Keys: Keys managed by AWS (e.g. SSE-S3) that live and are maintained in an AWS managed account.
b) AWS Managed Keys: Keys that are created and managed by AWS on a user's behalf in one's AWS account for a given service.
   (e.g. aws/rds, aws/ebs). Keys are automatically rotated every year.
c) Customer Managed Keys: Keys that are created and managed by users in their own AWS account. Keys can be rotated 
   automatically or on-demand. Unlike AWS managed keys, customer keys occur expenses on creation and usage. 
```

##### KMS Key Policy
```
Key policies are used to specify who can use and manage keys. While a default policy grants the root user full access
to the key, users can also create custom policies to define who can access the key, who can administer the key, and to
setup cross-account access.
```

### Secrets Storage
---
##### SSM Parameter Store
```
a) SSM Parameter Store is a service that provides secure storage for configuration data and secrets. Version tracking
   is supported and notifications can be integrated with EventBridge.
b) SSM Parameter Store can validate access to data through IAM permissions and can encrypt data with KMS keys.
```

##### AWS Secrets Manager
```
a) Secrets Manager is a service that helps manage secrets and database credentials, which can be encrypted with KMS keys. 
   As opposed to SSM parameter store, secrets and credentials can be rotated automatically based on a schedule.
b) Secrets manager can replicate secrets across multiple AWS regions and keep read replicas in sync. 
```

##### AWS Certificate Manager (ACM)
```
a) ACM is a service that deploys, manages, and rotates public and private TLS certificates. ACM can be integrated with 
   services such as Elastic Load Balancer, CloudFront, and API Gateway to load certificates.
b) When setting up public certificates, ACM can validate domain ownership through DNS validation or email validation.
   When using DNS validation, ACM will provide a CNAME record that users must add to their domain's DNS configuration.
   ACM will automatically check for these records and issue the certificate once validated. 
```