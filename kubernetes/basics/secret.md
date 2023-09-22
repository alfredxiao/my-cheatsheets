# Types
- generic: creates from literal, an env file, a file, or a directory
- docker-registry: for use with docker registry
- tls: TLS secret

# Create Generic Secret
`kubectl create secret generic creds --from-literal user=nigelpoulton --from-literal pwd=Password123`
```
apiVersion: v1
kind: Secret
metadata:
  name: tkb-secret
  labels:
    chapter: configmaps
type: Opaque
data:
  username: bmlnZWxwb3VsdG9u
  password: UGFzc3dvcmQxMjM=
```
or
```
stringData:
  username: user1
  password: password1
```

# Mount secret volume
```
apiVersion: v1
kind: Pod
metadata:
  name: secret-pod
  labels:
    topic: secrets
spec:
  volumes:
  - name: secret-vol
    secret:
      secretName: tkb-secret
  containers:
  - name: secret-ctr
    image: nginx
    volumeMounts:
    - name: secret-vol
      mountPath: "/etc/tkb"
      readOnly: true
```
