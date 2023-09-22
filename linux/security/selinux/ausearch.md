When SELinux denies an action, an Access Vector Cache (AVC) message is logged to the /var/log/audit/audit.log and /var/log/messages files or the journald daemon logs it.

- Use the ausearch utility to find any recent AVC messages and confirm that SELinux denies the action:
  - `ausearch -m AVC,USER_AVC -ts recent`
  - The -m option specifies what kind of information ausearch returns.The -ts option specifies the time stamp. For example -ts recent returns AVC messages from the last 10 minutes or -ts today returns messages from the whole day.
- Use the journalctl utility to view more information about the AVC message:
  - `journalctl -t setroubleshoot --since= [time]`
