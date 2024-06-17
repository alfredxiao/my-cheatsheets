# Create Cluster
- `docker swarm init --advertise-addr 172.31.31.158:2377 --listen-addr 172.31.31.158:2377`
- this node becomes first manager node
- also outputs command to join (as worker)
  - `docker swarm join --token TOKEN1 172.31.31.158:2377`

# Print Manager Join command
- `docker swarm join-token manager` (on first manager node)
- it outputs something like
 - `docker swarm join --token TOKEN2 172.31.31.158:2377`

# Join more manager Nodes
- `docker swarm join --token TOKEN2 172.31.31.158:2377 --advertise-addr 172.31.19.76:2377 --listen-addr 172.31.19.76:2377`
- `docker swarm join --token TOKEN2 172.31.31.158:2377 --advertise-addr 172.31.26.234:2377 --listen-addr 172.31.26.234:2377`

# Join more worker nodes
- `docker swarm join --token TOKEN1 172.31.31.158:2377 --advertise-addr 172.31.20.50:2377 --listen-addr 172.31.20.50:2377`
- `docker swarm join --token TOKEN1 172.31.31.158:2377 --advertise-addr 172.31.23.136:2377 --listen-addr 172.31.23.136:2377`

docker service create --name test \
   --network uber-net \
   --replicas 3 \
   alpine sleep infinity
