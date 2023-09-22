- append only
- compaction

The state of the cluster, networking, and other persistent information is kept in an etcd database, or, more accurately, a b+tree key-value store. Rather than finding and changing an entry, values are always appended to the end. Previous copies of the data are then marked for future removal by a compaction process. It works with curl and other HTTP libraries, and provides reliable watch queries.

Simultaneous requests to update a particular value all travel via the kube-apiserver, which then passes along the request to etcd in a series. The first request would update the database. The second request would no longer have the same version number as found in the object, in which case the kube-apiserver would reply with an error 409 to the requester. There is no logic past that response on the server side, meaning the client needs to expect this and act upon the denial to update.

There is a cp database along with possible followers. They communicate with each other on an ongoing basis to determine which will be master, and determine another in the event of failure. While very fast and potentially durable, there have been some hiccups with some features like whole cluster upgrades. The kubeadm cluster creation tool allows easy deployment of a multi-master cluster with stacked etcd or an external database cluster.


`curl -k  https://localhost:2379/version
curl: (35) error:14094412:SSL routines:ssl3_read_bytes:sslv3 alert bad certificate`

`sudo curl --cert /etc/kubernetes/pki/apiserver-etcd-client.crt --key /etc/kubernetes/pki/apiserver-etcd-client.key --cacert /etc/kubernetes/pki/etcd/ca.crt  https://localhost:2379/version`

{"etcdserver":"3.5.1","etcdcluster":"3.5.0"}

`sudo apt install etcd-client`
```
ETCDCTL_API=3 etcdctl --endpoints 127.0.0.1:2379  --cert=/etc/kubernetes/pki/etcd/server.crt  --key=/etc/kubernetes/pki/etcd/server.key   --cacert=/etc/kubernetes/pki/etcd/ca.crt  member list
```

`etcdctl get / --prefix --keys-only          # Get top-level keys`

# Backup & Restore
- backup
`sudo ETCDCTL_API=3 etcdctl --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/etcd-backup.db`
- restore (not yet used by etcd)
`sudo ETCDCTL_API=3 etcdctl --data-dir=/var/lib/from-backup snapshot restore  /opt/etcd-backup.db`

then `sudo vi /etc/kubernetes/manifests/etcd.yaml`
set `hostPath.path` from `/var/lib/etcd` as well as `--data-dir` argument to `/var/lib/from-backup`
