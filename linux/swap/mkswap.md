# File approach
- `sudo dd if=/dev/zero of=/swap-file count=2 bs=1GiB`
- `sudo chmod -v 0600 /swap-file` (recommended, not mandatory)
- `sudo chown -v root:root /swap-file`
- `sudo mkswap /swap-file`
- `sudo swapon /swap-file`
## Make it permanent
- `sudo vi /etc/fstab`
- `/swap-file swap swap defaults 0 0`
## Remove swap file
- `sudo swapoff -v /swapfile`
- remove entry from fstab
- `sudo rm /swapfile`

# Partition approach
## Make a swap partition
- `mkswap /dev/sda2`
## With a label
- `mkswap -L SWAP /dev/sda2` —> make a filesystem with a label
