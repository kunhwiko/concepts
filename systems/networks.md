### Getting Started with Networks
---
##### Networks
```
Networks
   a) System of links that interconnect computers to move data.
```

##### Protocols
```
Protocol
   a) St of rules and structures that define the syntax / semantics of how computers communicate.

IP
   a) Address of where packets come from and where they should be sent.

TCP
   a) Responsible for breaking data into packets, delivering / reassembling the packets, checks for corruption.

HTTP
   a) Set of rules for how request-response works in the web.
```

##### Network Devices
```
Network Devices
   a) Specialized computers focusing on I/O that forward messages by connecting through network links.

Switches 
   a) Connects L1 layers to form a network.
   b) Decides on a port based on a 'forwarding table' to determine where to send messages in an L2 network.
   c) Switches must become aware of different hosts in a given L2 network.
   d) Through broadcasting across the network, switches learn MAC addresses based on the source of the broadcast.

Routers
   a) Connects L2 networks to form L3 networks.
   b) Coordinate amongst themselves to decide on a 'forwarding table' and 'routing table' for each router.
```

##### Tables
```
Tables:
   a) Forwarding Table : Maps a MAC address to a port to forward packets.
   b) Routing Table    : Maps longest prefix matches on IPs to send packets over next hops (other L2 networks).
```

##### Ports
```
Ports: 
   a) Docking point for where information is received or sent.
   b) How multiple programs listen for new network connections on the same machine without collision.
   c) IP address is like a mailbox to an apartment complex, and ports are the specific apartment number.
```

### Internet
---
##### Internet Architecture 
```
Internet
   a) Networking infrastructure linking connected devices.

Circuit Switching:
   a) The process of establishing circuits, transferring data, and then terminating upon finish.
   b) Resource allocation is inefficient, but fast for large data transfers.
   c) Guarantees data transfer while connected.

Packet Switching
   a) Packet headers contain addresses, routing protocols compute packet hops, no resources are pre-allocated.
   b) No connection is required, minimal network assumptions.
   c) Easy to recover from errors.

Internet Characteristics
   a) Uses packet switching.
   b) Decentralized and fate sharing. 
   c) Provides heterogeneous services to different devices/networks.
      * Latency for calls vs high quality for video streaming.
        This is possible due to "layering", or mix and matching different protocols as needed.
   d) Curse of the Narrow Waist
      * Internet typically provides heterogeneous services by mix-matching various protocols.
        These protocols eventually agree on IP/TCP/HTTP. 
        Switching IP/TCP/HTTP can then cause a chain of problems for all other relying protocols. 
```

##### Fate Sharing vs Replication
```
Replication
   a) Networks hold replicas and are responsible for state.
   b) Fault tolerant only as long as replicas are fine.
   c) Concurrency / consistency issues might exist, and difficult to engineer.

Fate Sharing
   a) End hosts responsible for state.
   b) Fault toleranance is perfect as long as host is up.
   c) Hurts the end host performance, but ends up working fine in practice.
```

##### Layers
```
Layering
   a) Ability to mix and match different protocols.

Protocol Stack
   a) Set of protocols that are currently in use, can mix and match for different situations.
   b) Example: When using email, network can switch HTTP protocol layer for SMTP protocol layer.

OSI (formal layer): 
   Layer 7) Application 
   Layer 6) Presentation 
   Layer 5) Session
   Layer 4) Transport 
   Layer 3) Network
   Layer 2) Datalink
   Layer 1) Physical 

Internet (informal):
   Layer 7) Application
   Layer 4) Transport
   Layer 3) Internet
   Layer 1 & 2) Net Access/Physical

Encapsulation 
   Step 1) When sending over messages, start with the highest layer and move down to layer 1.
   Step 2) For each layer, the lower layer "wraps" the current message from higher levels to create protocol stacks.
   Step 3) Lower layers don't have to know what higher level layers actually do.
   Step 4) After sending over a message, switches and routers will decapsulate the message from lowest to highest layer.
   Step 5) The message is then recapsulated before resending and the steps are repeated.
```

### Layer 1 and Layer 2
---
##### Physical & Link Layer
```
Physical Layer
   a) Physical links that transfer message bits.

Link Layer
   a) Bundles physical layers into a LAN.
   b) takes stream of bits into frames that hold MAC src/dest address.
```

### Layer 3
---
##### Network Layer 
```
Network Layer 
   a) Bundles link layers to Internet.
   b) Takes frames and transmits as packets.
```

##### IPv4 vs IPv6
```
IPv4
   a) Consists of 32 bits that can be broken into two parts (network and host addresses).
   b) Network addresses are common addresses to groups of host addresses (geography / company).

IPv6: 
   a) More addresses, better functionality.
   b) Delayed as the entire network must become IPv6 compatible.
   c) Rely more on alternative solutions such as NATs for now.
```

##### Packets
```
Packets
   a) Small segments of data of a larger message.

Components
   a) IP Packet Header  : Holds the source and destination address.
   b) TCP Packet Header : Order of how packets should be reassembled.
   c) IP Packet Data    : Holds the data of the packet.
```

##### Classless Interdomain Routing (CIDR)
```
CIDR
   a) Example 1: 128.168.1.0/26
      * 26 bits for network addresses, 6 bits for host addresses.
      * Great if you have 50 or less computers in a common network.
   b) Example 2: 192.168.1.0/24
      * 24 bits for network addresses, 8 bits for host addresses (192.168.1.0 ~ 192.168.1.255).
   c) Routers do a longest prefix match on these "prefixes" to route packets.
```

##### Subnets
```
Subnets
   a) Networks inside networks.
   b) Uses "subnet mask" to send packet through the most efficient route.
      Send mail --> delivered to office --> delivered to department --> delivered to team --> delivered to person.
   c) For more: https://www.cloudflare.com/learning/network-layer/what-is-a-subnet/
```

##### Network Address Translation (NAT)
```
NAT Steps
   Step 1) Assign private IP addresses within common network.
   Step 2) NAT devices maps private IP addresses to a single public IP address.
   Step 3) Send to destination address with mapped public IP address as the source.
   Step 4) When host at destination tries to send a packet to a host within the network, use demultiplexing IDs (L4 protocol) to find host.

NAT Disadvantages
   a) Difficult to distinguish devices.
   b) Can't connect to hosts within the private network (address unknown) until they send a message first.
```

##### Discovery 
```
host name (www.github.com) --(DNS)--> IP Address --(ARP)--> MAC address 

Processes of Discovery 
   Step 1) Host begins only knowing its source MAC address.
   Step 2) Must discover its source IP address via DHCP.
   Step 3) Must discover the destination's IP address via DNS.
   Step 4) Longest prefix match on the IP to determine if the message should go to a router or a local machine.
   Step 5) Move to a router (another L2 network) or a local machine (current L2 network) via ARP. 

IP address 
   a) Determines "where" the machine is.
   b) IP addresses can vary based on location.
   c) IP addresses will efficiently narrow down from L3 to the correct L2 network, where computers can then be located through MAC addresses.

MAC address
   a) Physical unique address of a machine. 
   b) Useful for local communications.
   c) Determines "who" the machine is.
   d) Difficult for the entire Internet to keep track of where all MAC addresses are.

DNS Server
   a) Phonebook for finding the IP addresses of sites, DNS servers are replicated for availability, and caches popular addresses.

DHCP Server
   a) Host broadcasts a DHCP discovery ping --> DHCP servers respond with an offer response --> host sends a request message specifying the DHCP server it will use.
   b) DHCP helps a machine learn its own IP address, IP addresses of local DNS servers, gateway routers, and prefix length.

ARP Table
   a) Table used to translate IP addresses into MAC addresses.
```

### Layer 4
---
##### Transport Layer 
```
Transport Layer 
   a) Provides the means for transferring variable length data in a reliable fashion.
```

##### UDP
```
UDP 
   a) Wrapper around IP that provides lightweight, fast delivery of data.
   b) Connection-less protocol.
   c) Sends and receives chunks of data.
```

##### TCP
```
TCP 
   a) Extensive error checking, provides recovery options, and reliable delivery of data.
   b) Connection-oriented (must request a handshake with the destination host first).
   c) Sends and receives a stream of bytes in segments that must be reorganized upon delivery.
   d) Controls flow of packets (sending/receiving).

Handshake
   a) TCP sends requests by sending packets to destination server asking for a connection.

How does TCP ensure reliable delivery?
   a) checksum        : Checks for bit errors.
   b) ACK / NACK      : Acknowledges if receiver received the message.
   c) Sequence Number : Checks which packets receiver already received. 
   d) Timeout         : Resends if there is no response.
```

##### Flow Control
```
TCP Flow Control 
   a) End hosts control packet traffic.

Problem
   a) What if Netflix sends a 4K video, can a computer handle all the packets at once?

Solution
   a) TCP allows us to control the flow rate of packets by advertising an available window in our buffer.
   b) TCP sends more packets when receiving an acknowledgment and sends less when a packet loss occurs. 
```

##### Congestion Control
```
TCP Congestion Control 
   a) Routers control packet traffic.

Problem
   a) What if there are too many packets in the network?

Solution
   a) Routers tell hosts that the network is congested.
```

### Layer 7
---
##### HTTP 1
```
HTTP/1.0
   a) Wrapper on top of TCP/IP (before HTTP/3).
   b) Synchronous request/reply protocol (before HTTP/2).
   c) ASCII format (before HTTP/2).
   d) Stateless (server does not recognize same client) (before HTTP/1.1).
   e) Single request/reply per TCP connection (before HTTP/1.1).

HTTP/1.1
   a) Host header fields: Servers can distinguish different host addresses that have the same IP address through the header. 
                          This allows for multiple domains per web server. 
   b) Persistent connections: Maintains the same TCP connections across multiple requests/responses.
   c) Chunked transfer encoding: By transferring chunks, it is unnecessary to generate the full content and content size can be dynamic. 
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

##### HTTP 2 and 3
```
HTTP 2
   - binary framing: 
      - instead of sending stream of ASCII characters, messages are formatted into stream of packets 
      - this allows interleaving different messages based on priority 
   - compressed headers 
   - promises: send files that the client did not explicitly request, but might need as deemed by the server  

HTTP 3
   - moving to build a UDP based protocol: custom TLS handshake, custom congestion control 
```

### Getting Started with Security 
---
##### Encryptions
```
Man-in-the-Middle(MITM) Attack: malicious activity to intercept or alter IP packets in an HTTP connection

Symmetric Encryption
   - uses a single key to encrypt/decrypt data and is faster 
   - makes use of Advanced Encryption Standard (AES) algorithm 

Asymmetric Encryption
   - uses a public and private key and is slower  
   - anyone can encrypt with public key, only private key can decrypt messages  
```

##### TLS
```
HTTPS: HTTP over TLS 
Transport Layer Security (TLS): security protocol for secure communication 

TLS Handshake: process to establish a secure connection between clients and server 
   Step 1) 
      - Client sends "client hello"
      - Server responds with "server hello" + SSL certificate containing the public key
   Step 2) 
      - Client verifies SSL certificate and sends a "premaster secret" encrypted with the public key
      - Server decrypts the premaster secret using the private key
   Step 3) 
      - client hello, server hello, premaster secret are used to create temporary symmetric keys for the session
      - Symmetric keys are used to figure out whether a connection was established or has failed 

Purpose of SSL Certificates 
   - MITM attacks can intercept server hellos and public keys, send their own public key, and establish a connection with client 
   - SSL certificates guarantee where public keys come from
```

##### Authentication & Authorization
```
Authentication
   - Act of validating that users are whom they claim to be 

Authorization
   - Act of verifying if the user has the ability to perform a certain function
```
