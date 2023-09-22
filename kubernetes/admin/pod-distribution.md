
- Node affinity functions like the `nodeSelector` field but is more expressive and allows you to specify soft rules.
  - `requiredDuringSchedulingIgnoredDuringExecution` is like `nodeSelector`, but more expressive
  - `preferredDuringSchedulingIgnoredDuringExecution` allows soft rules
  - In the preceding types, `IgnoredDuringExecution` means that if the node labels change after Kubernetes schedules the Pod, the Pod continues to run.
- Inter-pod affinity/anti-affinity allows you to constrain Pods against labels on other Pods.
  - Inter-pod affinity and anti-affinity allow you to constrain which nodes your Pods can be scheduled on based on the labels of Pods already running on that node, instead of the node labels.

Types
- `spec.affinity.nodeAffinity`
- `spec.affinity.nodeAntiAffinity`
- `spec.affinity.podAffinity`
- `spec.affinity.podAntiAffinity`

Using topologySpreadConstraints
```
spec:
  topologySpreadConstraints:
  - maxSkew: <integer>
    topologyKey: <string>
    whenUnsatisfiable: <string>
    labelSelector: <object>
```
