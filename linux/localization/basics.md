# Unicode
- is logical mapping between a number and a character in languages
- does not yet cover every language in the world

# UTF
- is encoding scheme(s) that maps from a byte sequence into a unicode character.
## Encodings
- UTF-8: uses 1~4 bytes to represent a character
- UTF-16: uses 2 or 4 bytes
- UTF-32: 4 bytes

# Timezone
- `/etc/timezone` text file in Debian
- `/etc/localtime` binary file in Redhat
- but **not** expected to edit directly,
- To change the time zone for a Linux system, copy or link the appropriate time zone template file from the `/usr/share/zoneinfo` folder to the `/etc/timezone` or `/etc/localtime` location. The `/usr/share/zoneinfo` folder is divided into subfolders based on location.
- If you just need to change the time zone for a single session or program, instead of changing the system time zone you can set the time zone using the `TZ` environment variable. That overrides the system time zone for the current session.
  - e.g. `export TZ='America/New_York'`
  - e.g. `TZ=UTC mycommand`
