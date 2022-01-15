# rsyslog as journald client
- `/etc/rsyslog.conf`
```
$ModLoad imuxsock # provides support for local system logging [...]
$ModLoad imjournal # provides access to the systemd journal
```
# forward to syslog
- forwarded to `/run/systemd/journal/syslog`
- `ForwardToSyslog=yes` in `/etc/systemd/journald.conf`
- we cannot `systemctl reload systemd-journald` but have to use `restart`!

# modules
- `omjournal` is a module what writes to journal (from syslog side)
