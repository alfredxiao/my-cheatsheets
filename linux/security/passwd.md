`/etc/passwd`
* holds 'some' account information, each account has an entry here
* it used to hold user password, now it is in `/etc/shadow`
* id `0` is root
* system accounts has id below `1000`
* normal user account ids are from `1000`
* example: `sysadmin:x:1001:1001:System Administrator,,,,:/home/sysadmin:/bin/bash`
* fields
  * Name
  * Password (`x` means it is in /etc/shadow)
  * user id: 1001
  * group id: 1001
  * Comment/Description: full name or anything
  * Home Dir: /home/sysadmin
  * Shell: /bin/bash

# Change password for a user
- `sudo passwd user1`

# Status
- `sudo passwd -S user1`

# Force a user to change password
- `sudo passwd -e user1` (e for expire)

# Lock password
- `sudo passwd -l user1`

# Unlock password
- `sudo passwd -u user1`

# Set max days
- `sudo passwd -x user1`

# Set warning days (warning prior to expiry)
- `sudo passwd -w user1`

# Set inactive
- `sudo passwd -i user1`

# Script it
- `echo Temp@$$ | passwd --stdin test1`
