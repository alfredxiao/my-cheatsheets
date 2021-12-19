# Create a user with home directory
* `useradd -m user1`

# Create a user with a primary group
* `useradd -g group1 user1`

# Create a user with supplimentary groups
* `useradd -G group2,group3`

# Default values
* `useradd -D` to show default values
## Config files with default values
* `/etc/default/useradd`
* `/etc/login.defs`
  * To view its content: `grep -Ev '^#|^$' /etc/login.defs` which removes comments and empty lines
* `useradd -u new_uid username` to sepcify UID
* `useradd -G sales,research username` to specify supplementary group(s)
