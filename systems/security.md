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
Malicious activity to intercept or alter IP packets in an HTTP connection.
```

##### Symmetric Encryption
```
In symmetric encryption, the client and server uses a secret key to encrypt or decrypt messages.
The secret key is derived by both the client and server through a key exchange algorithm (e.g. AES algorithm).
This secret key is never trasmitted over the network.
```

##### Asymmetric Encryption
```
In asymmetric encryption, the client and the server uses public keys and private keys to encrypt or decrypt messages.
Public keys are used by any individual to encrypt messages, but only the host with the private key can decrypt messages.
Asymmetric encryption is used during the key exchange algorithm of symmetric encryption to help generate a shared secret key.
Asymmetric encryption tends to be slower than symmetric encryption for data transfer purposes.
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

Step 1) Client initiates a TCP handshake with the remote server, which should be listening on port 22.
Step 2) Server sends a public key to verify authenticity along with a list of supported encryption protocols.
Step 3) Client agrees on an encryption protocol that it can support and the connection is started.
Step 4) Client and server use a key exchange algorithm (i.e. Diffie–Hellman) to create a symmetric key to encrypt data.
        Public keys and private keys of both the client and the server are involved to compute the symmetric key. 
Step 5) Client sends login credentials that are encrypted with the symmetric key to the server for authentication.
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
     GRE headers are used to identify the packet is a GRE packet.
     IP headers are used to identify the IP addresses for the beginning and end of the tunnel. 

IPsec Tunneling
  a) Encrypts IP packets and adds an authentication step through encapsulation. 
  b) IPsec packets can also be used in conjunction with GRE tunneling.

SSH Tunneling
  a) Establishes an SSH connection to forward insecure data through an encrypted tunnel.
  b) Can be used to forward ports that are blocked by firewalls to a different port that is not blocked.
     More here: https://www.youtube.com/watch?v=AtuAdk4MwWw and https://goteleport.com/blog/ssh-tunneling-explained/
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
