* a utility for investigating network sockets
* meant to replace netstat
* `ss -l` display listening server sockets
* `ss -a` display all sockets (default: connected)
* `ss -i` display interface table
* `ss -s` display summary of socket usage
* `ss -e` display detailed socket info
* `ss -n` show IP not hostname
* `ss -p` display PID/program name for sockets
* `ss -t` only tcp sockets
* `ss -u` only udp sockets
* `ss --unix` shows unix domain sockets

* meant to replace netstat?
* The main reason a user would use the ss command is to view what connections are currently established between their local machine and remote machines, statistics about those connections, etc.
* `ss -s` for summary
