# Basics
* this allows for any size of boot loader program
* also the ability to store multiple boot loader programs for multiple operating systems.
* Requires **64bit** OS
* Prevents unauthorised OS from booting on the system
  * e.g. somebody with live USB, after booted, takes data of the original disk away
* UEFI firmware
  * is aware of partitions, bootloaders, and operating systems
  * has knowledge of a file system and supports FATs and VFAT
  * is capable of interpreting GPT and MBR (for backward compatibility)
  * is capable of accessing files on the ESP partition
  * can execute code in particular format (.efi)

# Where to find boot loader
* no more boot sector, but a special disk partition
  * called EFI System Partition (ESP)
* one partition out of the partition entries in GPT partition table is marked with boot flag, which is an ESP partition
  - Has to be `VFAT` or `FAT` variants like `FAT16` and `FAT32`

# efibootmgr
- `efibootmgr -v` shows EFI boot entries
- a typical entry for installed OS
  - `Boot0002* Fedora        HD(1,800,61800,6d98f360-cb3e-4727-8fed-5ce0c040365d)File(\EFI\fedora\grubx64.efi)`
  - Here, it is not just saying "boot from this disk", but "boot this specific bootloader in this specific location on this specific disk"

# UEFI enabled Operating System
- This is the mechanism the UEFI spec provides for operating systems to make themselves available for booting: the operating system is intended to install a bootloader which loads the OS kernel and so on to an EFI system partition, and add an entry to the UEFI boot manager configuration with a name - obviously, this will usually be derived from the operating system's name - and the location of the bootloader (in EFI executable format) that is intended for loading that operating system.

# Which partition
* ESP is typically mounted in the `/boot/efi` directory, and the boot loader files are typically stored using the `.efi` filename extension
* if a partition is ESP partition, you should see `/boot/efi` mounted by
  * `fdisk -l`

# Find out if your linux boots from EFI
- `ls /sys/firmware/efi`
  - 'no such file' -> bios
  - several file/dirs -> efi
- `efibootmgr -v`
  - shows boot menu entries
  - or, in non EFI boot mode, shows error message
