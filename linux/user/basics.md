
System Accounts
* root's UID is 0
* regular UID is greater than 500 or 1000
* There are accounts not designed for users to log in.
  * UID = 1~499
  * called `system accounts`
  * are designed to provide accounts for services running on the system
  * One system account that is an exception to this rule is the user nfsnobody, which has a UID of 65534
* No home directory
* shell is `/usr/sbin/nologin`
* password in shadow file is `*`

Group Accounts
* defined in `/etc/group`
* example: `mail:x:12:mail,postfix`
* fields  
  * group name
  * password (holder)
  * group id
  * member list
* primary user memberships defined `/etc/passwd`
