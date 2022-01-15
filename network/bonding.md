Network bonding is a process of combing or joining two or more network interfaces together into a single interface. Network bonding offers performance improvements and redundancy by increasing the network throughput and bandwidth. If one interface is down or unplugged the other one will work. It can be used in situations where you need fault tolerance, redundancy or load balancing networks. In Linux, a special kernel module named bonding is used to connect multiple network interfaces into a single interface. Two or more network interfaces can be connected into a single logical “bonded” interface. The behavior of the bonded interfaces depends on the type of bonding method used.

# Load bonding module
`sudo modprobe bonding` (this is not persistent)

# Create bonding interface
`sudo ip link add bond0 type bond mode 4`

# Add backing interfaces to bonding interface
`sudo ip link set eth0 master bond0`
`sudo ip link set eth1 master bond0`

# Modes
## mode 0 (balance-rr)
- Round-robin policy.
- Transmits packets in sequential order from the first available slave through the last. This mode provides load balancing and fault tolerance.
## mode 1 (active-backup)
- Active-backup policy.
- Establishes that only one slave in the bond is active. A different slave becomes active if, and only if, the active slave fails. The bond's MAC address is externally visible on only one port (network adapter) to avoid confusing the switch. This mode provides fault tolerance.
## mode 2 (balance-xor)
- Transmits based on the selected transmit hash policy, which can be altered via the xmit_hash_policy option. This mode provides load balancing and fault tolerance.
- This sounds like a sticky approach?
## mode 3 (broadcast)
- Transmits everything on all slave interfaces. This mode provides fault tolerance.
## mode 4 (802.3ad)
- IEEE 802.3ad Dynamic link aggregation policy.
- Creates aggregation groups that share the same speed and duplex settings. Utilizes all slaves in the active aggregator according to the 802.3ad specification.
- Dynamic Link Aggregation requires a switch (that connects to two or more slave interfaces) that supports IEEE 802.3ad.
# mode 5 (balance-tlb)
- Adaptive transmit load balancing.
- Establishes channel bonding that does not require any special switch support.
- The outgoing traffic is distributed according to the current load (computed relative to the speed) on each slave. Incoming traffic is received by the current slave. If the receiving slave fails, another slave takes over the MAC address of the failed receiving slave.
# mode 6 (balance-alb)
- Adaptive load balancing.
- Includes balance-transmit load balancing plus receive-load balancing for IPv4 traffic, and does not require any special switch support. The receive-load balancing is achieved by ARP negotiation. The bonding driver intercepts the ARP replies sent by the local system on their way out and overwrites the source hardware address with the unique hardware address of one of the slaves in the bond. Thus, different peers use different hardware addresses for the server.
