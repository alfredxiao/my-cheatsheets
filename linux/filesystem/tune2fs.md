# Basics
- adjust tunable filesystem parameters on `ext2/3/4` filesystems

# Display info of a filesystem
- `tune2fs -l /dev/sda2`

# Change Label (Volume Name)
- `tune2fs -L LABEL1 /dev/sda2`

# Find Label or Volume Name
- `tune2fs -l /dev/sda2 | grep volume`

# Find Check interval
- `tune2fs -l /dev/sda2 | grep interval`

# Find mount count setting
- `tune2fs -l /dev/sda2 | grep -i count`

# Set interval between checks
- `tune2fs -i 2m /dev/sda2` setting interval to 2 months
- set the value to `-1` to disable check by interval

# Set max mount counts
- `tune2fs -c 100 /dev/sda2`
- once mount count reaches specified number, file system will be checked by `e2fsck`
- set the value to `-1` or `0` to disable check by max mount count

# Set mount count
- `tune2fs -C 20 /dev/sda2`
- setting this value greater than value set by `-c` will cause file system check happen at next reboot

# Adjust percentage of reserved blocks
- `tune2fs -m 1 /dev/nvme1n1p1`  1 means 1%

# Convert `ext2` to `ext3`
- `tune2fs -j /dev/nvme1n1p1`  `-j` means enable journaling
- or `tune2fs -O has_journal /dev/nvme1n1p1`
- note unmount is not needed prior

# Convert `ext3` to `ext2`
- first `umount /dev/nvme1n1p1`
- `tune2fs -O ^has_journal /dev/nvme1n1p1`

# Adjust various mount options
- `tune2fs -o ^acl /dev/nvme1n1p1`
- other attributes like `user_xattr`
