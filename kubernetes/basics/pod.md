- Pod is the atomic unit of scheduling
- every container runs inside a Pod
- one Pod runs one or more containers
- Pod only run containers
- containers in the same Pod share network stack (IP address), volumes, IPC namespace, memory, and more
- containers in the same Pod can talk with `localhost`
- Pod is the minimum unit of scaling, you don't scale by adding more containers to existing Pod
- the deployment of a Pod is an atomic operation, meaning a Pod is only ready for service when all its containers are up and running. The entire Pod is up, or it fails
- a single Pod can be scheduled to a single node only, all containers in the same Pod runs on the same node
- Pods are mortal. if they die unexpectedly, we don't bring them back to life, K8s starts a new one in its place with new ID and IP address.
- Pods are immutable, meaning once they are created, you never change their configuration

# Label
- Let you group Pods and associate them with other objects

# Annotation
- Let you add experimental features and integration with 3rd-party tools

# Probe
- Let you test the health and status of Pods, supporting advanced scheduling, updates, etc.

# Affinity/Anti-Affinity
- Let you control over where Pods run

# Termination Control
- Let you gracefully terminate Pods (and containers within)

# Security Policies
- Let you enforce security features

# Resource Requests and Limits
- Let you specify allowed usage of CPU, memory and disk IO

# Static Pods
- add a pod yml file to `/etc/kubernetes/manifests/` which is default value for `staticPodPath` in `/var/lib/kubelet/config.yaml`

# Dynamic Pods
- via controllers, e.g. Deployments, StatefulSets

# Pod is
Seriously though, a Pod is a collection of resources that containers running inside of it inherit and share. These resources are actually Linux kernel namespaces, and include the following:
- netnamespace:IPaddress,portrange,routingtable...
- pidnamespace:isolatedprocesstree
- mntnamespace:filesystemsandvolumes...
- UTSnamespace:Hostname
- IPCnamespace:Unixdomainsocketsandsharedmemory

# Intra Pod communication
- containers talk to each other via `localhost` with different port numbers

# The Pod network
- every Pod gets its own unique IP addresses that’s fully routable on an internal Kubernetes network called the pod network.
- The pod network is flat, meaning every Pod can talk directly to every other Pod without the need for complex port mappings.
- In a default out-of-the-box cluster, the Pod network is wide open from a security perspective. You should use Network Policies to lock down access.

# Atomicity
- Pod deployment is an atomic operation. This means it’s all-or-nothing – deployment either succeeds or it doesn’t.
- You’ll never have a scenario where a partially deployed Pod is servicing requests.
- Only after all a Pod’s resources are running and ready will it start servicing requests.

# restartPolicy
- `pod.spec.restartPolicy` possible values `Always`, `OnFailure`, `Never`, default is `Always`
  - `OnFailure` restarts it only when exit code is non-zero, `Always` does not care about exit code

# `pod.spec.terminationGracePeriodSeconds`
- the periods after which processes running in the containers of a terminating Pod are killed
- defaults to 30 seconds

# `pod.spec.activeDeadlineSeconds`
- after which a running Pod will be stopped if not yet terminated

# Force deletion
- `kubectl delete pod mypod  --force --grace-period=0`

# Long vs Short lived
- Kubernetes has several controllers for different types of long-lived and short-lived workloads.
- `Deployments`, `StatefulSets`, and `DaemonSets` are examples of controllers designed for long-lived Pods.
  - `Always` is the default restart policy and appropriate for most long-lived Pods.
- `Jobs` and `CronJobs` are examples designed for short-lived Pods.
  - Appropriate container restart policies for short-lived Pods will usually be `Never` or `OnFailure`.

# `spec.imagePullSecrets`
- help to download container images from private registries

# Multi-container Patterns
## Sidecar Pattern
- It has a main application container and a sidecar container.
- It’s the job of the sidecar to augment or perform a secondary task for the main application container.
- An increasingly important user of the sidecar model is the service mesh. At a high level, service meshes inject sidecar containers into application Pods, and the sidecars do things like encrypt traffic and expose telemetry and metrics.
## Adapter Pattern
- helper container takes non- standardized output from the main container and rejigs it into a format required by an external system
- e.g. a helper container converts NGINX logs into a format accepted by Prometheus
## Ambassador Pattern
- helper container brokers connectivity to an external system
## Init Pattern
- an init container that’s guaranteed to start and complete before your main app container.
- It’s also guaranteed to only run once.

# Pod Lifecycle
- `Pending`	The Pod has been accepted by the Kubernetes cluster, but one or more of the containers has not been set up and made ready to run. This includes time a Pod spends waiting to be scheduled as well as the time spent downloading container images over the network.
- `Running`	The Pod has been bound to a node, and **all** of the containers have been **created**. At least one container is still running, or is in the process of starting or restarting.
- `Succeeded`	All containers in the Pod have **terminated** in success, and will not be restarted.
Failed	All containers in the Pod have terminated, and at least one container has terminated in failure. That is, the container either exited with non-zero status or was terminated by the system.
- `Unknown`	For some reason the state of the Pod could not be obtained. This phase typically occurs due to an error in communicating with the node where the Pod should be running.

# Container States
- `Waiting` A container in the Waiting state is still running the operations it requires in order to complete start up: for example, pulling the container image from a container image registry, or applying Secret data.
- `Running` indicates that a container is executing without issues.
- `Terminated` A container in the Terminated state began execution and then either ran to completion or failed for some reason.

# Example
```
apiVersion: v1
kind: Pod
metadata:
  name: hello-nginx
  labels:
    version: v1
spec:
  containers:
  - name: hello-nginx
    image: nginx
    ports:
    - containerPort: 80
```

# Probes (for Containers)
- The `kubelet` uses **liveness** probes to know when to **restart** a container. For example, liveness probes could catch a deadlock, where an application is running, but unable to make progress. Restarting a container in such a state can help to make the application more available despite bugs.
  - liveness via http
```
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-http
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/liveness
    args:
    - /server
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
        httpHeaders:
        - name: Custom-Header
          value: Awesome
      initialDelaySeconds: 3
      periodSeconds: 3
```
- The `kubelet` uses **readiness** probes to know when a container is ready to **start accepting** traffic. A Pod is considered ready when all of its containers are ready. One use of this signal is to control which Pods are used as backends for Services. When a Pod is not ready, it is removed from Service load balancers.
  - uses `spec.containers.readinessProbe`
- The `kubelet` uses **startup** probes to know when a container application **has started**. If such a probe is configured, it disables liveness and readiness checks until it succeeds, making sure those probes don't interfere with the application startup. This can be used to adopt liveness checks on slow starting containers, avoiding them getting killed by the `kubelet` before they are up and running
  - uses `spec.containers.startupProbe`
## Properties of a probe
- `initialDelaySeconds` Number of seconds after the container has started before liveness or readiness probes are initiated. Defaults to 0 seconds. Minimum value is 0.
- `periodSeconds`: How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
- `timeoutSeconds`: Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.
- `successThreshold`: Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup Probes. Minimum value is 1.
- `failureThreshold`: When a probe fails, Kubernetes will try failureThreshold times before giving up. Giving up in case of liveness probe means restarting the container. In case of readiness probe the Pod will be marked Unready. Defaults to 3. Minimum value is 1.

# `spec.readinessGate`
- If readiness gate is added to a Pod, then readiness of a Pod gets determined by readiness of **all** containers **and** status of **all** readiness gate conditions.
```
spec:
  readinessGates:
    - conditionType: "www.example.com/some-gate-1"
```
- Application/container got to use K8s client library to inform K8s that this condition is `True` or `False`

# command vs args
- `pod.spec.containers.command` serves the purpose of ENTRYPOINT
- `pod.spec.containers.args` serves the purpose of CMD
```
spec:
  containers:
  - args:
    - /bin/sh
    - -c
    - while true; do date; sleep 10; done
    image: busybox
    name: mypod
```
```
spec:
  containers:
  - command: ["/bin/sh"]
    args: ["-c", "while true; do date; sleep 10; done"]
    image: busybox
    name: mypod
```

# Init containers
- `spec.initContainers` defines one or more containers that Kubernetes guarantees will **run and complete**
before main app container starts.
- init containers are run sequentically (whereas app containers in parallel)
```
apiVersion: v1
kind: Pod
metadata:
  name: business-app
spec:
  initContainers:
  - name: configurer
    image: busybox:1.32.0
    command: ['sh', '-c', 'echo Configuring application... && \
              mkdir -p /usr/shared/app && echo -e "{\"dbConfig\": \
              {\"host\":\"localhost\",\"port\":5432,\"dbName\":\"customers\"}}" \
              > /usr/shared/app/config.json']
    volumeMounts:
    - name: configdir
      mountPath: "/usr/shared/app"
  containers:
  - image: bmuschko/nodejs-read-config:1.0.0
    name: web
    ports:
    - containerPort: 8080
    volumeMounts:
    - name: configdir
      mountPath: "/usr/shared/app"
  volumes:
  - name: configdir
    emptyDir: {}
```

# Sidecar containers
```
apiVersion: v1
kind: Pod
metadata:
  name: webserver
spec:
  containers:
  - name: nginx
    image: nginx
    volumeMounts:
    - name: logs-vol
      mountPath: /var/log/nginx
  - name: sidecar
    image: busybox
    command: ["sh","-c","while true; do if [ \"$(cat /var/log/nginx/error.log \
              | grep 'error')\" != \"\" ]; then echo 'Error discovered!'; fi; \
              sleep 10; done"]
    volumeMounts:
    - name: logs-vol
      mountPath: /var/log/nginx
  volumes:
  - name: logs-vol
    emptyDir: {}
```

# Sidecar Adapter Pattern
```
apiVersion: v1
kind: Pod
metadata:
  name: adapter
spec:
  containers:
  - args:
    - /bin/sh
    - -c
    - 'while true; do echo "$(date) | $(du -sh ~)" >> /var/logs/diskspace.txt; sleep 5; done;'
    image: busybox
    name: app
    volumeMounts:
      - name: config-volume
        mountPath: /var/logs
  - image: busybox
    name: transformer
    args:
    - /bin/sh
    - -c
    - 'sleep 20; while true; do while read LINE; do echo "$LINE" | cut -f2 -d"|" >> $(date +%Y-%m-%d-%H-%M-%S)-transformed.txt; done < /var/logs/diskspace.txt; sleep 20; done;'
    volumeMounts:
    - name: config-volume
      mountPath: /var/logs
  volumes:
  - name: config-volume
    emptyDir: {}
```

# Sidecar Ambassador Pattern
```
apiVersion: v1
kind: Pod
metadata:
  name: rate-limiter
spec:
  containers:
  - name: business-app
    image: bmuschko/nodejs-business-app:1.0.0
    ports:
    - containerPort: 8080
  - name: ambassador
    image: bmuschko/nodejs-ambassador:1.0.0
    ports:
    - containerPort: 8081
```

# `pod.spec.nodeSelector`
- selects nodes
```
spec:
  containers:
    - name: nginx
      image: nginx
  nodeSelector:
    accelerator: label-of-nodes # the selection label
```

# Run imperatively
- `kubectl run nginx1 --image=nginx`
- `kubectl run busybox --image=busybox -it --rm --restart=Never -- /bin/sh -c 'echo hello world'`
  - `--rm` will remove it once completed
- `kubectl run nginx1 --image=nginx --serviceaccount sa1`

# Start the nginx pod using the default command, but use custom arguments (arg1 .. argN) for that command
`kubectl run nginx --image=nginx -- <arg1> <arg2> ... <argN>`

# Start the nginx pod using a different command and custom arguments
`kubectl run nginx --image=nginx --command -- <cmd> <arg1> ... <argN>`
