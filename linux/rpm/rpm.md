# RPM vs DEB
- RPM file content is similar to DEB

# Rpm database
- The RPM database holds information about all the RPM packages installed on your system. You can use this database to query what is installed, to help determine if you have the latest versions of software, and to verify that your system is properly set up, at least from a packaging point of view.
- `/var/lib/rpm`
- run `rpm —rebuilddb` to repair

# RPM vs Yum
- rpm does NOT handle dependencies like yum does

# List installed software
* `rpm -qa` query all

# Install a package
* `rpm -i mypackage.rpm`
* `rpm -ivh` for verbose and a hash mark - the '#####' progress bar
- add `--replacepkgs` option to kind of 'reconfigure'

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
- `rpm -qi openssh-server` display info on a package
  - `rpm -qpi openssh-server.rpm`
- `rpm -ql openssh-server` lists files in a package
  - `rpm -qpl openssh-server.rpm`

# Which packages owns a file
- `rpm -qf /usr/bin/yacc` what package contains/provides this yacc file
- or `rpm -q --whatprovides /usr/bin/yacc`

# What facilities a package provides
- `rpm -q --provides mypackage.rpm`

# What config files a package provides
- `rpm -qc mypackage`

# What docs avaialble
- `rpm -qd mypackage`

# What script will execute when the package is installed
- `rpm -qp --scripts mypackage.rpm`

# Verify a package (its installation on this host)
* `rpm -V mypackage.rpm`
* `rpm -V openssh-server` would report on a missing config file (for example `/etc/ssh/sshd_config`)

# Convert rpm to cpio
- `rpm2cpio xyz.rpm > xyz.cpio` converts to a cpio file
  - then run `cpio -idv < xyz.cpio`
- `rpm2cpio xyz.rpm | cpio -idmv` extracts its content to current folder

# Check Signature
- `rpm --checksig`
