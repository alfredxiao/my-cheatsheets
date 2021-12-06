`/` or `the root filesystem`
* is the top-level directory of the filesystem. It must contain all of the files **required to boot the Linux system** before other filesystems are mounted. It must include all of the required executables and libraries required to boot the remaining filesystems.
* While the root filesystem and its contents are considered essential or necessary to boot the system, the /var, /usr and /usr/local directories are deemed non-essential for the boot process.

`/bin`
* Essential binaries like the ls, cp, and rm commands, and be a part of the root filesystem

`/boot`
* Files necessary to boot the system, such as the Linux kernel and associated configuration files

`/dev`
* Files that represent hardware devices and other special files, such as the /dev/null and /dev/zero files

`/etc`
* Essential host configurations files such as the /etc/hosts or /etc/passwd files

`/home`
* User home directories

`/lib`
* Essential libraries to support the executable files in the /bin and /sbin directories

`/lib64`
* Essential libraries built for a specific architecture. For example, the /lib64 directory for 64-bit AMD/Intel x86 compatible processors

`/media`
* Mount point for removable media mounted automatically

`/mnt`
* Mount point for temporarily mounting filesystems manually

`/opt`
* Optional third-party software installation location

`/proc`
* Virtual filesystem for the kernel to report process information, as well as other information

`/root`
* Home directory of the root user

`/sbin`
* Essential system binaries primarily used by the root user

`/sys`
*	Virtual filesystem for information about hardware devices connected to the system
* diff between `/dev` ??

`/srv`
* Location where site-specific services may be hosted

`/tmp`
* Directory where all users are allowed to create temporary files and that is supposed to be cleared at boot time (but often is not)

`/usr`
* Second hierarchy
* Non-essential files for multi-user use
* The /usr directory is intended to hold software for use by multiple users. The /usr directory is sometimes shared over the network and mounted as read-only.

`/usr/local`
* Third hierarchy
* Files for software not originating from distribution
* The /usr/local hierarchy is for installation of software that does not originate with the distribution. Often this directory is used for software that is compiled from the source code

`/var`
* Fourth hierarchy
* Files that change over time

`/var/cache`
*	Files used for caching application data

`/var/log`
*	Most log files

`/var/lock`
* Lock files for shared resources

`/var/spool`
*	Spool files for printing and mail

`/var/tmp`
* Temporary files to be preserved between reboots

`User-Specific Binaries`
* The binary directories that are intended to be used by non-privileged users include:
  * `/bin`
  * `/usr/bin`
  * `/usr/local/bin`
* Sometimes third-party software also store their executable files in directories such as:
  * `/usr/local/application/bin`
  * `/opt/application/bin`

`Root-Restricted Binaries`
*  On the other hand, the sbin directories are primarily intended to be used by the system administrator (the root user). These usually include:

  * /sbin
  * /usr/sbin
  * /usr/local/sbin
  Some third-party administrative applications could also use directories such as:

  * /usr/local/application/sbin
  * /opt/application/sbin

`Software Application Directories`
* For Debian-derived distributions, you can execute the `dpkg -L packagename` command to get the list of file locations.
* In Red Hat-derived distributions, you can run the `rpm -ql packagename` command for * the list of the locations of the files that belong to that application.
The data for the application may be stored in one of the following subdirectories:
  * /usr/share
  * /usr/lib
  * /opt/application
  * /var/lib
* The file related to documentation may be stored in one of the following subdirectories:
  * /usr/share/doc
  * /usr/share/man
  * /usr/share/info

`Library Directories`
* The libraries that support the essential binary programs found in the /bin and /sbin directories are typically located in either /lib or /lib64.

* To support the /usr/bin and /usr/sbin executables, the /usr/lib and /usr/lib64 library directories are typically used.

* For supporting applications that are not distributed with the operating system, the /usr/local/lib and /opt/application/lib library directories are frequently used
