# RPM vs DEB
- RPM file content is similar to DEB

# Rpm database
- The RPM database holds information about all the RPM packages installed on your system. You can use this database to query what is installed, to help determine if you have the latest versions of software, and to verify that your system is properly set up, at least from a packaging point of view.
- `/var/lib/rpm`
- run `rpm —rebuilddb` to repair

# RPM vs Yum
- rpm does NOT handle dependencies like yum does

# List installed software
* `rpm -qa`

# Install a package
* `rpm -i mypackage.rpm`
* `rpm -ivh` for verbose and a hash mark - the '#####' progress bar

# Remove a package
- `rpm -e mypackage`

# Update a package
* `rpm -U mypackage.rpm` will install no matter if an earlier version is installed

# Freshen a package
- `rpm -F package.rpm` to Freshen the package only if an earlier version is already installed

# Query dependencies
* `rpm -qR mypackage.rpm`
* `rpm -qR my-installed-package`

# Info about a packages
- `rpm -qpi` display info on a package
- `rpm -qpl`lists files in a package

# What package contains this file
- `rpm -qf /usr/bin/yacc` what package contains/provides this yacc file
- or `rpm -q --whatprovides /usr/bin/yacc`

# What facilities a package provides
- `rpm -q --provides mypackage.rpm`

# What config files a package provides
- `rpm -qc mypackage`

# Verify a package (its installation on this host)
* `rpm -V mypackage.rpm`
* `rpm -V zsh` would report on a missing config file (for example `/etc/zprofile`)

# Convert rpm to cpio
- `rpm2cpio xyz.rpm > xyz.cpio` converts to a cpio file
  - then run `cpio -idv < xyz.cpio`
- `rpm2cpio xyz.rpm | cpio -idmv` extracts its content to current folder
