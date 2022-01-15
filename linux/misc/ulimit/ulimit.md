* `ulimit -Ha` display hard limits
* `ulimit -Sa` display soft limits
* `su - tom -c "ulimit -Ha"` display hard limits of user tom
* `su - tom --shell /bin/bash -c "ulimit -Ha"` same as above but explicit on running which shell
* `cat /proc/PID/limits` display ulimit for a process (as shell access can be blocked for some user accounts)
* These limits are properties of a process. The limits apply independently for each process, **NOT** for each user account that may have many processes running.

* `ulimit -Sn 2048` set soft limit of open files
* `ulimit -n 2048` set **hard** limit of open files, or with `-Hn`

# Options
```
-c     size of core files created
-d     process data segment size
-e     scheduling priority ("nice")
-f     file size allowed to be written
-i     number of pending signals
-l     size of memory that can be locked
-m     resident set size
-n     number of open file descriptors
-p     pipe size in 512k blocks
-q     number of bytes in POSIX message queues
-r     real-time scheduling priority
-s     stack size
-t     amount of cpu time in seconds
-u     number of processes the user can run simultaneously
-v     amount of virtual memory available
-x     number of file locks
```
