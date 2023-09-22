projectId:
projectNumber:

region: australia-southeast1
zone: australia-southeast1-a
PUBLIC_ID (controller): 35.197.179.105


# Create VMs
gcloud compute networks create kubernetes-cluster --subnet-mode custom

gcloud compute networks subnets create kubernetes  --network kubernetes-cluster  --range 10.240.0.0/24

gcloud compute firewall-rules create kubernetes-cluster-allow-internal --allow tcp,udp,icmp  --network kubernetes-cluster --source-ranges 10.240.0.0/24,10.244.0.0/16

gcloud compute firewall-rules create kubernetes-cluster-allow-external --allow tcp:22,tcp:6443,icmp --network kubernetes-cluster  --source-ranges 0.0.0.0/0

gcloud compute addresses create kubernetes-controller --region $(gcloud config get-value compute/region)

PUBLIC_IP=$(gcloud compute addresses describe kubernetes-controller --region $(gcloud config get-value compute/region) --format 'value(address)')

gcloud compute instances create controller --async --boot-disk-size 200GB --can-ip-forward --image-family ubuntu-2004-lts --image-project ubuntu-os-cloud --machine-type n1-standard-1 --private-network-ip 10.240.0.10 --scopes compute-rw,storage-ro,service-management,service-control,logging-write,monitoring --subnet kubernetes --address $PUBLIC_IP

for i in 0 1; do \
gcloud compute instances create worker-${i} \
  --async \
  --boot-disk-size 200GB \
  --can-ip-forward \
  --image-family ubuntu-2004-lts \
  --image-project ubuntu-os-cloud \
  --machine-type n1-standard-1 \
  --private-network-ip 10.240.0.2${i} \
  --scopes compute-rw,storage-ro,service-management,service-control,logging-write,monitoring \
  --subnet kubernetes; \
done

# Install Docker (for controller and nodes)
gcloud compute ssh controller (worker-0, worker-1)

sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

apt-cache madison docker-ce

sudo apt-get update && sudo apt-get install -y docker-ce=5:20.10.14~3-0~ubuntu-focal  docker-ce-cli=5:20.10.14~3-0~ubuntu-focal

sudo apt-mark hold containerd.io docker-ce docker-ce-cli

cat <<EOF | sudo tee /etc/docker/daemon.json {
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"  
}
EOF


sudo mkdir -p /etc/systemd/system/docker.service.d

sudo systemctl daemon-reload
sudo systemctl restart docker
sudo systemctl enable docker

# Install cmd tools
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -

cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF

sudo apt-get update

apt-cache madison kubeadm

sudo apt-get install -y kubelet=1.23.5-00 kubeadm=1.23.5-00 kubectl=1.23.5-00

sudo apt-mark hold kubelet kubeadm kubectl

# Initialize Control Plane
(Run on control plane)
gcloud config set compute/zone australia-southeast1-a

KUBERNETES_PUBLIC_ADDRESS=$(gcloud compute instances describe controller \
 --zone $(gcloud config get-value compute/zone) \
 --format='get(networkInterfaces[0].accessConfigs[0].natIP)')

sudo kubeadm init --pod-network-cidr=10.244.0.0/16  --ignore-preflight-errors=NumCPU --apiserver-cert-extra-sans=$KUBERNETES_PUBLIC_ADDRESS

-- which generates
```
kubeadm join 10.240.0.10:6443 --token 9u4s86.ra0b09g6miz1aacm \
	--discovery-token-ca-cert-hash sha256:d8774b677dbcdcbd1f9e6269e23e48c2f718011296d7ed5d90c197714b23ea39
```

kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml

# Join Nodes with Cluster
Run above `kubeadm join` command

# Config kubectl on local dev
`gcloud compute scp controller:~/.kube/config kubeconfig`

KUBERNETES_PUBLIC_ADDRESS=$(gcloud compute instances describe controller \
--zone $(gcloud config get-value compute/zone) \
--format='get(networkInterfaces[0].accessConfigs[0].natIP)')

sed -i "s/10.240.0.10/$KUBERNETES_PUBLIC_ADDRESS/" kubeconfig


# Volume (nfs)
gcloud filestore instances create nfs-server \
  --project=$(gcloud config get-value project) \
  --zone=$(gcloud config get-value compute/zone) \
  --tier=STANDARD \
  --file-share=name="vol1",capacity=1TB \
  --network=name="kubernetes-cluster"

  gcloud filestore instances describe nfs-server \
  --project=$(gcloud config get-value project) \
  --zone=$(gcloud config get-value compute/zone)

## install driver on nodes
gcloud compute ssh worker-0
sudo apt-get -y update
sudo apt-get -y install nfs-common
sudo umount /mnt/nfs
sudo mkdir /mnt/nfs
sudo mount 10.158.99.26:/vol1 /mnt/nfs  # IP from above filestore describe output
ls /mnt/nfs
sudo umount /mnt/nfs

# Define PV
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-nfs
spec:
  # capabilities
  accessModes:
  - ReadWriteOnce
  - ReadOnlyMany
  - ReadWriteMany
  capacity:
    storage: 1Ti
  volumeMode: Filesystem
  # implementation
  nfs:
    path: /vol1
    server: 10.158.99.26
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pv-1
spec:
  storageClassName: ""
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 100Gi
    limits:
      storage: 200Gi
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app: pg
  name: pg
spec:
  serviceName: pg
  selector:
    matchLabels:
      app: pg
  template:
    metadata:
      labels:
app: pg spec:
      containers:
      - image: postgres
        name: pg
        env:
        - name: PGDATA
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app: pg
  name: pg
spec:
  serviceName: pg
  selector:
    matchLabels:
      app: pg
  template:
    metadata:
      labels:
        app: pg
    spec:
      containers:
      - image: postgres
        name: pg
        env:
        - name: PGDATA
          value: /var/lib/postgresql/data/db-files
        - name: POSTGRES_PASSWORD
          value: password001
        - name: POSTGRES_HOST_AUTH_METHOD
          value: trust        volumeMounts:
        - name: data
         mountPath: /var/lib/postgresql/data
        volumes:
        - name: data
          persistentVolumeClaim:
            claimName: pv-1
            readOnly: false
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app: box
  name: box
spec:
  serviceName: box
  selector:
    matchLabels:
      app: box
  template:
    metadata:
      labels:
        app: box
    spec:
      containers:
      - image: busybox
        name: box
        volumeMounts:
        - name: data
          mountPath: /mnt
        command: ['sh','-c','sleep 3000']
      volumes:
      - name: data
        persistentVolumeClaim:
          claimName: pv-1
          readOnly: false
```

## Delete
gcloud filestore instances delete nfs-server \ --project=$(gcloud config get-value project) \ --zone=$(gcloud config get-value compute/zone)

# Auto provisioned PV
gcloud container clusters create "kluster" \
  --project $(gcloud config get-value project) \
  --zone $(gcloud config get-value compute/zone) \
  --cluster-version "1.21.5-gke.1805" \
  --machine-type "g1-small" \
  --image-type "COS" \
  --disk-type "pd-standard" \
  --disk-size "30" \
  --num-nodes "1"

gcloud container clusters get-credentials kluster \
  --zone $(gcloud config get-value compute/zone) \
  --project $(gcloud config get-value project)
```
# auto-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: disk-rwo-10g
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```
## Remove the cluster
gcloud container clusters delete "kluster" \ 
  --project $(gcloud config get-value project) \
  --zone $(gcloud config get-value compute/zone)
