# syslog
- In the mid-1980s Eric Allman defined a protocol for logging events from his Sendmail mail application called syslog. The syslog protocol quickly became a de facto standard for logging both system and application events in Unix, and it made its way to the
Linux world.

# Facility
- The type of event is defined as a facility value
- `0` kern
- `1` user
- `2` mail
- ...

# Severity
- `0/emerg` The event causes the system to be unusable
- `1/alert` An event that requires immediate attention
- `2/crit` An event that is critical but doesn’t require immediate attention
- `3/err` An error condition that allows the system or application to continue
- `4/warning` A non-normal warning condition in the system or application
- `5/notice` A normal but significant condition message
- `6/info` An informational message from the system
- `7/debug` Debugging messages for developers










# older method of doing system logging is two daemons
* `syslogd`
* `klogd`

# Recent distributions, these two are combined into a single daemon
* `rsyslogd`

# Even newer systems based on systemd, the daemon is named
* `journald`
  * allows both text and binary output
* viewing tool is `journalctl`
