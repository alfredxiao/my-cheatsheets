originally `tcpd`

# Features
- checked on permissions files (hosts.allow and hosts.deny)
- Logging – Connections that are monitored by TCPD are reported through the syslog facility.
- Access Control – TCPD supports a simple form of access control that is based on pattern matching. You can even hook the execution of shell commands/script when a pattern matches.
- Host Name Verification – TCPD verifies the client host name that is returned by the address->name DNS server by looking at the host name and address that are returned by the name->address DNS server.

- If a service can employ TCP Wrappers, it will have the `libwrap` library compiled with it.
- to check
  - `ldd $(which sshd) | grep libwrap`

# `/etc/hosts.allow` and `/etc/hosts.deny`
- Because access is allowed if the remote system’s address is not found in either file, it is
best to employ the ALL wildcard in the `/etc/hosts.deny` file: `ALL: ALL`
- Be aware that some distributions use `PARANOID` instead of `ALL` for the address wildcard.
- To list individual IP addresses in the hosts.allow file, you specify them separated by commas:
`sshd: 172.243.24.15, 172.243.24.16, 172.243.24.17`
- or simply `sshd: 172.243.24.`
