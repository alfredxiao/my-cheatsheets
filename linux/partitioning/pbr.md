- The partition boot record is the first block of any **bootable** partition.

- It is also sometimes referred to as the Volume Boot Record (`VBR`).

- The MBR (or the GRUB chainloader) loads the PBR into memory and transfers control to it.

- Whereas the MBR code is operating-system independent, the code in the PBR is supplied by and works with a particular OS.

- In general, the job of the code in the PBR is to load a slightly larger "next-step" program.

- The code in the PBR can't read a file system's directory structure to locate this next piece of code to load. Such file-system-specific code would be too large to fit in the single 512 sector Instead a specific disk address for the next "chunk" is stored within the PPR. 
