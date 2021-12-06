* `ulimit -Ha` display hard limits
* `ulimit -Sa` display soft limits
* `ulimit -Sn 2048` set soft limit of open files
* `su - tom -c "ulimit -Ha"` display hard limits of user tom
* `su - tom --shell /bin/bash -c "ulimit -Ha"` same as above but explicit on running which shell
* `cat /proc/PID/limits` display ulimit for a process (as shell access can be blocked for some user accounts)
