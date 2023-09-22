# emptyDir
- lasts for the life of the Pod, even if the Container terminates and restarts
  - An emptyDir volume is first created when a Pod is assigned to a node, and exists as long as that Pod is running on that node.
  - meaning, as long as Pod is not deleted, and even if you terminate the container (kill the container process) which gets restarted, the files in this `emptyDir` volume is still there

```
spec:
  containers:
  - name: redis
    image: redis
    volumeMounts:
    - name: redis-storage
      mountPath: /data/redis
  volumes:
  - name: redis-storage
    emptyDir: {}
```
- As the name says, the emptyDir volume is initially empty. All containers in the Pod can read and write the same files in the `emptyDir` volume, though that volume can be mounted at the same or different paths in each container. When a Pod is removed from a node for any reason, the data in the `emptyDir` is deleted permanently.

- Depending on your environment, `emptyDir` volumes are stored on whatever medium that backs the node such as disk or SSD, or network storage. However, if you set the `emptyDir.medium` field to `"Memory"`, Kubernetes mounts a `tmpfs` (RAM-backed filesystem) for you instead. While `tmpfs` is very fast, be aware that unlike disks, `tmpfs` is cleared on node reboot and any files you write count against your container's memory limit.

- `emptyDir.medium` The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory.
