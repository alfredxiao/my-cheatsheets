# Overview
* it displays a list of library files required for the specific application.

# Usage
* `ldd /usr/bin/echo`

# `ldd` vs `readelf -d`
* `ldd` shows all libraries including transitive ones
* `readelf -d aprogram | grep NEEDED` shows only direct dependencies
