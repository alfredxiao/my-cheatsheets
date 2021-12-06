# `fsck`
- target need to be unmounted first!
## by label
- `fsck -r LABEL=SRV`
## by partition
- `fsck -r /dev/sda1`
## Force a check (meaning don't do too quick lazy check)
- `fsck -f /dev/sda1` force check even though it seems clean
## repair without prompting
- `fsck -p /dev/sda1`


# `e2fsck`
- for ext2/3/4 file system checking
