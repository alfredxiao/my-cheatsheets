1. setup local registry
2. setup containerd on nodes to talk to the registry
3. setup containerd to use http instead of https
[plugins."io.containerd.grpc.v1.cri".registry.mirrors]
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."10.109.49.204:5000"]
    endpoint = ["http://10.109.49.204:5000"]

  `sudo systemctl daemon-reload`
  `sudo systemctl restart containerd`
4. Go build, use static building `CGO_ENABLED=0`
5. Build image (docker or podman), push to registry

`podman tag localhost/my-scheduler 10.109.49.204:5000/my-scheduler`
`podman push 10.109.49.204:5000/my-scheduler`

6. Run a Deployment of the scheduler with service account that has permission assigned
7. Run pod with `spec.schedulerName` set to the name the schduler program uses
