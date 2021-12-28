If you have any legacy X11 applications that will not support Wayland, do not despair. There is the `XWayland` software, which is available in the Weston package. XWayland allows X-dependent applications to run on the X server and display via a Wayland session.

In order to disable Wayland in GNOME Shell, you edit the `/etc/gdm3/custom.conf` file and set `WaylandEnable` to `false`. 
