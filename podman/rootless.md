# Steps to run rootless (with host volume mount), runs as non-root inside container
```
mkdir -pv /home/student/local/mysql

# Create context
sudo semanage fcontext -a -t container_file_t '/home/student/local/mysql(/.*)?'
# Apply it
sudo restorecon -R /home/student/local/mysql

# Verify
ls -ldZ /home/student/local/mysql

# Run as non-root on the host, chown the directory to be mounted into the container,
# Using 27 as uid inside the container
podman unshare chown 27:27 /home/student/local/mysql
# this is similar to 'sudo chown 100026:100026 /home/student/local/mysql'

podman run --name persist-db  -d -v /home/student/local/mysql:/var/lib/mysql/data  -e MYSQL_USER=user1 -e MYSQL_PASSWORD=mypa55  -e MYSQL_DATABASE=items -e MYSQL_ROOT_PASSWORD=r00tpa55  registry.redhat.io/rhel8/mysql-80:1

# Check host
ls -ld /home/student/local/mysql/items
## owned by 100026 100026
ps -ef | grep mysqld
## started by User 100026

# Check container
ls -ld /var/lib/mysql
## owned by User 27 (mysql)
ps -ef | grep mysqld
## started by User 27 (mysql)

# Check Dockerfile
USER 27

```
If you run the processes within the container as a different non-root UID, however, then those processes will run as that UID. If they escape the container, they would only have world access to content in your home directory.

# Rootless with `root`
When not specifying a user as runner of processes inside a container, it would be a `root` (uid=0) inside the container, which maps back to your regular user in the host who runs `podman run` command to start the container.
This setup also means that the processes inside of the container are running as the user’s UID. If the container process escaped the container, the process would have full access to files in your home directory based on UID separation. SELinux would still block the access, but I have heard that some people disable SELinux. Doing so means that the escaped process could read the secrets in your home directory, like ~/.ssh and ~/.gpg, or other information that you would prefer the container not have access to.

- if it is the `root` of the host that runs `podman run`, the `root` inside container is mapped to the `root` of host

Podman defaults to mapping root in the container to your current UID (e.g. 1000) and then maps ranges of allocated UIDs/GIDs in /etc/subuid and /etc/subgid starting at 1.

By default, rootless Podman runs as root within the container. This policy means that the processes in the container have the default list of namespaced capabilities which allow the processes to act like root inside of the user namespace, including changing their UID and chowning files to different UIDs that are mapped into the user namespace.

# Summary
Running podman as |	With container process running as |	The actual UID visible on the host is
root	| root |	0
root	| non-root	| 0
non-root	| root	| Your UID
non-root	| non-root	| A non-root UID
