# Basics
- It can keep accurate time even on systems that have busy networks or that are down for periods of time, and even on virtualized systems.

# Config
- `/etc/chrony/chrony.conf` or `/etc/chrony.conf`
```
pool ntp.ubuntu.com iburst maxsources 4
pool 0.ubuntu.pool.ntp.org iburst maxsources 1
pool 1.ubuntu.pool.ntp.org iburst maxsources 1
pool 2.ubuntu.pool.ntp.org iburst maxsources 2

rtcsync
```
- The maxsources parameter is one exception to this similarity. It designates the maximum number of time servers from the designated source. If one of those servers goes down, chrony finds another one to take its place.
- `rtcsync` is for automatically sync with hwclock
