- Deployments are for stateless apps
- One Deployment, one Pod template
  - meaning if you have front-end and backend-end two Pod templates, you need two Deployments

- `ReplicaSets` manage Pods and bring self-healing and scaling.
- `Deployments` manage `ReplicaSets` and add rollouts and rollbacks.
- Think of `Deployments` as managing `ReplicaSets`, and `ReplicaSets` as managing `Pods`. Put them together, and you’ve got a great way to deploy and manage stateless applications on Kubernetes.

- `spec.minReadySeconds` indicates the number of seconds the controller should wait after a Pod starts without failing to consider the Pod is ready.

- `spec.revisionHistoryLimit` tells Kubernetes how many older versions/ReplicaSets to keep. Keeping more gives you more rollback options, but keeping too many can bloat the object. This can be a problem on large clusters with lots of software releases.

- `spec.progressDeadlineSeconds` tells Kubernetes how long to wait during a rollout for each new replica to come online. The example sets a 5 minute deadline, meaning each new replica has 5 minutes to come up before Kubernetes considers the rollout stalled. To be clear, the clock is reset after each new replica comes up, meaning each step in the rollout gets its own 5 minute window.

- `spec.strategy.type`: Can be `"Recreate"` or `"RollingUpdate"`. Default is RollingUpdate.
  - `"Recreate"` Kill all existing pods before creating new ones.
  - `"RollingUpdate"` Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.

- The Deployment’s label selector is immutable, so you can’t change it once it’s deployed.

# Create from CLI
`kubectl create deployment my-deploy --image=nginx:1.14.2 --replicas=2`

# Create from file
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deploy
  labels:
    app: my-deploy
spec:
  replicas: 1
  selector:
    matchLabels:
      app: my-deploy
  template:
    metadata:
      labels:
        app: my-deploy
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
```

# Rollout
- `kubectl apply -f deploy.v2.yml --record`
- `kubectl rollout status deployment hello-deploy`
- `kubectl rollout history deployment hello-deploy`
- `kubectl rollout pause deploy hello-deploy`
- `kubectl rollout resume deploy hello-deploy`

# Detail of one revision
- `kubectl rollout history deployments hello-deploy --revision=2`
- `kubectl rollout history deployments hello-deploy --revision=2 -o yaml`

# Rollback
- `kubectl rollout undo deployment hello-deploy --to-revision=1` this is imperative way

# Scale
`kubectl scale deployment my-deploy --replicas=5`

# Autoscale
- `kubectl autoscale deployment my-deploy --cpu-percent=70 --min=2 --max=8`
  - this will creates a HPA with same name as the deployment

# Update image
- `kubectl set image deployment my-deploy nginx=nginx:1.19.2`
- `kubectl set image deploy deployment1 nginx=nginx:1.19.8`
  - first `nginx` above is the name of the container

# Set Resource
- `ubectl set resources deployment nginx --requests=cpu=0.1,memory=1Gi`
