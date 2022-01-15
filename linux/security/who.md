* who are logged in
* example: `sysadmin 	pts/0       2013-10-11 09:59 (:0.0)`
* fields
  * username
  * terminal
    * `tty` indicates a local login
    * `pts` indicates a remote login, namely pseudo terminal
  * date: when logged in  
  * location: could be CLI, a terminal in GUI, a remote machine
* `who -b` for boot time
* `who -u` for logged in users remote IP and their shell PID
