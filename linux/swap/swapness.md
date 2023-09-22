- Swappiness is a Linux kernel property that defines how often the system will use the swap space. Swappiness can have a value between 0 and 100. A low value will make the kernel to try to avoid swapping whenever possible, while a higher value will make the kernel to use the swap space more aggressively.

# check current value
- `cat /proc/sys/vm/swappiness`

# Update
- `sudo sysctl vm.swappiness=10`
- `/etc/sysctl.conf`
  - `vm.swappiness=10`
