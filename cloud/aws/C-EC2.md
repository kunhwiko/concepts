### EC2
---
##### EC2 Properties
```
All EC2 properties are listed here: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_EC2.html
```

### Security
---
##### Security Group
```
Security group acts as a virtual firewall for EC2 instances and filters incoming and outgoing traffic based on a set of
rules. When an instance is launched, one or more security groups can be specified. If not specified, a default security
group for the VPC will be used. When EC2 decides whether to allow traffic to reach the instance, it evaulates all rules 
from all security groups associated with the instance.
```

##### Rule Exceptions
```
If a request is sent from an instance, the response for that request is allowed to reach the instance regardless of 
inbound security group rules. Similarly, responses to allowed inbound traffic are allowed to leave the instance
regardless of outbound rules.
```