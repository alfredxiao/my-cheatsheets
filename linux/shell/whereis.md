* locate the binary, source, or manual page for a command
* it searches for commands, system files (like `/etc/passwd`) as well as man pages
  * it does not search for regular user files
* similar to `which`, but which only searches for command, not manual page or source

# Trick
- `whereis mycmd | tr ' ' '\n'`
