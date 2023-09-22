# Add user/group permission
- `setfacl -m u:superman:w -m u:robin:w jlpayroll.txt`
- or `setfacl -m u:bob:6 file1`  6 means `rw`
- or `setfacl -m bob:rw file1`
- or `setfacl -m g:group1:rw file1`
- note that only file owner can do this, or root

# Set maximum permissions (in ACLs)
- `setfacl -m m::r file1` set maximum permissions to `r` only, after this
```
$ getfacl file1
# file: file1
# owner: alfred
# group: alfred
user::rw-
user:user1:rw-			#effective:r--
user:user2:rw-			#effective:r--
group::rw-			    #effective:r--
group:users:rw-			#effective:r--
mask::r--
other::r--
```

# Remove a permission
- `setfacl -x u:user2 file1`
- `setfacl -x g:users file1`

# Clear all permissions
- `-b`

# Set Default Permissions on a Directory
- `setfacl -m d:g:users:rw dir1`
- such that files created within `dir1` have same permission inherited

# Mount option
- To set ACLs (Access Control Lists) on files, use the setfacl and getfacl tools. In order to use these commands, and ACLs in general, the filesystem must be mounted with the `acl` mount option.
