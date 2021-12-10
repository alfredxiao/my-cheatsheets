# Basics
- `GPT` = GUID Partition Table
- Supports 128 partitions
- partition size up to ZB
- provides backup partition table

# Location of GPT Record
- The GPT header is in LBA 1
- The GPT header has a pointer to the partition table (Partition Entry Array), which is typically at LBA 2.
- LBA 0 still reserved for MBR

# Size
- Each entry on the partition table has a size of 128 bytes.
- UEFI Spec stipulates a minimum of 16,384 bytes allocated for partition entries
  - which is 128x128, meaning it supports 128 partition entries

# Hibrid MBR (LBA 0 + GPT)
- In operating systems that support GPT-based boot through BIOS services rather than EFI, the first sector may also still be used to store the first stage of the bootloader code, but **modified** to recognize GPT partitions.

# Boot Flag
- Like the MBR, UEFI marks one of the partitions with the boot flag, but this is only the EFI partition, never any of the OS partitions
