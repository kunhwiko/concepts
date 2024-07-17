### Basic Networking Components
---
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
c) For broadcast frames, the destination MAC address will typically look like ffff.ffff.ffff.
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
  * Ability to mix and match different protocols.

Protocol Stack
  * Set of protocols that are currently in use, but can be mixed and matched for different situations. For example, in 
    the case of emails, network can switch HTTP protocol layer for SMTP protocol layer.
```

##### Encapsulation
```
Step 1) When sending over messages, headers are "encapsulated" starting from highest to lowest layer.
Step 2) For each layer, the lower layer wraps the message coming from higher levels to create protocol stacks.
        Lower layers do not have to actually care about what these higher level layers do.
Step 3) After sending over a message, switches and routers will decapsulate the message from lowest to highest layer.
        As an example, a router at this time will identify from the layer 2 header that the message was intended for itself.
        It will then discard the L2 header and pass for Layer 3 header validation.
Step 4) The message is then recapsulated before resending and the steps are repeated.
        As an example, a router at this time will write a new layer 2 header specifying a new src/dest MAC address.
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
  * As data is transmitted over network wires, it decays as it travels.

Solution
  * Repeaters regenerate signals in between network traffic. This allows for network communication across greater 
    distances.
```

##### Hub
```
Problem Statement
  * If a new host joins the network, it needs to be connected with all the existing hosts in that network. Directly 
    connecting the new host to existing hosts through wires works, but is not a scalable solution.

Solution
  * Hubs act as a multi-port repeater that all existing and new hosts in the network can connect to.
  * Data sent from one host is broadcasted to all hosts connected to the hub.
```

### Layer 2
---
##### Data Link Layer
```
a) Data link layer bundles physical layers into a Local Area Network (LAN).
b) Data link layer is responsible for correct network hops and putting/receiving bits into the physical layer.
c) Data link layer abstracts stream of 'bits' into 'frames' and use MAC addresses as an address scheme. The header for 
   Layer 2 holds the src/dest MAC addresses.
d) Technologies include network interface cards (NICs), wifi access cards, bridges, and switches.
```

##### Network Interface Card (NIC)
```
NICs are a hardware component that allows for network connectivity. Each NIC hold a unique MAC address for the host 
machine.
```

##### Bridge
```
Problem Statement
  * Hubs cause all hosts connected to the hub to know about data that they do not need to be involved in.

Solution
  * Bridges sit between two hubs and have two ports each for one of the hubs.
  * Bridges know which hosts are on what side of either port.
  * When a host sends data and the destination is connected on the same hub, the bridge will prevent transmitting data 
    to the other hub. 
```

##### Switch
```
Problem Statement
  * All hosts on one side of a bridge still receive data that they might not be involved in.
  * All hosts on both sides of the bridge will receive data if the source and destination are on opposite sides of the 
    bridge.

Solution
  * Switches are a combination of hubs and bridges that help to connect L1 networks to form an L2 network. All hosts in 
    the same L2 network will share a common IP address space (i.e. prefix). 
  * Switches have multiple ports and know which hosts are on each port through a 'MAC address table'. Switches will use
    this table to determine which host to forward requests to.

MAC Address Table
  * A table that maps ports to MAC addresses of connected hosts.

Switch Functionality
  * Learn: Switches update their MAC address table with a <src-mac>:<port> mapping when a new frame passes the switch.
  * Forward: Use mapping on MAC address table to send frame on the appropriate port.
  * Flood: When a destination MAC address is not found on the MAC address table, the switch duplicates the frame to all 
           hosts except to the receiving port. Irrelevant hosts will drop the request and only the relevant host will 
           send a response back, which again causes an update on the MAC address table.
```

### Virtualization of Layer 2
---
##### Virtual Local Area Network (VLAN)
```
Problem Statement
  * Before virtualization, isolated switches were required for each and every network that needed to be isolated.

Solution
  * Switches can be logically separated through virtualization. Ports on a switch can be grouped into isolated 
    "mini-switches".
  * VLANs allow a single physical switch to be split into multiple virtual switches. VLANs allow a single virtual switch 
    to be extended across other physical switches.
```

##### Trunk Ports
```
Problem Statement
  * Assume that VLAN 1 and VLAN 2 are extended across the same 2 physical switches. This would normally require 2 wire 
    connections, one between ports for VLAN 1 and another between ports for VLAN 2. With more VLANs across physical 
    switches, this becomes difficult to scale.

Solution
  * Trunk ports allow data to flow from multiple VLANs in a single physical wire.
  * When data flows into a single wire, it becomes difficult to know which VLAN the data is intended for. "VLAN tags" 
    are added on top of existing L2 and L3 headers to distinguish which VLAN the packet is intended for.

Access Ports
  * Links that carry data just for a single VLAN.
  * When data flows into access ports, the network knows that the data is intended for a single VLAN.

Native VLAN
  * If a packet goes through a trunk port but does not have a VLAN tag, it will be sent to the Native VLAN as a default.
    This means that when sending a packet through a trunk port, the packet does not need a VLAN tag if the packet is 
    already intended for the Native VLAN. 

More info here: https://www.youtube.com/watch?v=MmwF1oHOvmg
```

##### Virtual Extensible Local Area Network (VXLAN)
```
Problem Statement
  * VLAN tags only support up to a maximum of 4096 (12 bits) following 820.1Q standards.

Solution
  * VXLANs encapsulate frames with a VXLAN header into UDP packets to resolve the inability for VLANs to be routed out 
    of L2 networks.
  * VXLANs are identified by a 24 bit VXLAN network identifier (VNI), allowing up to 16,777,216 VLANs.
```

##### Virtual Ethernet (VETH)
```
Virtualization of ethernet that act as tunnels between network namespaces without the need of assigning physical 
hardware. Virtual ethernet devices are typically created in pairs and are connected via a bridge.
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
a) Routers connect L2 networks to form L3 networks and are assigned an IP address from each L2 network it is connected to.
   This allows routers to facilitate communication between networks and serve as a 'gateway', or a means for traffic to travel outside an L2 network.
b) Routers are knowledgable of connected L2 networks through a 'routing table'.
c) Routers typically have multiple NICs and MAC addresses as each network interface (e.g. ports) requires a MAC address.
d) Routers help form the IP address hierarchy in networks and subnets. 
e) Routers provide a control point for network traffic where security, filtering, redirecting can be enforced.

Routing Table
  a) A table that maps IP prefixes to next hops.
     Both hosts and routers have a routing table.
  b) Routes on routing tables are updated through the following:
       * Direct Connection: Routers are aware of the networks they are directly attached to.
       * Static Routes: Routers are aware of routes that administrators manually seed in.
       * Dynamic Routes: Routers learn routes from other routers.

Routing
  * If there is one match on the routing table, packets are sent to the corresponding next hop.
  * If there are multiple matches on the routing table, packets are sent to next hop that has the longest subnet mask.
  * If there are no matches on the routing table, packets are sent to a default entry (i.e. match on 0.0.0.0/0) or dropped.

More info here: https://www.youtube.com/watch?v=H7-NR3Q3BeI&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=2
```

##### Maximum Transmission Unit (MTU)
```
MTU
  * Determines the largest data packet that can be accepted. Packets larger than the limit go through "fragmentation", 
    which can lead to additional network latency.
  * Larger MTU size means more data can fit in, resulting in a faster and more efficient transmission.
  * Larger packets are prone to corruption and delays, so packet arrival could end up being relatively slow. If an error 
    occurs, larger packets also take longer to retransmit.

Fragmentation
  * Routers check the size of each IP packet and the MTU of the next router to receive the packet. If the packet is too 
    big, packets are broken up with copies of the packet header.
```

##### IPv4 vs IPv6
```
IPv4
  * Consists of 32 bits that can be broken into two parts (network and host addresses).
  * Network addresses are common addresses to groups of host addresses (geography / company).

IPv6
  * More addresses and better functionality compared to IPv4.
  * Release for IPv6 is delayed as the entire network must become IPv6 compatible.
  * Rely more on alternative solutions such as NATs for now.
```

### Discovery
---
##### Process of Discovery
```
host name (www.github.com) --(DNS)--> IP Address --(ARP)--> MAC address 

Step 1) Host begins only knowing its source MAC address.
Step 2) Host looks for a DHCP server to discover its own IP address, IP address of DNS servers, IP address of default 
        gateway, and subnet mask.
Step 3) Host discovers the destination's IP address via DNS.
Step 4) Subnet mask determines if the message should go through a gateway router or if the destination is in the same L2 
        network. This is determined through entries in the host's routing table (i.e. 0.0.0.0/0 will map to IP address 
        of default gateway).
Step 5) Host fires an ARP broadcast to discover the destination MAC address of the default gateway or destination.
Step 6) If packet goes through default gateway, the L2 header is rewritten (i.e. L3 header is preserved) and the above 
        process is repeated.
```

##### Dynamic Host Configuration Protocol (DHCP)
```
DHCP servers helps a machine learn important information that is crucial to get packets into the Internet:
  * Host's IP address
  * DNS server's IP address
  * Default gateway's IP address
  * Subnet mask

Step 1) Host broadcasts a DHCP discovery ping.
Step 2) DHCP servers respond with an offer response.
Step 3) Host sends a request message specifying the DHCP server it will use.
```

##### Domain Name System (DNS)
```
DNS servers act as a phonebook for translating domain names or email addresses to IP addresses. A local DNS server does
a recursive lookup that starts with the root server, then the top level domain server, and ends with the authoritative 
nameserver (e.g. . --> .com --> google.com --> mail.google.com). The authoritative nameserver usually holds the definitive 
information about specific domains. DNS servers can be replicated for high availability and local DNS servers can cache 
IP addresses after a successful recursive lookup.
```

##### DNS Zone
```
A DNS zone is a portion of the DNS namespace that a specific organization or administrator manages. Each DNS server for
a zone has a zone file that contains records for the zone in plain text format.
```

##### Fully Qualified Domain Name (FQDN)
```
mail.google.com
  * com: top-level domain
  * google: second-level domain
  * google.com: domain name, root domain
  * mail: subdomain
  * mail.google.com: fully qualified domain name
```

##### Records
```
A     : Mapping of a hostname to one or more IPv4 addresses.
AAAA  : Mapping of a hostname to one or more IPv6 addresses.
CNAME : Mapping of one hostname to another hostname. Cannot map root domains to another hostname.
NS    : Mapping to the nameserver responsible for a particular domain.
SOA   : Record of administrative information in regards to the DNS zone.
``` 

##### Address Resolution Protocol (ARP)
```
Protocol used to translate IP addresses into MAC addresses. Hosts and switches preserve these mappings on ARP tables.

Step 1) When a client sends a request, it will know the destination IP but not the MAC address.
        This means the request cannot be sent as the L2 header is incomplete.
Step 2) Client fires an ARP request that sends a broadcast frame that holds the source's IP and MAC address.
        This broadcast is meant to discover the host with the destination IP address.
        Layer 2 header for the broadcast will carry a destination MAC address of ffff.ffff.ffff.
Step 3) Destination host receives the ARP broadcast and updates its ARP table with a <src-ip>:<src-mac> entry.
Step 4) Destination host fires an ARP response as a unicast that holds its IP and MAC address.
Step 5) Client receives the ARP response and updates its ARP table.
Step 6) Client is able to send its original request as it is now able to complete an L2 header.
```

### Subnets
---
##### Classless Interdomain Routing (CIDR)
```
a) Example 1: 192.168.1.0/24
     * 24 bits for network address (i.e. network's unique identifer). 
     * 8 bits for host address (i.e. device identifier in the network).
     * IP addresses in this CIDR range from 192.168.1.0 ~ 192.168.1.255
b) Routers perform a longest prefix match on these "prefixes" to route packets to the next hop.
```

##### Subnets
```
a) Subnets are a logical subdivision from the larger network that helps to build a scalable network infrastructure. This
   logical division helps to group and manage similar applications separately. A use case for this is better
   demonstrated here: https://www.youtube.com/watch?v=zmxLg4jV0ts&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=11.
b) Subnets allow for better security management. For instance, subnets can be configured such that only backend 
   applications might access a database subnet while only frontend applications can access the backend subnet but not 
   the database subnet.
```

##### Network and Braodcast Address
```
a) The xxx.xxx.xxx.0 in a subnet is reserved as the network address used to identify the subnet as a whole rather than
   an individual device on the network.
b) The last IP address in a subnet is reserved as the broadcast address.
```

##### Subnet Mask
```
32-bit number created by setting network bits to 1s and host bits to 0s (e.g 255.255.255.0). This helps to separate
the IP address into the network and host addresses. This number helps to determine whether a destination IP address lies 
in the same L2 network.
```

##### Network Address Translation (NAT)
```
NAT Steps
  Step 1) Assign private IP addresses within a common network.
  Step 2) NAT devices map private IP addresses to a single public IP address.
  Step 3) The mapped public IP address is used as the source IP when sending to the destination.
  Step 4) When host at destination tries to send a packet to the source within the subnet, demultiplexing IDs (L4 
          protocol) are used to find the original source.

NAT Disadvantages
  * Difficult to distinguish devices within the subnet.
  * Cannot connect to hosts within the private network (i.e. IP address is unknown) from the outside until a packet has 
    been trasmitted.
```

##### SNAT vs DNAT
```
SNAT: Refers to source NAT, which translates the source IP address.
DNAT: Refers to destination NAT, which translates the destination IP address.
```

##### Border Gateway Protocol (BGP)
```
BGP is a protocol responsible for picking the most efficient route to forward data to. This is done by autonomous 
systems constantly sharing routing information and next hops through TCP/IP peering sessions.
```
