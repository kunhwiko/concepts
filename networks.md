### Networking & Security 
---
##### Getting Started with Networks 
```
Networks: system of links that interconnect computers to move data 

Internet: networking infrastructure linking connected devices 

Protocol: set of rules and structures that define the syntax / semantics of how computers communicate 
  1) IP: address of where packets come from and where they should be sent 
  2) TCP: responsible for breaking data into packets, delivering / reassembling the packets, checks for corruption 
  3) HTTP: set of rules for how request-response works in the web 

Network Devices: specialized computers focusing on I/O that forward messages by connecting through network links 
  1) Switches: 
  		- decides on a port based on a 'forwarding table' to determine where to send messages in an L2 network 
  		- switches must become aware of different hosts in a given L2 network
  		- through broadcasting across the network, switches learn MAC addresses based on the source of the broadcast   
  		- connects L1 layers to form a network 
  2) Routers:
  		- coordinate amongst themselves to decide on a 'forwarding table' for each router 
  		- coordinate amongst themselves to decide on a 'routing table' for each router
  		- connects L2 networks to form L3 networks 

Tables:
  1) Forwarding Table: maps a MAC address to a port to forward packets 
  2) Routing Table: maps longest prefix matches on IPs to send packets over next hops (other L2 networks) 

Ports : 
  1) docking point for where information is received or sent
  2) how multiple programs listen for new network connections on the same machine without collision  
  3) IP address is like a mailbox to an apartment complex, and ports are the specific apt number
```

##### Internet Architecture 
```
Circuit Switching:
  1) the process of establishing circuits, transferring data, and then terminating upon finish 
  2) resource allocation is inefficient, but fast for large data transfers 
  3) guarantees data transfer while connected 

Packet Switching :
  1) packet headers contain addresses, routing protocols compute packet hops, no resources are pre-allocated 
  2) no connection is required, minimal network assumptions 
  3) easy to recover from errors 

Internet Characteristics
  1) uses packet switching 
  2) decentralized  
  3) Fate Sharing 
  4) provides heterogeneous services to different devices/networks
      - latency for calls vs high quality for video streaming
      - possible due to "layering", or mix and matching different protocols as needed  

Fate Sharing vs Replication 
  1) Replication 
    a) networks hold replicas and are responsible for state 
    b) fault tolerant only as long as replicas are fine 
    c) concurrency / consistency issues might exist 
    d) hard to engineer 

  2) Fate Sharing
    a) end hosts responsible for state
    b) fault tolerant is perfect as long as host is up 
    c) hurts the end host performance, but ends up working fine 

Curse of the Narrow Waist 
  - the Internet provides heterogeneous services and mixes and matches various protocols
  - those protocols eventually had to agree on IP/TCP/HTTP 
  - since various protocols rely on IP/TCP/HTTP, switching / changing IP/TCP/HTTP causes a chain of problems   
```

##### Layers 
```
Layering: ability to mix and match different protocols 

Protocol Stack : set of protocols that are currently in use, can mix and match for different situations
ex) when using email, network can switch HTTP protocol layer for SMTP protocol layer 

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
  1) when sending over messages, start with the highest layer and move down to layer 1
  2) for each layer, the lower layer "wraps" the current message from higher levels to create protocol stacks 
  3) lower layers don't have to know what higher level layers actually do 
  4) after sending over a message, switches and routers will decapsulate the message from lowest to highest layer 
  5) the message is then recapsulated before resending and the steps are repeated  
```

##### Lower Level Layers 
```
Layer 1: Physical Layer
  a) concerned with how signals are used to transfer message bits 
  b) network: physical links 
  c) message: bits 

Layer 2: Link Layer 
  a) takes physical layers and connects them 
  b) network: LAN 
  c) message: frames (holds the MAC src and dest address)
  d) framing: convert stream of bits into messages 
```

##### Network Layer 
```
Layer 3: Network Layer 
  a) takes link layers and connects them 
  b) network: Internet 
  c) message: packets 

IPv4: 32 bits broken into two parts (network and host addresses)
  - network addresses are common to groups of host addresses (geography / company)

Classless Interdomain Routing 
  - 128.168.1.0/26 --> 26 bits for network addresses, 6 bits for host addresses (great if you have 50 computers in a common network)
  - 192.168.1.0/24 --> 8 bits for host addresses (192.168.1.0 ~ 192.168.1.255) 
  - routers do a longest prefix match on these "prefixes" to route packets   

Packets : small segments of data of a larger message 
  * IP Packet Header : holds the source and destination address
  * TCP Packet Header : order of how packets should be reassembled
  * IP Packet Data : holds the data of the packet  

IPv6: has more addresses, better functionality, but has been delayed as the entire network must become IPv6 compatible 
  - Tunneling (not popular) : encapsulating IPv6 packets as IPv4 packets to carry over IPv4 networks
  - Network Address Translation (popular) : tries to solve having too few address in IPv4
      Steps 
        a) within a common network, assign private IP addresses
        b) NAT devices maps the private IP addresses to a single public IP address
        c) send to destination address with mapped public IP address as the source 
        d) when host at destination tries to send a packet to a host within the network, use demultiplexing IDs (L4 protocol) to find host 
      Disadvantages 
        a) more difficult to distinguish devices 
        b) can't connect to hosts within the private network (address unknown) until they send a message first   
```

##### Discovery 
```
host name (www.github.com) --(DNS)--> IP Address --(ARP)--> MAC address 

Processes of Discovery 
  Step 1) Host begins only knowing its source MAC address 
  Step 2) Must discover its source IP address (DHCP)
  Step 3) Must discover the destination's IP address (DNS)
  Step 4) Longest prefix match on the IP to determine if the message should go to a router or a local machine 
  Step 5) Move to a router (another L2 network) or a local machine (current L2 network) (ARP) 

IP address 
  a) determines "where" the machine is 
  b) IP addresses can vary based on your location 
  c) IP addresses will efficiently narrow down from L3 to the correct L2 network, where computers can then be located through MAC addresses  

MAC address
  a) physical unique address of a machine 
  b) useful for local communications 
  c) determines "who" the machine is 
  d) difficult for the entire Internet to keep track of where all MAC addresses are 

DNS Server: phonebook for finding the IP addresses of sites, DNS servers are replicated for availability, and caches popular addresses 

DHCP Server: 
  1) host broadcasts a DHCP discovery network, DHCP servers respond with an offer response, host sends a request message specifying the DHCP it wants
  2) DHCP helps a machine learn its own IP address, IP addresses of local DNS servers, gateway routers, and prefix length 

ARP Table: table used to translate IP addresses into MAC addresses 
```

##### Transport Layer 
```
Layer 4: Transport Layer 
  a) provides the means for transferring variable length data in a reliable fashion 
  b) on the host and typically not on the layer 
  c) message: UDP / TCP  

UDP 
  a) a simply wrapper to IP that provides lightweight, fast delivery of data 
  b) connection-less protocol
  c) sends and receives chunks of data 

TCP 
  a) extensive error checking, provides recovery options, and reliable delivery of data 
  b) connection-oriented (must request a handshake with the destination host first)
  c) Handshake: TCP sends requests by sending packets to destination server asking for a connection
  d) sends and receives a stream of bytes in segments that must be reorganized upon delivery 
```

##### Getting Started with Security 
```
Man-in-the-Middle(MITM) Attack : malicious activity to intercept or alter IP packets in an HTTP connection

Symmetric Encryption
  1) uses a single key to encrypt/decrypt data and is faster 
  2) makes use of Advanced Encryption Standard (AES) algorithm 

Asymmetric Encryption
  1) uses a public and private key and is slower  
  2) anyone can encrypt with public key, only private key can decrypt messages  
  
HTTPS : HTTP over TLS 
Transport Layer Security (TLS) : security protocol for secure communication 

TLS Handshake : process to establish a secure connection between clients and server 
  Step 1) 
    1) Client sends "client hello"
    2) Server responds with "server hello" + SSL certificate containing the public key
  Step 2) 
    1) Client verifies SSL certificate and sends a "premaster secret" encrypted with the public key
    2) Server decrypts the premaster secret using the private key
  Step 3) 
    1) client hello, server hello, premaster secret are used to create temporary symmetric keys for the session
    2) Symmetric keys are used to figure out whether a connection was established or has failed 

Purpose of SSL Certificates 
  1) MITM attacks can intercept server hellos and public keys, send their own public key, and establish a connection with client 
  2) SSL certificates guarantee where public keys come from
```
