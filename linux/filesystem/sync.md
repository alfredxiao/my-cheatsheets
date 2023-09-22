# Basics
- to flush the buffer cache out to the storage media (force all unwritten data out to the device drivers and, subsequently, to the storage device)

# Sync everything?
- `sudo sync`

# Sync a mounted partition
- `sudo sync /dev/sda1`

# Sync a folder
- `sudo sync /path/to/dir1`

# Sync on the data of a file
- `sudo sync -d /path/to/file1`

# Sync a file system containing the file/dir
- `sudo sync -f /path/to/file1`
