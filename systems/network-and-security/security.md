### Security Fundamentals 
---
##### Authentication (AuthN)
```
Act of validating that users are whom they claim to be.
```

##### Authorization (AuthZ)
```
Act of verifying if the user has the ability to perform a certain function.
```

### Encryption
---
##### Man In the Middle (MITM) Attack
```
Malicious activity to intercept or alter packets in a network before they reach the intended receiver.
```

##### Symmetric Encryption
```
In symmetric encryption, the client and server uses the same secret key to encrypt or decrypt messages. The secret key 
is derived by both the client and server through a key exchange algorithm (e.g. AES algorithm). This secret key should 
never be trasmitted over the network.
```

##### Asymmetric Encryption
```
a) In asymmetric encryption, the client and the server uses public and private keys to encrypt or decrypt messages. 
   Public keys are used by anyone to encrypt messages, but only the host with the private key can decrypt messages.
b) Asymmetric encryption is used during the key exchange algorithm of symmetric encryption to help generate a shared 
   secret key. 
c) Asymmetric encryption tends to be slower than symmetric encryption for data transfer purposes.
```

### TLS
---
##### TLS Handshake - RSA Key Exchange Algorithm
```
Process to establish a secure connection between a client and server. The RSA key exchange algorithm was used for TLS
before version 1.3 and the steps are as follows:

Step 1) Client initiates a "client hello" to the server. This message includes the TLS version that the client supports,
        list of cipher suites, and a string of random bytes called the "client random".
Step 2) Server responds with a "server hello" containing the server's SSL certificate, server's public key, a selected 
        cipher suite, and a string of random bytes called the "server random". 
Step 3) Client verifies the server's SSL certificate with the certificate authority that signed it. This verifies that
        the server can be entrusted.
Step 4) Client sends an additional string of random bytes called the "premaster secret". This secret is encrypted with
        the server's public key that the client had previously received.
Step 5) Server decrypts the premaster secret using the private key. This secret is enough to prevent MITM attacks, but
        it is used alongside client/server randoms to prevent replay attacks. Randoms represent a state that is unique
        per connection, preventing replay attacks as the attacks require a new connection.
Step 6) Both client and server will generate sessions keys using the client random, server random, and premaster secret.
        These session keys are temporary symmetric keys and should be identical.
Step 7) The client and server share a "finished" message encrypted with the session key and the handshake is complete. 
```

##### TLS Handshake 1.3
```
Step 1) Client initiates a "client hello" to the server with the TLS version, client random, list of cipher suites, and
        parameters to be used for calculating the premaster secret. The number of cipher suites supported in TLS 1.3 
        have vastly been reduced to the point that the client assumes that it already knows the server's preferred key 
        exchange method.
Step 2) Server uses the client random, server random, and premaster secret parameters to generate a master secret.
Step 3) Server sends back a "server hello" with the SSL certificate, public key, server random, and selected cipher 
        suite. Since the server already has a master secret, it also sends a "finished" message.
Step 4) Client verifies the SSL ceritificate and CA that signed it. Once verified, it will generate the master secret
        using the client random, server random, and premaster secret parameters. 
Step 5) Client sends a "finished" message and a session for symmetric encryption should have been established. 
```

##### SSL Certificates
```
a) Hackers can impersonate the identity of servers and reroute traffic to their servers. SSL certificates are used to 
   guarantee that the server is who they claim to be. Self-signed ceritificates are insecure and cannot be trusted, but
   ceritificates signed by trusted certificate authorities (a.k.a. CAs) can be trusted.
b) Similar to server certificates, servers may verify the identity of clients through client certificates. 
c) Similar to server certificates, CAs have root certificates to validate that their public key is legitimiate. These
   CA public keys are pre-built into web browsers.  
```

##### Certificate Signing Request (CSR)
```
A CSR can be created to get a certificate signed by a root CA. After the root CA reviews the legitimacy of the cert, it
will be signed. 
```

##### Server Name Identification (SNI)
```
When a server with a single IP address hosts multiple websites each with its own SSL certificates, the server may be
unaware of which SSL certificate to present back to the client. This is because TLS handshakes occur before the client
indicates which website it's connecting to. SNI is an extension of the TLS protocol that allows the client to indicate
the hostname it's connecting to during the TLS handshake.
```

### Secure Shell Protocol (SSH)
---
##### Port Forwarding
```
a) Port forwarding maps packets with a certain port number to be sent to a specified IP address and port number. 
b) Port forwarding allows hosts in the Internet to connect to a specific host sitting within a private network.
```

##### Secure Shell Protocol (SSH)
```
SSH is a secure means to connect to a remote server without risking privacy from MITM attacks.

Step 1) Client initiates a TCP handshake with the remote server, usually listening on port 22.
Step 2) Client and server agree on an SSH protocol version and encryption/key exchange algorithm that both can support.
Step 3) Server sends its public key to the client. The client will verify the server's public key against a list of 
        trusted hosts that it has stored locally. If the server's public key is not in the list, the client will prompt
        the user to accept or reject the connection.
Step 4) Client and server use a key exchange algorithm to create a symmetric key. Using Diffie–Hellman as an example, 
        the client and server first agree on a shared prime number and a generator number. The client and server use their 
        respective private keys along with the shared prime number and generator to each compute a new public key.
Step 5) Private keys along with the new public keys and shared prime number are used to compute a new key. The client and
        server will have arrived at the same value and this new key will be used as a symmetric key:
        - https://www.practicalnetworking.net/series/cryptography/diffie-hellman/.
Step 6) The server digitally signs the session key with its private key and sends it to the client. The client can verify
        the signature using the server's public key (e.g. RSA algorithm). If the signature is valid, the client will use 
        the session key to encrypt all future communication with the server: 
        - https://www.encryptionconsulting.com/education-center/what-is-rsa/
Step 7) After a connection is encrypted, the client sends credentials (e.g. password, public key auth) for authentication.
```

### Virtual Private Network (VPN)
---
##### VPN
```
a) VPNs use client side proxying to hide the IP address of the client.
b) VPNs use tunneling to encrypt data throughout the network session.
c) VPN providers typically have zero log retention policies.
```

##### Tunneling
```
Tunneling is the process of moving private network communications across the public network through encapsulation.
Encapsulation helps to encrypt data or to traverse a shorter network that normally cannot be crossed (e.g. IPv6 packets through IPv4 networks).
Packets will be decrypted once they reach either end of the tunnel.
```

##### Tunneling Protocols
```
GRE Tunneling
  a) Encapsulates packets with a GRE and IP header to traverse a previously unsupported network and potentially achieve less hops.
     GRE headers are used to encapsulate packets as GRE packet.
     IP headers are used to identify the IP addresses for the beginning and end of the tunnel. 

IPsec Tunneling
  a) Encrypts IP packets and adds an authentication step through encapsulation. 
  b) IPsec packets can also be used in conjunction with GRE tunneling.

SSH Tunneling
  a) SSH tunneling establishes an SSH connection to forward insecure data through an encrypted tunnel.
  b) SSH tunneling be used to forward ports that are blocked by firewalls to a different port that is not blocked.
  c) More information here: https://goteleport.com/blog/ssh-tunneling-explained/ and https://www.youtube.com/watch?v=AtuAdk4MwWw.
```

### Cross Origin Resource Sharing (CORS)
---
##### Same Origin Policy
```
Security measure that allows a web page to only interact with resources of the same origin. An origin is defined as
the combination of scheme, domain, and port. 
```

##### CORS
```
CORS is a browser-side HTTP header based security mechanism that allows a server to indicate which origins other than 
its own are permitted to access resources. 

Step 1) When a browser makes a request to https://example.com to fetch an index.html file, the file might require 
        additional image data to be fetched from https://service.com to completely render the website. 
Step 2) Before an actual request is made to https://service.com, the browser may send a preflight HTTP request with an
        origin header (Origin: https://example.com) to check if the server will permit the actual request.
Step 3) If permitted, the server will send back an access-control-allow-origin header.
Step 4) The browser will continue to make the actual request to https://service.com to fetch the necessary images.
```