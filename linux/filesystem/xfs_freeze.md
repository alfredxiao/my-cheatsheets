To freeze a XFS file system the basic syntax is:

`xfs_freeze -f /mount/point`

The `-f` flag requests that the specified XFS filesystem should be set to a state of frozen, immediately stopping any modifications from being made. When this option is selected, all ongoing transactions in the file system are allowed to complete. Any new write system calls are halted.

To unfreeze a XFS file system the basic syntax is:

`xfs_freeze -u /mount/point`

You can also use the "xfs_freeze" utility to freeze or unfreeze an `ext3`, `ext4` and `btrfs`, file system.
