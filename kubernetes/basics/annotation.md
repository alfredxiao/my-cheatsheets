
# Via YAML
```
apiVersion: v1
kind: Pod
metadata:
  name: annotated-pod
  annotations:
    commit: 866a8dc
    author: 'Benjamin Muschko'
    branch: 'bm/bugfix
```

# Query Annotations
- `kubectl annotate pod nginx1 --list`
- `kubectl describe po nginx1 | grep -i 'annotations'`

# Set/Update annotations
$ `kubectl annotate pod annotated-pod oncall='800-555-1212'`
$ `kubectl annotate pod annotated-pod oncall='800-555-2000' --overwrite`
## remove an annotation
`kubectl annotate pod annotated-pod oncall-`
## Remove annotations multiple pods
`kubectl annotate pod pod1 pod2 oncall-`
`kubectl annotate pod pod{1..3} oncall-`

## update via selector
- `kubectl annotate pod -l "ver in (v1,v2)" anno1=123`

## Annotate multiple pods
`kubectl annotate po nginx1 nginx2 nginx3 description='my description'`
or
`kubectl annotate po nginx{1..3} description='my description'`
