# Display Manager
- The **display manager** is a bit of code that provides the GUI login screen for your Linux desktop. After you log in to a GUI desktop, the display manager turns control over to the **window manager**. When you log out of the desktop, the **display manager** is given control again to display the login screen and wait for another login.
- Some desktop environments comes with display manager, some display manager implementations are not part of a desktop environment, but essentially any display manager can be used in any desktop environment.
- Uses the `X Display Manager Control Protocol` (`XDMCP`) to handle the graphical login process

# Display Managers
- `XDM`, one of the display manager imlementations, is the basic display manager software avail- able for Linux
  - `/etc/X11/xdm/xdm-config`
- `KDM`
- `GDM`
- `LightDM`

# Window Manager
- is yet another client of X Server
- it manages the creation, movement, and destruction of windows on a GUI desktop
- also controls the appearance of the windows it generates. This includes the functional decorative aspects of the windows, such as the look of buttons, sliders, window frames, pop-up menus, and more.

# Display Server
- is same terminology as 'X Server'

# Desktop Environments
- GNOME
- KDE (Plasma)
- Cinnamon
- Xfce
