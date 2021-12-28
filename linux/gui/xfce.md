# Install Xfce
1. `yum install epel-release -y` EPEL is the repo to install packages from
2. `yum groupinstall "Server with GUI" -y`
3. `yum groupinstall "Xfce" -y`
4. `systemctl set-default graphical.target`
5. `systemctl isolate graphical.target`
