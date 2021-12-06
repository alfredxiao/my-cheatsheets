# Three timestamps associated with a file
* `atime`: last time file content was read
* `mtime`: last time file content was modified
* `ctime`: last time file inode, or metadata (e.g. permission flags) is changed. Note ctime always changes when mtime changes

# Find files
* `find /tmp -mtime +14` finds files older than two weeks

# with touch
- `touch afile`: updates both atime and mtime
- `touch -a afile`: only change atime
- `touch -m afile`: only change mtime (and ctime)
