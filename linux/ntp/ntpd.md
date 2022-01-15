# config
- `/etc/ntp.conf`
- default port is `123` (UDP)

# iburst
```
server 0.centos.pool.ntp.org iburst
server 1.centos.pool.ntp.org iburst
server 2.centos.pool.ntp.org iburst
server 3.centos.pool.ntp.org iburst
```
- iburst helps to speed up the initial time synchronization by sending several queries in the first minute
