check to see if virtual machine extensions are on
- `cat /proc/cpuinfo | grep vmx`
- if not enabled, will have to use `qemu`
- and you will need to run x86 images if you are running virtual on top of virtual

# Tools
- `virt-install` install new image
- `virt-clone` clones existing image
- `virt-manager` GUI manager
- `virsh` command line interface for managing 
