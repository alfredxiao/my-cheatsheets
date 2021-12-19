# approach 1
1. `mkdir /etc/systemd/system/xxx.service.d`
2. `vi /etc/systemd/system/xxx.service.d/override.conf`
3. `systemctl daemon-reload`
4. `systemctl restart xxx.service`

# approach 2
1. `systemctl edit xxx.service`
2. which results in `/xxx.service.d/override.conf`

# approach 3
1. `systemctl edit --full xxx.service`
2. which results in file `/etc/systemd/system/xxx.service`
