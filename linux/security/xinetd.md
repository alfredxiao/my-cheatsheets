
# Features
- access control for TCP and UDP
  - allows you to permit or block connections according to the requestor’s IP address
- event logging that records connection success and failure
- access control based on time segments
- concurrent server limit - DoS attack defense
- log file size limit
- allocation of services to different interfaces, enabling certain services to be limited to the private network
- proxy function including network address translation

# Configuration File
- `/etc/xinetd.conf` and `/etc/xinetd.d/`

# Example config
- `/etc/xinetd.d/sshd`
```
service ssh
{
 disable     = no
 socket_type = stream
 protocol    = tcp
 port        = 22
 wait        = no
 user        = root
 server      = /usr/sbin/sshd
 server_args = -i
}
```

# Config Settings
- `<only_from | no_access> = <address | address range>`
  - can use CIDR notation
- `cps`: connection per second.
  - e.g. `50 10` meaning max 50 connection per second
  - If the rate of incoming connections is higher than this, the service will be temporarily disabled. The second argument is the number of seconds to wait before re-enabling the service after it has been disabled.
  - Purpose of this is defend against DOS
- `instances` Determines the number of service processes that can be simultaneously active for a service.
- `max_load` Determines the one-minute load average that when reached, the net- work service stops accepting connections, until the load level drops.
- `per_source` Specifies the maximum number of network server processes that can be started per source IP address.
