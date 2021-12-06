# Start
* `nice -n xx mycmd`
  * `xx` range from -20~19, the lower, the higher priority, defaults to 0

# Modify
* `renice -n nn -p PID`
* `renice -n nn -u user1`
* `renice -n nn -p PID1 PID2`
* **note** decreasing nice value requires `sudo`
