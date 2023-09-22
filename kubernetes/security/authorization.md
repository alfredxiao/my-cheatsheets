* Kubernetes authorizes API requests using the **API server**.
* When multiple authorization modules are configured, each is checked in sequence. If any authorizer approves or denies a request, that decision is **immediately returned and no other authorizer is consulted**. If all modules have no opinion on the request, then the request is denied. A deny returns an HTTP status code 403

They can be configured as kube-apiserver startup options:
- `/etc/kubernetes/manifests/kube-apiserver.yaml`
`--authorization-mode=Node,RBAC`
`--authorization-mode=Webhook`

The Node authorization is dedicated for kubelet to communicate to the kube-apiserver such that it does not need to be allowed via RBAC. All non-kubelet traffic would then be checked via RBAC.

# `Node` Mode
- On node machines, `/etc/kubernetes/kubelet.conf` defines
```
users:
- name: default-auth
  user:
    client-certificate: /var/lib/kubelet/pki/kubelet-client-current.pem
    client-key: /var/lib/kubelet/pki/kubelet-client-current.pem
```
- Content of `client-certificate`
`Subject: O = system:nodes, CN = system:node:<node_name_which_defaults_to_hostname>`
  - or overridden via the kubelet option `--hostname-override`
- In order to be authorized by the Node authorizer, kubelets must use a credential that identifies them as being in the `system:nodes` group, with a username of `system:node:<nodeName>`.
- To enable the Node authorizer, start the apiserver (on control plane) with `--authorization-mode=Node`.
- To limit the API objects kubelets are able to write, enable the NodeRestriction admission plugin by starting the apiserver with `--enable-admission-plugins=...,NodeRestriction,...`

# `RBAC` Mode
## Role and Binding
An RBAC `Role` or `ClusterRole` contains rules that represent a set of permissions. Permissions are purely additive (there are no "deny" rules).

A `Role` always sets permissions within a particular namespace; when you create a `Role`, you have to specify the namespace it belongs in.

`ClusterRole`, by contrast, is a non-namespaced resource. The resources have different names (`Role` and `ClusterRole`) because a Kubernetes object always has to be either namespaced or not namespaced; it can't be both.

A role binding grants the permissions defined in a role to a user or set of users. It holds a list of subjects (users, groups, or service accounts), and a reference to the role being granted. A `RoleBinding` grants permissions within a specific namespace whereas a `ClusterRoleBinding` grants that access cluster-wide.

`ClusterRole`s have several uses. You can use a ClusterRole to:
- define permissions on namespaced resources
  - namespaced resources (like Pods)
  - Granting (via ClusterRoleBinding)
    - be granted within individual namespace(s) (to ServiceAccounts)
    - granted across all namespaces (to User/Group)
- define permissions on cluster-scoped resources
  - cluster-scoped resources (like nodes)
- non-resource endpoints (like /healthz)

- A `RoleBinding` may reference any `Role` in the same namespace.
- A `RoleBinding` can reference a `ClusterRole` and bind that `ClusterRole` to the namespace of the `RoleBinding`.
- If you want to bind a `ClusterRole` to all the namespaces in your cluster, you use a `ClusterRoleBinding`.

Role:
- ns1:role1
ClusterRole:
- cr1
User
- user1
ServiceAccount:
- ns1:sa1
- ns2:sa2
RoleBinding
- namespace=ns1,role=ns1:role1,subject=user1
- namespace=ns1,role=ns1:role1,subject=ns1:sa1
- namespace=ns1,role=ns1:role1,subject=ns2:sa2
- namespace=ns1,role=cr1,subject=user1
- namespace=ns1,role=cr1,subject=ns1:sa1
- namespace=ns1,role=cr1,subject=ns2:sa2
ClusterRoleBinding
- role=cr1,subject=user1
- role=cr1,subject=ns1:sa1
- role=cr1,subject=ns2:sa2


## update
After you create a binding, you cannot change the Role or ClusterRole that it refers to. If you try to change a binding's roleRef, you get a validation error. If you do want to change the roleRef for a binding, you need to remove the binding object and create a replacement.

## Reconcile
The `kubectl auth reconcile` command-line utility creates or updates a manifest file containing RBAC objects, and handles deleting and recreating binding objects if required to change the role they refer to.

## Referring to resource
`GET /api/v1/namespaces/{namespace}/pods/{name}/log`
```
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: pod-and-pod-logs-reader
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list"]
```

below explicitly specify certain names
```
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: configmap-updater
rules:
- apiGroups: [""]
  #
  # at the HTTP level, the name of the resource for accessing ConfigMap
  # objects is "configmaps"
  resources: ["configmaps"]
  resourceNames: ["my-configmap"]
  verbs: ["update", "get"]
```

## Aggregation
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: monitoring
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.example.com/aggregate-to-monitoring: "true"
rules: [] # The control plane automatically fills in the rules
```

## Referring to subjects

A RoleBinding or ClusterRoleBinding binds a role to subjects. Subjects can be `groups`, `users` or `ServiceAccounts`.

Kubernetes represents usernames as strings. These can be: plain names, such as "alice"; email-style names, like "bob@example.com"; or numeric user IDs represented as a string. It is up to you as a cluster administrator to configure the authentication modules so that authentication produces usernames in the format you want.

ServiceAccounts have names prefixed with `system:serviceaccount:`, and belong to groups that have names prefixed with `system:serviceaccounts:`.

Note:
`system:serviceaccount`: (singular) is the prefix for service account usernames.
`system:serviceaccounts`: (plural) is the prefix for service account groups.

- For a user named alice@example.com:
```
subjects:
- kind: User
  name: "alice@example.com"
  apiGroup: rbac.authorization.k8s.io
```

For a group named frontend-admins:
```
subjects:
- kind: Group
  name: "frontend-admins"
  apiGroup: rbac.authorization.k8s.io
```
For the default service account in the "kube-system" namespace:
```
subjects:
- kind: ServiceAccount
  name: default
  namespace: kube-system
```

For all service accounts in the "qa" namespace:
```
subjects:
- kind: Group
  name: system:serviceaccounts:qa
  apiGroup: rbac.authorization.k8s.io
```

For all service accounts in any namespace:
```
subjects:
- kind: Group
  name: system:serviceaccounts
  apiGroup: rbac.authorization.k8s.io
```

For all authenticated users:
```
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
```

For all unauthenticated users:
```
subjects:
- kind: Group
  name: system:unauthenticated
  apiGroup: rbac.authorization.k8s.io
```

For all users:
```
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
- kind: Group
  name: system:unauthenticated
  apiGroup: rbac.authorization.k8s.io
```

# `Webhook` mode
When specified, mode Webhook causes Kubernetes to query an outside REST service when determining user privileges


# Admission Control
The last step in completing an API request is one or more admission controls.

Admission controllers are pieces of software that can access the content request payload. They can modify the content or validate it, and potentially deny the request.
For example, the ResourceQuota controller will ensure that the object created does not violate any of the existing quotas.

# Security Context
Pods and containers within pods can be given specific security constraints to limit what processes running in containers can do. For example, the UID of the process, the Linux capabilities, and the filesystem group can be limited.

Clusters installed using `kubeadm` allow pods any possible elevation in privilege by default. For example, a pod could control the nodes networking configuration, disable SELinux, override root, and more. These abilities are almost always limited by cluster administrators.

This security limitation is called a security context. It can be defined for the entire pod or per container, and is represented as additional sections in the resources manifests. A notable difference is that Linux capabilities are set at the container level.

For example, if you want to enforce a policy that containers cannot run their process as the root user, you can add a pod security context like the one below:
```
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  securityContext:
    runAsNonRoot: true
  containers:
  - image: nginx
    name: nginx
```
Then, when you create this Pod, you will see a warning that the container is trying to run as root and that it is not allowed. Hence, the Pod will never run.

## Features
A security context defines privilege and access control settings for a Pod or Container. Security context settings include, but are not limited to:
- Discretionary Access Control: Permission to access an object, like a file, is based on user ID (UID) and group ID (GID).
- Security Enhanced Linux (SELinux): Objects are assigned security labels.
- Running as privileged or unprivileged.
- Linux Capabilities: Give a process some privileges, but not all the privileges of the root user.
- AppArmor: Use program profiles to restrict the capabilities of individual programs.
- Seccomp: Filter a process's system calls.
- `allowPrivilegeEscalation`: Controls whether a process can gain more privileges than its parent process. This bool directly controls whether the no_new_privs flag gets set on the container process. allowPrivilegeEscalation is always true when the container:
  - is run as privileged, or
  - has CAP_SYS_ADMIN
- `readOnlyRootFilesystem`: Mounts the container's root filesystem as read-only.

- `kubectl auth can-i --list --as johndoe` to list all permissions
- `kubectl auth can-i list pods --as johndoe` to answer yes/no
- `kubectl auth can-i list pods --as=system:serviceaccount:space1:default -n practice`
