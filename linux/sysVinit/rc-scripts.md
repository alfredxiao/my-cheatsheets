# RC or Init Scripts
## Location
- RedHat: `/etc/rc.d`
- Debian: `/etc/init.d`
## Contents:
- `rc.sysinit` —> house keeping before entering run level
- `rc?.d` —> scripts for run levels
- `rc.local` —> after run level scripts finish (can be customised to run tasks that do not have init script)
- `rc` —> conductor of run level orchestra
### Example Contents inside `/etc/rc.d/rc3.d`
- K20nfs kill those services
- S25netfs Start those services
- Note those xx numbers are ordered by their numbers
