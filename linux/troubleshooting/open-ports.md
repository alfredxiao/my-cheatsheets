# What ports opened by a process
1. `ls -l /proc/PID/fd` will see socket and inode value
2. `cat /proc/PID/net/tcp grep INODE_VALUE` the second column is port number in hex code

# What processes opens a port
* `ss -tulp` list listening ports with program name (both tcp and udp)
* `ss -tuna` list both listening and non-listening ports (both tcp and udp)
* `netstat -tulpn` lists listening ports and processes
* `lsof -i` shows all open connections
* `lsof -i -P` shows port number rather than service name
* `lsof -i :3506` what opens port 3506
* `lsof -i -sTCP:LISTEN` shows LISTENING TCP ports
* `nmap host1` scan
* `fuser 80/tcp`
* `fuser 39287/tcp` will work if no matter this is the client socket or server socket port number
