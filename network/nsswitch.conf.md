`/etc/nsswitch.conf`
- is used by the GNU C Library and certain other applications to determine the sources from which to obtain name-service information in a range of categories, and in what order.
- Each category of information is identified by a database name.
- The order in which the services to be queried are listed determines the order in which they are queried when attempting to resolve a name. The query-order list is in the service description in the /etc/nsswitch.conf file.
- The services are queried from left to right and by default searching stops when a resolution is successful.

# Example databases
- `passwd`
- `shadow`
- `group`
- `hosts`

# Example services
- `files`
- `dns`
- **note** each service is backed by a library file `libnss_SERVICE.so`, e.g. `/usr/lib64/libnss_files.so`

# `files`
- The following files are read when "files" source is specified for respective databases:
```
aliases     /etc/aliases
ethers      /etc/ethers
group       /etc/group
hosts       /etc/hosts
initgroups  /etc/group
netgroup    /etc/netgroup
networks    /etc/networks
passwd      /etc/passwd
protocols   /etc/protocols
publickey   /etc/publickey
rpc         /etc/rpc
services    /etc/services
shadow      /etc/shadow
```

# Actions
- You are able to control the lookup behavior more precisely using “action items” that describe what action to take given the result of the previous lookup attempt. Action items appear between service specifications, and are enclosed within square brackets, [ ].
- e.g. `hosts: dns [!UNAVAIL=return] files` This means that we should use the hosts file only if our name server is unavailable for some reason.
