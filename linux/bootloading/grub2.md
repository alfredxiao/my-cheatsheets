- Works with GPT using UEFI
- boot.img on first 512 bytes
- Stage 1:
  - boot.img
  - Then
    - GPT Header (to indicate disk is GPT style)
    - Partition Entry Array (to indicate locations of partitions)
- Stage 1.5
  - core.img
  - from the ESP partition
- Stage 2
  - `/boot/grub2`
    - grubenv
    - themes

# Editing Config
- **NOT** like grub legacy, in GRUB2, we do not edit config files directly
- Example config (`grub.cfg`), not for direct editing
```
menuentry "CentOS Linux" { [...]
  set root=(hd1,1)
  linux16 /vmlinuz[...]
  initrd /initramfs[...]
}
menuentry "Windows" {
  set root=(hd0,1) [...]
}
```
- Location of the (generated) config file
  - `/boot/grub/grub.cfg` non EFI
  - `/boot/efi/EFI/centos/grub.cfg` EFI
  - `/etc/default/grub` for your editing, same location for Debian/Redhat

# Index!
- GRUB2 still uses `0` to identify the first drive
- BUT it uses `1` to identity the first partition, whereas GRUB Legacy uses `0` for this!

# Commands
- RedHat based commands: `grub2-*`
- Debian based commands: `grub-*`
## `grub2-editenv`/`grub-editenv`
- `grub2-editenv list`
  - to view default boot entry for grub config
  - (redhat)
## `grub2-mkconfig` (Redhat) / `grub-mkconfig` (Debian)
  - `grub2-mkconfig -o /boot/grub2/grub.cfg`
  - generates `/boot/grub2/grub.cfg` from `/etc/default/grub` config file
  - should you want to custom the generation, take a look at `/etc/grub.d/`
- On Debian, commands and paths have no `2` in them
## `update-grub`
- available on debian as well, which does same thing as `grub2-mkconfig`
- without the need to have `-o` option

# Kernel parameters
- `ro` mount root file system as readonly
- `mem` Set the total amount of available system memory
- `systemd.unit=` boot into specified target

# Reinstall
- `grub2-install /device`
