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

##### IP Address
```
a) Determines "where" the machine is.
b) IP addresses are 32 bits and are hierarchically designed by location/teams through subnetting.
c) IP addresses will efficiently narrow down from L3 to the correct L2 network, where computers can then be located through MAC addresses.
```

##### MAC Address
```
a) Physical unique address of a machine.
   It is difficult for the Internet to keep track of where all MAC addresses are.  
b) Determines "who" the machine is.
```

##### Ports
```
a) Docking point for where information is received or sent.
b) Means for how multiple programs can listen for new network connections on the same machine without collision.
c) IP address is like a mailbox to an apartment complex, and ports are the specific apartment number.
```

### Internet Fundamentals
---
##### Circuit Switching vs Packet Sharing
```
Circuit Switching
  a) The process of establishing circuits, transferring data, and then terminating upon finish.
  b) Resource allocation is inefficient, but fast for large data transfers.
  c) Guarantees data transfer while connected.

Packet Switching
  a) Packet headers contain addresses, routing protocols compute packet hops, and no resources are pre-allocated.
  b) No connection is required, minimal network assumptions.
  c) Easy to recover from errors.
```

##### Replication vs Fate Sharing
```
Replication
  a) Networks are responsible for state and holding replicas.
  b) Fault tolerant only as long as replicas are fine.
  c) Concurrency / consistency issues might exist, and is difficult to engineer.

Fate Sharing
  a) End hosts are responsible for state.
  b) Fault toleranance is only perfect as long as host is up.
  c) Hurts the end host performance, but ends up working fine in practice.
```

##### Internet 
```
a) Internet is defined as a networking infrastructure that links various connected devices.
b) Internet uses packet switching.
c) Internet is decentralized and is fate sharing.
d) Internet provides heterogeneous services to different devices/networks.
     * Latency for calls vs high quality for video streaming.
     * This is possible due to network "layering", or mix and matching different protocols as needed.
e) Internet has a "Curse of the Narrow Waist"
     * All heterogeneous services eventually agree on IP/TCP/HTTP. 
       Therefore, switching IP/TCP/HTTP can cause a chain of problems for all other dependent protocols. 
```

### OSI Model
---
##### Network Layers
```
OSI (formal layer): 
  Layer 7) Application 
  Layer 6) Presentation 
  Layer 5) Session
  Layer 4) Transport 
  Layer 3) Network
  Layer 2) Datalink
  Layer 1) Physical 

Internet (informal):
  Layer 7)     Application
  Layer 4)     Transport
  Layer 3)     Internet
  Layer 1 & 2) Net Access/Physical
```

##### Layering
```
Layering
  a) Ability to mix and match different protocols.

Protocol Stack
   a) Set of protocols that are currently in use, but can be mixed and matched for different situations.
   b) For emails, network can switch HTTP protocol layer for SMTP protocol layer.
```

##### Encapsulation
```
Step 1) When sending over messages, start with the highest layer and move down to layer 1.
Step 2) For each layer, the lower layer "wraps" the current message from higher levels to create protocol stacks.
Step 3) Lower layers don't have to know what higher level layers actually do.
Step 4) After sending over a message, switches and routers will decapsulate the message from lowest to highest layer.
Step 5) The message is then recapsulated before resending and the steps are repeated.
```

### Layer 1 & 2
---
##### Physical Layer
```
Physical links that transfer message bits.
```

##### Link Layer
```
a) Bundles physical layers into a Local Area Network (LAN).
b) Abstracts stream of bits from the Physical Layer into frames that hold MAC src/dest address.
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

##### Bridge
```
Problem Statement
  a) Hubs cause all hosts connected to the hub to know about data that they do not need to be involved in.

Solution
  a) Bridges sit between two hubs and have two ports each for one of the hubs.
  b) Bridges know which hosts are one what side of either port.
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
a) Bundles link layers to Internet.
b) Takes frames and transmits as packets.
```

##### Router
```
a) Routers connect L2 networks to form L3 networks and help facilitate communication between networks.
b) Routers are assigned an IP address from each L2 network it is connected to.
   This allows the router to serve as a network 'gateway', or a means for traffic to travel outside an L2 network.
c) Routers help form the IP address hierarchy in networks and subnets. 
d) Routers are knowledgable of the L2 networks connected to itself through a 'routing table'.
e) Routers provide a control point for network traffic where security, filtering, redirecting can be enforced.

Routing Table
  a) Routing table holds information about all the networks a router is aware about.
  b) Routing table maps longest prefix matches on IPs to send packets over to next hops (other L2 networks).

More info here: https://www.youtube.com/watch?v=H7-NR3Q3BeI&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=2
```

##### Packets
```
Small segments of data of a larger message.

Components
  a) IP Packet Header  : Holds the source and destination address.
  b) TCP Packet Header : Order of how packets should be reassembled.
  c) IP Packet Data    : Holds the data of the packet.
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

##### DNS Server
```
a) DNS servers act as a phonebook for finding the IP addresses of various sites.
b) DNS servers are replicated for availability, and caches popular addresses.
```

##### DHCP Server
```
a) Host broadcasts a DHCP discovery ping --> DHCP servers respond with an offer response --> host sends a request message specifying the DHCP server it will use.
b) DHCP helps a machine learn its own IP address, IP addresses of local DNS servers, gateway routers, and prefix length.
```

##### A and CNAME Records
```
A Records
  a) A mapping of a hostname to one or more IP addresses.

CNAME Records
  a) An alias from one domain to another.
``` 

##### ARP Table
```
Table used to translate IP addresses into MAC addresses.
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

Subnet Masking
  a) Means to send packets to a smaller subnet in a time efficient way.
     Send mail --> route to office --> route to department --> route to team --> route to person.
     The above does not waste time trying to get the mail to the individual person directly.
  b) For more: https://www.cloudflare.com/learning/network-layer/what-is-a-subnet/
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

### Layer 4
---
##### Transport Layer 
```
Provides the means for transferring variable length data in a reliable fashion.
```

##### UDP
```
a) UDP is a wrapper around IP that provides lightweight, fast delivery of data.
b) UDP is a connection-less protocol.
c) UDP sends and receives chunks of data.
```

##### TCP
```
a) TCP does extensive error checking, provides recovery options, and reliable delivery of data.
b) TCP is connection-oriented (i.e. must request a handshake with the destination first).
c) TCP sends and receives a stream of bytes in segments that must be reorganized upon delivery.
d) TCP controls flow of packets (i.e. sending/receiving).
e) TCP ensures reliable delivery through the following:
     * Checksum        : Checks for bit errors.
     * ACK / NACK      : Acknowledges if receiver received the message.
     * Sequence Number : Checks which packets receiver already received. 
     * Timeout         : Resends if there is no response.
```

##### TCP Handshake
```
TCP sends requests by sending packets to destination server asking for a connection.
```

##### Flow Control
```
End hosts control packet traffic.

Problem Statement
  a) What if Netflix sends a 4K video, can a computer handle all the packets at once?

Solution
  a) TCP allows us to control the flow rate of packets by advertising an available window in our buffer.
  b) TCP sends more packets when receiving an acknowledgment and sends less when a packet loss occurs. 
```

##### Congestion Control
```
Routers control packet traffic.

Problem Statement
  a) What if there are too many packets in the network?

Solution
  a) Routers tell hosts that the network is congested.
```

### Layer 7
---
##### HTTP 1.0
```
a) Wrapper on top of TCP/IP (before HTTP/3).
b) Synchronous request/reply protocol (before HTTP/2).
c) ASCII format (before HTTP/2).
d) Stateless (server does not recognize same client) (before HTTP/1.1).
e) Single request/reply per TCP connection (before HTTP/1.1).
```

##### HTTP 1.1
```
a) Host header fields
     * Servers can distinguish different host addresses that have the same IP address through the header. 
       This allows for multiple domains per web server. 
b) Persistent connections
     * Maintains the same TCP connections across multiple requests/responses.
c) Chunked transfer encoding
     * By transferring chunks, it is unnecessary to generate the full content and content size can be dynamic. 
       For each chunk, a new header can be sent instead of having to buffer the whole data before sending the initial header.
       This strategy makes use of the advantages of persistent connections.
```

##### Cookies
```
Pros of Stateless Application
  a) Improves scalability.
  b) Improves resource availability on server side.

Cons of Stateless Application
  a) Some applications need persistent state (e.g. shopping carts, usage tracking).

Cookies
  a) Sets a key-value pair that a website can store in the browser.
     The cookie is sent from the browser in subsequent client requests so the server can recognize clients.
```

##### HTTP 2
```
a) Binary Framing
     * Instead of sending stream of ASCII characters, messages are formatted into stream of packets.
       This allows interleaving different messages based on priority.
b) Compressed headers
c) Promises
     * Send files that the client did not explicitly request, but might need as deemed by the server.
```

##### HTTP 3
```
a) UDP based protocol.
b) Custom TLS handshake protocol.
c) Custom congestion control.
```

### Security Fundamentals 
---
##### Encryptions
```
Man-in-the-Middle(MITM) Attack
  a) Malicious activity to intercept or alter IP packets in an HTTP connection.

Symmetric Encryption
  a) Uses a single key to encrypt/decrypt data and is faster.
  b) <akes use of Advanced Encryption Standard (AES) algorithm.

Asymmetric Encryption
  a) Uses a public and private key and is slower.
  b) Anyone can encrypt with public key, only private key can decrypt messages.
```

##### Authentication (AuthN)
```
Act of validating that users are whom they claim to be.
```

##### Authorization
```
Act of verifying if the user has the ability to perform a certain function.
```

### TLS
---
##### HTTPS
```
HTTP over TLS.
```

##### TLS
```
Security protocol for secure communication.
```

##### TLS Handshake
```
Process to establish a secure connection between clients and server.

Step 1) Client sends "client hello".
        Server responds with "server hello" + SSL certificate containing the public key.
Step 2) Client verifies SSL certificate and sends a "premaster secret" encrypted with the public key.
        Server decrypts the premaster secret using the private key.
Step 3) Client hello, server hello, premaster secret are used to create temporary symmetric keys for the session.
        Symmetric keys are used to figure out whether a connection was established or has failed.
```

##### SSL Certificates
```
a) MITM attacks can intercept server hellos and public keys, send their own public key, and establish a connection with client.
b) SSL certificates guarantee where public keys come from.
```


