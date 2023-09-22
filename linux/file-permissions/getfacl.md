
# ls output
- an extra `+` in addition to `rwxrwxrwx` bits
```
ls -l
total 4
-rw-rw-r--+ 1 alfred alfred   0 Jan 30 10:34 file1
```

- `getfacl jlpayroll.txt`
# output
```
getfacl file1
# file: file1
# owner: alfred
# group: alfred
user::rw-
user:user1:r--
group::rw-
mask::rw-
other::r--
```
- **NOTE** `mask` value above is the maximum permissions allowed in all ACLs for this file; and this value is the combined permissions from the owner user and owner group regular permissions.

# Disputes with regular permissions
- Generally speaking the `getfacl` is final result
- running `chmod` can change the `mask` in ACLs thus effectively changes permissions of the ACLs, but run `getfacl` again you will get new final result
- Simply put, don't just trust regular permission bit when there are ACLs associated
