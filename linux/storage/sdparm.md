list or change SCSI/SATA device parameters and send simple SCSI commands.

- `sdparm -a /dev/sda` list all known fields for a device
- `sdpar --command=CMD /dev/sda` to send a SCSI command to a device, e.g. eject, start, stop
