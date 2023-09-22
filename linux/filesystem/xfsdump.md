# Basics
- install via package name `xfsdump`

# Full backup
- `xfsdump -l 0 -f /path/to/outfile.dump /dev/nvme1n1p5`
- then enter two labels as prompted
- level 1~9 is for incremental backup

# Restore
- `xfsrestore -f /path/to/file.dump /mnt/xfs`
- `xfsrestore -f /path/to/file.dump /mnt/ext3` actually it allows you to restore to non-xfs file system

# Selective (Interactive) Restore
- `xfsrestore -i -f /path/to/file.dump /mnt/ext4`
- you need to `add` files you want and then `extract`
