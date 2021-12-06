`dd if=/dev/urandom of=file5m bs=5MB count=1`

# create a bootable USB drive
- `dd if=boot.img of=/dev/sdc`

# Backup MBR
- `dd if=/dev/xvda of=/tmp/mbr.img bs=512 count=1`

# Create random file
- `dd if=/dev/urandom of=afile bs=1024k count=10`

# Copy an entire drive to another drive
* prerequisite: neither is mounted
* `dd if=/dev/sdb of=/dev/sdc status=progress`

# erase an entire drive (shreding)
- `dd if=/dev/zero of=/dev/sdc status=progress`
- to really shred a disk, this needs to be run more than 10 times
