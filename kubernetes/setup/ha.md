stacked etcd
- there are multiple cp (control plane) nodes
- each node runs an etcd
- there is a load-balancer that fronts the multiple api-servers

external etcd
- there are multiple nodes dedicated for running etcd
- there are multiple separate nodes running control plane
- also there is a lobad balancer 
