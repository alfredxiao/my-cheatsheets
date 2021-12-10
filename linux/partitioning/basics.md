two types of partition tables
* MBR
* GPT (GUID Partitioning Table)

MBR and GPT are types of partition structure

`MBR`
* MBR is a special boot sector at the beginning of a disk, containing information about the OS boot loader and logical disk partition information.

`GPT`
* GPT is another type that uses GUID or globally unique identifier to define its partitions and is the newer standard.

To tell partiton table type:
* `parted -l`
* `parted /dev/sda print`
* `gdisk -l /dev/sda`
* `fdisk -l`

Tools
* MBR
  * `fdisk`
  * `cfdisk`
  * `sfdisk`
* GPT
  * `gdisk`
  * `cgdisk`
  * `sgdisk`
* Both
  * `parted`
  * `gparted` (GUI Tool)

# Commonly used disk partition type id
- `83` : linux partition
- `82` : linux swap partition
- `8E` : linux LVM Volume
- `EF` : EFI partition

# Directories on their own partitions (generally)
- `/var`
- `/boot`
