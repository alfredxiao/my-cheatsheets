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

# Comparing ext2/3/4
- Ext2 has no journaling
- Ext3 is Ext2+Journaling, its size limitation is same as ext2
- Ext4 has journaling, and supports huge file and file system size.

# BTRFS
- B-tree Filesystem, from Oracle
## Features
- online FS expansion and reduction (without unmounting)
- readonly snapshots for faster backup
- more efficient handling of small files and directory indexes
- online defragmentation

# ReiserFS
- Offers several improvements over ext4

# ZFS
- from Sun, known as Zettabyte Filesystem, originally for Solaris
- each file has a checksum, hence easy to tell if a file is corrupted
- drive pooling, FS snapshots, filesystem striping

# XFS
- from Silicon Graphics
- Similar to ext4, with dynamic allocation and other features
- can expand on the fly, but cannot be reduced (where unmount is needed)
- handles large files very well but does suffer performance with many small files

# JFS
- From IBM, originally for AIX
- offers good performance across small and large files, small CPU footprint
- Many helpful features, but lacks extensive Linux production testing

# Swap
- Used to format a drive, but not technically a filesystem
- Used for virtual memory and does not have a viewable structure

# FAT/FAT32/exFAT
- from Microsoft
- No journaling, but supported by most OSes
- good choice for USB drives, or other media shared between systems
- exFAT supports larger file and partition sizes
