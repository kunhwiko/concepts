### KMS
---
##### Key Management Service (KMS)
```
a) KMS is a service that manages encryption keys. Symmetric keys can both encrypt and decrypt data, but users cannot get
   direct access to these keys and must invoke KMS APIs to use. Asymmetric keys are split into public and private key 
   pairs, and only the public key is accessible to use for encryption outside AWS or for users who cannot invoke KMS.   
b) KMS is fully integrated with IAM for authorization and usage can be audited by CloudTrail. While KMS is typically
   region specific, a key can be configured to be copied across regions in an AWS account.
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