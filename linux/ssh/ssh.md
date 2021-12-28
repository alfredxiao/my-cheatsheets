* `ssh -Y user@host` to be able to run GUI program remotely
  * requires `X11Forwarding yes` on remote server's `/etc/ssh/sshd_config` file
* `ssh-copy-id` basically adds an entry to `~/.ssh/authorized_keys` file in the remote machine
* if you have a password protected id key file and wants to avoid entering password
  * run `ssh-agent bash` wraps a bash and this handles authentication for keys that requires password
  * then `ssh-add` to adds the passwords to `ssh-agent`
