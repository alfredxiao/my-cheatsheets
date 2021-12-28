# Remote X11
## X11 via Non-SSH
1. Enable X Server listening on 6000
  - e.g. in Debian 11, set `xserver-allow-tcp=true` in `/etc/lightdm/lightdm.conf` then `systemctl restart lightdm`
2. Login machine (`gmachine`) with GUI desktop
3. Run `xhost +appmachine` where (appmachine is the machine name or IP for the machine where the application process runs on)
4. Gain a login shell into `appmachine` (can be via SSH, or direct console shell or GUI shell)
5. Within this shell, run `export DISPLAY=gmachine:0.0`
6. Within this shell, run you GUI application like `xclock`
7. The GUI is displayed on `gmachine`

## X11 via SSH
# X11 via SSH
1. `vi /etc/ssh/sshd_config` on `appmachine`
2. set `X11Forwarding yes`
3. `systemctl restart sshd` on `appmachine`
4. `xauth generate :0 .` on `gmachine`
4. `ssh -Y user1@IP1`
5. `xeyes` (within the ssh session)
