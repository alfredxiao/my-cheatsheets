
The IPv6 protocol also provides for two different types of host addresses:
- Link local addresses
- Global addresses

IPv6 link-local addresses are addresses that can be used to communicate with nodes (hosts and routers) on an attached link. Packets with those addresses are not forwarded by routers. At least, they should not be. There have been cases where routers would happily forward packets with a link-local source address.

The **link local address** uses a default network address of `fe80::` and then derives the host part of the address from the media access control (MAC) address built into the network card. This ensures that any IPv6 device can automatically communicate with any other IPv6 device on a local network without any configuration.

The **IPv6 global address** works similarly to the original IP version: each network is assigned a unique network address, and each host on the network must have a unique host address.
