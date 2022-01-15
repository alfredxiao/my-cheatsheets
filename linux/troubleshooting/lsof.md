# open files by a user
* `lsof -u auser` or `lsof -u usera,userb`-->
* `lsof -u ^auser` --> open files not by a user

# by file/dir
* `lsof /dev/sda1` by device
* `lsof +D /dir` shows all open files in a directory
* `lsof /path/to/myfile` by file

# by process id
* `lsof -p PID` or `lsof -p PID1,PID2`
* `lsof -p ^PID`

# by process name
* `lsof -c httpd` by process name e.g. httpd
* `lsof -c ttp` by process name substring

# by port and connection
* `lsof -i [46][protocol][@hostname|hostaddr][:service|port]` open files by a network connection
* `lsof -i :23` shows connections with port number 23
* `lsof -i tcp:23` shows connections with port number 23
* `lsof -i -sTCP:LISTEN` shows LISTENING TCP ports
* `lsof -i -sTCP:ESTABLISHED` shows ESTABLISHED TCP connections
* `lsof -iTCP` shows TCP connections (including listening and others)
* `lsof -iUDP` shows UDP connections
* `lsof -i6` shows IPv6 connections

# multiple criteria
- `lsof -a -u user1 -i` all network activity for a specific user
- `-a` means `AND` of all conditions (otherwise they are ORed)

# find processes by user
- `lsof -t -u user1`
- `kill $(lsof -t -u user1)` kill all processes by a user

# FD column of lsof output
- `cwd` current working directory
- `ltx` shared library text (code and data)
- `mem` memory mapped file
- `mmap` memory mapped device
- `pd` parent directory
- `rtd` root directory
- `txt` program text (code and data)
- `{digit}r` read-only
- `{digit}w` write-only
- `{digit}u` read/write
