# Create File system
- `mkfs -t ext4`

# Create with a label
- `mkfs.ext4 -L MYLABEL /dev/sda1`

# mke2fs
- is mkfs for ext*
- configured by `/etc/mke2fs.conf`
- `mke2fs -t ext4 -L MYLABEL /dev/sdb1`
