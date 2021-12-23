# partitioning
- Partition and choose type `fd`/`fd00` for Linux raid

# Install software
- `yum install -y mdadm`

# Create raid device
- `mdadm -C /dev/md0 -l raid5 -n 3 /dev/sdg1 /dev/sdg2 /dev/sdg3 -x 3 /dev/sdh1 /dev/sdh2 /dev/sdh3`
- or `mdadm -C /dev/md0 -l raid5 -n 3 /dev/nvme1n1p[1-3] -x 3 /dev/nvme2n1p[1-3]`
- `-C` to name the device to be created
- `-l` to specify RAID level
- `-n` to specify number of active devices used for the RAID device
- `-x` to specify number of spare devices used for the RAID device

# View stat
- `ll /dev/md*`
- `cat /proc/mdstat`
- `mdadm -D /dev/md0`

# Make this persistent (across reboots)
- `mdadm -D -s -v > /etc/mdadm.conf`

# Create Filesystem and Mount
- `mkfs -t ext4 /dev/md0`
- `mkdir /mnt/data`
- `mount -t ext4 /dev/md0 /mnt/data`

# add mount to `fstab`
- `blkid` to find UUID for `/dev/md0`
- `UUID=xxxxxxxx /mnt/data ext4 defaults 0 0`
- `umount /mnt/data` then `mount -a`


----
# Fail a device
- `mdadm -f /dev/md0 /dev/sdg1`
# Remove a failed device
- `mdadm -r /dev/md0 /dev/sdg1`
# Add more spare
- `mdadm -a /dev/md0 /dev/sdj1`
- remember to rebuild `mdadm.conf`
