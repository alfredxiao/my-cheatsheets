# Default mode
- `diff file1 file2`
```
2d1
< Brian     # '<' means this line is from file1, '2d' means delete this second line from file1 to match file2
4a4,5
> Dwight        
> Drax      # '>' means these two lines are from file2, `4a4,5` means after line 4 in file1, add two lines to match file 2
6c7
< Frado     # line 6 from file1 has to change to match line 7 in file2
---
> Fredo     # line 7 in file2
```

# Context Format
- `diff -c file1 file2`
```
*** file1	2022-01-22 22:39:29.193161680 +1100  # '***' below refers to file 1
--- file2	2022-01-22 22:42:18.034727148 +1100  # '---' below refers to file2
***************
*** 1,6 ****    # file1 content below
  Alfred
- Brian         # '-' means this line should be deleted to match file2
  Chad
  Derek
  Eric
! Frado         # '!' means this line should be changed to match file2
--- 1,7 ----    # file2 content below
  Alfred
  Chad
  Derek
+ Dwight        # means file1 should add this line to match file2
+ Drax          # same as above
  Eric
! Fredo         # means file1 should update a line to match file on this line
```

# Unified Format
- `diff -u file1 file2`
```
--- file1.txt	2022-01-22 22:39:29.193161680 +1100
+++ file2.txt	2022-01-22 22:42:18.034727148 +1100
@@ -1,6 +1,7 @@
 Alfred
-Brian
 Chad
 Derek
+Dwight
+Drax
 Eric
-Frado
+Fredo
```

# Generate an Apply `ed` script
## Generate
- `diff -e file1 file2  > change.ed`
- `echo "w" >> change.ed`
## Apply
- `ed - file1 < change.ed`

# Create and Apply patch file
## Create
- `diff -u working.file tobe.file > change.patch`
# Apply
- `patch -u working.file -i change.patch`

# Comparing Directories
- `diff dir1/ dir2/`
- output will show files only exist in either, and file differences (between those having same file name under same path)
