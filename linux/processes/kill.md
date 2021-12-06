# Kill a job
* `kill %n`

# Kill a process
* `kill PID`

# Kill with force
* `kill -9 PID`
* or `kill -s KILL PID`
* **note** signal TERM is 15, KILL is 9

# Kill with any sig number
* `kill -3 PID` (in JVM, this prints thread dump)

# Killall
* `killall httpd`

# show available signals
* `kill -l`
