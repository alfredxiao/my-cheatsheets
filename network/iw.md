# `iw`
* similar to `ip`
* meant to replace `iwconfig` and `iwlist`
* `iw list` --> list all wireless devices
* `iw event` --> monitor events from the kernel
## `iw dev`
* `iw dev` -> view available wireless interfaces
* `iw dev wlan0 link` display link info
* `iw dev wlan0 info` show info for an interface
## access point
* `iw wlan0 scan` --> scan for available SSIDs
* `iw dev wlan0 connect SSID` connect to an SSID
* `iw dev wlan0 disconnect` disconnect from a wireless network
