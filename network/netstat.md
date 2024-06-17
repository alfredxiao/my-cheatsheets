* `netstat -r` -> route table, same as `route` command
* replaced by `ss`
* `-l` display listening server sockets
* `-a` display all sockets (default: connected)
* `-i` display interface table
* `-s` display networking statistics
  - `-su` stats for udp
  - `-st` stats for tcp
* `-e` display detailed socket info
* `-v` verbose
* `-n` show IP not hostname
* `-p` display PID/program name for sockets
* `-t` only tcp sockets (if no `-a` and no `-l`, then lists only ESTABLISHED sockets)
* `-u` only udp sockets (if no `-a` and no `-l`, then lists only ESTABLISHED sockets)
* `-i` for interface activity details
  - `RX-OK` packets received
  - `TX-OK` packets transmitted
- `-c` for continuously update the output
  - `-ci` to monitor interface activity 

* `netstat -tln`: Currently open tcp ports
  * `-t` for tcp
  * `-l` for listening
  * `-n` for numeric
