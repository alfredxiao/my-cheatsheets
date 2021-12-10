# Default umask value
- `umask`
- `0022` means
  - new file created are assigned `644` (subtracted from `666`)
  - new dir created are assigned `755` (subtracted from `777`)

# Set umask value
- `umask 027`
- `umask u=rwx,g=,o=` --> permission is `700`, meaning umask set to `077`
- **note** this is not persistent across reboots, but one can add this statement to user profile bash script
