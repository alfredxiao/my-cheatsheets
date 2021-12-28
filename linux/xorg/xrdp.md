- supports RDP (Remote Desktop Protocol)
- Xrdp provides only the server side of an RDP connection
- The package name on Linux is `xrdp`
- standard RDP port (TCP `3389`)
- The Xrdp server handles mouse movements and keystrokes as well as with audio transfers and mounting of local client drives on the remote system.

# Server setup
- `sudo apt install xrdp`
- `sudo usermod -aG ssl-cert xrdp`
- `sudo systemctl stop xrdp`
- `sudo systemctl start xrdp`

# Client
- You could run Microsoft Remote Desktop on Windows or Mac
- or run `remmina` on Linux

# Server Config
- server config `/etc/xrdp/xrdp.ini`
- `security_layer`
  - `negotiate` (default) Sets the security method to be the highest the client can use. This is problematic if the connection is over a public network and the client must use the standard RDP security method.
  - `tls`
