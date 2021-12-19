When cloning VM, be care of not duplicating below items should you run both original and the cloned VMs.
- Hostname
- Mac Address
- IP address (if static)
- Machine ID (`/etc/machine-id`)

# To deduplicate (run on the cloned VM)
- `rm /etc/machine-id`
- `rm /var/lib/dbus/machine-id`
- `dbus-uuidgen --ensure`
