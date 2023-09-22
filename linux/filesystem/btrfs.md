# Show Info
- `btrfs filesystem show /mnt/btrfs`

# Add an extra partition to running btrfs file system
- meaning no need to unmount prior
- `btrfs device add /dev/nvme2n1p6 /mnt/btrfs`

# Balance (if filesystem consist of multiple devices)
- `btrfs balance start --full-balance /mnt/btrfs`

# Convert ext filesystem to btrfs
- `umount /mnt/ext3`
- `btrfs-convert /dev/nvme1n1p3`
