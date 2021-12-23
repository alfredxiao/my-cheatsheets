`mkfs -t ext4 /dev/<VGNAME>/<LVNAME>`
`mkdir /mnt/newdir`
`mount -t ext4 /dev/mapper/<VGNAME>-<LVNAME> /mnt/newdir`
or `add to fstab` then `mount -a`:
`/dev/mapper/<VGNAME>-<LVNAME> /mnt/newdir ext4 defaults 0 0`


# Resize
- xfs can only grow in size
- `lvextend -L 250M /dev/<VGNAME>/<LVNAME>`
- or `lvextend -L +50M /dev/<VGNAME>/<LVNAME>`
- after resizing LV, resize fs (online resizing)
- `resize2fs /dev/mapper/<VGNAME>-<LVNAME>`
- `xfs_growfs` to resize xfs
- or combine fs resizing into lvextend
  - `lvextend -r -L 250M /dev/<VGNAME>/<LVNAME>`
- shrinke (will need to umount)
  - `lvreduce -r -L 160M /dev/<VGNAME>/<LVNAME>`
