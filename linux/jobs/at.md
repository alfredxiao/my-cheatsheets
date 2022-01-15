# Basics
- `at` is for scheduling one-off tasks
- `atd`, runs in the background and checks the job queue for jobs to run. Most Linux distributions start this automatically at boot time.
- The `atd` command checks a special directory on the system (usually /var/spool/at) for jobs submitted using the at command. By default, the `atd` command checks this directory every 60 seconds
- `at` is interactive and asks for user input for what command to run

# Permissions
- `/etc/at.allow` contains the names of the users who may use the `at` command to submit jobs.
- `/etc/at.deny` contains the names of users who are not allowed to use the `at` command to submit jobs.

# Time format
- `at now`
- `at now + 2 minutes`
- `at now + 7 days`
- `at teatime` meaning 4pm
- `at noon` today or tomorrow's noon
- `at now next hour` exactly 60 minutes from now
- `at now next day` exactly same time tomorrow
- `at 17:00 tomorrow`
- `at 4:45pm` today or tomorrow
- `at 3:00 Dec 28, 2018`
  - date part can be any of `MMDD[CC]YY`, `MM/DD/[CC]YY`, `DD.MM.[CC]YY` or `[CC]YY-MM-DD`
- `at -t 202005111321.32` [[CC]YY]MMDDhhmm[.ss]

# View job queues
- `atq`

# Cancel a job
- `atrm n` n is the job number from `atq`

# Submit a job without user interaction
- `echo "command_to_be_run" | at 09:00`
- or
```
at 09:00 <<END
command_to_be_run
END
```

# Read command from a file
- `at 09:00 -f /home/linuxize/script.sh`

# Notification
- if there is output (e.g. stdout from the job process), it is sent to the user's mail
## Suppress this mail
- `at 09:00 -M`
## Send this mail even if no output
- `at 09:00 -m`

# Queue/niceness
- defaults to `a` queue, `batch` jobs in `b` queue
- higher queue has higher niceness value
- `at -q h 09:00` specify queue `h`
