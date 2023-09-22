# Prerequisites
## Unique hostname, MAC address, and product_uuid for every node
### hostname
`hostname`
### MAC address
`ip link`/`ifconfig -a`
### product_uuid
`sudo cat /sys/class/dmi/id/product_uuid`
## Port Numbers are available
### Control Plane
- api-server: 6443
- etcd: 2379-2380
- kubelet: 10248,10250
- kube-scheduler: 1025x
- kube-controller-manager: 1025x
### Worder Nodes
- kubelet: 10250
- node ports: 30000-32767

# Swap Off
- Need to turn off all swap devices before initalizing the master.
`sudo swapoff -a`

# Container Runtime Prerequisites
```
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# sysctl params required by setup, params persist across reboots
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

# Apply sysctl params without reboot
sudo sysctl --system
```

# Install containerd
## Update Repository
```
sudo apt-get update && sudo apt-get install -y \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg \
  lsb-release
```
### Add Repository
```
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
```
## Install
`sudo apt-get install containerd.io`

## Configure and Start
```
sudo mkdir -p /etc/containerd
containerd config default | sudo tee /etc/containerd/config.toml

sudo systemctl restart containerd
## Using systemd cgroup driver

sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/g' /etc/containerd/config.toml

sudo systemctl restart containerd
```

# Install Kubernetes packages
## Add Repository for Kubernetes packages
```
sudo bash -c 'echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list'

# Add a GPG key for the package, and update with new repo.
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo apt-get update
```
## Install
```
# Install kubeadm, kubelet, and kubectl.
sudo apt-get install -y kubeadm kubelet kubectl
# To install a specific version
# sudo apt-get install -y kubeadm=1.24.0-00 kubelet=1.24.0-00 kubectl=1.24.0-00

sudo apt-mark hold kubelet kubeadm kubectl
```

# Initialize Control Plane
```
# Initialize the master using the following command.
sudo kubeadm init --pod-network-cidr=192.168.0.0/16 | tee kubeadm-init.out
```

# Node Join
`sudo kubeadm join 172.31.40.44:6443 --token i7z9dp.f30g0rt9pd3eeo7t --discovery-token-ca-cert-hash sha256:3af324505583b8cc6a1cf02a51bba2e4c227f02aa51577445deefa15dabe4000`

# Status
- At this point, `kubectl get nodes` shows they are not ready
- Also, `coredns-*` pods are Pending, because network plugin not yet run
- But `kube-proxy` pods are running already

# Install CNI (need not to be run as root) - on master node
`kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml`
