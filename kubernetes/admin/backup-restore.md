# Save
`sudo ETCDCTL_API=3 etcdctl  --endpoints=https://[127.0.0.1]:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save snapshot.db`
- cacert, cert, key can be obtained from `etcd` ps command parameters, or the `etcd` pod

# Stop etcd
`sudo mv /etc/kubernetes/manifests/etcd.yaml ~/etcd.yaml`

# Restore
```
sudo ETCDCTL_API=3 etcdctl snapshot restore ./snapshot.db --data-dir /var/lib/etcd-new-dir
```

# Update etcd.yaml
update `hostPath.path` from `/var/lib/etcd` to `/var/lib/etcd-new-dir`

# Start etcd
`sudo mv ~/etcd.yaml /etc/kubernetes/manifests/etcd.yaml`
