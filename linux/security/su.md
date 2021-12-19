# Basics
* to switch to anyuser or superuser
* to run a command with substitute user and group ID

# switch to a user
- `su user1`
- `su` means `su root`

# switch with login (run login init scripts)
- `su -` or `su -l` or `su --login`

# Run a command as anyuser
- `su user1 -c 'touch user1file'`
