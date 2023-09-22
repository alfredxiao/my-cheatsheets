# Controller/Master
-  refers to the machine that runs control plane

# Control Plane
- refers to multiple components as below
- these components can run on single server, multiple servers
## api-server
- All communication, between all components, must go through the API server.
## etcd
- highly available state store
- is the only stateful part of the control plane
## scheduler
- select available node to run containers
## controller-manager
- runs multiple utilities in a single process
- carry out a variety of automation related tasks
- it spawns all independent controllers and monitors them
- those controllers monitor cluster components and respond to events
## cloud-controller-manager
- interfacing with different cloud platforms

# Worker Node
## kubelet
- agent that interacts with control Plane
## container runtime
- not part of K8s itself
- e.g. Docker, containerd
## kube-proxy
- network proxy

# Container Standards
## Container Runtime
- both containerd and CRI-O implements K8s CRI
- containerd is from Docker
- CRI-O is from Redhat
## OCI (Open Container Initiative)
- provides specifications for images and containers
- provides `runc`, a tool for creating and runing container processes
## Kubernetes CRI (Container Runtime Interface)
- defines an API between K8s and container runtime
## Kubernetes CSI (Container Storage Interface)
## Kubernetes CNI (Container Network Interface)


# High Availability
- When using multiple control planes for HA, you will likely need to communicate with the K8s API through a load balancer
- Two patterns for etcd
  - stacked etcd: each control plane has its own etcd
  - external etcd:
