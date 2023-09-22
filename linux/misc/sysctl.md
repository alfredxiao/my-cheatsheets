# View all
- `sysctl -a`

# View single
- `sysctl net.ipv4.ip_forward`

# Update single
- `sysctl -w net.ipv4.ip_forward=1`

# Edit via config file
- `echo 0 > /proc/sys/net/ipv4/ip_forward`

# Permament
- `echo 'net.ipv4.ip_forward=1' >> /etc/sysctl.conf`
- or make the change in `/etc/sysctl.d`
