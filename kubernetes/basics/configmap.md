# Create from literal
`kubectl create configmap testmap1 --from-literal shortname=AOS --from-literal longname="Agents of Shield"`

# Create form env file
`kubectl create configmap db-config --from-env-file=config.env`

# Create form normal file
`kubectl create cm testmap2 --from-file cmfile.txt`

# Create form normal directory
`kubectl create cm testmap3 --from-file config-dir`

# Create from yaml
```
kind: ConfigMap
apiVersion: v1
metadata:
  name: multimap
data:
  given: Nigel
  family: Poulton
```

# Mount as ENV
## Load Individuals
```
apiVersion: v1
kind: Pod
metadata:
  labels:
    chapter: configmaps
  name: envpod
spec:
  containers:
    - name: ctr1
      image: busybox
      command: ["sleep"]
      args: ["infinity"]
      env:
        - name: FIRSTNAME
          valueFrom:
            configMapKeyRef:
              name: multimap
              key: given
```
## Load a whole map
```
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - image: nginx:1.9.0
    name: app
    envFrom:
    - configMapRef:
        name: my-config-map
```

# Mount as Volume
```
apiVersion: v1
kind: Pod
metadata:
  name: cmvol
spec:
  volumes:
    - name: volmap
      configMap:
        name: multimap
  containers:
    - name: ctr
      image: nginx
      volumeMounts:
        - name: volmap
          mountPath: /etc/name
```

# Reflect changes
- only when using mount volume mode changes will be reflected in the Pod

# Edit a configmap
`kubectl edit cm mymap1`

# Update
`kubectl set env deployment nginx --from=configmap/vars --keys="var1,var2"`
`kubectl set env deployment nginx --from=configmap/vars` this uses all keys
`kubectl set env deployment nginx --from=secret/passwords  --keys="pass1"`
`kubectl set env deployment nginx --from=secret/passwords` this uses all keys
