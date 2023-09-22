- Services bring stable IP addresses and DNS names to the unstable world of Pods.
- Stable IP, stble DNS name, stable port
- Every time you create a Service, Kubernetes automatically creates an associated Endpoints object. The Endpoints object is used to store a dynamic list of healthy Pods matching the Service’s label selector.
  - `kubectl get endpoints`
  - `kubectl get endpointslices`


# ClusterIP
- A ClusterIP Service has a **stable** virtual IP address that is **only accessible from inside the cluster**.
- the ClusterIP is registered against the name of the Service in the cluster’s internal DNS service.
- All Pods in the cluster are pre-programmed to use the cluster’s DNS service, meaning all Pods can convert Service names to ClusterIPs.
- in below example, Pods can access `my-service:8080`
```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: ClusterIP
  ports:
  - port: 8080
    protocol: TCP
  selector:
    app: hello-world
```

# NodePort
- it extends ClusterIP
- it enables external access via dedicated port on every cluster node
- in below example, Pods can access `my-service:8080`
  - Clients connecting from outside the cluster can send traffic to **any cluster node** on port `30050`.
  - which can be redirected to any Pod (matching selector) as it does basic load-balancing
```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: NodePort
  ports:
  - port: 8080
    nodePort: 30050
  selector:
    app: hello-world
```
- or, can imperatively create it like `kubectl expose deployment my-deploy --type=NodePort`

# LoadBalancer
# ExternalName
- maps a service to a DNS name



# list endpoints
- `kubectl get endpoints myservice`

# Create Service imperatively
- `kubectl create service clusterip myservice --tcp=80:80`
- `kubectl create service nodeport my-svc --tcp=8084:80 --node-port=30080`
- `kubectl run nginx --image=nginx --restart=Never --port=80 --expose`

- `k create deployment ndeploy --image=nginx --replicas=2`
- `k expose deployment ndeploy --port=80 --target-port=80 --name=my-service --type=NodePort`

# update a service
- `k patch service nginx -p '{"spec":{"type":"NodePort"}}'`

# Headless
- each pod gets a DNS record such that other pods can access them via the pod's own unique DNS name
- When learning about headless Services, it can be useful to visualize a Service object with a head and a tail. The head is the stable IP address, and the tail is the list of Pods it sends traffic to. Therefore, a headless Service is a Service object without an IP address.
- The sole purpose of a headless Service is to create DNS SRV records for Pods that match its label selector. Clients then need to know to use DNS to reach Pods instead of using the Service’s ClusterIP. This is why a headless Service doesn’t need a ClusterIP.
```
apiVersion: v1
kind: Service
metadata:
  name: default-subdomain
spec:
  selector:
    app: busybox
  clusterIP: None
  ports:
  - name: foo # Actually, no port is needed.
    port: 1234
    targetPort: 1234
---
apiVersion: v1
kind: Pod
metadata:
  name: busybox1
  labels:
    app: busybox
spec:
  hostname: busybox-1
  subdomain: default-subdomain
  containers:
  - image: busybox:1.28
    command:
      - sleep
      - "3600"
    name: busybox
---
apiVersion: v1
kind: Pod
metadata:
  name: busybox2
  labels:
    app: busybox
spec:
  hostname: busybox-2
  subdomain: default-subdomain
  containers:
  - image: busybox:1.28
    command:
      - sleep
      - "3600"
    name: busybox
# busybox-1/2.default-subdomain.default.svc.cluster.local maps to individual IP
# default-subdomain.default.svc.cluster.local Round Robin to one of the two IPs
```
- when used with statefulset, The `spec.serviceName` field designates the governing Service of a statefulset

# Reference port name
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webapp-deployment
spec:
  replicas: 2
  selector:
    matchLabels:
      app: webapp
  template:
    metadata:
      labels:
        app: webapp
    spec:
      containers:
      - name: webapp
        image: katacoda/docker-http-server:latest
        ports:
        - name: webport
          containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: webapp-service
  labels:
    app: webapp
spec:
  type: ClusterIP
  ports:
  - port: 8080
    targetPort: webport
  selector:
    app: webapp
```
