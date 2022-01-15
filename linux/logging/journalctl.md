
# how many entries
- `sudo journalctl` then `G`

# Most recent n entries
- `journalctl -n 20`

# no wrapping, no pager
- `journalctl --no-pager`

# jump to end
- `journalctl -e`

# Reverse order
- `journalctl -r`

# show messages from current user or kernel
  `journalctl --user`
  `journalctl --system`

# Inspect boot sessions
## list out all boot sessions
- `sudo journalctl --list-boots`
## inspect entries from a boot session
- `sudo journalctl -b 0` where `0` output from above `--list-boots`
- `sudo journalctl -b -1` -1 means previous one (if exists)

# Kernel messages
- `sudo journalctl -k`

# a Unit only
- `sudo journalctl -u dbus.service` or `--unit`

# Follow mode
- `sudo journalctl -u dbus.service --follow`

# Filter by tag
- `sudo journalctl -t myapp`

# By fields
- `sudo journalctl _SYSTEMD_CGROUP=/user.slice/user-1001.slice/session-1.scope`
- `sudo journalctl _UID=xxxx`
- `sudo journalctl _PID=xxx`
- `man systemd.journal-fields` to see all available fields

# Filter by datetime range
## Since
- `sudo journalctl --since=today` for entries from today, or `-S`
- `sudo journalctl --since=-5m` entries from last 5 minutes
## Until
- `sudo journalctl --since="2020-07-07 16:31" --until="2020-07-08 16:25"` or `-U`


# Filter by priority (severity)
- `journalctl -p "emerg".."crit"`

# add explanation to entires
- `journalctl -x`

# Output JSON
- `journalctl -o json` oneline json
- `journalctl -o json-pretty` pretty json

# Specify a file or directory for journal files
- `journalctl -D /path/to/mydir`
- `journalctl --file=/path/to/myfile`

# Administrative
- `journalctl --list-catalog` for message catalog
- `journalctl --disk-usage` for storage space used
- `sudo journalctl --system --disk-usage` space used for the entire system
# cleanup old entries
- `sudo journalctl --vacuum-size=10M` removes archived journal files until the disk space they use falls below the specified size (specified with the usual "K", "M", "G", "T" suffixes)
- `sudo journalctl --vacuum-time=?` you designate the oldest journal entries allowed, and the rest are deleted. e.g. `10months` `(s, min, h, days, months, weeks, or years)`
# flush entries from memory to disk
  `sudo journalctl --sync`
  `sudo journalctl --flush`
