- Ingress is all about accessing multiple web applications through a single LoadBalancer Service.
- It uses `host-based` and `path-based` routing to send traffic to the correct backend Service
- Ingress is defined in the networking.k8s.io API sub-group as a v1 object and is based on the usual two constructs.
  1. Acontroller
  2. Anobjectspec

On the topic of routing, Ingress operates at layer 7 of the OSI model, also known as the `application layer`. This means it has awareness of HTTP headers, and can inspect them and forward traffic based on hostnames and paths.

- you need to install ingress to allow functioning of an Ingress object

# Ingress
- single load balancer to serve multiple serices, saves cost
- install `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.2/deploy/static/provider/cloud/deploy.yaml`
```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: webapp-ingress
spec:
  ingressClassName: nginx
  defaultBackend:
    service:
      name: webapp
      port:
        number: 80
```
```
spec:
  ingressClassName: nginx
  rules:
  - host: webapp.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: webapp
            port:
              number: 80
  - host: echo.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: echo
            port:
              number: 3000
```
