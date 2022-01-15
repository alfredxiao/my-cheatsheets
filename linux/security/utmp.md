# Files
- `/var/run/utmp`
- `/var/log/wtmp`
- `/var/log/btmp`

- The 'u' stands for user. utmp gives information about who is on the system.
- The 'w' in wtmp probably comes from 'who', it acts as a historical utmp
- The 'b' comes from 'bad', btmp records the bad login attempts.

# `utmpdump` to view raw records
- `utmpdump /var/run/utmp`
- `utmpdump /var/log/wtmp`
- `utmpdump /var/log/btmp`
