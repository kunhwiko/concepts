### EC2 Basics
---
##### Elastic Compute Cloud (EC2)
```
EC2 provides scalable computing capacity in AWS. It allows to launch necessary virtual servers, scale based on
requirements, configure security/networking, and manage storage.
```

##### EC2 Properties & APIs
```
All EC2 properties are listed here:
  * https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_EC2.html
All EC2 APIs are listed here: 
  * https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Welcome.html
```

### EC2 Purchase Options
---
##### On-Demand Instance & Capacity Reservation
```
a) Typically used for short-term workloads and are a pay per second after the first minute model.
b) There is an option for "Capacity Reservation", which reserves on-demand capacity in a specific zone for any duration.
```

##### Reserved Instances
```
a) Model to reserve a certain quantity of instances of specific attributes (e.g. type, OS) with a reservation period of 
   1 or 3 years at a discounted price. Instances can be reserved at either a regional or zonal scope.
b) Payment options can vary from paying upfront, partially upfront, or none upfront for further discounts.
c) Excess reserved instances can be bought or sold on the Reserved Instance Marketplace.
d) There is an option for "Convertible Reserved Instances" which allow for flexibility to change instance types, family, 
   OS, and zone/region scope.
```

##### Savings Plan
```
a) 1 or 3 year commitment to a certain amount of usage for a discounted price.
```

##### Spot Instance & Spot Fleet
```
a) Model to use unused EC2 capacity at a discounted price but can be interrupted at any time. Customers specify a max 
   spot price and are assigned instances that match that price. If the current spot prices exceed the customer specified 
   max spot price, the instances will be terminated.
b) Spot fleets allow customers to specify price constraints, max capacity, set of possible launch pools (e.g. instance 
   type, OS, zone), and strategy where AWS will then optimally bring up a fleet of spot instances from those pools:
     * Lowest Price: Choose from lowest priced pool that has available capacity
     * Diversified: Distribute from different pools
     * Capacity Optimized: Choose from pool with instances based on capacity remaining (i.e. low chance of interruption) 
     * Price Capacity Optimized: Selects pools with highest capacity available, and then selects pools with lowest price 
```

##### Dedicated Host & Instance
```
a) Dedicated Hosts is a model to book an entire physical server, typically due to compliance requirements. In this model,
   AWS allows users to control instance placements and provides full visibility about the server (socket, cores, host 
   ID) which could be required for custom software licenses.
b) There is an option for "Dedicated Instances" which is used to book an entire dedicated hardware. The physical server 
   will vary per instance startup as AWS will look for hardware that doesn't have VMs running. EC2 instances in the same
   account may share the same dedicated hardware.
```

### EC2 Configurations
---
##### User Data
```
Script that is used when EC2 instances bootstrap. The user data script is executed as root user (i.e. sudo rights).
```

##### Key Pair
```
Key pairs allow users to connect to instances securely via SSH. Note that the instance's security group needs to allow
for incoming SSH traffic.
```

##### Instance Profiles
```
Instance profiles are similar to IAM users but are intended for EC2 instances. They define "who" the EC2 instances are
and can assume at most 1 IAM role that defines what privileges the profile has. 
```

### Security
---
##### Security Group
```
a) Security group acts as a virtual firewall for EC2 and filters incoming and outgoing traffic based on a set of rules 
   (e.g. allowed protocols, port range on instance, source/destination CIDR, source/destination security groups). 
b) When an instance is launched, one or more security groups can be specified. If not specified, a default security 
   group per VPC will be used. When EC2 decides whether to allow traffic to reach the instance, it evaulates all rules 
   from all security groups associated with the instance.
```

##### Security Rule Exceptions
```
If a request is sent from an instance, the response for that request is allowed to reach the instance regardless of 
inbound security group rules. Similarly, responses to allowed inbound traffic are allowed to leave the instance
regardless of outbound rules.
```