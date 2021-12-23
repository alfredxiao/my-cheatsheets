# Basics
- The `/etc/security/access.conf` file specifies combinations for which a **login** will be either accepted or refused.
  - (user/group, host),
  - (user/group, network/netmask)
  - or (user/group, tty)

# Format
- `permission:users/groups:origins`
- `@group1` for groups

# Keywords
- `ALL`
- `NONE`
- `EXCEPT`
- `LOCAL` means logging in via physical terminal

# Examples
- `-:user1:ALL`
- `-:root:ALL EXCEPT LOCAL`
- `+:ALL EXCEPT root: 10.1.1.`

# Prerequisite
- `account required pam_access.so` added to `/etc/pam.d/login` for terminal login
- `account required pam_access.so` added to `/etc/pam.d/sshd` for SSH login
