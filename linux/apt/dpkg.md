# Install
- `dpkg -i xyz.deb`

# Remove a package
- `dpkg -r xyz`

# Dependencies
- dpkg does **NOT** handle dependencies for you, they need to be installed already prior

# View info
- `dpkg —info xyz.deb` or `-I`
- `dpkg -—status xyz` or `-s`

# View contents
* `dpkg -c mypackage.deb` or `--contents`
* `dpkg -L mypackage`

# List installed
- `dpkg -l  xyz`
- `dpkg -l` list all installed packages

# Remove Package
- `dpkg -P xyz` or `--purge`
  * removes package as well as configurations
- `dpkg -r ` or `--remove`
  * removes package but leaves configurations

# Search by fuzzy match
- `dpkg -S xyz`  —> Search through the database for a file specified and lists out any mentions of the specified file or string

# Re-configure
- `dpkg-reconfigure` —> modify a package by running the app’s config tool again
