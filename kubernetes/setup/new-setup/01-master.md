
`sudo -i`

```
# Need to turn off all swap devices before initalizing the master.
sudo swapoff -a

# Update
apt-get update && apt-get upgrade -y

# Follow instructions to install docker.
apt-get install -y docker.io

cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

mkdir -p /etc/systemd/system/docker.service.d

# Restart docker.
systemctl daemon-reload
systemctl restart docker

# Add new repo for kubernetes.
sudo bash -c 'echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list'

# Add a GPG key for the package, and update with new repo.
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo apt-get update

# Install kubeadm, kubelet, and kubectl.
apt-get install -y kubeadm kubelet kubectl
# To install a specific version
# apt-get install -y kubeadm=1.22.11-00 kubelet=1.22.11-00 kubectl=1.22.11-00

sudo apt-mark hold kubelet kubeadm kubectl

# Initialize the master using the following command.
kubeadm init --pod-network-cidr=192.168.0.0/16 | tee kubeadm-init.out

# Install Calico CNI
# Follows intrustions to use Calico for Network Policy. Download calico.yaml file. Looking for the CALICO_IPv4POOL_CIDR. Note that the default is 192.168.0.0/16. If you are using a different pod CIDR, change it accordingly in the calico.yaml file.
# wget https://docs.projectcalico.org/manifests/calico.yaml
# cat calico.yaml | grep -a2 CIDR
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
# Before this CNI, the master node is NotReady, after a couple of minutes afer CNI installation, the master node becomes Ready

# Save the kubeadm join output to use when adding workers to the cluster.
# sudo kubeadm join 10.0.2.15:6443 --token m3jpro.pvufj1envk6mx3g5 --discovery-token-ca-cert-hash sha256:822885260222721f04296d96d490ddf9de568cd3507b735a1c21a5185755042e
```
