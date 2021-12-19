# Prerequisite
- `write` on the directory (to modify the directory)
- and `execute` on a directory (to stat() the file's inode)
- and `execute` on all parent directories
- **NOTE** a user needs no permissions on a file nor be the file's owner to delete it!

# Lesson Learned
* The permissions of all parent directories must be considered before considering the permissions on a specific file.
  * When the command `more /data/abc.txt` is executed, the following permissions are checked: `x` permission on the `/` directory, `x` permission on the `/data` directory and `r` permission on the `abc.txt` file.
* The `w` permission allows a user to delete files from a directory, but only if the user also has `x` permission on the directory.
  * If a user has `wx` on directory, without `w` on a file inside this directory, the user CAN still delete the file (`rm` command prompts though)
* The `x` permission is required to "get into" a directory, but the `r` permission on the directory is not necessary unless you want to list the directory's contents.
