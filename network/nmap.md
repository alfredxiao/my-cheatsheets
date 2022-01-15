# Fast mode
* `nmap -F 10.0.1.10`

# Additional aggressive scan
* `nmap -A -F 10.0.1.10` -> discover host OS info and service info

# Ping scan (no port scan)
* `nmap -sn 10.0.1.10`

# TCP Scan
* `nmap -sT 10.0.1.10`

# UDP Scan
* `nmap -sU 10.0.1.10`

# Stealthy scan
* `nmap -sS 10.0.1.10` (sends only SYN)

# Specify Port range
* `nmap -p 20-1024 10.0.1.10` scan between port range

# Scan IPv6
- `nmap -sU -6 ::1`
