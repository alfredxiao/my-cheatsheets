# Types of Disks
* HDD: Magnetic, slow
* SSD: IC, fast

# Types of Connections (between disk and system)
* PATA
* SATA
* SCSI

# Device file
* Linux kernel assigns a connected drive a file in the /dev directory. That file is called a **raw device**, as it provides a path directly to the drive from the Linux system. Any data written to the file is written to the drive, and reading the file reads data directly from the drive.
* The devices are all? block devices
* `/dev/hda` and `/dev/hdb` are reserved for the master and slave devices on PATA devices
* `/dev/hdc` for CD/DVD drive connected to the second IDE channel
* `/dev/fd0` and etc. are for floppy drive
* `/dev/sdx` From Linux kernel 2.4 onwards, most storage devices now identified as if they were SCSI devices, regardless of their hardware type. IDE, SSD and USB block devices will be prefixed by `sd`.
* SD cards, the paths `/dev/mmcblk0p1`, `/dev/mmcblk0p2`, etc. are used for the first and second partitions of the device identified first and `/dev/mmcblk1p1`, `/dev/mmcblk1p2`, etc.

# Partition
* a drive is partitioned into multiple partitions
* Linux creates /dev files for each separate disk partition

# Udev
* Removable devices are handled by the `udev` subsystem, which creates the corresponding devices in `/dev`
* In current linux distributions, `udev` is responsible for the identification and configuration of the devices already present during machine power-up (`coldplug` detection) and the devices identified while the system is running (`hotplug` detection).
* `Udev` relies on `SysFS`, the pseudo filesystem for hardware related information mounted in `/sys`.
* As new devices are detected, udev searches for a matching rule in the pre-defined rules stored in the directory `/etc/udev/rules.d/`. The most important rules are provided by the distribution, but new ones can be added for specific cases.

# LVM
* uses device mapper

# RAID 
