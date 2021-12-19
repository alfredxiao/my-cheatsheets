# What file systems are available
* `cat /proc/filesystems`

# display information about mounted file systems
* `mount | column -t`

# Types
# Two types
  * Non-journaling: e.g. `ext2`
  * Journaling: e.g. `ext3` `ext4`
# Journaling
  * Since under the neath, a single file operation (like delete a file) is implemented as multiple steps as multiple storage objects are involved (e.g. inode, dentry), the file system can end up inconsistent should crash occurs before all steps finishes.
  * Journaling is a mechanism built to allow recording of those steps and being able to replay them when the system comes back online, ensuring the 'single file operation' is atomic.
# Non-journaling
  * requires the use of tools for checking and repairing

* Each filesystem has its own method for indexing files and directories and tracking file access.

# Steps for adding new disk
- create partition
- create file system
- mount it manually or added to fstab to make it persistent

# Kernel modules
- `ls /lib/modules/$(uname -r)/kernel/fs/*/*ko` lists modules related to filesystem
