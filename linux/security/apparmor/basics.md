# Bascis
- Similar to SELinux
- is kernel enhancement used to confine applications to a select set of resources
- it is path based
- Common on Debian based systems
- install tools and profiles
  - `sudo apt install -y apparmor-utils`
  - `sudo apt install apparmor-profiles`
- **NOTE** AppArmor runs as a service, not like SELinux

# Config
- `/etc/apparmor.d`

# Modes
- enforce
- complain (similar to `Permissive` in SELinux?)

# Profile
- A Profile is used to specify the capabilities of a specific program. A Profile can be enabled or disabled. For any Profile that is enabled it can be placed into one of two modes:
  - `Complain` – the Profile is active but not following the specified rules. A log is kept of all broken rules
  - `Enforce` – the Profile is active and the rules are active.
