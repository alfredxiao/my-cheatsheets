# Physical Volume
- `pvcreate /dev/sde /dev/sdf1`
- you can provide one or multiple devices, or partitions

# Volume Group
- `vgcreate <VGNAME> /dev/sde /dev/sdf`
## to extend
- `vgextend <VGNAME> /dev/sdg`
## to remove
- `vgreduce <VGNAME> /dev/sdg`

# Logical Volume
- `lvcreate -n <LVNAME> -L 3G <VGNAME>`
- `lvcreate -n <LVNAME> -l 50 <VGNAME>` create by number of Logical Extend
- `lvcreate -n <LVNAME> -l 30%VG <VGNAME>` create by percentage
