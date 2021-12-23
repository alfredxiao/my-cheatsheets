Is RAID 0 plus RAID 1 setup.

Four disks required.
First a RAID 0 controller, then two RAID 1 controllers.
Data is stripped by RAID 0 controller, any piece of data goes to only one of the two RAID 1 controllers, which writes to two disks, so that it always has two copies.
50% of storage is available for data.
