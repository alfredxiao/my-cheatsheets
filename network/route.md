`route`
* display routing table
* `route -n` print numeric values only
* being replaced by `ip route show`

# `route`
* `route` -> shows route table
* `route add -net 10.0.2.0/24 gw 10.0.2.10 eth1` --> add a route
* `route del -net 10.0.2.0/24 gw 10.0.2.10 eth1`
* `route add default gw 10.0.2.10` --> add a default gateway
* `route add -host 10.0.2.10 reject` -> block destination route for a host
* `route add -net 10.0.2.0 netmask 255.255.255.0 reject` --> block a network

# `0.0.0.0`
- `0.0.0.0` has the specific meaning "unspecified". This roughly translates to "there is none" in the context of a gateway. Of course, this assumes that the network is locally connected, as there is no intermediate hop.
