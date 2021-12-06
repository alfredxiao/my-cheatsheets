# Grep process by name
* `pgrep ssh`

# show pid as well as process name
* `pgrep -l ssh` could show process name `ssh` or `sshd`, etc.
* `pgrep -a ssh` shows the process command line as well

# Process linked with a Terminal
* `pgrep -t tty2`
