* `/etc/systemd/journald.conf`

* `Storage=`
  * `none`
  * `volatile` saved in RAM and in /run/log/journal
  * `persistent` saved in /var/log/journal persistently
  * `auto` like persistent, ONLY if /var/log/journal folder exists, otherwise fallback to volatile
