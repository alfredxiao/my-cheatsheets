* `su - user1`
* `visudo`
* `visudo -f /etc/sudoers.d/myextra`
  ```
  Cmnd_Alias WEB = /bin/systemctl restart httpd.service, /bin/systemctl reload httpd.service

  user2 ALL=WEB
  ```
