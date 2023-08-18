### EC2 Basics
---
##### EC2 Properties
```
All EC2 properties are listed here: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_EC2.html.
```

##### EC2 APIs
```
All EC2 APIs are listed here: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Welcome.html.
```

### Security
---
##### Security Group
```
a) Security group acts as a virtual firewall for EC2 instances and filters incoming and outgoing traffic based on a set 
   of rules (e.g. allowed protocols, ports, source/destination CIDR, source/destination security groups). 
b) When an instance is launched, one or more security groups can be specified. If not specified, a default security 
   group for the VPC will be used. When EC2 decides whether to allow traffic to reach the instance, it evaulates all 
   rules from all security groups associated with the instance.
```

##### Security Rule Exceptions
```
If a request is sent from an instance, the response for that request is allowed to reach the instance regardless of 
inbound security group rules. Similarly, responses to allowed inbound traffic are allowed to leave the instance
regardless of outbound rules.
```