- swap file: performance overhead
- swap partition: less overhead
- tools can be used: fdisk, gdisk, parted
    - specified partition type/id: 82(00)
    - after create swap partition, one has to make a swap filesystem on it, also list it in fstab to make it persistent across restarts, also need to ‘swapon’ it.
- mkswap
- fstab
    - you can have multiple entires of swap partitions
    - avoid using /dev/sda2 in fstab, better use UUID, or label
    - ‘LABEL=SWAP swap swap defaults 0 0’
- swapon
- swapoff
    - swapoff -L SWAP —> turn off swap by label

Swap Space
- swap partition  (very common setup)
- swap file (rather than dedicated partition)
    - Much slower than using dedicated partition (this an old saying!?)
- Sizing
    - old ‘rule of thumb’ is 1.5~2x the size of RAM
    - nowadays, it is up to you, maybe <50%, maybe zero

# Priority
- the higher the number, the higher the swap partition or file gets used
