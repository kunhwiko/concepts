### Layer 3
---
##### Network Layer
```
a) Network layer bundles data link layers to the Internet.
b) Network layer abstracts 'frames' into 'packets' and use IP addresses as an address scheme. The header for Layer 3 
   holds the src/dest IP addresses.
```

##### Router
```
a) Routers connect L2 networks to form L3 networks and are assigned an IP address from each L2 network it is connected to.
   This allows routers to facilitate communication between networks and serve as a 'gateway', a means for traffic to 
   travel outside an L2 network.
b) Routers are knowledgable of connected L2 networks through a 'routing table'.
c) Routers typically have multiple NICs and MAC addresses as each network interface (e.g. ports) requires a MAC address.
d) Routers help form the IP address hierarchy in networks and subnets. 
e) Routers provide a control point for network traffic where security, filtering, redirecting can be enforced.

Routing Table
  a) A table that maps IP prefixes to next hops. Both hosts and routers have a routing table.
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

##### Unicast vs Anycast IP
```
Unicast IP: One server holds one IP address
Anycast IP: Multiple services hold the same IP address. The closest server to the client will respond to the request.
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

##### Network and Broadcast Address
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