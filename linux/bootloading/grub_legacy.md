# Basics
- Runs on MBR
- The good thing about GRUB is that is has knowledge of Linux file systems
  - so your `/boot` partition can be, e.g. `ext3`

# Stages
## Stage 1
- Primary boot loader is read into memory by BIOS from MBR
- The primary boot loader exists on less than 512 bytes with MBR
- It is capable of loading either stage 1.5 or stage 2 boot loader
## State 1.5
- This stage is necessary for in some hardware and disk settings
  - e.g. when the `/boot` partition is above the 1024 cylinder head of the drive, or when using LBA mode
- Stage 1.5 boot loader is read into memory
- The stage 1.5 boot loader is found either on the `/boot` partition or on a small part of the MBR.
## Stage 2
- The stage 2, or secondary boot loader is read into memory from `/boot` partition
- The secondary boot loader looks for `/boot/grub`, and files `grub.conf` or `menu.lst`
- The secondary boot loader displays GRUB menu that allows user to select which kernel or operating system to boot, pass arguments to the kernel or look at system parameters.
- Another file is `device.map`, which indicates which drive has linux kernel to be loaded
- The secondary boot loader reads the operating system or kernel into memory
## Kernel Stage
- Then kernel image and init RAM disk are loaded into memory
- Control then passed to kernel
- Kernel starts with minimal hardware setup and decompress itself which then loaded into high memory
- RAM disk is mounted, it is a temporary root file system in RAM
  - and it allows kernel to fully boot without having to mount any physical disks
  - it contains necessary modules/drivers to interface with peripherals
  - it contains tools like `insmod` to load modules into kernel
- Once kernel is booted, real root file system is pivoted and initrd is unmounted and the real root file system is mounted
- Kernel loads kernel modules
## Init Stage
- First user-space (`init`) is started


# Config
- `grub.conf` or `menu.lst`
  - `kernel vmlinux.img root=….   ` —> Kernel image
  - `initrd initramfs.img`  —> init RAM image
  - provides multiple options (each can be a different version of kernel)

# grub-install
- to install:
  - `grub-install /dev/sda1`
- where the location (`/dev/sda1`) is the path to boot partition, can be found via
  - `df -h`
  - `findmnt /boot`
  - `lsblk`
- or alternatively `grub-install '(hd0)'`
  - where `(hd0)` is found
    - by running `grub` to start the GRUB 'shell'
    - `find /grub/stage1` which tells `(hd0,0)`
- **Don’t** really do it when it is already installed
