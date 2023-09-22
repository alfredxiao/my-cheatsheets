- reload profile status
- loads AppArmor profiles into the kernel

# Add
- `sudo apparmor_parser -a /etc/apparmor.d/usr.bin.firefox`

# Remove
- `sudo apparmor_parser -R /etc/apparmor.d/usr.bin.firefox`
