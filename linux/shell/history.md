`history` and `!`
* `history 10`
* `!!` to run last command
* `!ls` to run last `ls` command (with same arguments)
* `!-2` to run previous command's previous command
* `!123` to run command numbered 123 (in history output)
* `$HISTSIZE` to control how many entries
* `$HISTFILE` points to the file with history records, typically `.bash_history`
* `$HISTFILESIZE` determine how many lines the .bash_history file  will contain

* to wipe out history
  * `history -c` then `history -w`
