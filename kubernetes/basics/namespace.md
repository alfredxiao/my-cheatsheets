- `kubectl get namespace` or `get namespaces`
- `kubectl create namespace myspace`

- Namespaces partition a Kubernetes cluster and are designed as an easy way to apply `quotas` and `policies` to groups of objects.
- They’re **not** designed for strong workload isolation.
- If you need strong workload isolation, the current method is to use multiple clusters.

- It’s important to understand that most Kubernetes objects are deployed to a Namespace.
These objects are said to be namespaced and include common objects such as `Pods`, `Services` and `Deployments`.
- Other objects exist outside of Namespaces and include `nodes` and `PodSecurityPolicies`.

# namespaces
- The `default` Namespace is where newly created objects go unless you explicitly specify otherwise.
- `kube-system` is where DNS, the metrics server, and other control plane components run.
- `kube-public` is for objects that need to be readable by anyone. And last but not least,
- `kube-node-lease` is used for node heartbeat and managing node leases.

# filter by namespace
- `kubectl get pods --namespace myspace` or `get pods -n myspace`
- `kubectl get pods --all-namespaces` or `-A`

# Use a namespace
- `kubectl config set-context --current --namespace <namespace>`
- this avoids adding `-n` all the time

# Describe
- `kubectl describe ns <namespace>`

# List
- `kubectl get alll -n <namespace>`

# Namespaced resources
- not all resources are namespaced
## In a namespace
`kubectl api-resources --namespaced=true`

## Not in a namespace
`kubectl api-resources --namespaced=false`
e.g. PV, storageclasses, ingressclasses
