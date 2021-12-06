* install
  `sudo apt install systemd-container`
  `sudo apt install debootstrap`
* Create a container
  `sudo debootstrap --arch=amd64 buster /var/lib/machines/buster1`
* list instances
  `machinectl list-images`
* start the instance
  `sudo systemd-nspawn -D /var/lib/machines/buster1 --machine buster1`
* disconnect from the container
  `Ctrl-]]]`
* start the instance and boot it
  `sudo systemd-nspawn -D /var/lib/machines/buster1/ --machine buster1 -b`
* start up an instance
  `sudo machinectl start buster1`
* view status
  `machinectl status buster1`
  `systemctl status systemd-nspawn@buster1.service`
* connect to an instance
  `sudo machinectl login buster1`
* stop an instance
  `sudo machinectl stop buster1`
* create a service for the container
  `sudo systemctl edit --full --force buster1.service`
  ```
  [Unit]
Description=Buster1 Container

[Service]
LimitNOFILE=100000
ExecStart=/usr/bin/systemd-nspawn --machine=buster1 --directory=/var/lib/machines/buster1/ -b --network-ipvlan=eth0
Restart=always

[Install]
Also=dbus.service
WantedBy=default.target
  ```
