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

# Commands
- RedHat based commands: `grub2-*`
- Debian based commands: `grub-*`
## `grub2-editenv`/`grub-editenv`
- `grub2-editenv list`
  - to view default boot entry for grub config
  - (redhat)
## `grub2-mkconfig`/`grub-mkconfig`
  - generates `/boot/grub2/grub.cfg` from `/etc/default/grub` config file
  - no args needed when running it
  - should you want to custom the generation, take a look at `/etc/grub.d/`
- On Debian, commands have no ‘2’ in them
## `update-grub`
- available on debian as well, which does same thing
