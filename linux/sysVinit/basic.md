# Find out which init is running
- `which init`
- `ls -l /sbin/init`
- `ps -p 1`

The `init` program implementations
* `SysV`:
  * A service manager based on the SysVinit standard controls which daemons and resources will be available by employing the concept of runlevels. Runlevels are numbered 0 to 6 and are designed by the distribution maintainers to fulfill specific purposes. The only runlevel definitions shared between all distributions are the runlevels 0, 1 and 6.
* `systemd`
  * systemd is a modern system and services manager with a compatibility layer for the SysV commands and runlevels. systemd has a concurrent structure, employs sockets and D-Bus for service activation, on-demand daemon execution, process monitoring with cgroups, snapshot support, system session recovery, mount point control and a dependency-based service control. In recent years most major Linux distributions have gradually adopted systemd as their default system manager.
* `Upstart`
  * Upstart is a substitute to init. The focus of Upstart is to speed up the boot process by parallelizing the loading process of system services. Upstart was used by Ubuntu based distributions in past releases, but today gave way to systemd.

  # check current runlevel
  * `runlevel`
    * Example output `N 3`
      * `N` means previous runlevel not available
  * `who -r`
  * `chkconfig`

  # `/etc/inittab`
  * defines Runlevels and services/actions
  * `id:5:initdefault:` means default runlevel is 5

  # Changing runlevel
  * `telinit NEXT_LEVEL`
    * if you change to 6, will reboot the system
    * if you change from 5 to 3, for example, will close GUI and back to CLI

  # Scripts for runlevels
  * Usually stored in `/etc/init.d/` where it links to `/etc/rc*.d/`
  * The script responsible starts/stops services is file `/etc/rc.d/rc`
  * Each level has a directory namely `/etc/rc0.d/`, etc.
    * has symbolic links to those scripts in `/etc/init.d/`
    * Symbolic links starts with either `K` or `S`
      * `K` means the service will be killed when entering the runlevel (kill)
      * `S` means the service will be started when entering the runlevel (start)

  # `rc.local`
  * you can put extra initialisation stuff here that is run after **all** other init scripts are finished
