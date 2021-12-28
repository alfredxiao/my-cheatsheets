* `who` who is currently logged in
* `w` who is currently logged in, plus process being run, and cpu usage
* a user's command history
  * `~/.bash_history`
  * `/var/log/auth.log | grep username` (Debian)
  * `/var/log/secure | grep username` (Redhat)
* `last` who recently logged in
  * `pts` login device when connecting through the network or a console (e.g. ssh).
  * `tty` serial or console connections (text mode)
  * `:0` "local:display #0": X11 server, used for graphical login (e.g. gdm)
* `last -f /var/log/btmp` who recently had failed login attempts
  * `btmp` records only failed login attempts
  * `/var/run/utmp` and
  * `wtmp` complete picture of user logins at which terminals, logouts, system events, etc.
* `utmpdump path-to-?tmp` display raw records
* `ps -u user1` list processes started by a user
* `pkill -u user1` kills all processed started by a user
* `who -u` lists logged in user's source and PID (sshd)
  * kill that process to kick the user out
* `ps -u user1 t` lists process with terminal info for a user
  * where we can match up PID and TTY
  * `echo hello | write user1 pts/0`
* `lsof -u user1` open files by a user
* `lsof /adir` open files in a dir
* `find / -perm -u+s` find files have SUID bit on
  * or `find / -perm -04000`
  * `find / -perm -g+s` find files with SGID bit on
* `/etc/security/limits.conf`
