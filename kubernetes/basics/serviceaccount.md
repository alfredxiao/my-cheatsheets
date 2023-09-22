# Create
- `kubectl create sa myaccount`

# Assign to a Pod
- `kubectl run nginx --image=nginx --restart=Never --serviceaccount=custom`
```
spec:
  serviceAccountName: custom
```
