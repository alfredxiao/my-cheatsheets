# Basics
- derived from `/proc`
- pid 1 is `init`/`systemd` - parent of all user space processes
- pid 2 is `kthreadd` - parent of all kernel processes

# View process tree
* `ps -xawf` to view ps tree
* `ps -ef --forest` to view ps tree
* `ps -eH` H for hierarchy

# Scope
* `ps -e` for all active processes

# Format
* `ps -f` for more columns
* `ps -F` for even more columns
  * `SZ` estimated swap size
  * `RSS` real memory usage
  * `PSR` processors assigned
* specify columns
  * `ps -eo pid,user,start_time,cmd`

# BSD Style
* `ps -aux`
  * similar to `ps -eF` but cpu and mem displayed as percentage

# sorting
* `ps -ef --sort user` sort by user

# by a user
* `ps -u user1`
* `ps -u 1001`

# by pid
* `ps nnn`
* `ps --pid nnn`
* `ps p nnn`

`sleep`
* `sleep 88` 88 seconds?

`ps -o`
* `ps -o pid,tty,time,%cpu,cmd` defines what columns to output

sort
* `ps -o pid,tty,time,%mem,cmd --sort %mem` or `+%mem`
* in descending order: `-%mem`

`pgrep`
