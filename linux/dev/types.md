There are two types of device files in Linux, based on how Linux transfers data to the device:
* Character device files: Transfer data one character at a time. This method is often used for serial devices such as terminals and USB devices.
* Block device files: Transfer data in large blocks of data. This method is often used for high-speed data transfer devices such as hard drives and network cards.
* This can be told by `ls` and whether an entry starts with `c` or `b`

# device mapper
* Besides device files, Linux also provides a system called the device mapper. The device mapper function is performed by the Linux kernel. It maps physical block devices to virtual block devices. These virtual block devices allow the system to intercept the data written to or read from the physical device and perform some type of operation on them.
* The device mapper creates virtual devices in the /dev/mapper directory. These files are links to the physical block device files in the /dev directory.
