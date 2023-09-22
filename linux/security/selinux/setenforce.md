Open the file /etc/selinux/config  
Change option SELINUX from disabled to enforcing

$ setenforce
usage:  setenforce [ Enforcing | Permissive | 1 | 0 ]
$ #To Set mode to Permissive
$ setenforce Permissive
