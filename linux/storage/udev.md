Another feature of the `udev` application is that it also creates persistent device files for storage devices. When you add or remove a removable storage device, the /dev name assigned to it may change, depending on what devices are connected at any given time. That can make it difficult for applications to find the same storage device each time.
To solve that problem, the `udev` application uses the /dev/disk directory to create links to the /dev storage device files based on unique attributes of the drive. `udev` creates four separate directories for storing links:
* `/dev/disk/by-id`: Links storage devices by their manufacturer make, model, and serial number
* `/dev/disk/by-label`: Links storage devices by the label assigned to them /dev/disk/by-path Links storage devices by the physical hardware port they are connected to
* `/dev/disk/by-uuid`: Links storage devices by the 128-bit universally unique identifier (UUID) assigned to the device
With the `udev` device links, you can specifically reference a storage device by a permanent identifier rather than where or when it was plugged into the Linux system.
