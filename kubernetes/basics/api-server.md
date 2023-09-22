The kube-apiserver is central to the operation of the Kubernetes cluster.

All calls, both internal and external traffic, are handled via this agent.
All actions are accepted and validated by this agent, and it is the **only agent** which connects to the etcd database. As a result, it acts as a master process for the entire cluster, and acts as a frontend of the cluster's shared state. Each API call goes through three steps: authentication, authorization, and several admission controllers.
