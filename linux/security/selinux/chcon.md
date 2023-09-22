- The 'chcon' command may be used to change SELinux security context of a file or files/directories in a similar way to how 'chown' or 'chmod' may be used to change the ownership or standard file permissions of a file.

# Change Type
- `chcon -v --type=httpd_sys_content_t /html`
## Recursive
- `chcon -Rv --type=httpd_sys_content_t /html`
  - includes sub directories

- Modifying security contexts in this manner will persist between system reboots but only until the modified portion of the filesystem is relabeled.
- a permanent way is `semanage fcontext -a -t httpd_sys_content_t "/html(/.*)?" `
