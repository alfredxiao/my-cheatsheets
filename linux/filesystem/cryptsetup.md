# Basics
- install `cryptsetup` package

# Setps
1. `cryptsetup luksFormat /dev/nvme2n1p1`
  - you will need to setup a passphrase
2. `cryptsetup luksOpen /dev/nvme2n1p1 mysecretdata`
  - you will need to enter above passphrase
3. `mkfs -t ext4 /dev/mapper/mysecretdata`
4. `mkdir /mnt/mydata`
5. `mount -v /dev/mapper/mysecretdata /mnt/mydata`
6. `restorecon -vRF /mnt/mydata`
  - `mount -v -o remount /dev/mapper/mysecretdata /mnt/mydata` to verify there is no SELinux label mistach warning
- below to close/stop using it
7. `umount /mnt/mydata`
8. `cryptsetup luksClose /dev/nvme2n1p1`

# Use a key file instead
1. `dd if=/dev/random of=/root/luks_key bs=4096 count=1`
2. `restorecon -vF /root/luks_key`
3. `cryptsetup luksAddKey /dev/nvme2n1p1 /root/luks_key`
-below steps to close it
4. `umount /mnt/mydata`
5. `cryptsetup -v luksClose mysecretdata`
6. `cryptsetup -v luksOpen /dev/nvme2n1p1 mysecretdata --key-file /root/luks_key`

# Add to cryptab and fstab after setting up key file
1. `lsblk -f | grep nvme2n1p1` to grab the UUID
2. `vi /etc/crypttab`
```
mysecretdata UUID=4306b5bd-d105-45d7-bbce-185b673a7be3 /root/luks_key luks
```
3. `vi /etc/fstab`
```
/dev/mapper/mysecretdata /mnt/mydata xfs defaults 0 0
```
