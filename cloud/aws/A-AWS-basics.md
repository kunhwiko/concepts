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
should only be used for tasks that require root user credentials (e.g. create IAM user, delete AWS account).  
```

##### Region
```
Region (e.g. us-west-1) represents a cluster of data centers in a certain geographical area and consists of multiple
zones (e.g. us-west-1a). Most AWS services are region scoped and some services might not be available in some regions.
```

##### Availability Zone
```
Availability zones are one or more data centers with redundant power, networking, and connectivity. Each zone is isolated
from one another to reduce blast damages caused by disasters.
```

### Account Protection
---
##### Multi Factor Authentication (MFA)
```
Uses a combination of a known password and a token on a personal device to login to an AWS console. The use of MFA is
recommended for login via the root user and IAM users.
```

##### Access Keys
```
a) While passwords and MFA are used to protect login to the AWS console, access keys can be used to secure access for 
   AWS CLI and SDK. The access key comes with an access key ID and a secret key which acts as a username and password.
b) When using the AWS CLI or SDK with an access key, the commands will match the IAM permissions of the user for which 
   the access key had been created from.
```

##### 

### ARN
---
##### Amazon Resource Name (ARN)
```
ARNs uniquely identify AWS resources. These identifiers are typically in the following format:
  * arn:partition:service:region:account-id:resource-id
  * arn:partition:service:region:account-id:resource-type:resource-id
  * arn:partition:service:region:account-id:resource-type/resource-id
```
