# How to Check a Program is PAM-aware
- `sudo ldd /usr/sbin/sshd | grep libpam.so`
  - if output shows the program uses libpam.so, then yes, it is PAM aware

# Location of Module Object Files
- `/lib/security/` or `/lib64/security`
- `/usr/lib/x86_64-linux-gnu/security`

# Architectural Parties
- Application
  - Uses pam library API, offload authentication work to PAM API and PAM modules
- PAM Config Files
  - Configures what modules are used for what program
- Linux PAM Library (part of linux core)
- PAM Modules (pluggable, can install additional modules)
  - as `.so` files dynamically loaded as application is not coded into linux PAM library, not module libraries

# Config
- `/etc/pam.d`
- or `/etc/pam.conf`

# Config Format
## `pam.conf`
- `service type control-flag module module-arguments`
## files under `/pam.d`
- `type control-flag module module-arguments`
## Type
- `authentication`: authenticate a user and set up user credentials.
- `account`: provide services for account verification: has the user’s password expired?; is this user permitted access to the requested service? etc.
- `password`: are responsible for updating user passwords and work together with authentication modules.
- `session`: manage actions performed at the beginning of a session and end of a session.
## Control Flags
| Flag | Module outcome to Overall outcome | This Module Failure Terminates Overall Process |
|------|-------------------------------------------|-------------------------------------------------|
| `required` | Module Fails means overall Fails | No |
| `requisite` |  Module Fails means overall Fails | Yes |
| `sufficient` | If preceding modules all succeed, the success of this module leads to an immediate overall success | No |
| `optional` | ? | No unless this is the only module |
- **NOTE** that this keyword is not an indicator to the module, but to the underlying PAM library.

# References
 - https://www.linuxjournal.com/article/5940
 - https://www.redhat.com/sysadmin/pam-configuration-file
