# what process has this file open
`fuser /path/to/myfile`

# what processes using files in a directory
- `fuser -v /path/dir1`  v for verbose

# kill the found process
- `fuser -k -KILL some-huge-file.txt`
  - `-KILL` is to use SIGKILL
  - just `-k` will use SIGTERM

# what process has files open in this file system
`fuser /mnt/mountpoint` or `fuser -m /mnt/mountpoint`
- output
  - `PIDc` where `c` means current directory

# Give verbose info
- `fuser -v /mnt/mountpoint`

# what process using a protocol and port
- `fuser -vn tcp 22`
- Note that you must specify both the protocol and the port with the fuser command.
