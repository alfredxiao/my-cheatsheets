The project claims the “r” stands for “rocket fast.” Speed is the focus of the rsyslog project, and the `rsyslogd` application has quickly become the standard logging package for many Linux distributions.

# Config
- `/etc/rsyslog.conf` or `/etc/rsyslog.d/`
## Format
- `facility.priority action`
- `-/var/log/maillog` : `-` means allow caching of contents synced later to file
- `cron.none   /var/log/messages` : do **NOT** write cron log to this file
- `cron.crit ...` logs only messages of `CRIT` severity level **and** higher
- `cron.=crit ...` logs only messages of `CRIT` severity level and **NOT** higher
- `cron.*` means whatever severity level from the `cron` facility
- `auth,cron.crit` means `CRIT` messages from both `auth` and `cron` facility
- `auth.crit;cron.notice` means `CRIT` from `auth` and NOTICE from `cron`
- `*.emerg` means whatever facility at `emerg` severity level
- `~` as destination means discard it
- `*` as destination means all destinations

# Example config
```
local6.crit /var/log/mydemo.crit.log
local6.* /var/log/mydemo.general.log
:programname, isequal, "mydemo"   /var/log/mydemo.log
```

# Actions
- Forward to a regular file
- Pipe the message to an application
- Display the message on a terminal or the system console
- Send the message to a remote host
- Send the message to a list of users
- Send the message to all logged-in users

# Send to remote host
- `TCP|UDP[(z#)]HOST:[PORT#]`
- `*.* @@(z9)loghost.ivytech.edu:6514`
