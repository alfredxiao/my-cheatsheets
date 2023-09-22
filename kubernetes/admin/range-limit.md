# Default cpu requests and limits for a Namespace
If you create a Pod within a namespace that has a default CPU limit, and any container in that Pod does not specify its own CPU limit, then the control plane assigns the default CPU limit to that container.

# Create RangeLimit
```
apiVersion: v1
kind: LimitRange
metadata:
  name: cpu-limit-range
  namespace: ns-with-default-limit
spec:
  limits:
  - default:
      cpu: 1
    defaultRequest:
      cpu: 0.5
    type: Container
```

- If you create a pod without specifying cpu request or limit, they get assigned the settings from `LimitRnge` in the same `Namespace` (if there is a LimitRange)
- If you create a pod specifying a limit but not the request, the request get assigned to match its own limit value you specified. (This is the case regardless of the existence of `LimitRange`)
- If you create a pod specifying a request but not the limit, the limit is assigned from `LimitRange` in the same `Namespace`

# Memory
- behave similar to CPU described above
