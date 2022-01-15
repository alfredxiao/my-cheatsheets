- logrotate runs as cron job, not daemon
- The logrotate utility does more than rotate log files. It can also compress, delete, and if desired, mail a log file to a designated account.

# configuration
- location `/etc/logrotate.conf`
- location `/etc/logrotate.d/` with config file name matching log file name, file content similar to below `/var/log/wtmp` section

```
# rotate log files weekly
weekly

# keep 4 weeks worth of backlogs (as rotate max 4 times, and rotate weekly)
rotate 4

# create new (empty) log files after rotating old ones
create

# use date as a suffix of the rotated file
dateext

# uncomment this if you want your log files compressed
#compress

# RPM packages drop log rotation information into this directory
include /etc/logrotate.d

# no packages own wtmp and btmp -- we'll rotate them here
/var/log/wtmp {
    monthly
    create 0664 root utmp
        minsize 1M
    rotate 1
}
```
- If the `dateext` option is not employed, the rotated log files will have a number as their file extension, with the biggest number indicating the oldest log file
## Other config options
- `hourly` this requires modification to cron job configuration
- `weekly n` Log file is rotated weekly on the n day of the week, where 0 is equal to Sunday. 7 is a special number that indicates the log file is rotated every 7 days, regardless of the current day of the week.
- `monthly`
- `size n` Rotates log file based on size and not time, where n indicates the file’s size that triggers a rotation (n followed by nothing or k assumes kilobytes, M indicates megabytes, and G denotes gigabytes).
- `rotate n` Log files rotated more than n times are either deleted or mailed, depending on other directives. If n equals `0`, rotated files are **deleted**, instead of rotated.
- `dateformat FORMAT-STRING` Modify the dateext setting’s date string using the format-string specification
- `missingok` If log file is missing, do not issue an error message and continue on to the next log file.
- `notifempty` If the log file is empty, do not rotate this log file, and continue on to the next log file.

# Status file
- `/var/lib/logrotate/logrotate.status` (Redhat) or `/var/lib/logrotate/status` (Debian)

# Run for a config
- `logrotate -f -v /etc/logrotate.d/mylog`
