# Basics
- Point in time copy

# Create Snapshot LV
- `lvcreate -L 50M -s -n <SNAPNAME> /dev/<VGNAME>/<LVNAME>`

# Info
- `lvdisplay /dev/<VGNAME>/<SNAPNAME>`

# Increase size
- `lvextend -L +50M /dev/<VGNAME>/<SNAPNAME>`

# Remove
- `lvremove /dev/<VGNAME>/<SNAPNAME>`

# Merge
1. `umount /mnt/themountpoint`
2. `lvconvert --merge <VGNAME>/<SNAPNAME>`
- note after merging, the snapshop LV is gone
3. `mount /dev/mapper/<VGNAME>-<LVNAME> /mnt/themountpoint`
