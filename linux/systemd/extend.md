# approach 1
* mkdir /etc/systemd/system/xxx.service.d
* vi /etc/systemd/system/xxx.service.d/override.conf
* systemctl daemon-reload
* systemctl restart xxx.service

# approach 2
* systemctl edit xxx.service
* which results in `/xxx.service.d/override.conf`

# approach 3
* systemctl edit --full xxx.service
* which results in file `/etc/systemd/system/xxx.service`
