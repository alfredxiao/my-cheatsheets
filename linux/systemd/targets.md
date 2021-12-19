targets are not exclusive
  * unless you say `systemctl isolate abc.target`

* `systemctl get-default` to show default target
* `systemctl set-default abc.target` to set default target
* `systemctl list-units -t target`
* `systemctl start abc.target` to start a target
  * similar to `init x` in runlevel world
* `systemctl isolate abc.target` to start a target with
 minimal dependencies
  * similar to `init x`

# Special targets
- `rescue.target` similar to single user mode in SysV mode, allows only root to login, no network, mounts all local filesystems
- `emergency.target` even more restricted than rescue target. it mounts only the root filesystem, and mounts it as read-only.
