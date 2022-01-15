# find out current run level
- `runlevel`

# Defined Runlevels
* A Runlevel is a high level system state, the definitions of each runlevel can vary from system to system, but levels 0/1/6 are common
* `0`: System shutdown, can safely remove power
* `1`: Single user mode, without network
* `2`: multi user mode, without network
* `3`: multi user mode, with network
* `4`: custom
* `5`: multi user mode, graphical mode login
* `6`: reboot
* Varies slightly across distributions

# Redhat based
- `0~6` except 4

# Debian based
- only `0/1/2/6` are valid run levels
- `1` single user
- `2` multi user mode with GUI

# Changing Runlevel
- View current runlevel: `runlevel`
    - e.g. `N 5` —> Previously N (n/a), Now 5
- Change run level: `telinit 3`
- Set runlevel DURING boot process: press anykey to activate GRUB menu, type `a` or `e` or something to edit startup command, append with `3`
