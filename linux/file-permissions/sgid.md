# SGID on a file
* similar to `setuid`, it allows a user to run an executable binary file in a manner that provides them additional (temporary) group access.
- Example:
  - `ls -l /usr/bin/wall` which has sgid set, and this file owned by `tty` group
  - when any user runs `wall` is in the `tty` group during current session

# SGID on a directory
* the group ownership of any further files and directories created within will be set to the same group
* further, sub directories created will also have the SGID bit set on.

# to setgid
* `chmod g+s file/dir`
* `chmod 2775 file/dir` basically plus 2000
