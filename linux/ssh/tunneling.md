# X11 Forwarding
1. Make sure `X11Forwarding yes` is on remote hosts `/etc/ssh/sshd_config`
2. Client connect via `ssh -X user1@host1`
3. Run `xclock` for example

# TCP Forwarding (from local to destination via ssh server)
- Syntax: `ssh -L [LOCAL_IP:]LOCAL_PORT:DESTINATION:DESTINATION_PORT [USER@]SSH_SERVER`
- Example: `ssh -L 3336:db001.host:3306 user@pub001.host`
- Example (multiple): `ssh -L 3336:db001.host:3306 3337:db002.host:3306 user@pub001.host`
- When the destination host is the same as the SSH server, instead of specifying the destination host IP or hostname, you can use `localhost`

# TCP Forwarding (from ssh server to local to destination)
- Syntax: `ssh -R [REMOTE:]REMOTE_PORT:DESTINATION:DESTINATION_PORT [USER@]SSH_SERVER`
- Example: `ssh -R 8080:127.0.0.1:3000 -N -f user@remote.host`
  - ssh server listens on `8080`, any connection to that port will be forwarded to this machine(ssh client)'s `3000`
  - `-N` for not execute command
  - `-f` for background running

# TCP Forwarding (SOCKS proxy)
- Syntax: `ssh -D [LOCAL_IP:]LOCAL_PORT [USER@]SSH_SERVER`
- Example: `ssh -D 9090 -N -f user@remote.host`
  - Then an application can use this `9090` as SOCKS proxy to connect to any host/port via ssh server
