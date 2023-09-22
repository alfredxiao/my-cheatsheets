# Permissions
Permission | File    | Directory
Read       | Read    | List files
Write      | Write   | create, move (rename), modify attributes of, and delete files within the directory
Execute    | Execute | cd working directory to this as long as this permission is set on all its parent directories as well

# Syntax
* `chmod [u|g|o|a]+[r|w|x] afile` to add a permission
* `chmod [u|g|o|a]-[r|w|x] afile` to remove a permission
* `chmod [u|g|o|a]=[rwx] afile` to set permission
* `chmod nnn afile` where n=0~7
* `chmod ug+x afile`, `chmod go=rx`
- `chmod u+x,g+rw afile`

# Add a permission
- `chmod u+r file1`
- use `u+w` or `u+x`
- use `g+w` for others
- use `o+r` for others
- use `a+r` for all

# Remove a permission
- `chmod u-r file1`

# Set exactly a permission
- `chmod o=r file1`

# Combine permissions
- `chmod ug=rwx file1`

# Using octal codes
- `chmod 664 file1`

# Soft Links
- a soft link always appears to have permission bits like `lrwxrwxrwx`
- `chmod` on a soft link will actually have effect on the original file

# Hard Links
- a hard link always share same permission bits with original file
- `chmod` or `chown` on the hard link or the original file yield the same outcome
