# Config file locations
- Debian-based `/etc/network/interfaces` file
- Red Hat–based `/etc/sysconfig/network-scripts` directory
- OpenSUSE `/etc/sysconfig/network` file

# Debian
## Static IP
```
auto enp0s5
iface enp0s5  inet static
 address 192.168.2.236
 netmask 255.255.255.0
 gateway 192.168.2.254
 dns-search sweet.home
 dns-nameservers 192.168.2.254 8.8.8.8
```
- Lines beginning with the word `auto` are brought up at boot time
## DHCP
```
auto eth0
iface eth0 inet dhcp
```
## Restart Network Service
- `systemctl restart networking.service`

# Ubuntu
- `/etc/netplan`
- yaml file
```
network:
  version: 2
  ethernets:
    ens5:
      dhcp4: no
      addresses: [10.9.8.7/24]
      gateway: 10.9.8.1
      nameservers:
        addresses: [8.8.8.8,8.8.4.4]
      dhcp6: true
      match:
        macaddress: 02:b9:28:48:de:22
      set-name: ens5
```

# CentOS
## Static IP
```
$ cat /etc/sysconfig/network-scripts/ifcfg-enp0s8
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=none
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME=enp0s8
UUID=3b77534a-e86b-4ceb-abaf-051804059a97
DEVICE=enp0s8
ONBOOT=yes
IPADDR=192.168.56.31
PREFIX=24
```
## DHCP
```
DEVICE="eth0"
ONBOOT=yes
NETBOOT=yes
UUID="41171a6f-bce1-44de-8a6e-cf5e782f8bd6"
IPV6INIT=yes
BOOTPROTO=dhcp
HWADDR="00:08:a2:0a:ba:b8"
TYPE=Ethernet
NAME="eth0"
DNS1="8.8.8.8"
DNS2="8.8.4.4"
```
## Restart Network Service
- `systemctl restart network`

## `/etc/sysconfig/network` config file
```
NETWORKING=yes
HOSTNAME=mysystem
GATEWAY=192.168.1.254
IPV6FORWARDING=yes
IPV6_AUTOCONF=no
IPV6_AUTOTUNNEL=no
IPV6_DEFAULTGW=2003:aef0::23d1::0a10:0001
IPV6_DEFAULTDEV=eth0
```
## `/etc/hostname` file
- specify hostname
