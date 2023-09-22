- Every node runs a service called the kubelet that registers it with the cluster and communicates with the API server.
- It watches the API server for new work tasks and maintains a reporting channel.

# Config files
## Context with control plane
- `/etc/kubernetes/kubelet.conf`, with api server endpoint, the local keys, etc.
## Config for local Node
- `/var/lib/kubelet/config.yaml`
  - e.g. `staticPodPath: /etc/kubernetes/manifests`
- `/var/lib/kubelet/kubeadm-flags.env`
  - `KUBELET_KUBEADM_ARGS="--flag1=value1 --flag2=value2 ..."` are flags for when kubelet starts
