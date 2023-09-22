Because of this there are 4 different RBAC combinations and 3 valid ones:

- `Role` + `RoleBinding` (available in single Namespace, applied in single Namespace)
- `ClusterRole` + `ClusterRoleBinding` (available cluster-wide, applied cluster-wide)
- `ClusterRole` + `RoleBinding` (available cluster-wide, applied in single Namespace)
- `Role` + `ClusterRoleBinding` (NOT POSSIBLE: available in single Namespace, applied cluster-wide)
