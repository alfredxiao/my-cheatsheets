# Basics
- available on systemd enabled systems
- The timedatectl command provides one-stop shopping for all of the time informa- tion, including the hardware clock, called RTC; the date information; and the time zone information.

# display date and time settings
- `timedatectl`

# set both system time and RTC to the same time, thus persistent across next session
- `timedatectl set-time '2019-12-23 20:09:02'`
    - this will not work if NTP is on

# Other commands
- `timedatectl list-timezones`

# Set timezone/time (across reboots)
- `timedatectl set-timezone Australia/Melbourne`
- `timedatectl set-time '2015-12-01'`
- `timedatectl set-time '2020-10-11 17:51:00'`

# Set NTP on
- `timedatectl set-ntp yes`
