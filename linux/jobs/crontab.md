# Format
```
# ┌───────────── minute (0 - 59)
# │ ┌───────────── hour (0 - 23)
# │ │ ┌───────────── day of the month (1 - 31)
# │ │ │ ┌───────────── month (1 - 12)
# │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
# │ │ │ │ │                                   7 is also Sunday on some systems)
# │ │ │ │ │
# │ │ │ │ │
# * * * * * <command to execute>
```
- `* * * * * username <command to execute>` this format is used for system wide crontab at `/etc/crontab`
- `*/5` on the `minute` field means every 5 minutes
- `09-17` on the `hour` field means from 09 to 17

# Location
- `/etc/crontab` for system
- `/etc/cron.d/` for system
- `/etc/cron.[hourly|daily|weekly|monthly]` for system, and always executed as root
- `/var/spool/cron/` or `/var/spool/crontab/` for individuals

# Commands
## List current cron entries
- `crontab -l`
- `crontab -u username -l` view entries from another user
## Edit cron jobs
- `crontab -e`
- `crontab -u user1 -e`
- which creates a cron file `/var/spool/cron/user1`
## Remove cron entries
- `crontab -r` removes entries for current user

# Permissions
- `/etc/cron.deny` usernames listed here cannot create cron jobs, i.e. cannot run `crontab -e`
- `/etc/cron.allow` usernames listed here allowed to create cron jobs.

# Environment
```
# crontab -e
SHELL=/bin/bash
MAILTO=root@example.com
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin
# below are cron job entries
...
```
- The environment must be set to whatever is necessary for a given user because cron does not provide an environment of any kind.
- The `SHELL` variable specifies the shell to use when commands are executed. This example specifies the Bash shell.
- The `MAILTO` variable sets the email address where cron job results will be sent.
- The third line sets up the `PATH` for the environment. 
