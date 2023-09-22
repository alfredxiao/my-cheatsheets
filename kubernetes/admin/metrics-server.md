# Install
1. `wget https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml`
2. if using self-signed certs, add `- --kubelet-insecure-tls` after the line `- --secure-port=4443`
3. `kubectl apply -f components.yaml`

# top
- `kubectl top nodes`
- `kubectl top pod memory-demo --namespace=mem-example`
- `kubectl top pod --containers=true`
- `kubectl top pod --sort-by cpu`
