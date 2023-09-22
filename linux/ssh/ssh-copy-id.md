dry run
`ssh-copy-id -n user1@host1`

actually same as `cat YOUR_PUBKEY >> ~/.ssh/authorized_keys`

# Copy it to a remmote machine
- `cat id_rsa.pub | ssh user1@10.10.10.10 "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"`
