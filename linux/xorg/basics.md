X.org is an implementation of X Window System, Wayland is a newer one.

# X Windows System
- X11
  - is the protocol between X clients and X server
## X Server
- Server side implementation of X11 protocol
- most popular and de-fecto is X.org implementation
- X Server responsible for drawing/display
- Applications are clients, they tell X Server to draw
- When running a GUI application remotely, X Server runs on your local computer, application/client runs on remote machine
## Window Manager
- Is just another client
- Windows, Bars, Titles, etc.
## Display Manager
    - Session management
- How does system knows to run a display manager
    - sysvinit: last line in /etc/inittab
    - systemd:
- `$DISPLAY`
    - host:X-Server-NUM.Monitor-NUM
    - e.g. “:0.0”

# Check
- `echo $WAYLAND_DISPLAY` if it shows something, you are using Wayland, otherwise X11
## or
1. `loginctl` note the session number
2. `loginctl show-session <SESSION_NUMBER> -p Type` 

## Test programs
- `xeyes`
- `xclock`
