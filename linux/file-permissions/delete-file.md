- To delete a file requires both write (to modify the directory itself) and execute (to stat() the file's inode) on a directory.  

- Note a user needs no permissions on a file nor be the file's owner to delete it!
