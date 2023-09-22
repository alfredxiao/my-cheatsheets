# CPU Resource units
- 1 CPU unit = 1 Physical CPU Core (or 1 Virtual Core)
- `spec.containers[].resources.requests.cpu` it is allowed to use float number
- `0.1` = `100m` (where m means milli, as 1=1000 millis)
- so this is always in absolute amount, no matter how many cores the actual machine has

# Memory Resource units
- can use an integer number, e.g. 1293938223
- or use `e` notation, e.g. `123e5`
- or use `k,M,G,T,P,E` notations, where `k`=1000, `M`=1000k, etc.
- or use `Ki,Mi,Gi,Ti,Pi,Ei` notations, where `Ki`=1024, `Mi`=1024Ki
- **Be careful** of casing, `m` means milli, so `500m` means 0.5 byte

# Imposing resource quota (on containers)
```
apiVersion: v1
kind: Pod
metadata:
  name: frontend
spec:
  containers:
  - name: app
    image: images.my-company.example/app:v4
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
```

# Pod scheduling
- Each node has its maximum capacity
- Even if actual CPU/Memory usage is low on a node, if adding a new Pod to it would exceed its capacity, the Pod won't be placed into this node
- If that process asking for more memory than the limit specified, Linux kernel OOM subsystem would stop the process. If it is the container's PID 1, and the container is marked as restartable, Kubernetes restarts the container.
- The memory limit for the Pod or container can also apply to pages in memory backed volumes, such as an `emptyDir`. The `kubelet` tracks `tmpfs` `emptyDir` volumes as container memory use, rather than as local ephemeral storage.
- If a container exceeds its memory request and the node that it runs on becomes short of memory overall, it is likely that the Pod the container belongs to will be evicted.
- A container might or might not be allowed to exceed its CPU limit for extended periods of time. However, container runtimes **don't** terminate Pods or containers for excessive CPU usage.

# Namespace resource quota
- You can configure resource quotas to limit the total amount of resources that a namespace can consume. Kubernetes enforces quotas for objects in particular namespace when there is a ResourceQuota in that namespace
- Once you set a quota for a namesapce, you **must** specify request or limit for your pod containers in this namespace
```
apiVersion: v1
kind: ResourceQuota
metadata:
  name: awesome-quota
  nmespace: ns-quota
spec:
  hard:
    pods: 2
    requests.cpu: "1"
    requests.memory: 1024m
    limits.cpu: "4"
    limits.memory: 4096m
```
`kubectl create -f awesome-quota.yaml --namespace=myspace`
`kubectl describe resourcequota awesome-quota --namespace=myspace`

# Imperative
- `kubectl create quota myrq --hard=cpu=1,memory=1G,pods=2 -n myspace` (if namespace not specified, the quota is applied in the `default` namesapce)
`kubectl set resources deployment nginx --requests=cpu=0.1,memory=1Gi`
