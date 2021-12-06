node network: your linux host network
pod network:
service network: not really a network

why a clusterIp or nodePort can forward and load balance requests to actual pod?
it is kube-proxy that relies on either iptables or ipvs, which rewrites packet destination IP at a node interface level, that makes this magic happen
