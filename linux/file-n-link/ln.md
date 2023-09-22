# Create a hard link
- `ln original.file link.file`
- **note** hard link works only if they are on the same file system, and does not work for directory

# display inode number
- `ls -i`

# display hard link count
- `ls -iog`
```
787322 -rw-r--r-- 2 28 Dec  4 23:30 hlink
787322 -rw-r--r-- 2 28 Dec  4 23:30 original
784952 lrwxrwxrwx 1  8 Dec  4 23:29 slink -> original
```

# create soft link
- `ln -s src tgt`

# overwrite existing
- `ln -sf src tgt`
