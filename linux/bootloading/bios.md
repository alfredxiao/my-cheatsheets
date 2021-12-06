# Boot loader
* run by `BIOS`
* its job is to run full OS

# Where to find boot loader
* Most BIOS setup allows finding boot loader from
  * internal hard drive
  * external hard drive
  * CD/DVD
  * USB memory stick
  * Network server (HTTP, FTP, etc)

# MBR
* is the first sector on the first hard drive partition on the system
* There is only one MBR for the computer system
* BIOS finds boot loader from MBR

# Boot loader
* Very small, 512B
* its job is finding actual OS kernel file, which has no size limitation
* the boot loader does not need to point to OS kernel file directly, it can point to a a proxy program
  * that allows users to interactive choose which OS to load
  * e.g. GRUB
