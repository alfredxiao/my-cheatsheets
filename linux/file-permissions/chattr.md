changes file attributes

# Attributes
- `(i)` immutable
- `(a)` append only
- `(A)` no atime updates
- `(u)` undeletable
- `(c)` compressed
- `(C)` no copy on write
- `(d)` no dump
- `(j)` data journalling
- `(s)` secure deletion
- `(S)` synchronous updates
- `(D)` synchronous directory updates
- `(e)` extent format
- `(P)` project hierarchy
- `(t)` no tail-merging
- `(T)` top of directory hierarchy
- **Note** not all attributes are honoured by all filesystems

# Operators
- `+` : Adds the attribute to the existing attribute of the files.
- `–` : Removes the attribute to the existing attribute of the files.
- `=` : Keep the existing attributes that the files have.

# Setting flag
- `sudo chattr +i file1`
- `sudo` is required, as even owner cannot change attributes

# Prerequisite
- file system mounted with `ext_attr`, which enables the use of extended attributes. This eature is supported by ext2, ext3, and ext4.
