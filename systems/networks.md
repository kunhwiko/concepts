### Networking & Security 
---
##### Getting Started with Networks 
```
Networks: system of links that interconnect computers to move data 

Internet: networking infrastructure linking connected devices 

Protocol: set of rules and structures that define the syntax / semantics of how computers communicate 
   - IP: address of where packets come from and where they should be sent 
   - TCP: responsible for breaking data into packets, delivering / reassembling the packets, checks for corruption 
   - HTTP: set of rules for how request-response works in the web 

Network Devices: specialized computers focusing on I/O that forward messages by connecting through network links 
   - Switches: 
      - connects L1 layers to form a network 
      - decides on a port based on a 'forwarding table' to determine where to send messages in an L2 network 
  	  - switches must become aware of different hosts in a given L2 network
  	  - through broadcasting across the network, switches learn MAC addresses based on the source of the broadcast   
  2) Routers:
  		- connects L2 networks to form L3 networks 
  		- coordinate amongst themselves to decide on a 'forwarding table' and 'routing table' for each router 

Ports: 
   - docking point for where information is received or sent
   - how multiple programs listen for new network connections on the same machine without collision  
   - IP address is like a mailbox to an apartment complex, and ports are the specific apt number

Tables:
   - Forwarding Table: maps a MAC address to a port to forward packets 
   - Routing Table: maps longest prefix matches on IPs to send packets over next hops (other L2 networks) 
```

##### Internet Architecture 
```
Circuit Switching:
   - the process of establishing circuits, transferring data, and then terminating upon finish 
   - resource allocation is inefficient, but fast for large data transfers 
   - guarantees data transfer while connected 

Packet Switching :
   - packet headers contain addresses, routing protocols compute packet hops, no resources are pre-allocated 
   - no connection is required, minimal network assumptions 
   - easy to recover from errors 

Internet Characteristics
   - uses packet switching 
   - decentralized  
   - Fate Sharing 
   - provides heterogeneous services to different devices/networks
      - latency for calls vs high quality for video streaming
      - possible due to "layering", or mix and matching different protocols as needed  

Fate Sharing vs Replication 
   - Replication 
      - networks hold replicas and are responsible for state 
      - fault tolerant only as long as replicas are fine 
      - concurrency / consistency issues might exist 
      - hard to engineer 

   - Fate Sharing
      - end hosts responsible for state
      - fault tolerant is perfect as long as host is up 
      - hurts the end host performance, but ends up working fine 

Curse of the Narrow Waist 
   1) Internet provides heterogeneous services by mix-matching various protocols
   2) those protocols eventually agree on IP/TCP/HTTP 
   3) since various protocols rely on IP/TCP/HTTP, switching / changing IP/TCP/HTTP causes a chain of problems   
```

##### Layers 
```
Layering: ability to mix and match different protocols 

Protocol Stack: set of protocols that are currently in use, can mix and match for different situations
ex) when using email, network can switch HTTP protocol layer for SMTP protocol layer 

OSI (formal layer): 
   Layer 7) Application 
   Layer 6) Presentation 
   Layer 5) Session
   Layer 4) Transport 
   Layer 3) Network
   Layer 2) Datalink
   Layer 1) Physical --> physical links that transfer message bits 

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
   - physical links that transfer message bits 

Layer 2: Link Layer 
   - bundles physical layers into a LAN
   - takes stream of bits into frames that hold MAC src/dest address
```

##### Network Layer 
```
Layer 3: Network Layer 
   - bundles link layers to Internet
   - takes frames and transmits as packets

Packets : small segments of data of a larger message 
   - IP Packet Header  : holds the source and destination address
   - TCP Packet Header : order of how packets should be reassembled
   - IP Packet Data    : holds the data of the packet  

IPv4: 
   - 32 bits broken into two parts (network and host addresses)
  -  network addresses are common to groups of host addresses (geography / company)

Classless Interdomain Routing 
  - 128.168.1.0/26 
     - 26 bits for network addresses, 6 bits for host addresses
     - great if you have 50 or less computers in a common network
  - 192.168.1.0/24
     - 8 bits for host addresses (192.168.1.0 ~ 192.168.1.255) 
  - routers do a longest prefix match on these "prefixes" to route packets   

Subnets
   - networks inside network
   - uses "subnet mask" to send packet through the most efficient route
   - send mail --> delivered to office --> delivered to department --> delivered to team --> delivered to person 
   - for more: https://www.cloudflare.com/learning/network-layer/what-is-a-subnet/

IPv6: 
   - more addresses, better functionality
   - delayed as the entire network must become IPv6 compatible 
  
Tunneling
   - encapsulating IPv6 packets as IPv4 packets to carry over IPv4 networks
   - not a popular strategy

NAT
   Steps 
      a) assign private IP addresses within common network
      b) NAT devices maps private IP addresses to a single public IP address
      c) send to destination address with mapped public IP address as the source 
      d) when host at destination tries to send a packet to a host within the network, use demultiplexing IDs (L4 protocol) to find host 
   Disadvantages 
      a) difficult to distinguish devices
      b) can't connect to hosts within the private network (address unknown) until they send a message first
```

##### Discovery 
```
host name (www.github.com) --(DNS)--> IP Address --(ARP)--> MAC address 

Processes of Discovery 
   1) Host begins only knowing its source MAC address 
   2) Must discover its source IP address (DHCP)
   3) Must discover the destination's IP address (DNS)
   4) Longest prefix match on the IP to determine if the message should go to a router or a local machine 
   5) Move to a router (another L2 network) or a local machine (current L2 network) (ARP) 

IP address 
   - determines "where" the machine is 
   - IP addresses can vary based on location 
   - IP addresses will efficiently narrow down from L3 to the correct L2 network, where computers can then be located through MAC addresses  

MAC address
  a) physical unique address of a machine 
  b) useful for local communications 
  c) determines "who" the machine is 
  d) difficult for the entire Internet to keep track of where all MAC addresses are 

DNS Server: phonebook for finding the IP addresses of sites, DNS servers are replicated for availability, and caches popular addresses 

DHCP Server: 
   - host broadcasts a DHCP discovery network, DHCP servers respond with an offer response, host sends a request message specifying the DHCP it wants
   - DHCP helps a machine learn its own IP address, IP addresses of local DNS servers, gateway routers, and prefix length 

ARP Table: table used to translate IP addresses into MAC addresses 
```

##### Transport Layer 
```
Layer 4: Transport Layer 
   - provides the means for transferring variable length data in a reliable fashion 
   - on the host and typically not on the layer 
   - message: UDP / TCP  

UDP 
   - a simply wrapper to IP that provides lightweight, fast delivery of data 
   - connection-less protocol
   - sends and receives chunks of data 

TCP 
   - extensive error checking, provides recovery options, and reliable delivery of data 
   - connection-oriented (must request a handshake with the destination host first)
   - Handshake: TCP sends requests by sending packets to destination server asking for a connection
   - sends and receives a stream of bytes in segments that must be reorganized upon delivery 
   - controls flow of packets (sending/receiving) 

How does TCP ensure reliable delivery?
   - checksum: checks for bit errors 
   - ack/nack: acknowledges if receiver received the message 
   - sequence number: checks which packets receiver already received 
   - timeout: resends if there is no response 

TCP Flow Control: end hosts control packet traffic 
   Problem: What if Netflix sends a 4K video, can a computer handle all the packets at once?
      - TCP allows us to control the flow rate of packets by advertising an available window in our buffer
      - TCP sends more packets when receiving an acknowledgment and sends less when a packet loss occurs. 

TCP Congestion Control: routers controlling packet traffic 
   Problem: What if there are too many packets in the network? 
      - routers tell hosts that the network is congested 
```

##### HTTP (Application Layer)
```
Base Structure of HTTP
   - synchronous request/reply protocol (before HTTP/2)
   - on top of TCP/IP (before HTTP/3)
   - stateless (server does not recognize same client) (before HTTP/1.1)
   - ASCII format (before HTTP/2)

Stateless
   Pros: improves scalability / resource availability on server side 
   Cons: some applications need persistent state (shopping carts, usage tracking)
   Cookies:
      - sets a key-value pair that a website can store in the browser 
      - the cookie is sent in subsequent requests so server can recognize clients 

HTTP 1.1
   - persistent connection: maintain TCP connections across multiple requests 
   - understands state (server recognizes client) 

Reducing page load times 
   - CDNs 

HTTP 2
   - binary framing: 
      - instead of sending stream of ASCII characters, messages are formatted into stream of packets 
      - this allows interleaving different messages based on priority 
   - compressed headers 
   - promises: send files that the client did not explicitly request, but might need as deemed by the server  

HTTP 3
   - moving to build a UDP based protocol: custom TLS handshake, custom congestion control 
```


##### Getting Started with Security 
```
Man-in-the-Middle(MITM) Attack: malicious activity to intercept or alter IP packets in an HTTP connection

Symmetric Encryption
   - uses a single key to encrypt/decrypt data and is faster 
   - makes use of Advanced Encryption Standard (AES) algorithm 

Asymmetric Encryption
   - uses a public and private key and is slower  
   - anyone can encrypt with public key, only private key can decrypt messages  
  
HTTPS: HTTP over TLS 
Transport Layer Security (TLS): security protocol for secure communication 

TLS Handshake: process to establish a secure connection between clients and server 
   1) 
      - Client sends "client hello"
      - Server responds with "server hello" + SSL certificate containing the public key
   2) 
      - Client verifies SSL certificate and sends a "premaster secret" encrypted with the public key
      - Server decrypts the premaster secret using the private key
   3) 
      - client hello, server hello, premaster secret are used to create temporary symmetric keys for the session
      - Symmetric keys are used to figure out whether a connection was established or has failed 

Purpose of SSL Certificates 
   - MITM attacks can intercept server hellos and public keys, send their own public key, and establish a connection with client 
   - SSL certificates guarantee where public keys come from
```
