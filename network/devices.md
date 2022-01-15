
# Hub
- works at Network Access/Interface Layer
- Hubs do **not** perform packet filtering or addressing functions
- Hubs just send data packets to **all** connected devices
- is passive device, does not have software, not intelligent

# Repeater
- works at Network Access/Interface Layer
- connects two segments of a network cable
- it regenerates the signals to proper amplitudes and sends them to the other segment

# Switch
- has software, is intelligent, is active device
- The switch uses the MAC address to identify which attached device outgoing packets are being sent from and where to deliver incoming packets.
- When a device sends a packet to another device, it enters the switch and the switch reads its header to determine what to do with it. It matches the destination address or addresses and sends the packet out through the appropriate ports that leads to the destination devices.

# Bridge
- works at Network Access/Interface Layer
- A bridge is a repeater, with add on the functionality of filtering content by reading the MAC addresses of source and destination.
- A network bridge is a device that divides a network into segments.
- Each segment represent a separate collision domain, so the number of collisions on the network is reduced.
- Each collision domain has its own separate bandwidth, so a bridge also improves the network performance.
- It inspects incoming traffic and decide whether to forward it or filter it. Each incoming Ethernet frame is inspected for destination **MAC** address. If the bridge determines that the destination host is on another segment of the network, it forwards the frame to that segment.

# Router
- works at Network Layer
- connected to more than one networks
