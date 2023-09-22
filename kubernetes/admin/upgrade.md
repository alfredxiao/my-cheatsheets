# First, upgrade primary control plane node
1. Drain the control plane node
- `kubectl drain cp-node-name --ignore-daemonsets`
2. Upgrade `kubeadm` on the control plane node
- `sudo apt-get update && sudo apt-get install -y --allow-change-held-packages kubeadm=1.23.8-00`
- `kubeadm version` to verify it is installed

**for primary control plane node**
3. Plan the upgrade
- `sudo kubeadm upgrade plan v1.23.8`
4. Apply the upgrade
- `sudo kubeadm upgrade apply v1.23.8`
**for other control plane nodes**
- `sudo kubeadm upgrade node`

- to verify version `kubectl version` where Server Version should be the upgraded-to version
5. Upgrade `kubelet` and `kubectl` on the control plane node
- `sudo apt-get update && sudo apt-get install -y --allow-change-held-packages kubectl=1.23.8-00 kubelet=1.23.8-00`
- `sudo systemctl daemon-reload`
- `sudo systemctl restart kubelet`
6. Uncordon the control plane node
- `kubectl uncordon cp-node-name`

# Second, upgrade any additional control nodes

# Third, upgrade worker nodes
1. Drain the node
- `kubectl drain worker-node-name`
2. Upgrade kubeadm
- `sudo apt-get update && sudo apt-get install -y --allow-change-held-packages kubeadm=1.23.8-00`
- `kubeadm version` to verify it is installed
3. Upgrade the kubelet configuration
- `sudo kubeadm upgrade node`
4. Upgrade kubelet and kubectl
- `sudo apt-get update && sudo apt-get install -y --allow-change-held-packages kubelet=1.23.8-00 kubectl=1.23.8-00`
- `sudo systemctl daemon-reload`
- `sudo systemctl restart kubelet`
5. Uncordon the node
- `kubectl uncordon worker-node-name`
