- This permission does not affect individual files. However, at the directory level, it restricts file deletion.
- Only the owner (and root) of a file can remove the file within that directory.
- Only the owner can modify the file content in that directory
- only the owner can rename, move the files in that directory

# Example
- `ls -ld /tmp` which has this bit on
- meaning files created by user1 cannot be removed by user2

# Display
- Notice that `t` implies both execute permission and sticky bit. A `T` would be sticky bit only.

# To set both sticky bit and execute permissions:
* `chmod o+t adir`
* `chmod 1775 adir` or plus `1000` on base value
