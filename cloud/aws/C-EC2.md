### Overview
---
##### Elastic Compute Cloud (EC2)
```
EC2 provides scalable computing capacity in AWS. It allows users to launch virtual servers based on OS, CPU, RAM, 
storage (e.g. network-attached vs hardware based), network (e.g. speed of network card), and security (e.g. public access)
requirements.
```

##### Instance State
```
a) Stop
- When stopped, instances retain their instance IDs, private IPv4 address, and data from attached EBS volumes.
- When stopped, instances do not retain the host machine, data from local disk or RAM, and public IPv4 address.   
  
b) Terminate
- When terminated, the instance along with all previous data including EBS volumes are deleted.

c) Hibernate
- When hibernated, the state of RAM is saved into the root EBS volume and the root volume is persisted. This leads to
  being able to persist previous state and faster boot times as loading the preserved RAM means the OS can expedite 
  initialization tasks (e.g. hardware detection, filesystem checks etc.).
- The EBS volume needs to have enough space to be able to store the state of RAM.
- When hibernated, instances do not retain the host machine and public IPv4 address. 
```

##### Placement Groups
```
a) Cluster  
- Packs instances close together inside an availability zone to achieve low-latency network communication.

b) Partition
- Spreads instances across logical partitions such that groups of instances in one partition do not share underlying
  hardware with groups of instances in different partitions.
  
c) Spread
- Strictly places individiual instances across distinct hardware to reduce correlated failures. The max number of
  instances for each group (i.e. rack) per available zone is 7. 
```

### Purchase Options
---
##### On-Demand Instance & Capacity Reservation
```
Typically used for short-term workloads and are a pay per second after the first minute model. There is an option for 
"Capacity Reservation", which reserves on-demand capacity at no discounts in a specific zone for any duration.
```

##### Reserved Instances
```
Reserved instances is a plan to reserve a certain quantity of instances of specific attributes (e.g. type) with a 
reservation period of 1 or 3 years at a discounted price. Instances can be reserved at either a regional or zonal scope.

Payment options for reserved instances can vary from paying upfront, partially upfront, or none upfront. Excess reserved 
instances can also be bought or sold on the Reserved Instance Marketplace.

The option "Convertible Reserved Instances" allows for flexibility to change instance types, family, OS, and the 
zone/region scope.
```

##### Savings Plan
```
1 or 3 year commitment to a certain amount of usage on a particular instance family and region for a discounted price. 
Any usage beyond the commitment will be charged at the on-demand price.
```

##### Spot Instance & Spot Fleet
```
Spot instances utilize unused EC2 capacity at a discounted price but can be interrupted at any time. Customers specify 
a max spot price and are assigned instances that match that price. If the current spot prices exceed the customer 
specified max spot price, the instances will be terminated.

Spot Fleets allow customers to specify price constraints, max capacity, set of possible launch pools (e.g. instance 
type, OS, zone), and strategy where AWS will optimally bring up a fleet of spot instances based on specified options:
- Lowest Price: Choose from lowest priced pool that has available capacity
- Diversified: Distribute from different pools
- Capacity Optimized: Choose from pool with instances based on capacity remaining (i.e. lowest chance of interruption) 
- Price Capacity Optimized: Diversifies pools across low priced pools with high capacity availability
```

##### Dedicated Host & Instance
```
Dedicated Hosts is a model to book a physical server, typically due to compliance requirements. In this model, AWS allows 
users to control instance placements and provides full visibility about the server (socket, cores, host ID).

Dedicated Instances is a model used to book dedicated hardware. AWS will look for hardware that doesn't have VMs running, 
so the physical server may vary per instance startup. EC2 instances in the same account may share the same dedicated hardware.
```

### Configurations
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

### Networking
---
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

### Security
---
##### Security Group
```
Security groups act as a virtual firewall for EC2 and allowlists inbound and outbound traffic based on a set of rules
(e.g. allowed protocols, port range on instance, source/destination CIDR or source/destination security groups). Security 
groups only contain allow rules.

When an instance is launched, one or more security groups can be specified. If not specified, a default security group 
per VPC will be used. When EC2 decides whether to allow traffic to reach the instance, it evaluates all rules from all 
security groups associated with the instance.

All inbound traffic is blocked by default and all outbound traffic is allowed by default. If a request is sent from an 
instance, the response for that request is allowed to reach the same instance regardless of inbound security group rules. 
Similarly, responses to allowed inbound traffic are allowed to leave the same instance regardless of outbound rules.
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
a) EBS is a SAN (storage area network) based storage that connects multiple storage devices to instances through a 
   dedicated fiber network. Due to this direct networking, security groups are not required for EBS volumes. 
b) EBS supports various volume types such as general purpose, I/O optimized, and throughput optimized.
c) EBS volumes are bound to an availability zone and can be dynamically attached to instances in the same zone. EBS 
   volumes are attached to one instance at a time, but some volume types support multiple attachments.
d) Snapshots can be taken to make a backup of EBS volumes at a certain point in time and can be used to create volumes.
   These snapshots can be copied across zones and regions.
e) If encryption is enabled, data at rest, in-transit data, snapshots, and volumes created by snapshots are encrypted.
```

##### Elastic File System (EFS)
```
a) EFS is a network attached storage that can be mounted to multiple EC2 instances in multiple zones. Network is attached 
   through a VPC, so security groups are used to control access.
b) EFS scales automatically with a pay-per-use model, so capacity does not need to be pre-provisoned. In general, EFS 
   has a higher billing cost than EBS. It is possible to set lifecycle management policies to move files to lower cost 
   storage classes for files that are infrequently accessed (e.g. infrequent-access tier, archive tier).
c) EFS supports various volume types such as standard and infrequent access.
```

##### Amazon File Server (FSx)
```
a) FSx allows users to launch third party file systems such as Windows File Server (primarily for Windows containers), 
   Lustre (typically used for high performance computing), NetApp ONTAP, and OpenZFS on AWS.
b) Lustre can be deployed as a scratch file system where data is not replicated but has higher burst performance, or as a
   persistent file system where data is replicated within the same zone.
```

##### Instance Store
```
Instance stores are non-network based hardware disks and generally have better I/O performance than network based 
volumes. These volumes however are terminated when an instance is stopped.
```

### Load Balancing and High Availability
---
##### Elastic Load Balancing (ELB)
```
AWS provides a managed load balancing service that can automatically scale based on traffic. AWS provides the following
load balancer types for elastic load balancing (ELB):
  * Application Load Balancer (ALB)
  * Network Load Balancer (NLB)
  * Gateway Load Balancer (GWLB)
```

##### Features of Elastic Load Balancers
```
a) Connection Drain: Otherwise known as deregistration delay, is a feature that allows designated instances to complete 
   in-flight requests but causes the ELB to stop sending new requests to deregistered instances.
b) Sticky Sessions: Otherwise known as session affinity, is a feature that allows a client's session to be bound to a 
   specific instance within a target group. For ALBs, session affinity is implemented in the form of cookies that are 
   configured at the target group level and can be application-based or duration-based.
c) Cross Zone Load Balancing: A feature that allows ELBs to distribute traffic evenly across all registered instances
   even if those instances are in a different availability zone.
d) TLS Encyrption: Allows connections to be encrypted between the client and the load balancer. ELBs can host multiple
   SSL certificates that can be managed by AWS Certificate Manager (ACM). To support multiple SSL certificates, SNI is
   used as part of the TLS handshake.
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
c) ALBs come with a hostname that can be connected via DNS resolution. The IP address of load balancers can change.
d) ALBs come with security groups that can be referenced and whitelisted by security groups of EC2 instances.
```

##### Network Load Balancer (NLB)
```
a) Layer 4 load balancer that is highly performant compared to ALB (e.g. latency performance). As a layer 4 load 
   balancer, NLBs support routing to various target groups based on ports and protocols, but do not support routing 
   based on URLs and custom rerouting logic. It also does not support authentication at the load balancer layer.
b) Like ALBs, NLBs support routing to target groups, come with security groups, and are assigned a hostname.
c) NLBs have a single static IP per availability zone and can be assigned Elastic IPs.
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
