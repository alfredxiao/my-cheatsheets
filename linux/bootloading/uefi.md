# Where to find boot loader
* not more boot sector, but a special disk partition
  * called EFI System Partition (ESP)

# Which partition
* ESP is typically mounted in the `/boot/efi` directory, and the boot loader files are typically stored using the `.efi` filename extension
* if a partition is ESP partiton, you should see `/boot/efi` mounted by
  * `fdisk -l`

# ESP partition
- Has to be VFAT or FAT32
