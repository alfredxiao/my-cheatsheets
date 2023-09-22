to drain a node
`kubectl drain node-name`
- add `--ignore-daemonsets` if there are daemonsets running on the node
  - the draining process won't remove those daemonsets, but without this parameter, it won't start draining
- add `--force` if there are pods without controllers running on the node, e.g. static pods, and pods started by `kubectl run`

after draining
- `Ready,SchedulingDisabled` is the new status

bring it back
`kubectl uncordon node-name`
- note that those pods evicted and started again on other nodes during the draining process won't automatically be scheduled back on this node that is brought back available for scheduling
