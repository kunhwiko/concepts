### Networking & Security 
---
##### Networking Terms 
```
Networks : system of links that interconnect computers to move data 

Internet : networking infrastructure linking connected devices 

Protocol : set of rules and structures that defines the syntax/semantics of how computers communicate 
  1) IP : address of where packets come from and where they should be sent  
    * IPv4, IPv6
  2) TCP : responsible for breaking data into packets, delivering/reassembling the packets, checks for corruption
  3) DNS Server : phonebook for finding the IP address of sites 
  4) HTTP : set of rules for how request-response works in the web 

Ports : 
  1) docking point for where information is received or sent
  2) how multiple programs listen for new network connections on the same machine without collision  
  3) IP address is like a mailbox to an apartment complex, and ports are the specific apt number

Types of Devices 
  1) Switches : 
       - machines that decide on a port based on a 'forwarding table' and where to send the message within an L2 network 
       - connects L1 layers to form a network 
       - switches must be aware of the different hosts in a given L2 network 
     (helps connect L1 layers and form a network)
  2) Routers : 
       - machines that coordinate amongst themselves to decide on a 'routing table' for each machine 
       - connects L2 networks to enable messages to travel beyond a certain L2 network (helps form the Internet)
  --> Switches and Routers make up modern Ethernet connections 
```

##### Internet Architecture
```
Circuit Switching : 
  1) the process of establishing circuits, transferring data, and then terminating upon finish 
  2) resource allocation is inefficient, but potentially better for large data transfers 
  3) guarantees data transfer while connected 

Packets : small segments of data of a larger message 
  * IP Packet Header : holds the source and destination address
  * TCP Packet Header : order of how packets should be reassembled
  * IP Packet Data : holds the data of the packet 

Packet Switching :
  1) packet headers contain addresses, routing protocols compute packet hops, no resources are pre-allocated 
  2) no connection is required, minimal network assumptions 
  3) easy to recover from errors 

Internet Characteristics 
  1) uses packet switching
  2) decentralized / handles many users using existing networks
  3) Fate Sharing  
  4) provides heterogeneous services to different interfaces/devices/networks (low latency for calls, high quality for video streaming)

Fate Sharing vs Replication
  1) Replication
    a) the networks hold replicas & are responsible for state 
    b) fault tolerant only as long as replicas are fine 
    c) concurrency / consistency issues exist 

  2) Fate Sharing
    a) end hosts responsible for state
    b) fault tolerant is perfect as long as host is up 
    c) hurts the end host performance      

Handshake : TCP sends requests by sending packets to destination server asking for a connection

Curse of the Narrow Waist : different protocols (FTP, HTTP, SMTP) all rely on IP, so switching/changing IP can be a huge problem  
```

##### Layers 
```
Layering : ability to mix and match different protocols 

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
  2) for each layer, the lower layer "wraps" the current message from higher levels 
  3) lower layers don't have to know what higher level layers actually do 
  4) after sending over a message, switches and routers will decapsulate the message, then recapsulate before resending
  5) this step is repeated until we reach the end user  
```

##### Low Level Layers
```
Physical Layer
  a) concerned with how signals are used to transfer message bits 
  b) network : physical bits  
  c) message : bits 

Data Link Layer
  a) connects physical layers 
  b) network : LAN
  c) message : frames 

Internet Layer 
  a) connects data link layers 
  b) network : Internet 
  c) message : packets 
```

##### Packets 
```
IPv4 : 32 bits broken into two parts (network and host address), network addresses are addresses common to some group of host addresses (geography, company)

Classless Interdomain Routing
  - 128.168.1.0/26 --> 26 bits for network addresses, 6 bits for host addresses (great if you have 50 computers in a common network)
  - 192.168.1.0/24 --> 8 bits for host addresses (192.168.1.0 ~ 192.168.1.255)   

IPv6 : has more addresses, better functionality, but has been delayed as the entire network must become IPv6 compatible 
  - Tunneling (not popular) : encapsulating IPv6 packets as IPv4 packets to carry over IPv4 networks 
  - Network Address Translation (popular) : tries to solve having too few address in IPv4
      Steps 
        a) within a common network, assign private IP addresses
        b) NAT devices maps the private IP addresses to a single public IP address
        c) send to destination address with mapped public IP address
        d) when host at destination tries to send a packet to a host within the network, use demultiplexing IDs (L4 protocol) to find host 
      Disadvantages 
        a) more difficult to distinguish devices 
        b) can't connect to hosts within the private network (address unknown) until they send a message first 
```

##### Security Fundamentals 
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
