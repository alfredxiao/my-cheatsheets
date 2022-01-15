* if you have a password protected id key file and wants to avoid entering password
1. run `ssh-agent /bin/bash` wraps a bash and this handles authentication for keys that requires password
2. then `ssh-add ~/.ssh/id_rsa` to adds the passwords to `ssh-agent`
3. then `ssh user1@host1`

- However, if you exit the agent session and start it up again, you must readd the key and reenter the passphrase.
