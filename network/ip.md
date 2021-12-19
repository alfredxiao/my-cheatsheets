# `ip`
* from `iproute2` package
* replacing `ifconfig`
* does extra work of several other legacy commands such as `route` and `arp`.

## `ip addr`
* `ip addr show`
  * `ip a s`
* `ip addr help`
* `ip link help`
* `ip addr show dev em1`
  * show info about a device
* `ip addr add IP/CIDR dev em1`
  * add an address to a device
* `ip addr del IP/CIDR dev em1`
  * remove an address from a device
* `ip addr add IP/CIDR boradcast 192.168.1.255 dev em1`
  * add an IP and specific broadcast address to a device
## `ip link`
  * show info about all interfaces
* `ip link show dev em1`
  * show info about one device
* `ip -s link`
  * show interface statistics
  * `ip -s link show dev eth1`
* `ip link set`
  * alter the status of an interface
  * `ip link set em1 mtu nnnn`
  * `ip link set em1 promisc on`
    * set a network interface to promiscuous mode
  * `ip link set em1 up/down`
    * bring a device up or down

## `ip neigh`
* `ip neigh show` --> arp cache
* `ip neigh show dev eth0`
* `ip neigh add 192.168.1.10 lladdr 0e:0f:20:c3:c5:89 dev eth1`
* `ip neigh del 192.168.1.10 dev eth1`

## `ip route`
* `ip route show` display routing table
* `ip route add 10.0.2.0/24 via 10.0.2.10 dev eth1` add a route
* `ip route del 10.0.2.0/24 via 10.0.2.10 dev eth1` remove a route
* `ip route add default via 10.0.2.10` add a default gateway
* `ip route add prohibit 10.0.2.0/24` block destination route, respond icmp
* `ip route add blackhole 10.0.2.0/24` block destination route, drop
