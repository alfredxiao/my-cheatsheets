# Basics
- Keep in mind that the **order** of your rules matter.
- Use the `-A` option to **append** the new rule to the end of a chain.
- If you want to put it somewhere else in the chain, you can use the `-I` option which allows you to specify the position of the new rule (or place it at the beginning of the chain by not specifying a rule number).
- iptables rules are ephemeral, which means they need to be manually saved for them to persist after a reboot.

# List Chains
- `iptables -S` exmple output:
```
-P INPUT ACCEPT
-P FORWARD DROP
-P OUTPUT ACCEPT
```
where there are commands to set the default policy for these chains, default policy can be either `ACCEPT`, `DROP` (silently ignore), or `REJECT` (disallow and send back an error)
- or `iptables -L | grep policy`

# Set default policy for a chain
- `iptables --policy INPUT ACCEPT`
- or `iptables -P INPUT ACCEPT`

# List Rules
- `iptables -L`
- `iptables -L INPUT` rules for a chain

# Accept incoming connection from a single IP
- `iptables -A INPUT -s 10.10.10.10 -j ACCEPT`
- `-A` append
- `-s` source
- `-j` action, can be `ACCEPT`, `DROP` or `REJECT`

# Drop incoming connection from IP range
- `iptables -A INPUT -s 10.10.10.0/24 -j DROP`

# Reject incoming connection from IP at an interface
- `iptables -A INPUT -i eth0 -s 10.10.10.10 -j DROP`

# Drop outgoing connection for an source IP on destination port
- `iptables -A OUTPUT -p tcp --dport ssh -s 10.10.10.10 -j REJECT`
- `-p` protocol, or `--protocol`
- `--dport` destination port

# Allow outgoing to an IP only
- `iptables -A OUTPUT -d 192.168.1.1 -j ACCEPT`
- `-d` destination IP

# Delete a rule
- `iptables -D INPUT -s 127.0.0.1 -p tcp -dport 111 -j ACCEPT `
- `-D` delete as in contrast to `-A` and `-I`
- `iptables -D INPUT 4` delete by position number

# Flush all rules
- `iptables --flush`

# Clear all rules
- `iptables -F`

# Save rules across reboots
- Ubuntu: sudo `/sbin/iptables-save`
- RedHat / Centos: `/sbin/service iptables save`
- Others: `/etc/init.d/iptables save`

# For IPv6
- use `ip6tables`

# Specifying interface
- `sudo iptables -A INPUT -i lo -j ACCEPT`
  - `-i`, or `--in-interface` for in-interface, via which interface the packet incomes
- `sudo iptables -A OUTPUT -o lo -j ACCEPT`
  - `-o`, or `--out-interface` for out-interface, via which interface the packet outgoes

# Allow return traffic (the stateful part)
- `iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT`
- `iptables -A OUTPUT -m conntrack --ctstate ESTABLISHED -j ACCEPT`

# Allow forward
- `iptables -A FORWARD -i eth1 -o eth0 -j ACCEPT`
  - allow forward packets from eth1 forwarded to eth0

# Allowing SSH from anywhere
- `sudo iptables -A INPUT -p tcp --dport 22 -m conntrack --ctstate NEW,ESTABLISHED -j ACCEPT`
- `sudo iptables -A OUTPUT -p tcp --sport 22 -m conntrack --ctstate ESTABLISHED -j ACCEP`
- The second command, which allows the outgoing traffic of established SSH connections, is only necessary if the OUTPUT policy is not set to ACCEPT.

# Allowing SSH from specific IP
- `sudo iptables -A INPUT -p tcp -s 203.0.113.0/24 --dport 22 -m conntrack --ctstate NEW,ESTABLISHED -j ACCEPT`
- `sudo iptables -A OUTPUT -p tcp --sport 22 -m conntrack --ctstate ESTABLISHED -j ACCEPT`
or
- `iptables -A INPUT -p tcp --dport ssh -s 10.10.10.10 -m state --state NEW,ESTABLISHED -j ACCEPT`
- `iptables -A OUTPUT -p tcp --sport 22 -d 10.10.10.10 -m state --state ESTABLISHED -j ACCEPT`

# Allowing outgoing SSH
- `sudo iptables -A OUTPUT -p tcp --dport 22 -m conntrack --ctstate NEW,ESTABLISHED -j ACCEPT`
- `sudo iptables -A INPUT -p tcp --sport 22 -m conntrack --ctstate ESTABLISHED -j ACCEPT`

# DROP vs REJECT
- REJECT: ping -> Destination Unreachable
- DROP: ping -> Timed out!
