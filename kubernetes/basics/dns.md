Kubelet sets up file `/etc/resolv.conf` for each pod
`search <namespace>.svc.cluster.local svc.cluster.local cluster.local` is used to expand DNS queries
 - search name `test` means search for `test.svc.cluster.local` and `test.svc.cluster.local` and `test.cluster.local`

 What objects get DNS records?
- Services
- Pods

Types of DNS Records
- A: IPv4 ip for hostname
- AAAA: IPv6 ip for hostname
- SRV: defining the location (i.e. the hostname and port number) of servers for specified services

# Pods
## A/AAAA records
- In general a pod has the following DNS resolution:
`pod-ip-address.my-namespace.pod.cluster-domain-name` where cluster-domain-name defaults to `cluster.local`.
e.g. `172-17-0-3.default.pod.cluster.local`

Any pods exposed by a Service (regular and headless) have the following DNS resolution available:
`pod-ip-address.service-name.my-namespace.svc.cluster-domain-name`
e.g. `192-168-126-44.pod-headless-service.default.svc.cluster.local`
- The Pod spec also has an optional `subdomain` field which can be used to specify its subdomain. For example, a Pod with `hostname` set to "foo", and `subdomain` set to "bar", in namespace "my-namespace", will have the fully qualified domain name (FQDN) `"foo.bar.my-namespace.svc.cluster-domain.example"`.

# Service
## A/AAAA records
- "Normal" (not headless) Services are assigned a DNS A or AAAA record, for a name of the form
  - `my-svc.my-namespace.svc.cluster-domain.example`. This resolves to the cluster IP of the Service.
  - e.g. `pod-regular-service.default.svc.cluster.local`

- "Headless" (without a cluster IP) Services are also assigned a DNS A or AAAA record, for a name of the form
  - `my-svc.my-namespace.svc.cluster-domain.example`. Unlike normal Services, this resolves to the set of IPs of the pods selected by the Service.
  - e.g. `pod-headless-service.default.svc.cluster.local`
  - Clients are expected to consume the set or else use standard round-robin selection from the set.

## SRV records
For each named port, the SRV record would have the form
`_my-port-name._my-port-protocol.my-svc.my-namespace.svc.cluster-domain.example`.
- `dig _foo._udp.pod-regular-service.default.svc.cluster.local SRV`
- For a regular service, this resolves to the port number and the domain name:
  `my-svc.my-namespace.svc.cluster-domain.example`.
- For a headless service, this resolves to multiple answers, one for each pod that is backing the service, and contains the port number and the domain name of the pod of the form
  - `auto-generated-name.my-svc.my-namespace.svc.cluster-domain.example`

## Governing Service
- if a statefulset's `spec.serviceName` is set to be same name as another headless service's name
  - then A/AAAA records would be created for each individual pod (of the statefulset)
    - e.g. `statefulset-name-*.headless-service-name.default.svc.cluster.local`

## headless service having same name as subdomain
- If there exists a headless service in the same namespace as the pod and with the **same name** as the `spec.subdomain` of the pod, the cluster's DNS Server also returns an A or AAAA record for the Pod's fully qualified hostname.
 - For example, given a Pod with the hostname set to `busybox-1` and the `spec.subdomain` set to `default-subdomain`, and a headless Service named `default-subdomain` in the same namespace, the pod will see its own FQDN as `busybox-1.default-subdomain.my-namespace.svc.cluster-domain.example`.
 - DNS serves an A or AAAA record at that name, pointing to the Pod's IP. Both pods `busybox1` and `busybox2` can have their distinct A or AAAA records.
