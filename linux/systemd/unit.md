# [Unit]
- `After` start this unit after the other, often used together with `Requires`
- `Before` start this before the other
- `Requires` must start the other as well - if the other cannot start, don't start this
- `Wants` better start the other as well - if the other cannot start, this still starts
- `Conflicts` must not start together with the other - if the other starts, don't start this
- `Description` describes this unit
- `Documentation` provides URIs pointing to a URL, man page, etc.
- `AllowIsolate` whether allow `systemctl isolate` on a target unit

# [Service]
- for service unit only
- `ExecStart` command to run to start this unit
- `ExecStop` command to run to stop this unit
- `ExecReload` command to run to reload this unit
- `Environment` sets environment variable substitutes, separated by a space
- `Environment File` Indicates a file that contains environment variable substitutes
- `RemainAfterExit` when `no` (default) `ExecStop` is called when the process started by ExecStart terminates. when `yes`, the service is left active even when the process started by ExecStart terminates.
- `Restart` Service is restarted when the process started by ExecStart terminates. values include `no` (default), `on-success`, `on-failure`, `on-abnormal`, `on-watchdog`, `on-abort`, or `always`.
- `Type`
  - `simple` (default) – starts the service immediately. It is expected that the main process of the service is defined in `ExecStart`.
  - `forking` – considers the service started up once the process forks and the parent has exited.
  - `oneshot` – similar to simple, but it is expected that the process has to exit before systemd starts follow-up units (useful for scripts that do a single job and then exit). You may want to set `RemainAfterExit=yes` as well so that systemd still considers the service as active after the process has exited.
  - `dbus` – similar to simple, but considers the service started up when the main process gains a D-Bus name.
  - `notify` – similar to simple, but considers the service started up only after it sends a special signal to systemd.
  - `idle` – similar to simple, but the actual execution of the service binary is delayed until all jobs are finished.

# [Install]
- `Also` Sets additional units that must be enabled or disabled for this service. Often the additional units are socket type units.
- `RequiredBy` Designates other units that require this service.
- `WantedBy` Designates which target unit manages this service.
