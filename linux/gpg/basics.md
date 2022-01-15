To begin using GPG, first need to geneate a key pair

# Generate Key Pair
- `gpg --gen-key`
- GPG will ask a series of questions, including your full name and email address, as well as a passphrase. You’ll need to remember the **email address** and the **passphrase** because the address identifies your public key, whereas the passphrase allows you to use your private key.
- After the keys are generated, they are stored in a file, which is called your keyring in the `~/.gnupg/` directory.

# Export public key
- `gpg --export EMAIL-ADDRESS > FILENAME.pub`
- or `gpg --armor --export EMAIL-ADDRESS > FILENAME.pub.ascii` for a non-binary file

# Import someone's public key
- `gpg --import FILENAME.pub`

# List keys
- `gpg --list-keys`

# Encrypt a file
- `gpg --out secret.file --recipient EMAIL-ADDRESS --encrypt input.file`

# Decrypt a file
- `gpg --out plain.file --decrypt secret.file`

# Encrypt and Sign a file
1. `gpg --out tosign.file --recipient EMAIL-ADDRESS --encrypt input.file`
2. `gpg --output signed.file --sign tosign.file`
  - if one wants to generate the signed file in ASCII, use `--clearsign` option instead

# Decrypt and Verify a Signed file
1. `gpg --verify signed.file` which should output `Good signature from`
2. `gpg --out message.gpg --decrypt signed.file`
3. `gpg --out plain.file --decrypt message.gpg`

# Detached signature
1. `gpg --output doc.sig --detach-sig doc`
2. `gpg --verify doc.sig doc` (both files must be present to verify)

# Revoking a key pair
1. `gpg --out key-revocation.asc --gen-revoke cb1234@ivytech.edu`
2. `gpg --import FILENAME.asc` (run by everyone who imported the key idenfied by this email)
