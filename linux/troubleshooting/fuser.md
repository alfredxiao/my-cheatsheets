what process has this file open
`fuser /path/to/myfile`

kill the found process (SIGKILL)
`fuser -k -KILL some-huge-file.txt`

what process has files open in this file system
`fuser -m /mnt/mountpoint`
