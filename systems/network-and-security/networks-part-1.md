### Basic Networking Components
---
##### Protocols
```
Set of rules and structures that define the semantics of how computers communicate.
Below are some examples of protocols.
  * IP   : Address of where packets come from and where they should be sent.
  * TCP  : Responsible for breaking data into packets, delivering / reassembling packets, and checking for corruption.
  * HTTP : Set of rules for how request-response works in the web.
```

##### MAC Address
```
a) Determines "who" the machine is.
b) Physical unique address of a machine represented in 48 bits. It is difficult for the Internet to keep track of where 
   all the MAC addresses are.
c) For broadcast frames, the destination MAC address will typically look like ffff.ffff.ffff.
```

##### IP Address
```
a) Determines "where" the machine is.
b) IP addresses are 32 bits and are hierarchically designed by location/teams through subnetting.
c) IP addresses will efficiently narrow down from L3 to the correct L2 network, where computers can then be located 
   through MAC addresses.
```

##### Ports
```
a) Docking point for where information is received or sent.
b) Means for how multiple programs can listen for new network connections on the same machine without collision.
c) IP address is like a mailbox to an apartment complex, and ports are the specific apartment number.
```

### OSI Model
---
##### Network Layers
```
Layer 7) Application 
Layer 6) Presentation 
Layer 5) Session
Layer 4) Transport
Layer 3) Network
Layer 2) Datalink
Layer 1) Physical 
```

##### Layering
```
Layering
  * Ability to mix and match different protocols.

Protocol Stack
  * Set of protocols that are currently in use, but can be mixed and matched for different situations. For example, in 
    the case of emails, network can switch HTTP protocol layer for SMTP protocol layer.
```

##### Encapsulation
```
Step 1) When sending over messages, headers are "encapsulated" starting from highest to lowest layer.
Step 2) For each layer, the lower layer wraps the message coming from higher levels to create protocol stacks. Lower 
        layers do not have to actually care about what these higher level layers do.
Step 3) After sending over a message, switches and routers will decapsulate the message from lowest to highest layer.
        As an example, a router at this time will identify from the layer 2 header that the message was intended for itself.
        It will then discard the L2 header and pass for Layer 3 header validation.
Step 4) The message is then recapsulated before resending and the steps are repeated. As an example, a router at this 
        time will write a new layer 2 header specifying a new src/dest MAC address.
```

### Layer 1
---
##### Physical Layer
```
a) Physical links that transfer message bits.
b) Technologies include coaxial cables, fiber cables, ethernet, wifi, repeaters, hubs.
```

##### Repeater
```
Problem Statement
  * As data is transmitted over network wires, it decays as it travels.

Solution
  * Repeaters regenerate signals in between network traffic. This allows for network communication across greater 
    distances.
```

##### Hub
```
Problem Statement
  * If a new host joins the network, it needs to be connected with all the existing hosts in that network. Directly 
    connecting the new host to existing hosts through wires works, but is not a scalable solution.

Solution
  * Hubs act as a multi-port repeater that all existing and new hosts in the network can connect to.
  * Data sent from one host is broadcasted to all hosts connected to the hub.
```

### Layer 2
---
##### Data Link Layer
```
a) Data link layer bundles physical layers into a Local Area Network (LAN).
b) Data link layer is responsible for correct network hops and putting/receiving bits into the physical layer.
c) Data link layer abstracts stream of 'bits' into 'frames' and use MAC addresses as an address scheme. The header for 
   Layer 2 holds the src/dest MAC addresses.
d) Technologies include network interface cards (NICs), wifi access cards, bridges, and switches.
```

##### Network Interface Card (NIC)
```
NICs are a hardware component that allows for network connectivity. Each NIC hold a unique MAC address for the host 
machine.
```

##### Bridge
```
Problem Statement
  * Hubs cause all hosts connected to the hub to know about data that they do not need to be involved in.

Solution
  * Bridges sit between two hubs and have two ports each for one of the hubs.
  * Bridges know which hosts are on what side of either port.
  * When a host sends data and the destination is connected on the same hub, the bridge will prevent transmitting data 
    to the other hub. 
```

##### Switch
```
Problem Statement
  * All hosts on one side of a bridge still receive data that they might not be involved in.
  * All hosts on both sides of the bridge will receive data if the source and destination are on opposite sides of the 
    bridge.

Solution
  * Switches are a combination of hubs and bridges that help to connect L1 networks to form an L2 network. All hosts in 
    the same L2 network will share a common IP address space (i.e. prefix). 
  * Switches have multiple ports and know which hosts are on each port through a 'MAC address table'. Switches will use
    this table to determine which host to forward requests to.

MAC Address Table
  * A table that maps ports to MAC addresses of connected hosts.

Switch Functionality
  * Learn: Switches update their MAC address table with a <src-mac>:<port> mapping when a new frame passes the switch.
  * Forward: Use mapping on MAC address table to send frame on the appropriate port.
  * Flood: When a destination MAC address is not found on the MAC address table, the switch duplicates the frame to all 
           hosts except to the receiving port. Irrelevant hosts will drop the request and only the relevant host will 
           send a response back, which again causes an update on the MAC address table.
```

### Virtualization of Layer 2
---
##### Virtual Local Area Network (VLAN)
```
Problem Statement
  * Before virtualization, isolated switches were required for each and every network that needed to be isolated.

Solution
  * Switches can be logically separated through virtualization. Ports on a switch can be grouped into isolated 
    "mini-switches".
  * VLANs allow a single physical switch to be split into multiple virtual switches. VLANs allow a single virtual switch 
    to be extended across other physical switches.
```

##### Trunk Ports
```
Problem Statement
  * Assume that VLAN 1 and VLAN 2 are extended across the same 2 physical switches. This would normally require 2 wire 
    connections, one between ports for VLAN 1 and another between ports for VLAN 2. With more VLANs across physical 
    switches, this becomes difficult to scale.

Solution
  * Trunk ports allow data to flow from multiple VLANs in a single physical wire.
  * When data flows into a single wire, it becomes difficult to know which VLAN the data is intended for. "VLAN tags" 
    are added on top of existing L2 and L3 headers to distinguish which VLAN the packet is intended for.

Access Ports
  * Links that carry data just for a single VLAN.
  * When data flows into access ports, the network knows that the data is intended for a single VLAN.

Native VLAN
  * If a packet goes through a trunk port but does not have a VLAN tag, it will be sent to the Native VLAN as a default.
    This means that when sending a packet through a trunk port, the packet does not need a VLAN tag if the packet is 
    already intended for the Native VLAN. 

More info here: https://www.youtube.com/watch?v=MmwF1oHOvmg
```

##### Virtual Extensible Local Area Network (VXLAN)
```
Problem Statement
  * VLAN tags only support up to a maximum of 4096 (12 bits) following 820.1Q standards.

Solution
  * VXLANs encapsulate frames with a VXLAN header into UDP packets to resolve the inability for VLANs to be routed out 
    of L2 networks.
  * VXLANs are identified by a 24 bit VXLAN network identifier (VNI), allowing up to 16,777,216 VLANs.
```

##### Virtual Ethernet (VETH)
```
Virtualization of ethernet that act as tunnels between network namespaces without the need of assigning physical 
hardware. Virtual ethernet devices are typically created in pairs and are connected via a bridge.
```
