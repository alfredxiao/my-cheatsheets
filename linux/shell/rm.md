# delete empty folder
- `rm -d abc`
- same as `rmdir dir1`

# force (do not ask)
- `rm -f file1`

# prompt before deleting
- `rm -i file1`

# Recursive delete files along with the directory
- `rm -r dir1`
  - may give warning and ask for confirmation depending on individual file (permission bits)

# force delete a directory and contents recursively and without prompt
- `rm -rf dir1`
