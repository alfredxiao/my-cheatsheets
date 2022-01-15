* `if` means `interface`
* is limited primarily to modification of networking parameters, and displaying the configuration details of networking components

- from `net-tools` package

`ifconfig enp0s3 down 10.0.2.10 netmask 255.255.255.0`

* `ifconfig -a` -> display info about all interfaces (including down ones)
* `ifconfig eth0` --> display info about one interface
* `ifconfig eth0 up/down` --> bring one interface up/down
* `ifconfig eth0 192.168.1.20` --> assign an IP to an interface
* `ifconfig eth0 netmask 255.255.255.0`
* `ifconfig eth0 broadcast 192.168.1.255`
* `ifconfig eth0 0.0.0.0` --> remove an IP address from a network interface
* `ifconfig eth0 mtu nnn`
* `ifconfig eth0 promisc`

# `ifup` and `ifdown`
- `ifup eth0`
