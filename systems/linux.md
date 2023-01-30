### Networking
---
##### Netfilter / IPTables
```
Netfilter
  a) Framework to configure packet filtering, create NAT or port translation rules, and manage traffic flow in the network.
  b) Allows various networking operations to be implemented in the form of handlers. 
     These handlers perform certain predefined actions for incoming and outgoing packets.

IP Tables
  a) Command line utilities that can configure set of "chains" on Linux firewall to filter or block network traffic.
     These chains are organized into tables based on the type of rules (e.g. filter, NAT, mangle, raw).
  b) Linux IP Tables are built as part of Netfilter modules. 

More here: https://www.digitalocean.com/community/tutorials/a-deep-dive-into-iptables-and-netfilter-architecture
```

### Resource Allocation
---
##### Control Groups (cgroups)
```
Linux kernel feature that limits, isolates, and monitors resource usage (CPU, memory, I/O, network) of a collection of processes.
```