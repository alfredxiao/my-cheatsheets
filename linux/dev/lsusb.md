* `lsusb`: Lists `USB` (Universal Serial Bus) devices currently connected to the machine. Although `USB` devices for almost any imaginable purpose exist, the `USB` interface is largely used to connect input devices — keyboards, pointing devices — and removable storage media.

* `lsusb` lists devices with their ids
* `lsusb -v -d 1781:0c9f` gives info about selected device
* `lsusb -s BUS_NUM:DEVICE_NUM` to list that device only
  * e.g. 'Bus 002 Device 003' -> '2:3'
