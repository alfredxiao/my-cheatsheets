# Edit `/etc/sudoers` file
* `visudo`

# Placing an extra sudoers config
* `visudo -f /etc/sudoers.d/myextra`
```
Cmnd_Alias WEB = /bin/systemctl restart httpd.service, /bin/systemctl reload httpd.service

user2 ALL=WEB
```

* do something as another user or superuser
* will result in an entry placed in a log file. Each entry includes the name of the user who executed the command, the command that was executed and the date and time of execution
  * /var/log/auth.log

# Invalidate credential cache
- `sudo -k`

# Check what current user can do with `sudo`
- `sudo -l`
- `sudo -ll` for long listing
- `sudo -U user1 -l` (run this as root)

# Validate credential cache
- `sudo -v`
- will prompt for password if cache expired (default to 5 mins), or will not prompt and will extend the cache liveness

# Execute as another user
- `sudo -u user1 touch user1file`

# Execute as another group
- `sudo -g group1 touch group1file`
