# all files (including hidden files)
- `ls -a`

# only list directories (if something is a directory)
- `ls -d something`

# Show time
- `ls -lu` —> show atime
- `ls -lc` —> show ctime (otherwise mtime)
- `ls --full-time `

# differentiate between file and directories
* `ls -F` appends a `/` to directory names

# Friendly size
* `ls -hl` gives better size indication for files

# Tree
* `ls -R` is the poor man's `tree` command

# sort by size descending
- `ls -lS`
- ascending: `ls -lSr`, `r` for reverse

# sort by timestamp descending
- `ls -lt`
- ascending: `ls -ltr`

# File types
* first character in the output donates file types
- `d` for directory,
- `-` for regular file (hard link is regular file),
- `l` for soft link
- `s` for socket
- `p` for pipe
- `b` for block file, `c` for character file

# Link count
* hard link count, as `1` and `2` above

# file size
* For directories, this value does not describe the total size of the directory, but rather how many bytes are reserved to keep track of the filenames in the directory. In other words, ignore this field for directories.

# timestamp
* Indicates the time that the file's contents were last modified.
* For directories, this timestamp indicates the last time a file was added or deleted from the directory.
  * file updated in a directory does not update this timestamp
