# Overview
* it displays a list of library files required for the specific application.

# Usage
* `ldd /usr/bin/echo`
- `ldd -v /usr/bin/echo` for verbose mode

# `ldd afile` vs `readelf -d afile`
  * `ldd` shows all libraries including transitive ones
  * `readelf -d aprogram | grep NEEDED` shows only direct dependencies

# Unused direct dependencies
- `ldd -u afile`
