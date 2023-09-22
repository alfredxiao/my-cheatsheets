* `ssh -Y user@host` to be able to run GUI program remotely
  * requires `X11Forwarding yes` on remote server's `/etc/ssh/sshd_config` file
* `ssh-copy-id` basically adds an entry to `~/.ssh/authorized_keys` file in the remote machine

- OpenSSH application keeps track of any previously connected hosts in the `~/.ssh/known_hosts` file. This data contains the remote servers’ public keys.

# Run a command
- `ssh user1@host1 ls /`

# Connect to different port
- `ssh -p 2222 user1@host1`

# Connect with an identity key file
- `ssh -i /path/to/private_key user1@host1`

# `/etc/ssh/ssh_know_hosts`
- This file is like the `known_hosts` file for all users (on the system)

# Key files
- server keys `/etc/ssh`
  - `ssh_host_*_key`
  - `ssh_host_*_key.pub`
- client keys `~/.ssh` (assuming rsa)
  - `id_rsa`
  - `id_rsa.pub`
- it is critical that the private key files are properly protected. Private key files should have a `0640` or `0600` (octal) permission setting and be root owned.
- However, public key files need to be world readable.

# algorithms
- rsa
- dsa
- ecdsa
- ed25519
