# Transient Timers
* `systemd-run --on-active=30sec /bin/touch /home/user1/file1`

# Calender Times
* `sudo systemctl edit --full --force blahwoof.timer`
```
[Unit]
Description=Run the blahwoof service every 5 seconds

[Timer]
OnCalendar=*-*-* *:*:0/5
AccuracySec=1s

[Install]
WantedBy=timers.target
```
* `sudo systemctl edit --full --force blahwoof.service`
```
[Unit]
Description=Write a quick entry to the systemd journal

[Service]
Type=oneshot
User=cloud_user
ExecStart="/home/cloud_user/journo.sh"

[Install]
WantedBy=blahwoof.timer
```
* `sudo systemctl enable blahwoof.timer`
* no need to enable blahwoof.service

# Monotonic Timers
