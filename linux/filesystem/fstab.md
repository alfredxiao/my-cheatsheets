For permanent storage devices, Linux maintains the `/etc/fstab` file to indicate which drive devices should be mounted to the virtual directory at boot time.

If you use the `eCryptfs` filesystem type on any partitions, they will appear in the `/etc/crypttab` file and will be mounted automatically at boot time. While the system is running, you can also view all currently mounted devices, whether they were mounted automatically by the system or manually by users, by viewing the `/etc/mtab` file.
