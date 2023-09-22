- not supported by all container runtime
- need to enable the feature gate

# Debugging with an ephemeral debug container
`kubectl run ephemeral-demo --image=k8s.gcr.io/pause:3.1 --restart=Never`
`kubectl debug -it ephemeral-demo --image=busybox:1.28 --target=ephemeral-demo`
  - where `ephemeral-demo` is an existing pod
  - this runs busybox as another container in the same pod with shared process namespace

# Debugging using a copy of the Pod
`kubectl run myapp --image=busybox:1.28 --restart=Never -- sleep 1d`
`kubectl debug myapp -it --image=ubuntu --share-processes --copy-to=myapp-debug`

# Copying a Pod while changing its command
`kubectl run --image=busybox:1.28 myapp -- false`
`kubectl debug myapp -it --copy-to=myapp-debug --container=myapp -- sh`
  - this changes command to an interactive shell

# Copying a Pod while changing container images
`kubectl run myapp --image=busybox:1.28 --restart=Never -- sleep 1d`
`kubectl debug myapp --copy-to=myapp-debug --set-image=*=ubuntu`
  -  `*=ubuntu` means change the image of all containers to ubuntu

# Debugging via a shell on the node
`kubectl debug node/mynode -it --image=ubuntu`
- The root filesystem of the Node will be mounted at `/host`
