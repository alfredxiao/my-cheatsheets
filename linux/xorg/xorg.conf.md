# X11 Configuration
- `/etc/X11/xorg.conf`
    - Primary config for X Server
    - `man 5 xorg.conf`
    - `/etc/X11/xorg.conf.d/`
        - Supplementary Xorg configuration files
- X/Xorg
    - either command can create a new configuration file for your X server
- xdpyinfo
    - Displays info about current X session and X server instance

# Tweek Configuration
1. Stop X Server via `systemd isolate multi-user.target` or `telinit 3`
2. `Xorg -configure` generates new config
3. Move config files to correct location
4. Restart X server

# Config File
Input Device: Configures the session’s keyboard and mouse
Monitor: Sets the session’s monitor configuration
Modes: Defines video modes
Device: Configures the session’s video card(s)
Screen: Sets the session’s screen resolution and color depth
Module: Denotes any modules that need to be loaded
Files: Sets file path names, if needed, for fonts, modules, and keyboard layout files Server Flags: Configures global X server options
Server Layout: Links together all the session’s input and output devices

# Troubleshooting
- `~/.session-errors` generated if any error
