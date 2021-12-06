# Config
* `/etc/yum.repos.d/` contains info about repositories
* `/etc/yum.conf` is global config
* `/var/cache/yum` caches latest repo info

# Install
- `yum install xyz` —> installs xyz and its dependencies
    - add `-y` to not prompt

# Update
- `yum update` —> upgrades all installed packages, **preserving** obsolete packages
- `yum update xyz` updates a single package

# Upgrade
- `yum upgrade` —> upgrades all installed packages, **removing** obsolete packages
    - which in effect is the same as `yum update —obsolete`

# Search package by name
- `yum search xyz` —> find available packages in local cache

# Info about a package
- `yum info xyz` —> info about specified package

# List packages
- `yum list installed` —> list all installed packages
- `yum list available` —> list all available packages

# List dependencies
* `yum deplist mypackage`

# Clean local yum database
- `yum clean all` —> cleans up all yum cache info and its local database (next time yum will download newer database)

# Remove packages
- `yum remove xyz` —> removes xyz, but **leaves dependencies behind**
- `yum autoremove` —> removes packages not used
- `yum whatprovides` —> what package provides specified file name
    - `yum whatprovides */httpd`
    - `yum whatprovides /usr/sbin/semanage`
    - similar to `rpm -q --whatprovides /etc/myfile`
- `yum provides mypackage` what facilities this packge provides
  - similar to `rpm -q --provides mypackage`
- `yum reinstall xyz` —> reinstalls a package

# Group Install
- `yum groupinstall “My Group”`
- `yum grouplist` to list available groups

# Shell
* `yum shell` enter the shell for yum

# Verify and reinstall
* `yum -V mypackage`
* `yum reinstall mypackage`

# List repos
* `yum repolist all`

# Other RPM Managers
- Zypper
  - used on SUSE
  - examples
    - `zypper repos`
    - `zypper install vim`
- DNF
  - Dandified yum
  - used on Fedora
  - future replacement for yum in RHEL
  - same command syntax as yum
