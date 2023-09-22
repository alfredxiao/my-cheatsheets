There are three classes:
When Kubernetes creates a Pod it assigns one of these QoS classes to the Pod:
- Guaranteed
- Burstable
- BestEffort

# Guaranteed
- Every container in a pod must **have specified equal** memory request and limit
- The same for CPU
- Applies to both init and app containers

# Burstable
- The Pod does not meet the criteria for QoS class Guaranteed.
- At least one Container in the Pod has a memory or CPU request or limit.

# BestEffort
- The Containers in the Pod must **not** have any memory or CPU limits or requests.

# Automatic Request Matching
- If a Container specifies its own memory limit, but does not specify a memory request, Kubernetes automatically assigns a memory request that matches the limit. 
- Similarly, if a Container specifies its own CPU limit, but does not specify a CPU request, Kubernetes automatically assigns a CPU request that matches the limit.

# Find out qosClass
- `kubectl get pods -o jsonpath="{range .items[*]}{.metadata.name} {.status.qosClass}{'\n'}"`

If you set a memory limit of 4GiB for that container, the kubelet (and container runtime) enforce the limit. The runtime prevents the container from using more than the configured resource limit. For example: when a process in the container tries to consume more than the allowed amount of memory, the system kernel terminates the process that attempted the allocation, with an out of memory (OOM) error.

If your app starts hitting your CPU limits, Kubernetes starts throttling your container, giving your app potentially worse performance. However, it won’t be terminated.
