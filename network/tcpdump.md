* `-D` --> list interfaces
* `-i any -c 5` --> on all(any) interfaces, capture 5 packets
* `-i any src 10.0.3.10` --> all traffic from a source host
* `-i eth0 -c 5 tcp -w out.pcap` --> capture tcp and write to a file
* `-r out.pcap` read it
