# Repo list
* `/etc/apt/sources.list`

# Download a package
* `apt-get download dos2unix` dos2unix is a package name

# Install
* `apt-get install xyz`

# Remove
- `apt-get remove xyz`
- **note** this does **NOT** remove those dependencies or configs
- `apt-get autoremove`
  - used to remove packages that were automatically installed to satisfy dependencies for other packages and are now no longer needed
## Purge
- `apt-get purge xyz` remove package and configs

# Update all packages
- `apt-get upgrade` update all installed packages
- `apt-get dist-upgrade` update all packages as well as kernel

# Update local cache
- `apt-get update`
