# DB File
- `mlocate.db`, which is located in the `/var/lib/mlocate/` directory

# Update
- `sudo updatedb`

# Locate
- `locate myfile`

# Locate by basename
- `locate -b myfile`
- matches the last part of a path
  - could be a directory name partial match
  - could be a filename partial match

# Locate by exact name
- `locate -b '\myfile'`
- `\` causes exact match rather than partial match
- `locate -b '\passwd' '\group'` multiple pattern search

# Ignore case
- `locate -i myfile`

# Display count only
- `locate -c myfile`


# Prevent some files from being found in `mlocate.db`
- update `/etc/updatedb.conf`, modify `PRUNEFS`, `PRUNENAMES`, or `PRUNEPATHS`
- consult `man updatedb.conf`
