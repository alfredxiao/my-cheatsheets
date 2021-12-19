* operate on a single file
* removes original file

- `gzip afile`
- `gunzip afile.gz`

# Info
- `gzip -l afile.gz`

# View
- `zcat afile.gz`

# Send output to stdout and not remove original file
- `gzip -c file1`

# multiple files
- `gzip -c file1 file2 > file3.gz`

# compression rate
- `gzip -9 file1` (from `-1` to `-9`, lowest to highest compression rate)
