- `$HOME/.ssh/authorized_keys`  for accounts login using keys rather than password

# Installation
- CentOS: `openssh`, `openssh-clients`, `openssh-server`
- Ubuntu: `openssh-server`, `openssh-client`


# Client Config
- `/etc/ssh/ssh_config`
- Contains OpenSSH client configurations. May be overridden by `ssh` command options or settings in the `~/.ssh/config` file.

# Server Config
- `/etc/ssh/sshd_config`
- Contains the OpenSSH daemon (sshd) configurations.
## Some config entries
- `AllowTcpForwarding`: Permits SSH port forwarding. (defaults to yes)
- `ForwardX11`: Permits X11 forwarding. (defaults to no)
- `PermitRootLogin`: Permits the root user to log in through an SSH connection. (Default is yes.) Typically, should be set to no.
