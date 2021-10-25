* `nmap -F 10.0.1.10`
* `nmap -A -F 10.0.1.10` -> discover host OS info and service info
* `nmap -sn 10.0.1.10` ping scan
* `nmap -sS 10.0.1.10` stealthy scan (sends only SYN)
* `nmap -sU -F 10.0.1.10` UDP scan
* `nmap -p 20-1024 10.0.1.10` scan between port range
