# Basics
- used with executable files.
- it tells kernel to run the program with the permission of the owner and not the user account actually running the program

# Display
- `rws` meaning `rwx` + `s`
- Same place, an uppercase `S` means suid is set but `x` bit is not set

# To set this flag on a file
* `chmod u+s an-exec-file`
* or `chmod 4775 an-exec-file` - plus `4000` from existing permission number

# Example
- `ls -l /usr/bin/passwd`
