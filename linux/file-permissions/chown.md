# Change owner
- `chown user1 file1`

# Change group
- `chown :group1 file1`
- or `chown .group1 file1`

# Change both and recursively
- `chown user1:group1 -R dir1`


# root
- change owner to someone else requires ‘root’ privilege
- otherwise you can create a malicious file and give it to others
