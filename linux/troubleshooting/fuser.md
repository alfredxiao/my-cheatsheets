# what process has this file open
`fuser /path/to/myfile`

# kill the found process (SIGKILL)
`fuser -k -KILL some-huge-file.txt`

# what process has files open in this file system
`fuser -m /mnt/mountpoint`

# what process using a protocol and port
- `fuser -vn tcp 22`
- Note that you must specify both the protocol and the port with the fuser command.
