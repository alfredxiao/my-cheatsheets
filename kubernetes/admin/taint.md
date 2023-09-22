Node affinity is a property of Pods that attracts them to a set of nodes (either as a preference or a hard requirement).
Taints are the opposite -- they allow a node to repel a set of pods.
- taints are applied to nodes

Tolerations and allow (but do not require) the pods to schedule onto nodes with matching taints.
- tolerations are applied to pods

# Add taint
`kubectl taint nodes node1 key1=value1:NoSchedule`

# Remove taint
`kubectl taint nodes node1 key1=value1:NoSchedule-`

# NoSchedule
if there is at least one un-ignored taint with effect NoSchedule then Kubernetes will not schedule the pod onto that node
- in other words, please do not schedule pods on it

# PreferNoSchedule
if there is no un-ignored taint with effect NoSchedule but there is at least one un-ignored taint with effect PreferNoSchedule then Kubernetes will try to not schedule the pod onto the node
- in other words, please try not to schedule pods on it

# NoExecute
if there is at least one un-ignored taint with effect NoExecute then the pod will be evicted from the node (if it is already running on the node), and will not be scheduled onto the node (if it is not yet running on the node).
- in other words, please make sure pods not running or scheduled on it

# Scheduler
Only the scheduler takes `tains/tolerations/affinity` into account when finding the correct node name. This means, if you specify `spec.nodeName` for a pod to master node, without any `tolerations`, the pod will be started in master node (by the kubelet on master node) without going through any scheduler.
