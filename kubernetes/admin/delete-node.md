In case of failure, the node needs to be deleted, you can just as easily remove the node and replace it with a new one, joining it to the cluster.

```
kubectl drain worker1 --ignore-daemonsets --delete-local-data
kubectl get nodes
kubectl delete node worker1
kubeadm token create --print-join-command
```
