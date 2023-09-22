- As it’s the main command-line tool, it’s important you use a version that’s no more than one minor version higher or lower than your cluster. For example, if your cluster is running Kubernetes `1.20.x`, your kubectl should be between `1.19.x` and `1.21.x`.

- At a high-level, `kubectl` converts user-friendly commands into HTTP REST requests with JSON content required by the API server.
- It uses a configuration file to know which cluster and API server endpoint to send commands to, and it takes care of sending authentication data with commands.

# Config File
- The `kubectl` configuration file is called config and lives in a hidden directory called kube in your home directory (`$HOME/.kube/config`).
- `kubectl config view` to view simplified version of config
- `kubectl config current-context` to view current context
- `kubectl config use-context` to switch context
## Contents
- Cluster: defines K8s clusters and their API server endpoint
- Users: Defines users and credentials
- Contexts: Defines combination of users and clusters

# Explain/Documentation
- `kubectl explain pods` or `explain pod`
- `kubectl explain pod.spec.affinity`

# Get Pod Details
- `kubectl get pods hello-pod -o yaml`
  - this includes desired state (under `spec` section) and observed state (under `status` section)
- `kubectl describe pods hello-pod`
- `kubectl get pods --show-labels` with Labels

# Logs
- `kubectl logs <pod>` (for single container pod, or first container anyway - first in yaml)
- `kubectl logs <pod> --container <container>` or `-c`
- `--all-containers=true` to get logs from all containers
- `kubectl logs mypod -f` to follow
- `kubectl logs mypod -p` or `--previous` for logs from previous exited pod instance
- `kubectl logs mypod -l app=nginx` to use selector for matching pods
- `kubectl logs mypod -l app=nginx --prefix` to prefix logs with pod name
- `--timestamps` displays timestamps of logs at the beginning of lines.
- `kubectl logs deployment/mydeploy` views a random pod from this deployment
- `kubectl logs --selector app=random-logger` selects multiple pods (by default return 10 lines from each pod)
- `kubectl logs --selector app=random-logger --tail=50` selects multiple pods and select 50 lines from each pod
- `kubectl logs service/bipedal --all-containers` view logs via service that fronts pods

# Exec
- `kubectl exec hello-pod -- ps aux`
- `kubectl exec -it hello-pod -- sh`

# Resource
- `kubectl api-resources`

# Delete
- `kubectl delete -f abc.yaml` could be a deployment or service or other yaml file

# Scale
- `kubectl scale deployment hello-deploy --replicas=10` this is imperative way

# Run
- `kubectl run nginx1 --image=nginx --labels=app=v1`
- `kubectl run hazelcast --image=hazelcast/hazelcast --restart=Never --port=5701 --env="DNS_DOMAIN=cluster" --labels="app=hazelcast,env=prod”`
- `kubectl run nginx1 --image=nginx --env PROFILE=dev --env USERNAME=user1`
- `kubectl run busybox --image=busybox --restart=Never --command -- env`
# Edit
- `kubectl edit pod mypod`

# Apply
- `kubectl apply -f pod.yaml`
# Create
- `kubectl create -f pod.yaml`
# Replace
- `kubectl replace -f pod.yaml`

# Events
- `kubectl get events`

# Debug with ephemeral containers
- `kubectl debug -it pod1 --image=busybox` where busybox is a ephemeral container for debugging purpose

# Drain
- `kubectl drain worker-0 --force`

# Uncordon
- `kubectl uncordon worker-0`

# cp
- `kubectl cp mypod:/file/in/pod file-on-host`
- `kubectl cp file-on-host mypod:/file/in/pod`
- `namespace/pod:/path/to/file` for namespaced file

# Patch
- `kubectl patch svc nginx -p '{"spec":{"type":"NodePort"}}'`

# cluster info
- `kubectl cluster-info`
- `kubectl cluster-info dump`

# Custom Columns
display container name and container image
`kubectl get pod  -o custom-columns=CONTAINER:.spec.containers[*].name,IMAGE:.spec.containers[*].image`
