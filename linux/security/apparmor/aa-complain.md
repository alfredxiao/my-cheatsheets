set an AppArmor security profile to complain mode

e.g.
- `aa-complain /usr/bin/man`
afterwards, run
`sudo systemctl reload apparmor.service` to have changes take effect
