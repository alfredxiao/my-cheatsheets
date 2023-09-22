# Basics
- functionality is similar to systemd automount
- install package `autofs`
- start the service `systemctl enable autofs --now`

# Config
- `/etc/sysconfig/autofs`
- `/etc/auto.master`

# Create an indirect automount
1. `vi /etc/auto.master`
```
/myautomounts /etc/auto.myautomounts
```
2. `mkdir /myautomounts`
3. `vi /etc/auto.myautomounts`
```
point1 -fstype=ext4 :/dev/nvme1n1p1
point2 -fstype=xfs,ro :/dev/nvme1n1p2
```
4. `systemctl restart autofs`

# Create direct automount
1. `vi /etc/auto.master`
```
/- /etc/auto.another --timeout=300
```
2. `vi /etc/auto.another`
```
/another/data -fstype=ext3 :/dev/nvme1n1p3
```
3. `systemctl restart autofs`
