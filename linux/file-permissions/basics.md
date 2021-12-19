# Directory permissions are trickier
## Read
- grants the ability to **read the names of files** in the directory,
- but **not** to find out any further information about them such as contents, file type, size, ownership, permissions
```
alfred@debian11:test$ ls -l
total 4
dr-------- 2 alfred alfred 4096 Dec 19 13:57 dir1

alfred@debian11:test$ ls dir1
ls: cannot access 'dir1/file1': Permission denied
file1

alfred@debian11:test$ ls -l dir1
ls: cannot access 'dir1/file1': Permission denied
total 0
-????????? ? ? ? ?            ? file1                                
```
* In traditional Unix filesystems, a directory was simply a list of **(name, inode number) pairs**.
* An inode number is an integer used as index into the filesystem's **inode table** where the rest of the file metadata is stored.
* The `r` permission on a directory allows to list the names in it, but **not** to access the information stored in the **inode table**, that is, getting file type, file length, file permissions etc, or opening the file. For that you need the `x` permission on the directory.
```
alfred@debian11:test$ chmod 500 dir1

alfred@debian11:test$ ls -l dir1
total 0
-rw-r--r-- 1 alfred alfred 0 Dec 19 13:57 file1
```

## Execute
- `x` permission enables access to inode metadata of the files or subdirectories, but `x` alone without `r` does not allow reading names of the files or subdirectories inside
- the user can "get into" the directory and use the directory in a pathname to access files and, potentially, subdirectories under this directory.
```
alfred@debian11:test$ chmod 100 dir1

alfred@debian11:test$ ls -l
total 4
d--x------ 2 alfred alfred 4096 Dec 19 13:57 dir1

alfred@debian11:test$ ls dir1
ls: cannot open directory 'dir1': Permission denied

alfred@debian11:test$ ls -l dir1
ls: cannot open directory 'dir1': Permission denied

alfred@debian11:test$ ls -l dir1/file1
-rw-r--r-- 1 alfred alfred 0 Dec 19 13:57 dir1/file1
```

## Write
- grants the ability to modify entries in the directory, which includes **creating** files, **deleting** files, and **renaming** files. Note that this requires that `x` is also set; without it, the write permission is meaningless for directories.
```
alfred@debian11:test$ chmod 100 dir1

alfred@debian11:test$ touch dir1/file1
touch: cannot touch 'dir1/file1': Permission denied

alfred@debian11:test$ chmod 200 dir1

alfred@debian11:test$ touch dir1/file1
touch: cannot touch 'dir1/file1': Permission denied

alfred@debian11:test$ chmod 300 dir1

alfred@debian11:test$ touch dir1/file1

alfred@debian11:test$  
```
