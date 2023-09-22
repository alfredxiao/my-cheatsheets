on a pod
```
spec:
  securityContext:
    runAsUser: 1000
    runAsGroup: 3000
    fsGroup: 2000
```
if `fsGroup` was not specified, files would be owned by the group specified by `runAsGroup`.
