# Steps to create a Mount unit
0. You do not need to have fstab entry for this before you start, you also do not need to create the mount point before you start.
1. `vi /etc/systemd/system/mnt-data.mount`
```
[Unit]
Description=Data mount

[Mount]
What=/dev/disk/by-uuid/e3c44394-e390-46f0-85e2-2cec2058020b
Where=/mnt/data
Type=xfs
Options=defaults

[Install]
WantedBy=multi-user.target
```
**NOTE** `Where=` must match unit name
2. `systemctl daemon-reload`
3. `systemctl enable mnt-data.mount --now`

# Steps to create Automount unit
0. You should create a Mount unit prior, which needs not be started or enabled
1. `vim /etc/systemd/system/mnt-backup.automount`
```
[Unit]
Description=Automount backup directory

[Automount]
Where=/mnt/backup
TimeoutIdleSec=600

[Install]
WantedBy=multi-user.target
```
2. `systemctl daemon-reload`
3. `systemctl enable mnt-backup.automount --now`
