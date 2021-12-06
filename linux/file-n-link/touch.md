# Changes timestamp
- `touch afile`: updates both atime and mtime
- `touch -a afile`: only change atime
- `touch -m afile`: only change mtime (and ctime)

# multiple files
* `touch a b c`

# Set to a provided timestamp
- `touch -a -m -t 201512180130.09 fileName.ext` YYYYMMDDhhmm.ss
  - **note** ctime will still be current time rather than the provided timestamp

# touch with timestamp of hours ago
- `touch -d '10 hours ago' abc`
