`/etc/shadow`
* contains password and other information about accounts
* only root can see it
* example line: `sysadmin:$6$c75ekQWF$.GpiZpFnIXLzkALjDpZXmjxZcIll14OvL2mFSIfnc1aU2cQ/221QL5AX5RjKXpXPJRQ0uVN35TY3/..c7v0.n0:16874:5:30:7:60:15050::`
* fields
  * username
  * password ( it is a one-way encryption)
    - A `!!` or `!` indicates a password has not been set for the account.
    - A `!` or an `*` indicates the account **cannot** use a password to log in.
    - A `!` in front of a password indicates the account has been **locked**.
  * last change: last time the password was changed
  * min: e.g. 5. Meaning min days between password change. zero means can change password anytime
  * max: e.g. 30. Meaning password will expire after 30 days.
  * warn: e.g. 7. Gives warning before password expires
  * inactive: e.g. 60. Once password expires, this give users the opportunity to change password at login time. Once this period is over, the account is inactivated, only root can set the user's password
  * expire: e.g. 15050. Number of days since 1/1/1970. When that day comes, lock the account, only root can reset password to unlock it.
  * reserved

- If your Linux system is still storing passwords in the `/etc/passwd` file, you need to move them to the `/etc/shadow` file. The `pwconv` utility can perform this operation for you.
