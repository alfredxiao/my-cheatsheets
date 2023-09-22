# Query
- `kubectl get pods --show-labels`

# Filter
- `kubectl get pods --selector=app=cassandra`
- `kubectl get pods --selector=app=cassandra,env=prod,tier=frontend`

# Add new Label
- `kubectl label pods foo label1=true`

# Update label
- `kubectl label pods foo label1=false --overwrite`

# Update all pods in the namespace
`kubectl label pods --all status=unhealthy`

# Update a pod identified by the type and name in "pod.json"
`kubectl label -f pod.json status=unhealthy`

# Update pod 'foo' only if the resource is unchanged from version 1.
`kubectl label pods foo status=unhealthy --resource-version=1`

# Update pod 'foo' by removing a label named 'bar' if it exists.
- Does not require the --overwrite flag.
- `kubectl label pods foo bar-`
- `kubectl label po nginx1 nginx2 nginx3 app-` for multiple pods
or
- `kubectl label po nginx{1..3} app-`
or
`kubectl label po -l app app-` via selector

# Notes
- label value is limited to 63 characters maximum

# Set labels
- `kubectl run labeled-pod --image=nginx --restart=Never --labels=tier=backend,env=dev`
```
metadta:
  labels:
    env: prod
```

# Select labels
- `kubectl get pods --selector="abc=123"` or `-l`
- `--selector="env!=prod"`
- `--selector="env in (dev,prod)"`
- `--selector="env notin (dev,test)"`
- `--selector="env in (dev,test),ver=2"`
```
spec:
  podSelector:
    matchLabels:
      tier: frontend
```

# Show label as column
`kubectl get pod -L app`
or
`kubectl get pod --label-columns=app`

# Update label using selector
`kubectl label pod -l "app in(v1,v2)" tier=web`
