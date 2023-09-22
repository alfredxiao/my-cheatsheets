When process namespace sharing is enabled, processes in a container are visible to all other containers in that pod.

You can use this feature to configure cooperating containers, such as a log handler sidecar container, or to troubleshoot container images that don't include debugging utilities like a shell.
```
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  shareProcessNamespace: true
  containers:
  - name: nginx
    image: nginx
  - name: shell
    image: busybox:1.28
    securityContext:
      capabilities:
        add:
        - SYS_PTRACE
    stdin: true
    tty: true

kubectl attach -it nginx -c shell

```
You can signal processes in other containers. For example, send `SIGHUP` to nginx to restart the worker process. This requires the `SYS_PTRACE` capability.
```
PID   USER     TIME  COMMAND
    1 65535     0:00 /pause
    8 root      0:00 nginx: master process nginx -g daemon off;
   14 root      0:00 sh
   45 101       0:00 nginx: worker process
   46 101       0:00 nginx: worker process
   47 root      0:00 ps
/ # kill -HUP 8
# this will restart (start new PID processes for worker process like 45, 46)
```

Container filesystems are visible to other containers in the pod through the `/proc/$pid/root` link
