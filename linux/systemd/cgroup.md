* in systemd, cgroups are synonymous to slices

* configure a service to a slice
  `vim /etc/systemd/system/xxx.service`
  ```
  [Service]
  Slice=user.slice
  CPUWeight=100
  ```

* to see what you can configure
  `man systemd.resource-control`

* systemd-cgtop
* systemd-cgls
