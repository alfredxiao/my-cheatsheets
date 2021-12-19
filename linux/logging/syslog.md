# older method of doing system logging is two daemons
* `syslogd`
* `klogd`

# Recent distributions, these two are combined into a single daemon
* `rsyslogd`

# Even newer systems based on systemd, the daemon is named
* `journald`
  * allows both text and binary output
* viewing tool is `journalctl`
