- does NOT do compression on its own

# Create archive
- `tar -cf mytar.tar myfolder`

# List content
- `tar -tf mytar.tar`

# Extract
- `tar -xf mytar.tar`

# Create gzipped archive
- `tar -czf mytar.tgz myfolder`
- `tar -czf mytar.tar.gz myfolder`

# Extract gzipped archive
- `tar -xzf mytar.tar.gz`

# Create bz2 archive
- `tar -cjf mytar.tar.bz2 myfolder`

# Extract bz2 archive
- `tar -xjf mytar.tar.bz2`

# Create xz archive
- `tar -cJf mytar.tar.xz myfolder`

# Extract xz archive
- `tar -xJf mytar.tar.xz`

# Create full backup with snapshot
- `tar -g full.snar -czvf full.tgz file*.in`
# Create incremental backup
- `tar -g full.snar -czvf inc.tgz file*.in`
