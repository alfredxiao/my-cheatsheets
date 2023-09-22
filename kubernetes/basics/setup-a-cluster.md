# Setup all machines
1. Set hostname
- on Master: `sudo hostnamectl set-hostname k8s-control`
- on Worker Nodes: `sudo hostnamectl set-hostname k8s-worker1` and k8s-worker2
2. add to `/etc/hosts`
```
172.31.25.171  k8s-control
172.31.23.2    k8s-worker1
172.31.21.169  k8s-worker2
```
3. add kernel modules
```
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF
```
4. load them (rather than wait for restart)
```
sudo modprobe overlay
sudo modprobe br_netfilter
```
5. setup networking rules
```
cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables  = 1
net.ipv4.ip_forward                 = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF
```
6. Apply the without reboot
`sudo sysctl --system`
7. Install
`sudo apt-get update && sudo apt-get install -y containerd`
8. Configure containerd
- `sudo mkdir -p /etc/containerd`
- `sudo containerd config default | sudo tee /etc/containerd/config.toml`
- `sudo systemctl restart containerd`
9. Disable swap: `sudo swapoff -a`
10. install more packages
- `sudo apt-get install -y apt-transport-https curl`
11. add PGP key
- `curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -`
12. add source list
```
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
```
then `sudo apt-get update`
13. Install kube packages
`sudo apt-get install -y kubelet=1.23.0-00 kubeadm=1.23.0-00 kubectl=1.23.0-00`
`sudo apt-mark hold kubelet kubeadm kubectl`

# Init the K8s cluster (mostly on controller)
1. `sudo kubeadm init --pod-network-cidr 192.168.0.0/16 --kubernetes-version 1.23.0`
2. init config file
```
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```
3. install CNI (non ROOT!)
`kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml`
4. Print join command
- `kubeadm token create --print-join-command`
5. Join (from worker nodes)
- `sudo {COPY_OUTPUT_FROM_PREVIOUS_STEP}`
