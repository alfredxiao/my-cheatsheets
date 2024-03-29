LVM (Logical Volume Manager)
- Allows to group/combine disks/partitions that are assembled into a single or multiple filesystems
    - Nearly any mount point EXCEPT /boot.  —> because GRUB can’t read LVM metadata
- Features
    - Flexibility - allow resizing of volumes
    - Snapshots - allow point-in-time backup copies
- Layers
    - Physical Volume: Your disks, e.g. 3 disks
    - Volume Group: a logical name for your 3 disks grouped together
    - Logical Volume: Logical portion that you VG is split into.
    - File system: one LV maps to one Filesystem
    - `mkfs.xfs /dev/mapper/<VGNAME>-<LVNAME>`  —> format a LV file system
- To check current LVM status
    - cat /etc/fstab
        - where lines starting with ‘/dev/mapper/‘ is LV
