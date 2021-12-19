`/etc/passwd`
* hols 'some' account information, each account has an entry here
* it used to hold user password, now it is in `/etc/shadow`
* example: `sysadmin:x:1001:1001:System Administrator,,,,:/home/sysadmin:/bin/bash`
* fields
  * Name
  * Password (x means it is in /etc/shadow)
  * user id: 1001
  * group id: 1001
  * Comment/Description: full name or anything
  * Home Dir: /home/sysadmin
  * Shell: /bin/bash
