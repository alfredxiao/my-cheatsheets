# create new group
- `groupadd group1`

# Providing group id
- `groupadd -g 1005 rnd` create group with group id specified (if not, will auto assign one)

# Create system group
* `groupadd -r sales` creates a system group (GID < 500 or 1000)
