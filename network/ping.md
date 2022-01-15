* `ping hostname/IP`
* `ping -c 5 hostname/IP` with specified counts
* `ping -s 100 hostname/IP` --> alter size of packets
* `ping -i 3 hostname/IP`  --> alter interval for sending packets
* `ping -q hostname/IP` --> quite - only shows summary
* `ping -w 5 hostname/IP` --> set a timeout of when to stop
* `ping -f hostname/IP` --> flood ping, send packets asap
* `ping -p ff hostname/IP` -> fill a packet with data, `ff` fills with 1
* `ping -b hostname/IP` --> send packets to a broadcast address
* `ping -t hostname/IP` --> limit the number of hops
* `ping -v hostname/IP` --> increase verbosity
* `ping -M do -s 1490 example.com` to now allow fragmentation and test maximum MTU in the path
