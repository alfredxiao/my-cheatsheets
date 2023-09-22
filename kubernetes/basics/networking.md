- Every Pod gets its own IP address.
- pods on a node can communicate with all pods on all nodes without NAT
- agents on a node (e.g. system daemons, kubelet) can communicate with all pods on that node

Note: For those platforms that support Pods running in the host network (e.g. Linux):
- pods in the host network of a node can communicate with all pods on all nodes without NAT

Pod IP Range
- if cluster initialised with `kubeadm`: `sudo kubeadm init --pod-network-cidr=10.168.0.0/16`
- or will be set the the CNI plugin
  - For weave-net, Check the weave pods logs using command `kubectl logs weave -n kube-system` and look for ipalloc-range
  - For calico, check the calico.yaml file and looking for `CALICO_IPV4POOL_CIDR`


Service IP Range
- from `kube-apiserver` command argument: `--service-cluster-ip-range=10.96.0.0/12`

Proxy mode
- how kube-proxy uses iptables or ipvs?
- `k -n kube-system logs kube-proxy-xxxxx` look for something like `"Unknown proxy mode, assuming iptables proxy" proxyMode=""` or `Using iptables Proxier`
