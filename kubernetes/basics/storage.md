- Container Storage Interface (CSI) which is an open standard aimed at providing a clean storage interface for container orchestrators such as Kubernetes.
- Each provider (a.k.a provisioner) needs a CSI plugin to expose their storage assets to Kubernetes. The plugin usually runs as a set of Pods in the `kube-system` Namespace.

# StorageClasses(SC)
- Before we can use PV and PVC, we need SC
- Before we can use SC, we need to install and configure CSI storage plugin in K8s cluster
- example
```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: slow
  annotations:
    storageclass.kubernetes.io/is-default-class: "true"
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-standard
reclaimPolicy: Retain
```
- Note
  - StorageClass objects are immutable – this means you can’t modify them after they’re deployed
  - `metadata.name` should be meaningful as it’s how you and other objects refer to the class
  - The terms provisioner and plugin are used interchangeably
  - The `parameters` block is for plugin-specific values
- You can configure as many StorageClasses as you need. However, each class can only relate to a single type of storage on a single back-end.
## Access Mode
### ReadWriteOnce (RWO)
- can only be bound by a **single** PVC, Block storage usually only supports `RWO`
- allows read-write by a single node
### ReadWriteMany (RWM)
- can be bound by **multiple** PVCs, usually supported by file and object storage
- allows read-write by many nodes
### ReadOnlyMany (ROM)
- can be bound by **multiple** PVCs
- allows read-only by multiple nodes
### Notes
- a PV can only be opened in one mode across PVCs
- for a PV to be used by several Pods, the access mode claimed must be the same by all the Pods; it is not possible that one Pod uses a ReadOnlyMany mode and another Pod uses ReadWriteMany on the same PV.
## Reclaim Policy
## Delete
- is most dangerous, but also is the default for PVs created dynamically via storage classes
- it delets the PV and associated storage resource, meaning data is lost
## Retain
-

# PersistentVolumes(PV)
- using hostPath
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: db-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /data/db
```
# PersistentVolumeClaims(PVC)
```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: db-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 512m
---
apiVersion: v1
kind: Pod
metadata:
  name: app-consuming-pvc
spec:
  volumes:
    - name: app-storage
      persistentVolumeClaim:
        claimName: db-pvc
  containers:
  - image: alpine
    name: app
    command: ["/bin/sh"]
    args: ["-c", "while true; do sleep 60; done;"]
    volumeMounts:
      - mountPath: "/mnt/data"
        name: app-storage
```

# hostPath (no PV)
```
apiVersion: v1
kind: Pod
metadata:
  name: hostpath-volume-test
spec:
  restartPolicy: OnFailure
  containers:
  - name: busybox
    image: busybox:stable
    command: ['sh', '-c', 'cat /data/data.txt']
    volumeMounts:
    - name: host-data
      mountPath: /data
  volumes:
  - name: host-data
    hostPath:
      path: /etc/hostPath
      type: Directory
```
- type can be `File` `FileOrCreate` `Directory` `DirectoryOrCreate`

# Released vs Available
- A PV is initial Available, but once claimed by a PVC, it becomes Bound, once PVC deleted, the PV becomes Released.
  - However, 'Released' state does not mean it is available for other PVC to consume since it still has data from previous PVC
- Run this to make it from 'Released' to 'Available'
  - `kubectl patch pv my-pv -p '{"spec":{"claimRef": null}}'`
