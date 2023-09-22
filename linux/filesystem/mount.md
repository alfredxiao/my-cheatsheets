# List current mounts
- `mount`
  - derived/same as /etc/mtab

# List by type
- `mount -t ext4` —> shows all ext4 mount points

# Mount a device at an endpoint
- `mount -t ext4 /dev/sdb1 /opt`

# Mount and set a label
- `mount -L OPT -t xfs -o rw,noexec /opt`

# Mount an ISO file
- `mount /myiso.iso -o ro,loop /media/myiso`

# Remount
- means mount without umount first
- `mount -o remount,ro /mnt/data`

# MOunt all from fstab
- `mount -a`
- **NOTE** this does not include those swap devices or files, which require `swapon -a`

# Swap
* To mount a partition created as a swap area, you don’t use the `mount` command but instead use the `swapon` command. The kernel will automatically flag the partition to be used as part of the virtual memory structure.
