### Basic Networking Components
---
##### Networks
```
System of links that interconnect computers to move data.
```

##### Protocols
```
Set of rules and structures that define the semantics of how computers communicate.
Below are some examples of protocols.
  * IP   : Address of where packets come from and where they should be sent.
  * TCP  : Responsible for breaking data into packets, delivering / reassembling packets, and checking for corruption.
  * HTTP : Set of rules for how request-response works in the web.
```

##### MAC Address
```
a) Determines "who" the machine is.
b) Physical unique address of a machine represented in 48 bits.
   It is difficult for the Internet to keep track of where all MAC addresses are.
c) For network broadcasts, the destination MAC address will typically look like ffff.ffff.ffff.
```

##### IP Address
```
a) Determines "where" the machine is.
b) IP addresses are 32 bits and are hierarchically designed by location/teams through subnetting.
c) IP addresses will efficiently narrow down from L3 to the correct L2 network, where computers can then be located through MAC addresses.
```

##### Ports
```
a) Docking point for where information is received or sent.
b) Means for how multiple programs can listen for new network connections on the same machine without collision.
c) IP address is like a mailbox to an apartment complex, and ports are the specific apartment number.
```

### OSI Model
---
##### Network Layers
```
Layer 7) Application 
Layer 6) Presentation 
Layer 5) Session
Layer 4) Transport
Layer 3) Network
Layer 2) Datalink
Layer 1) Physical 
```

##### Layering
```
Layering
  a) Ability to mix and match different protocols.

Protocol Stack
  a) Set of protocols that are currently in use, but can be mixed and matched for different situations.
     For example, in the case of emails, network can switch HTTP protocol layer for SMTP protocol layer.
```

##### Encapsulation
```
Step 1) When sending over messages, headers are 'encapsulated' starting from highest to lowest layer.
Step 2) For each layer, the lower layer "wraps" the message coming from higher levels to create protocol stacks.
        Lower layers do not have to actually care about what these higher level layers do.
Step 3) After sending over a message, switches and routers will decapsulate the message from lowest to highest layer.
        As an example, a switch might identify from the layer 2 header that the message was intended for that switch.
Step 4) The message is then recapsulated before resending and the steps are repeated.
        As an example, a switch could write a new layer 2 header specifying a new src/dest MAC address.
```

### Layer 1
---
##### Physical Layer
```
a) Physical links that transfer message bits.
b) Technologies include coaxial cables, fiber cables, ethernet, wifi, repeaters, hubs.
```

##### Repeater
```
Problem Statement
  a) As data is transmitted over network wires, it decays as it travels.

Solution
  a) Repeaters regenerate signals in between network traffic.
     This allows for network communication across greater distances.
```

##### Hub
```
Problem Statement
  a) If a new host joins the network, it needs to be connected with all the existing hosts in that network.
     Directly connecting the new host to existing hosts through wires works, but is not a scalable solution.

Solution
  a) Hubs act as a multi-port repeater that all existing and new hosts in the network can connect to.
  b) Data sent from one host is broadcasted to all hosts connected to the hub.
```

### Layer 2
---
##### Data Link Layer
```
a) Data link layer bundles physical layers into a Local Area Network (LAN).
b) Data link layer is responsible for correct network hops and putting/receiving bits into the physical layer.
c) Data link layer abstracts stream of 'bits' into 'frames' and use MAC addresses as an address scheme.
   The header for Layer 2 holds the src/dest MAC addresses.
d) Technologies include network interface cards (NICs), wifi access cards, bridges, and switches.
```

##### Network Interface Card (NIC)
```
a) NICs are a hardware component that allows for network connectivity.
b) NICs hold a unique MAC address for the host machine.  
```

##### Bridge
```
Problem Statement
  a) Hubs cause all hosts connected to the hub to know about data that they do not need to be involved in.

Solution
  a) Bridges sit between two hubs and have two ports each for one of the hubs.
  b) Bridges know which hosts are on what side of either port.
  c) When a host sends data and the destination is connected on the same hub, the bridge will prevent transmitting data to the other hub. 
```

##### Switch
```
Problem Statement
  a) All hosts on one side of a bridge still receive data that they might not be involved in.
  b) All hosts on both sides of the bridge will receive data if the source and destination are on opposite sides of the bridge.

Solution
  a) Switches are a combination of hosts and bridges that help to connect L1 networks to form an L2 network.
     All hosts in the same L2 network will share a common IP address space (prefix).  
  b) Switches have multiple ports and know which hosts are on each port through a 'forwarding table'.
  c) Hosts broadcast themselves to the network, and switches learn the existence of new MAC addresses through this broadcasting.

Forwarding Table
  a) Maps a MAC address to a port to forward packets.
```

##### Virtual Ethernet (veth)
```
a) Virtual implementation of Ethernet that allow for communications across logical partitions or different virtual machines without the need of assigning physical hardware.
b) Devices are typically created in pairs and are connected via a bridge.
```

### Layer 3
---
##### Network Layer 
```
a) Network layer bundles data link layers to Internet.
b) Network layer abstracts 'frames' into 'packets' and use IP addresses as an address scheme.
   The header for Layer 3 holds the src/dest IP addresses.
c) Technologies include routers.
```

##### Router
```
Routers
  a) Routers connect L2 networks to form L3 networks and are assigned an IP address from each L2 network it is connected to.
     This allows routers to facilitate communication between networks and serve as a 'gateway', or a means for traffic to travel outside an L2 network.
  b) Routers are knowledgable of the L2 networks connected to itself through a 'routing table'.
  c) Routers typically have multiple NICs and MAC addresses as each network interface (e.g. ports) requires a MAC address.
  d) Routers help form the IP address hierarchy in networks and subnets. 
  e) Routers provide a control point for network traffic where security, filtering, redirecting can be enforced.

Routing Table
  a) Routing table holds information about all the networks a router is aware about.
  b) Routing table maps longest prefix matches on IPs to send packets over to next hops (other L2 networks).

More info here: https://www.youtube.com/watch?v=H7-NR3Q3BeI&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=2
```

##### Maximum Transmission Unit (MTU)
```
MTU
  a) Determines the largest data packet that can be accepted. 
     Packets larger than the limit go through "fragmentation", which can lead to additional network latency.
  b) Larger MTU size means more data can fit in, resulting in a faster and more efficient transmission.
  c) Larger packets are prone to corruption and delays, so packet arrival could end up being relatively slow.
     If an error occurs, larger packets also take longer to retransmit.

Fragmentation
  a) Routers check the size of each IP packet and the MTU of the next router to receive the packet.
     If the packet is too big, packets are broken up with copies of the packet header.
```

##### IPv4 vs IPv6
```
IPv4
  a) Consists of 32 bits that can be broken into two parts (network and host addresses).
  b) Network addresses are common addresses to groups of host addresses (geography / company).

IPv6: 
  a) More addresses and better functionality compared to IPv4.
  b) Release for IPv6 is delayed as the entire network must become IPv6 compatible.
  c) Rely more on alternative solutions such as NATs for now.
```

### Discovery
---
##### Process of Discovery
```
host name (www.github.com) --(DNS)--> IP Address --(ARP)--> MAC address 

Step 1) Host begins only knowing its source MAC address.
Step 2) Host discovers its source IP address and DNS servers via DHCP.
Step 3) Host discover the destination's IP address via DNS.
Step 4) Perform a longest prefix match on the IP to determine if the message should go to a router or a local machine.
Step 5) Move to a router (another L2 network) or a local machine (current L2 network) via ARP. 
```

##### DHCP Server
```
a) Host broadcasts a DHCP discovery ping --> DHCP servers respond with an offer response --> host sends a request message specifying the DHCP server it will use.
b) DHCP helps a machine learn its own IP address, IP addresses of local DNS servers, gateway routers, and prefix length.
```

##### DNS Server
```
a) DNS servers act as a phonebook for finding the IP addresses of various sites.
b) DNS servers are replicated for availability, and caches popular addresses.
```

##### A and CNAME Records
```
A Records
  a) A mapping of a hostname to one or more IP addresses.

CNAME Records
  a) An alias from one domain to another.
``` 

##### Address Resolution Protocol (ARP)
```
Protocol used to translate IP addresses into MAC addresses. Hosts preserve these mappings on ARP tables.

Step 1) When a client sends a request, it will know the destination IP but not the MAC address.
        This means the request cannot be sent as the L2 header is incomplete.
Step 2) Client fires an ARP request that sends a broadcast that holds the source's IP and MAC address.
        This broadcast looks for a host with the destination IP address.
        Layer 2 header for the broadcast will carry a destination MAC address of 'broadcast MAC address'.
Step 3) Destination host receives the ARP broadcast and is able to update its ARP table with <src-ip>:<src-mac>.
Step 4) Destination host fires an ARP response that sends a unicast that holds its IP and MAC address.
Step 5) Client receives the ARP response and updates its ARP table.
Step 6) Client is able to send its original request as it is now able to complete an L2 header.
```

### Subnets
---
##### Classless Interdomain Routing (CIDR)
```
a) Example 1: 128.168.1.0/26
     * 26 bits for network addresses, 6 bits for host addresses.
     * Great if you have 50 or less computers in a common network.
b) Example 2: 192.168.1.0/24
     * 24 bits for network addresses, 8 bits for host addresses (192.168.1.0 ~ 192.168.1.255).
c) Routers do a longest prefix match on these "prefixes" to route packets to the next hop.
```

##### Subnets
```
Subnets
  a) Networks inside networks that represents a segmented piece of the larger network.
  b) Narrows down and groups various network devices into the same IP address range.
  c) Means to send packets in a time efficient way.
     Send mail --> route to office --> route to department --> route to team --> route to person.
     The above does not waste time trying to get the mail to the individual person directly.
  d) For more: https://www.cloudflare.com/learning/network-layer/what-is-a-subnet/  

Subnet Mask
  a) A number (e.g 255.255.255.0) that helps distinguish the network address and host address for an IP address.
  b) This number helps to determine whether a destination IP address lies in the same L2 network.
```

##### Network Address Translation (NAT)
```
NAT Steps
  Step 1) Assign private IP addresses within a common network.
  Step 2) NAT devices maps private IP addresses to a single public IP address.
  Step 3) Send to destination address with mapped public IP address as the source.
  Step 4) When host at destination tries to send a packet to the source within the subnet, use demultiplexing IDs (L4 protocol) to find the source.

NAT Disadvantages
  a) Difficult to distinguish devices within the subnet.
  b) Cannot connect to hosts within the private network (address unknown) from the outside until a packet has been trasmitted.
```