* `useradd -m user1` create a user with home directory
* `useradd -g group1 user1` create a user with a primary group
* `useradd -G group2,group3` create a user with supplimentary groups
* `su - user2`
* `usermod -L user1` lock a user's password
  * account is not locked completely, e.g. `su - user1` could work
  * `usermod -e 1 user1` to lock an account
  * `usermod -s /sbin/nologin user1` to disallow login shell
* `usermod -U -e "" user1` unlock a user's password and account
