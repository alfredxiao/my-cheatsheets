# Basics
- adjust tunable filesystem parameters on `ext2/3/4` filesystems

# Display info of a filesystem
- `tune2fs -l /dev/sda2`

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
- set the value to `-1` to disable check by max mount count

# Adjust percentage of reserved blocks
- `-m`

# Adjust various mount options
- `-o`
