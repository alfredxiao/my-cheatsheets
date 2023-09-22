command how kubelet runs `/usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/var/lib/kubelet/config.yaml ...`

`/var/lib/kubelet/config.yaml` contains `staticPodPath: /etc/kubernetes/manifests`
