# Create a user with home directory
* `useradd -m user1`

# Create a user with home directory, specified shell, full user name
- `useradd -m -c "Alfred Xiao" -s /bin/tcsh user1`
- `-c` means comment

# Create a system account (or service account)
- `useradd -r service1`
- further lock shell access by `usermod -s /sbin/nologin service1`

# Create a user with a primary group
* `useradd -g group1 user1`

# Create a user with supplimentary groups
* `useradd -G group2,group3 user1`

# Default values
* `useradd -D` to show default values, some came from `/etc/default/useradd`, some are defaults if not specified in this file
* `useradd -D -s /bin/bash` would modify file `/etc/default/useradd`
## Config files with default values
* `/etc/default/useradd`
  - `INACTIVE` —> number of days after a password has expired before the account is disabled
  - `SKEL` : The skeleton directory.
  - `SHELL`: User account default shell program.
* `/etc/login.defs`
  * To view its content: `grep -Ev '^#|^$' /etc/login.defs` which removes comments and empty lines
  * this file overrides `/etc/default/useradd`
  * Some Settings
    - `PASS_MAX_DAYS`: Number of days until a password change is required
    - `PASS_MIN_DAYS`: Number of days after a password is changed until the password may be changed again
    - `PASS_MIN_LENGTH`: Minimum number of characters required in password.
    - `PASS_WARN_AGE`: Number of days a warning is issued to the user prior to a password’s expiration.
    - `CREATE_HOME`: Default is no. If set to yes, a user account home directory is created.
    - `ENCRYPT_METHOD`: The method used to hash account passwords.
* `useradd -u new_uid username` to sepcify UID
* `useradd -G sales,research username` to specify supplementary group(s)
