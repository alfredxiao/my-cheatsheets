- turn on swap
- `swapon -a` —> turn on all swap partitions (for things listed in fstab)
- `swapon -U <THE_UUID>` —> turn on swap on partition identified by UUID
- `swapon -L SWAP` —> turn on swap on a partition identified by a label

# Summary
- `swapon -s` or `swapon --summary`
- same as `cat /proc/swaps`
