### AWS Accounts
---
##### AWS Accounts
```
An AWS account is required to be able to control AWS resources. An email address, password, AWS account name, phone 
number, and credit card information are required to create an AWS account. AWS billing is done per account. 
```

##### Root User
```
The root user is created when an AWS account is created. The root user has admin permissions for all AWS services, and 
it is recommended to use this user only for tasks that require root user credentials.  
```

##### Region
```
Region (e.g. us-west-1) is a geographical area that consists of multiple zones (e.g. us-west-1a).
```

### ARN
---
##### Amazon Resource Name (ARN)
```
ARNs uniquely identify AWS resources. These identifiers are typically in the following format:
  * arn:partition:service:region:account-id:resource-id
  * arn:partition:service:region:account-id:resource-type:resource-id
  * arn:partition:service:region:account-id:resource-type/resource-id
```