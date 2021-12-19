# Lock password
- `passwd -l user1`
- `usermod -L user1`
- note password locked account can still login via `su` or `ssh` key based authentication
  - `su - user1`

# Unlock password
- `passwd -u user1`
- `usermod -U user1`

# Lock account (Expiring an account)
- `chage -E 0 user1` (0 means 1970-01-01)
- or `usermod -e 0 user1`
- `usermod -s /sbin/nologin user1` to disallow a login shell

# Unlock account
- `chage -E -1 user1`

# Verify if an account is locked (have expriry date)
- `chage -l user1`

# Check login shell
- `grep ^user1 /etc/passwd`
- `getent passwd user1`
