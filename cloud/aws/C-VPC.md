### Observability
---
##### Flow Logs
```
Flow logs can be used to capture info on IP traffic that pass network interfaces in a VPC or subnet. These logs are
sent to AWS CloudWatch.
```

### Networking
---
##### AWS VPC
```
a) VPC is a virtual network that resembles a traditonal network that one would operator in their own data center. After
   the VPC is created:
     * Subnets can be created in different availability zones to deploy AWS resources.
     * Gateways can be created to connect VPCs to other networks.
     * Route tables can be used to determine where network traffic from a subnet or gateway should be directed.
     * VPC endpoints can be created to connect to AWS services privately without using a gateway or NAT device.
b) When VPCs are created, the maximum and minimum IPv4 CIDR block are /16 and /28. The following IP addresses in each
   subnet are reserved and cannot be allocated:
     * xxx.xxx.xxx.0: network address
     * xxx.xxx.xxx.1: reserved for VPC router
     * xxx.xxx.xxx.2: reserved for DNS server
     * xxx.xxx.xxx.3: reserved for future use 
     * last IP addr : broadcast address
```

##### Internet Gateway (IGW)
```
Internet gateways allow traffic into and from the public internet. To send traffic to the internet, subnet route tables 
should be configured to forward packets to the gateway. To receive traffic, instances within the subnet must have a 
public IP address. Note that network ACLs should not be blocking traffic.   
```

##### Virtual Private Gateway (VGW)
```
Virtual private gateways can be used to establish VPN endpoints to other networks (e.g. private data centers, corporate
network) for secure communication.
```

##### NAT Gateway
```
NAT gateways allow instances within a private subnet to communicate to the internet while keeping IP addresses hidden.
This is accomplished by translating private IP addresses to a single public IP address.
```

##### VPC Endpoints
```
VPC endpoints allow connection to AWS services without the need to create gateways of any kind. 
```

##### Route Tables
```
Route tables determine where network traffic should be sent. A non-deletable main table is created by default for the 
VPC and will be the default table for subnets that do not have a custom route table. Note that a subnet can only have 
one table at a time.   
```

##### VPC Peering
```
VPC peering can be configured to route traffic between resources in different VPCs. Note the following limitations:
  * If there are overlapping CIDR ranges, VPCs cannot be peered.
  * If VPC1 and VPC2 are peered, and VPC2 and VPC3 are peered, this does not mean VPC1 and VPC3 are peered.
```

### Security
---
##### Network Access Control List (ACL)
```
a) Allows or denies inbound or outbound traffic at the subnet level based on rules similar to security group rules. The
   network ACL will be triggered as traffic enters and leaves the subnet, not as it is routed within the subnet.
b) An ACL can be associated with multiple subnets but a subnet cannot be associated with more than one ACL at a time. If 
   not specified, a default network ACL for the VPC will be associated with the subnet. 
c) Network ACLs do not follow security rule exceptions, which means responses to allowed inbound traffic are subject
   to rules for outbound traffic.
```