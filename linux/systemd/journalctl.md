
* how many entries
  `sudo journalctl` then `G`
* no wrapping, no pager
  `journalctl --no-pager`
* show messages from current user or kernel
  `journalctl --user`
  `journalctl --system`
* create entries manually
  `echo "my message" | systemd-cat`
* list out all boot sessions
  `sudo journalctl --list-boots`
* inspect entries from a boot session
  `sudo journalctl -b 0`
  * `0` is the most recent boot
* View entries for a service
  `sudo journalctl -u dbus.service`
* to follow in real time
  `sudo journalctl -u dbus.service --follow`
* view a user session on the system
  `systemd-cgls`
* view journal entries from a user scope
  `sudo journalctl _SYSTEMD_CGROUP=/user.slice/user-1001.slice/session-1.scope`
* view entries from today
  `sudo journalctl --since=today`
* view entries from last 5 minutes
  `sudo journalctl --since=-5m`
* view entries over time range
  `sudo journalctl --since="2020-07-07 16:31" --until="2020-07-08 16:25"`
* view entries from a uid
  `sudo journalctl _UID=xxxx`
* view entries from a pid
  `sudo journalctl _PID=xxx`
* add explanation to entires
  `journalctl -x`
* inspect the message catalog entries available
  `journalctl --list-catalog`
* Find out how much space the current user journals take up
  `journalctl --disk-usage`
* find out the same for entire system
  `sudo journalctl --disk-usage`
* cleanup old entries
  `sudo journalctl --vacuum-size=10M`
  `sudo journalctl --vacuum-time=?`
* flush entries from memory to disk
  `sudo journalctl --sync`
  `sudo journalctl --flush`
