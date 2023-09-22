# Install
- install packages `qemu` `qemu-kvm` `libvirt-bin` (or `libvirt`) `bridge-utils` `virt-manager` `virtinst` (or `virt-install`)

# Enable KVM
- `sudo modprobe kvm`

# Create VM
- `virt-install --name=tiny --vcpus=1 --memory=1024 --cdrom=alpine-standard-3.15.0-x86.iso --disk size=1 --check disk_size=off`
  - add disk_size=off if default mount point for `/var/lib/libvirt/` has no enough space but symbloic link `/var/lib/libvirt/images` to another mount point with enough space
