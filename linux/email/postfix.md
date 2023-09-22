# Install
- `yum -y install postfix mailx mutt`

# Use the postconf command to configure Postfix so that it listens on all network interfaces:
- `postconf -e inet_interfaces=all`

# Enable service
- `systemctl enable postfix.service --now`

# Send a msg
- `mail -s 'test' cloud_user@server1 < /etc/hosts`

# Verify
- `tail /var/log/maillog`

# mutt
