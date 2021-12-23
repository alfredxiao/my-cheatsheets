# Physical Volume
- `pvremove /dev/sde`

# Logical Volume
- `lvremove /dev/<VGNAME>/<LVNAME>`
- or `lvremove <VGNAME>/<LVNAME>`

# Volume Group
- must remove logical volumes before removing volume group
- `vgremove <VGNAME>`
