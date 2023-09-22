# List booleans
- `semanage boolean -l`

# Display User Mapping
- `# semanage login -l`
```
Login Name                SELinux User              MLS/MCS Range            

__default__               unconfined_u              s0-s0:c0.c1023           
root                      unconfined_u              s0-s0:c0.c1023           
system_u                  system_u                  s0-s0:c0.c1023  
```
## List SELinux users
- `semanage user -l`

# Change file type
1. Use the semanage command to define the file type httpd_sys_content_t for the directory hierarchy:
  - `# /usr/sbin/semanage fcontext -a -t httpd_sys_content_t "/var/webcontent(/.*)?"`
    - This command adds the following entry to the file `/etc/selinux/targeted/contexts/files/file_contexts.local`:
    - `/var/webcontent(/.*)?     system_u:object_r:httpd_sys_content_t:s0`
2. Use the `restorecon` command to apply the new file type to the **entire directory hierarchy**.
  - `# /sbin/restorecon -R -v /var/webcontent`

# Restoring to default file type
1.  Use the semanage command to delete the file type definition for the directory hierarchy from the file `/etc/selinux/targeted/contexts/files/file_contexts.local`:
  - `# /usr/sbin/semanage fcontext -d "/var/webcontent(/.*)?"`
2.  Use the restorecon command to apply the default file type to the entire directory hierarchy.
  - `# /sbin/restorecon -R -v /var/webcontent`

# Map a regular user to a SELinux user
- `# semanage login -a -s user_u user1`
- `user_u` is existing SELinux user, `user1` is regular linux user

# View Policy
- `semanage fcontext -l  | grep '/var/www/html'`

# Move/Copy files
- **NOTE** `mv` command retains the existing context
- `cp` will make a new context because it is making a new file

# Allow a port number
- `semanage port -a -t http_port_t -p tcp 8080`
- `semanage port -l | grep http_port` query allowed ports

# Grant permissive mode to a particular domain
- `semanage permissive -a postfix_smtpd_t`
- `semanage permissive -d postfix_smtpd_t` to move back to enforcing mode for this domain
