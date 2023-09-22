A few important things to note about service discovery in Kubernetes:
1. Kubernetes uses its internal DNS as a service registry
2. All Kubernetes Service sautomatically register their details with DNS

This command lists the Things running the cluster DNS.
- `kubectl get pods -n kube-system -l k8s-app=kube-dns`
- `kubectl get deploy -n kube-system -l k8s-app=kube-dns`
- `kubectl get svc -n kube-system -l k8s-app=kube-dns`

# How service name is registered (in CoreDNS)
- CoreDNS runs on Kubernetes and implements a controller that watches the API server for new Service objects.
- Any time it observes one, it automatically creates the DNS records mapping the Service name to its ClusterIP.
- This means apps, and even Services, don’t need to perform their own service registration – the cluster DNS does it for them.

# Service Name vs DNS name
- It’s important to understand that the name registered in DNS for the Service is the value stored in its `metadata.name` property. This is why it’s important that Service names are valid DNS names and don’t include exotic characters.

# Routing to Pods
- The kubelet agent on every node is watching the API server for new Endpoints/EndpointSlice objects. When it sees one, it creates local networking rules to redirect ClusterIP traffic to Pod IPs.
- In modern Linux-based Kubernetes clusters, the technology used to create these rules is the Linux IP Virtual Server (IPVS). Older versions used iptables.

# Pods
- configured with `/etc/resolv.conf` to talk to Core DNS, where the `nameserver` entry in this conf file is set to the ClusterIP of `service/kube-dns`

# kube-proxy
- kube-proxy is a Pod-based Kubernetes-native app that implements a controller watching the API server for new Services and Endpoints objects. When it sees them, it creates local IPVS rules telling the node to intercept traffic destined for the Service’s ClusterIP and forward it to individual Pod IPs.
- This means that every time a node’s kernel processes traffic headed for an address on the service network, a trap occurs, and the traffic is redirected to the IP of a healthy Pod matching the Service’s label selector.

# Service Traffic Flow
1. Query DNS for Service Name
2. Receive ClusterIP
3. Send traffic to ClusterIP
4. No Route. So send to containers' default gateway
5. Forward to Node
6. No route. Send to Node's default gateway
7. Processed by Node's kernel
8. Trap (IPVS rule)
9. Rewrite IP destination field to Pod IP

# FQDN
- The format is `<object-name>.<namespace>.svc.cluster.local`, where the domain name for the cluster is usually `cluster.local`
  - e.g. For example, a Service called `ent` will have a fully qualified domain name (FQDN)  of `ent.default.svc.cluster.local`
- Objects can connect to Services in the local Namespace using short names
- Connect to remote namespaces requires FQDN
