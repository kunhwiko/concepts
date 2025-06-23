### Overview
---
##### Accounts
```
An AWS account is required to be able to control AWS resources. An email address, password, account name, phone number,
and credit card information are required to create an account. Billing is done per account. 
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

##### Edge Locations
```
Edge locations are AWS data centers designed to deliver services with the lowest latency possible. Services such as
CloudFront and Route53 cache data at these locations for faster responses.
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
While passwords and MFA are used to protect login to the AWS console, access and secret keys can be used to authenticate 
for the AWS CLI and SDK. The access key acts as a username and the secret key acts as a password.

When using the AWS CLI or SDK locally with an access key, the commands will match the IAM permissions of the user for 
which the access key had been created from.
```

##### IAM Credentials Report
```
Report that lists all of an account's users and the status of their various credentials (access key usage, access key 
age, MFA enablement, user creation time etc).
```

##### IAM Access Advisor
```
Shows the service permissions granted to a user and when those services were last accessed. This information can then 
be used to revise existing policies for the user. 
```

### Properties & APIs
---
##### Amazon Resource Name (ARN)
```
ARNs uniquely identify AWS resources. These identifiers are typically in the following format:
- arn:partition:service:region:account-id:resource-id
- arn:partition:service:region:account-id:resource-type:resource-id
- arn:partition:service:region:account-id:resource-type/resource-id
```

##### Properties
```
Refer to the CloudFormation template docs for the list of AWS properties. Below is an example for IAM:
- https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/AWS_IAM.html
```

##### API
```
Refer to the API reference docs for the list of AWS APIs. Below is an example for IAM APIs:
- https://docs.aws.amazon.com/IAM/latest/APIReference/API_Operations.html
```
