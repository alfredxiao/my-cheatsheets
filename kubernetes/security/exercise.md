## 1. Know how to configure authentication and authorization
### Authentication and authorization
- Create two namespaces
```bash
kubectl create ns production
kubectl create ns development
```
- Create a new user and password
```
sudo useradd DevDan -s /bin/bash
sudo passwd DevDan
```
- Generate private key then csr for DevDan
```
openssl genrsa -out DevDan.key 2048
openssl req -new -key DevDan.key -out DevDan.csr -subj "/CN=DevDan/O=development"
# To view the csr
openssl req -in DevDan.csr -text -noout
```
- Generate the self-signed certificate
```
sudo openssl x509 -req -in DevDan.csr -out DevDan.crt \
     -CA /etc/kubernetes/pki/ca.crt \
     -CAkey /etc/kubernetes/pki/ca.key \
     -CAcreateserial -days 45
# To view the certificate
openssl x509 -in DevDan.crt -text -noout
```
- Update the access config file to add user DevDan
```
kubectl config set-credentials DevDan --client-certificate=/home/cloud_user/devdan/DevDan.crt --client-key=/home/cloud_user/devdan/DevDan.key
```
- Create context name DevDan-context
```
kubectl config set-context DevDan-context --user=DevDan --cluster=kubernetes --namespace=development
kubectl config get-contexts
```
- Create role 'developer'
```
kubectl create role developer -n development --resource=pods,deployments,replicasets --verb="*" --dry-run=client -o yaml > role-dev.yaml
vim role-dev.yaml
kubectl create -f role-dev.yaml
```
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: developer
  namespace: development
rules:
- apiGroups: ["", "extensions", "apps"]
  resources: ["pods", "deployments", "replicasets"]
  verbs: ["*"]
```
- Create rolebinding 'developer-role-binding'
```
kubectl create rolebinding developer-role-binding --role=developer --user=DevDan -n development --dry-run=client -o yaml > developer-role-binding.yaml
kubectl create -f developer-role-binding.yaml
```
- In DevDan-context, create a new deployment nginx with image nginx, verify it exists, then delete it
```
kubectl --context=DevDan-context create deployment nginx --image=nginx
kubectl --context=DevDan-context get deployment
kubectl --context=DevDan-context delete deployment nginx
```

## 2. Understand Kubernetes security primitives
### Certificate API
- Create a new user DevTrang with password lfs258 (list all users with cat /etc/passwd)
- Create DevTrang.key and DevTrang.csr with openssl.
- Create a Certificate signing request (csr) object from DevTrang.csr to send to the Kubernetes API ([here](https://kubernetes.io/docs/tasks/tls/managing-tls-in-a-cluster/#create-a-certificate-signing-request-object-to-send-to-the-kubernetes-api)).
- Get csr and approve it.
- Download the certificate and use it.

- Create new user and password
  ```
  sudo useradd DevTrang -s /bin/bash
  sudo passwd DevTrang
  ```
- Create private key and csr
  ```
  openssl genrsa -out DevTrang.key 2048
  openssl req -new -key DevTrang.key -out DevTrang.csr --subj "/CN=DevTrang/O=development"
  ```
- Create csr object in Kubernetes
  - Search kubernetes.io for csr [here](https://kubernetes.io/docs/tasks/tls/managing-tls-in-a-cluster/#create-a-certificate-signing-request-object-to-send-to-the-kubernetes-api).
  - Copy and paste to terminal, edit to use DevTrang.csr
```
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: DevTrang
spec:
  request: $(cat DevTrang.csr | base64 | tr -d '\n')
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - client auth
EOF
```
- Get csr and approve it.
  ```
  kubectl get csr
  kubectl certificate approve DevTrang
  ```
- Download and use it.
  ```
  kubectl get csr DevTrang -o jsonpath='{.status.certificate}' | base64 --decode > DevTrang.crt
  ```

### Cluster roles and cluster role bindings
- A new user DevDan joined the team. He will be focusing on the nodes in the cluster:
  - Create the required ClusterRoles and ClusterRoleBindings
  - So he gets access to the nodes (Grant permission to list nodes)
- Create the required ClusterRoles and ClusterRoleBindings to allow DevDan access to Storage.
  - ClusterRole: storage-admin, Resource: persistentvolumes, Resource: storageclasses
  - ClusterRoleBinding: DevDan-storage-admin, ClusterRoleBinding Subject: DevDan, ClusterRoleBinding Role: storage-admin

```
kubectl create clusterrole DevDan-cluster-role --verb=list --resource=nodes --dry-run -o yaml > DevDan-cluster-role.yaml
kubectl create -f DevDan-cluster-role.yaml
kubectl create clusterrolebinding DevDan-role-binding --user=DevDan --clusterrole=DevDan-cluster-role
# Check the authorization
kubectl auth can-i --list --as DevDan
kubectl auth can-i list nodes --as DevDan

kubectl create clusterrole storage-admin --resource=persistentvolumes --resource=storageclasses --verb="*"
kubectl create clusterrolebinding DevDan-storage-admin --user=DevDan --clusterrole=storage-admin
# Check the authorization
kubectl auth can-i --list --as DevDan
```
