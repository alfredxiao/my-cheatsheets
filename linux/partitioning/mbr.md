# MBR Partitioning scheme
- Saves partition table on first sector of disk (LBA 0)
- Each partition entry (in the partition table) is 16 bytes, at most four entries.
  - So partition table is limited to 4 partitions.
  - with 1 of of 4 can be extended to 23 logical partitions
- each partition size limited to 2TB
- one of the four entries can be marked as bootable
- MBR contains the machine code instructions for booting the machine, called boot loader, along with the partition table.
  - First 446 bytes are the primary boot loader
  - Next 64 bytes are partition table
  - Last two bytes are magic number `0xAA55`

# Location
## Addressing
- traditional hard disk addressing scheme like CHS (`Cylinder-head-sector`) uses a tuple which defined the cylinder, head and sector of the block as it appear on the physical hard disk
  - this does not work on other devices like tapes, networked storage
- LBA (`logical block address`) is an abstraction layer to address by a single integer, pointing to a logical block, which generally is 512 bytes, or sometimes 4096 bytes (not sure)
## MBR Record
- Stored at LBA 0, i.e. the first 512 bytes
