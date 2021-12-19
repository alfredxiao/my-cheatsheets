# Basics
* is metadata about files
  * Metadata includes information about the file like the permissions, ownership, and timestamps.
  * Metadata does not include the file name or the contents of the file
* Every file on a partition has a unique identification number called an inode number,
  * To view inode number: `ls -i afile`
  * For each file, there is also an entry that is stored in a directory's data area (data block) that includes an association between an inode number and a file name.
* hard link  
  * `ln original linkfile`
  * `ls -l` displays file type as `-`
* soft link
  * `ln -s original linkfile`
  * `ls -l` displays file type as `l`
* To search for files having same inode number
  * `find / -inum 278772`
* links vs inodes
  ```
  sysadmin@localhost:~$ ls -i abc*                                                
  143131717 abc  143131717 abc.hard  143131718 abc.soft
  ```
  * hard links share same inode, hence share the same metadata, like permissions, ownership
  * soft links has its own inode, but effectively using the same permissions/ownership
* Pros/Cons
  * hard links cannot cross partitions (each partition has its own set of inodes)
  * hard links cannot point to directories
