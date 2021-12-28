The VNC server offers a GUI service at TCP port 5900 + n, where n equals the display number, usually 1 (port 5901).
On the command line you point the VNC client (called a viewer) to the VNC server’s hostname and TCP port.
Alternatively, you can use the display number instead of the whole TCP port number.

The client user is required to enter a predetermined password, which is for the VNC server and not Linux system authentication.


# Potential difficulties or concerns with VNC:
- The VNC server handles only mouse movements and keystrokes. It does not deal with file and audio transfer or printing services for the client.
- VNC, by default, does not provide traffic encryption, so you must employ another means of protection, such as tunneling through openSSH.
- The VNC server password is stored as plaintext in a server file.

# Install and setup VNC Server (TightVNC)
1. `sudo apt install tightvncserver`
2. `vncserver` (run as non-root)
  - **NOTE** this command accepts different options if you chose `tigervnc-standalone-server` in previous step
3. setup password, e.g. `vnc001`
4. optionally set view-only password, e.g. `vnc002`
5. This creates `~/.vnc/xstartup` start script
6. This also starts a process as `Xtightvnc :1 -desktop X -auth /home/alfred/.Xauthority ...` under your non-root user, and listens on port `5901`


# VNC Client
1. Run VNC Viewer (various software packages available on various platforms)
2. connect to `target-host-or-ip:1` which will translates to port `5901` for display `:1`, input your password like `vnc001`
  - e.g. run `vncviewer target-host-or-ip:1`
3. It probably does not show up a desktop environment. And If you run `DISPLAY=:1 xclock` the clock will be displayed on this client session.

# Make VNC Display Desktop
1. `vncserver -kill :1`
2. `touch ~/.Xresources`
3. add `startxfce4 &` to last line in `~/.vnc/xstartup`
4. Make sure there is no Xfce destkop running in `target-host-or-ip`
5. Connect VNC Viewer to `:1`
6. Login with password `vnc001`, it shows Xfce desktop with a user already logged in.

# Connect with SSH Tunneling
1. `ssh -L 5901:127.0.0.1:5901 -C -N -l alfred your_server_ip`
  - this opens up a local port 5901 but also use SSH to forward requests to 5901 on remote host.
2. Run VNC Viewer to connect to `localhost:1`

# add or change password
- `vncpasswd`
