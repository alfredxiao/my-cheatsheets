Module Parameters
* Can be made persistent in `/etc/modprobe.conf`
  * or in individual config file (`.conf`) in `/etc/modprobe.d/`

Module Blacklist
* Adding a line `blacklist amodule` to `/etc/modprobe.d/blacklist.conf` can cause kernel to not load the module.
