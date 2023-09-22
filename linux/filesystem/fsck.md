# Bascis
- target need to be unmounted first!
- `fsck` sis simply a front-end for the various filesystem checkers (fsck.fstype) available under Linux, which can be found via `ls /sbin/fs* -l`

# Check
- `fsck /dev/sda1`
- `fsck -r /dev/sda1` check and report

## Check all listed in `/etc/fstab`
- `fsck -A`

## by label
- `fsck -r LABEL=SRV`
## by partition
- `fsck -r /dev/sda1`
## by uuid
- `fsck -r UUID=xxx-xxxxxxx`
## Force a check (meaning don't do too quick lazy check)
- `fsck -f /dev/sda1` force check even though it seems clean

# repair without prompting
- `fsck -y /dev/sda1`

# Specify a filesystem type
- `fsck -t ext4 /dev/sda1` note older fsck assumes ext4 if not specified

# `e2fsck`
- for ext2/3/4 file system checking
