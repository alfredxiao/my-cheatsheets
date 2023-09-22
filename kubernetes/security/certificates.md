# CAs
## general purpose
- `/etc/kubernetes/pki/ca.crt` has same content on control plane node and worker nodes
  - `Issuer: CN = kubernetes, Subject: CN = kubernetes`
- on worker nodes: `/var/lib/kubelet/config.yaml` has `clientCAFile: /etc/kubernetes/pki/ca.crt`
## etcd
- `/etc/kubernetes/pki/etcd/ca.crt`
- `Issuer: CN = etcd-ca, Subject: CN = etcd-ca`
## front-proxy
- `/etc/kubernetes/pki/front-proxy-ca.crt`
- `Issuer: CN = front-proxy-ca, Subject: CN = front-proxy-ca`


# kube-controller-manager
- has its own identity defined in `/etc/kubernetes/controller-manager.conf`?
## Command line
--allocate-node-cidrs=true
--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf
--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf
--bind-address=127.0.0.1
`--client-ca-file`=`/etc/kubernetes/pki/ca.crt`
--cluster-cidr=10.168.0.0/16
--cluster-name=kubernetes
`--cluster-signing-cert-file`=`/etc/kubernetes/pki/ca.crt`
`--cluster-signing-key-file`=`/etc/kubernetes/pki/ca.key`
--controllers=*,bootstrapsigner,tokencleaner
--kubeconfig=`/etc/kubernetes/controller-manager.conf`
--leader-elect=true
--requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt
`--root-ca-file`=`/etc/kubernetes/pki/ca.crt`
`--service-account-private-key-file`=`/etc/kubernetes/pki/sa.key`
--service-cluster-ip-range=10.96.0.0/12 -
-use-service-account-credentials=true

# scheduler
- has its own identity defined in `/etc/kubernetes/scheduler.conf`

# apiserver
## https identity
- apiserver presents to its clients (e.g. kubectl, kubelet)
- `/etc/kubernetes/pki/apiserver.crt` via `--tls-cert-file` argument
- `Issuer: CN = kubernetes, Subject: CN = kube-apiserver`
## kube-apiserver-etcd-client
- apiserver presents to etcd
- `/etc/kubernetes/pki/apiserver-etcd-client.crt`
- `Issuer: CN = etcd-ca, Subject: O = system:masters, CN = kube-apiserver-etcd-client`
## kube-apiserver-kubelet-client
- `/etc/kubernetes/pki/apiserver-kubelet-client.crt`
- `Issuer: CN = kubernetes; Subject: O = system:masters, CN = kube-apiserver-kubelet-client`
- apiserver presents to kubelet
## Command line
kube-apiserver
  --advertise-address=172.31.6.93
  --allow-privileged=true
  --authorization-mode=Node,RBAC
  `--client-ca-file`=`/etc/kubernetes/pki/ca.crt`
  --enable-admission-plugins=NodeRestriction
  --enable-bootstrap-token-auth=true
  --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt
  `--etcd-certfile`=`/etc/kubernetes/pki/apiserver-etcd-client.crt`
  --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key
  --etcd-servers=https://127.0.0.1:2379
  `--kubelet-client-certificate`=`/etc/kubernetes/pki/apiserver-kubelet-client.crt`
  --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key
  --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
  --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt
  --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key
  --requestheader-allowed-names=front-proxy-client
  --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt
  --requestheader-extra-headers-prefix=X-Remote-Extra-
  --requestheader-group-headers=X-Remote-Group
  --requestheader-username-headers=X-Remote-User
  --secure-port=6443
  --service-account-issuer=https://kubernetes.default.svc.cluster.local
  `--service-account-key-file`=`/etc/kubernetes/pki/sa.pub`
  `--service-account-signing-key-file`=`/etc/kubernetes/pki/sa.key`
  --service-cluster-ip-range=10.96.0.0/12
  `--tls-cert-file`=`/etc/kubernetes/pki/apiserver.crt`
  --tls-private-key-file=/etc/kubernetes/pki/apiserver.key

# administrator
- inside `/etc/kubernetes/admin.conf` defines its identity
- `certificate-authority-data` is ?

# etcd
## Server Identity
- `/etc/kubernetes/pki/etcd/server.crt`
- `Issuer: CN = etcd-ca, Subject: CN = 1bfddccd781c.mylabserver.com`
- When running etcd cluster, etcd instances share the same Server Identity, but the identity has multiple IP/hostnames added to its SAN
## Peer Identity
- `/etc/kubernetes/pki/etcd/peer.crt`
- `Issuer: CN = etcd-ca; Subject: CN = 1bfddccd781c.mylabserver.com`
- Each etcd instance has its own peer identity
## Command line
--advertise-client-urls=https://172.31.6.93:2379
--cert-file=`/etc/kubernetes/pki/etcd/server.crt`
--client-cert-auth=true
--data-dir=/var/lib/etcd
--experimental-initial-corrupt-check=true
--initial-advertise-peer-urls=https://172.31.6.93:2380
--initial-cluster=1bfddccd781c.mylabserver.com=https://172.31.6.93:2380
--key-file=/etc/kubernetes/pki/etcd/server.key
--listen-client-urls=https://127.0.0.1:2379,https://172.31.6.93:2379
--listen-metrics-urls=http://127.0.0.1:2381
--listen-peer-urls=https://172.31.6.93:2380
--name=1bfddccd781c.mylabserver.com
--peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt
--peer-client-cert-auth=true
--peer-key-file=/etc/kubernetes/pki/etcd/peer.key
--peer-trusted-ca-file=`/etc/kubernetes/pki/etcd/ca.crt`
--snapshot-count=10000
--trusted-ca-file=`/etc/kubernetes/pki/etcd/ca.crt`

# etcdctl
- use etcd/ca.crt as cacert
- use etcd/healthcheck-client.crt as cert and etcd/healthcheck-client.key as key

- Kubernetes sets the user to be the Common Name field and the group to be the Organization field)
- serviceaccount private keys (which aren’t signed by a certificate authority)

# kubelet (each kubelet has its own identity)
- from `/etc/kubernetes/kubelet.conf`
- client-certificate: `/var/lib/kubelet/pki/kubelet-client-current.pem`
  - `Issuer: CN = kubernetes; Subject: O = system:nodes, CN = system:node:1bfddccd782c.mylabserver.com`
- client-key: `/var/lib/kubelet/pki/kubelet-client-current.pem`
- `certificate-authority-data` is the same as that in `/etc/kubernetes/admin.conf` on control plane node
## Command line
/usr/bin/kubelet
  --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf
  --kubeconfig=/etc/kubernetes/kubelet.conf
  --config=/var/lib/kubelet/config.yaml
  --container-runtime=remote
  --container-runtime-endpoint=unix:///var/run/containerd/containerd.sock
  --pod-infra-container-image=k8s.gcr.io/pause:3.7

# Commands
- `kubeadm certs check-expiration`

# References
- https://kubernetes.io/docs/setup/best-practices/certificates/
- https://jvns.ca/blog/2017/08/05/how-kubernetes-certificates-work/
