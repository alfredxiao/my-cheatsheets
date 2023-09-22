# Installation
- package name `quota`

# fstab
- add to options (or replace defaults): `usrquota,grpquota`
- remount via `sudo mount -o remount /mnt/data`
- confirm via `sudo cat /proc/mounts | grep /mnt/data`

# Commands
- `sudo quotacheck -cugm /mnt/data` to generate quota file
- `sudo quotaon -v /mnt/data` turn on quota
- `sudo edquota -u user1` to edit quota number
- `sudo edquota -g group1`
- `sudo setquota -u user2 100M 200M 0 0 /mnt/data` to set quota number (not edit)
- `quota -vs user1`
- `quota -vgs group1`
- `repquota -asgu` to report
