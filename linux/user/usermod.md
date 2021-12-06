
* `su - user2`
* `usermod -L user1` lock a user's password
  * account is not locked completely, e.g. `su - user1` could work
  * `usermod -e 1 user1` to lock an account
  * `usermod -s /sbin/nologin user1` to disallow login shell
* `usermod -U -e "" user1` unlock a user's password and account
