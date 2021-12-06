* displaying information about PCI buses in the system and devices connected to them
* `lspci -k` show devices along with the kernel driver and modules used
* `lspci -v` shows even more

* `lspci` : Shows all devices currently connected to the `PCI` (Peripheral Component Interconnect) bus. `PCI` devices can be either a component attached to the motherboard, like a disk controller, or an expansion card fitted into a PCI slot, like an external graphics card.

* `lspci` lists devices with their addresses
* `lspci -s 04:02.0 -v` gives details
  * Example: `lspci -s 04:02.0 -v`
  * `-s` means slot
* `lspci -s 04:02.0 -k` gives info about kernel module in use.
