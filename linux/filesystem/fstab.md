For permanent storage devices, Linux maintains the `/etc/fstab` file to indicate which drive devices should be mounted to the virtual directory at boot time.

If you use the `eCryptfs` filesystem type on any partitions, they will appear in the `/etc/crypttab` file and will be mounted automatically at boot time. While the system is running, you can also view all currently mounted devices, whether they were mounted automatically by the system or manually by users, by viewing the `/etc/mtab` file.



- `Device` – the first field specifies the mount device. These are usually device filenames. Most distributions now specify partitions by their labels or UUIDs.
- `Mount point` – the second field specifies the mount point, the directory where the partition or disk will be mounted. This should usually be an empty directory in another file system.
- `File system type` – the third field specifies the file system type.
Options – the fourth field specifies the mount options. Most file systems support several mount options, which modify how the kernel treats the file system. You may specify multiple mount options, separated by commas.
- `Backup operation` – the fifth field contains a `1` if the dump utility should back up a partition or a `0` if it shouldn’t. If you never use the dump backup program, you can ignore this option.
- `File system check order` – the sixth field specifies the order in which fsck checks the device/partition for errors at boot time. A `0` means that fsck should not check a file system. Higher numbers represent the check order. The root partition should have a value of `1` , and all others that need to be checked should have a value of `2`.
