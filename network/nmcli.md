# List all the available device
`nmcli dev status`

# List all the available connections
`nmcli con show`

# show configs of an interface
`nmcli con show eth2`

# create a new ethernet connection with static IP
`nmcli con add con-name eth9 type ethernet ifname enp0s8 ipv4.method manual ipv4.address 192.168.56.100/24 ipv4.gateway 192.168.56.1`

# create a new ethernet connection with DHCP
`nmcli con add con-name eth10 type ethernet ifname enp0s8 ipv4.method auto`

# start a connection
`nmcli con up eth9`
- same goes with `down` and `delete`

# Grab a single field
`nmcli -g ip4.address connection show eth1` grab the value only
`nmcli -f ipv4.dns,ipv4.addresses,ipv4.gateway con show eth1` prints both field name and value
