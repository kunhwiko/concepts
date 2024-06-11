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

##### Stop / Terminate / Hibernate
```
Stop
  * When stopped, instances retain their instance IDs, private IPv4 address, and data from attached EBS volumes.
  * When stopped, instances do not retain the host machine, data from local disk or RAM, and public IPv4 address.   
  
Terminate
  * When terminated, the instance along with all previous data including EBS volumes are deleted.

Hibernate
  * When hibernated, the state of RAM is saved into the root EBS volume and the root volume is persisted. This leads to
    being able to persist previous state and faster boot time as loading the preserved RAM means the OS can expedite 
    initialization tasks (e.g. hardware detection, filesystem checks etc.).
  * The EBS volume needs to have enough space to be able to store the state of RAM.
  * When hibernated, instances do not retain the host machine and public IPv4 address. 
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
1 or 3 year commitment to a certain amount of usage for a discounted price. Any usage beyond the commitment will be 
charged at the on-demand price.
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
##### Amazon Machine Image (AMI)
```
AMI is an image that provides info on how an instance should be launched. This includes one or more EBS snapshots, 
templates for the instance's root volume (e.g. OS), launch permissions to control which AWS accounts can use the AMI, 
and a block device mapping that specifies the volumes to attach to the instance during launch.
```

##### User Data
```
Script that is used when EC2 instances bootstrap. The user data script is executed as root user (i.e. sudo rights).
```

##### Key Pair
```
Key pairs allow users to connect to instances securely via SSH. Note that the instance's security group needs to allow
for incoming SSH traffic. Also note that SSH requires the instances to have a public IP unless a VPN exists.
```

##### Instance Profiles
```
Instance profiles are similar to IAM users but are intended for EC2 instances. They define "who" the EC2 instances are
and can assume at most 1 IAM role that defines what privileges the profile has. 
```

### Storage
---
##### Root Volume
```
When an instance is launched, a root volume is attached. Each instance has a single root volume that contains the OS and 
system files to boot the device. It is generally recommended to use EBS backed root volumes due to persistence.
```

##### Elastic Block Store (EBS)
```
a) EBS is a network based storage for EC2 instances and persists after instance termination. EBS volumes are bound to an
   availability zone and can be dynamically attached to instances in the same zone. EBS volumes are attached to one 
   instance at a time, but some volume types support multiple attachments.
b) Snapshots can be taken to make a backup of EBS volumes at a certain point in time and can be used to create volumes.
   These snapshots can be copied across zones and regions.
c) If encrypted, data at rest, any in-transit data, snapshots, and all volumes created by snapshots are encrypted.
d) EBS supports various volume types such as I/O optimized, throughput optimized etc.
```

##### Elastic File System (EFS)
```
a) EFS is a network based storage that can be mounted to multiple EC2 instances in multiple zones. EFS scales automatically
   with a pay-per-use model, so capacity does not need to be pre-provisoned. In general, EFS has a higher cost than EBS.
b) EFS supports various volume types such as I/O optimized, throughput optimized etc. It is possible to set lifecycle 
   management policies to move files to lower cost storage classes for files that are infrequently accessed (e.g. 
   infrequent-access tier, archive tier).
b) Security groups are used to control access to EFS.
```

##### Instance Store
```
Instance stores are non-network based hardware disks and generally have better I/O performance than network based 
volumes. These volumes however are terminated when an instance is stopped.
```

### Networking
---
##### Placement Groups
```
a) Cluster:   Packs instances close together inside an availability zone to achieve low-latency network communication.
b) Partition: Spreads instances across logical partitions such that groups of instances in one partition do not share 
              underlying hardware with groups of instances in different partitions.
c) Spread:    Strictly places individiual instances across distinct hardware to reduce correlated failures. The max 
              number of instances for each group (i.e. rack) per available zone is 7. 
```

##### Elastic IP
```
a) When an EC2 instance is stopped and then started, the public IP will change. If a fixed IP is required, an Elastic IP 
   needs to be configured. An Elastic IP is a public IPv4 IP address that one own's as long as it is not deleted.
b) Elastic IPs can be attached to a single instance or network interface at a time. 
```

##### Elastic Network Interface (ENI)
```
a) Logical component representing a virtual network card. An instance will come with a default non-detachable ENI and 
   can have multiple ENIs additionally attached.
b) ENIs are bound to a specific subnet and can be dynamically attached and detached onto different EC2 instances in 
   that subnet.
c) Each ENI can have a primary private IPv4 address from the IPv4 address range of the VPC. It can then have one or more 
   secondary private IPv4 addresses from the IPv4 address range of the VPC.
d) Each ENI can have one public IPv4 address, one MAC address, and one or more security groups. One Elastic IP can
   be associated with one of the private IPv4 addresses of the ENI.
```

### Scalability and High Availability
---
##### Elastic Load Balancing (ELB)
```
AWS provides a managed load balancing service that can automatically scale based on traffic. AWS provides the following
load balancer types for elastic load balancing (ELB):
  * Application Load Balancer (ALB)
  * Network Load Balancer (NLB)
  * Gateway Load Balancer (GWLB)
```

##### Target Groups
```
Target groups are a grouping of individually registered targets (e.g. EC2 instances, Lambda functions, private IPs). 
Each target group can perform health checks to ensure traffic is sent to healthy targets.
```

##### Application Load Balancer (ALB)
```
a) Layer 7 load balancer that can route traffic to various target groups. Operating at layer 7, ALBs support routing 
   based on host, path, query parameters, and headers and supports authentication at the load balancer level. 
b) As a layer 7 load balancer, ALB supports rerouting based on rule conditions (e.g. path, headers, Source IP) and can 
   perform custom logic on incoming traffic based on those rules (e.g. forward to traffic group, traffic redirect, 
   return error).         
c) ALBs come with a hostname that can be connected via DNS resolution. Note that underlying IPs can change.
d) ALBs come with security groups that can be referenced and whitelisted by security groups of EC2 instances.
```

##### Network Load Balancer (NLB)
```
a) Layer 4 load balancer that is highly performant compared to ALB (e.g. latency performance). As a layer 4 load 
   balancer, NLBs support routing to various target groups based on ports and protocols, but do not support routing 
   based on URLs and custom rerouting logic. It also does not support authentication at the load balancer layer.
b) Like ALBs, NLBs support routing to target groups, come with security groups, and are assigned a hostname.
c) NLBs have a single static IP per availability zone and supports assigning Elastic IPs.
```

##### Gateway Load Balancer (GWLB)
```
Layer 3 load balancer that is commonly used to send traffic to a target group of security applications to inspect 
incoming traffic (e.g. deep packet inspection, payload manipulation) before routing to the destination. Traffic that 
satisfies security requirements can be forwarded back to the GWLB and to the destination while those that do not meet 
requirements are dropped.  
```

##### Auto Scaling Group (ASG)
```
a) Means to automate the scaling of EC2 instances based on demand. ASGs can recreate instances that fail health checks
   and can also be integrated with CloudWatch to scale based on CloudWatch alarms. After an instance is scaled or
   terminated, there is generally a cooldown period before the next scaling action can occur.
b) ASGs can be associated with ELBs and can allow new instances to automatically be registered to target groups.
c) ASGs are configured with launch templates, which define attributes such as instance type, AMI, user data, EBS volumes, 
   SSH key pair, IAM roles, security groups, and subnet information.
d) ASGs have various scaling policies that allow scaling based on resource usage, alarms, time intervals, or predictions 
   based on historical data.
```

### Security
---
##### Security Group
```
a) Security group acts as a virtual firewall for EC2 and defines what inbound and outbound traffic are allowed based on
   rules (e.g. allowed protocols, port range on instance, source/destination CIDR or security groups). 
b) Security groups only contain allow rules. All inbound traffic is blocked by default and all outbound traffic is
   allowed by default.
c) When an instance is launched, one or more security groups can be specified. If not specified, a default security 
   group per VPC will be used. When EC2 decides whether to allow traffic to reach the instance, it evaluates all rules 
   from all security groups associated with the instance.
```

##### Security Rule Exceptions
```
If a request is sent from an instance, the response for that request is allowed to reach the instance regardless of 
inbound security group rules. Similarly, responses to allowed inbound traffic are allowed to leave the instance
regardless of outbound rules.
```

### 