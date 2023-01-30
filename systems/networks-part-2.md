### Layer 4
---
##### Transport Layer 
```
a) Provides the means for transferring variable length data in a reliable fashion to the right place.
b) Transport layer abstracts 'packets' into 'segments' and use ports as an address scheme.
   The header for Layer 3 holds the src/dest port numbers.
```

##### Significance of Ports on Layer 4
```
a) Ports determine where the server is listening on and where the client is awaiting for a response.
b) Ports act as a unique identifier to distinguish data streams and ensures various applications will receive the right data.
   As an example, each browser tab assigns a new random port on the client side to ensure the right data will be sent to that particular tab.
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


