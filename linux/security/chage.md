# Basics
* change user password expiry information (by modifying `/etc/shadow`)
* by default prompts interactively if no arguments provided

# View Info
- `chage -l user1`

# Set Account Expiry Date
- `chage -E 2020-10-25 user1`

# set number of days before warning password expiry
- `chage -W 14 user1`

# Force user to change password on next login
- `chage -d0 user1`
