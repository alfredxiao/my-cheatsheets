
# Reload Config
* `daemon-reload` to reload unit files etc.


# View unit file
- `systemctl cat abc.service`

# Remote host
* `-H user@host` to run against a remote host (via ssh)

* `systemctl -t help` Available unit types

# List unit files
* `systemctl list-unit-files`List available units
* `systemctl list-unit-files --type=service` Further filter by type
- `enabled` Service starts at system boot.
- `disabled` Service does not start at system boot.
- `static` Service starts if another unit depends on it. Can also be manually started

# List units
* `systemctl list-units` List active units (or have been active during current system session)
* `systemctl list-units --all` List all (not necessary active) units
* Further filter by type: `systemctl list-units --type=service`

# List failed units
- `systemctl --failed`

# Operation on a service unit
* `status` as in `systemctl status abc.service`
* `start`
* `stop`
* `restart`
* `enable` make it start automatically at boot time
* `disable` make it not start automatically at boot time
* `is-enabled`
* `is-active`
* `is-failed`
* `reload` reload the unit's own config, e.g. `httpd.conf`
* `mask` prevents the unit from starting
* `unmask` revert mask

# Misc
  * `isolate` ???
  * `get-default`
  * `set-default`
  * `default` jumps to default target
  * `daemon-reload` reload the unit's config file
  * `is-system-running` reports whether system is `running`, `degraded` (some units failed), `initializing`, etc.

# View dependencies
  * `systemctl list-dependencies nginx.service`: view dependency tree
  * `systemctl list-dependencies -all nginx.service`: view dependency tree and transitive dependencies
  * `systemctl show myunit.type` show configs managed by systemd
  * edit a unit file (without knowing its location): `systemctl edit myunit.type`

# Paths
  * `/lib/systemd/system`: For package-installed units/distribution wide configuration
  * `/etc/systemd/system`: For administrator-configured units/Local configuration
  * `/run/systemd/system`: For non-persistent runtime modifications/Runtime configuration
