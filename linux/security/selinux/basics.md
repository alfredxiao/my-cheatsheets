# Basics
- `SELinux` means `Security Enhanced Linux`
- is a kernel module
- common on Redhat based systems
- `/var/log/audit/audit.log` is its logs
- `/var/log/messages` for application access denied logs

# Config
- `/etc/selinux/config`
- `SELINUX=`
  - `enforcing`
  - `permissive`
  - `disabled`
- setting from `disabled` to others requires a reboot to take effect
  - SELinux will have to re-label the entire filesystem (effectively running `estorecon /`) because contexts are not set at all when SELinux is disabled.


# Contexts
- Under SELinux, all *file systems*, *files*, *directories*, *devices*, and *processes* have an associated **security context**.
- For files, SELinux stores a context label in the *extended attributes* of the file system. The context contains additional information about a system object:
  - the SELinux user
  - their role
  - their type
  - and the security level
- SELinux uses this context information to control access by processes, Linux users, and files.
- Sometimes also termed **SELinux Labels**

## Context Format
- You can specify the `-Z` option to certain commands (`ls`, `ps`, and `id`) to display the SELinux context with the following syntax:
  - `SELinux user:Role:Type:Level`
- Where the fields are as follows:
### SELinux user
- An SELinux user account compliments a regular Linux user account.
- SELinux maps every Linux user to an SELinux user identity that is used in the SELinux context for the processes in a user session.
### Role
- In the Role-Based Access Control (RBAC) security model, a role acts as an intermediary abstraction layer between SELinux process domains or file types and an SELinux user.
- Processes run in specific SELinux domains, and file system objects are assigned SELinux file types.
- SELinux users are authorized to perform specified roles, and roles are authorized for specified SELinux domains and file types.
- A user's role determines which process domains and file types he or she can access, and hence, which processes and files, he or she can access.
### Type
- A type defines an SELinux file type or an SELinux process domain.
- Processes are separated from each other by running in their own domains.
- This separation prevents processes from accessing files that other processes use, and prevents processes from accessing other processes.
- The SELinux policy rules define the access that process domains have to file types and to other process domains.
### Level
- A level is an attribute of Multilevel Security (**MLS**) and Multicategory Security (**MCS**). An MLS range is a pair of sensitivity levels, written as low_level-high_level. The range can be abbreviated as low_level if the levels are identical. For example, `s0` is the same as `s0-s0`. Each level has an optional set of security categories to which it applies. If the set is contiguous, it can be abbreviated. For example, `s0:c0.c3` is the same as `s0:c0,c1,c2,c3`

## Example Context
- To view the SELinux context of a directory, use the `ls` command with a `-Z` flag.
  - e.g. `system_u:object_r:httpd_sys_content_t:s0`
- Commands
  - `ls -Z`
  - `ps -Z`
  - `id -Z`

# Relabeling Complete Filesystem
- Sometimes it is necessary to relabel the complete filesystem although this should only be necessary when enabling SELinux after it has been disabled or when changing the SELinux policy from the default targeted policy to strict. To automatically relabel the complete filesystem upon reboot, do:
```
# touch /.autorelabel
# reboot
```
