# `split`
* splits a large file to smaller files, by line number, size etc.

# by line number
* `split -l 3 bigfile smaller`
  * splits every 3 lines into smaller files, named with prefix `smaller`
# by size
* `split -b 1000 bigfile smaller`
  * every 1000 bytes into smaller files

# by number of output files
* `split -n3 bigfile` splits into 3 files

# naming schema
* `split -d -n3 bigfile` use number naming scheme for output files

# combine them again
* `cat smaller* > combined`
