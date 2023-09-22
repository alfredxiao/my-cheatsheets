locaton `/etc/aliases`

Run the `newaliases` command to update the aliases database, `/etc/aliases.db`.
- newer distribution does not seem to have this .db file (no need, right?)
- for `postfix`, you can run `postalias /etc/postfix/aliases`

The format of the alias records in the `/etc/aliases` file is
`ALIAS-NAME: RECIPIENT1[,RECIPIENT2[,...]]`

in `postfix`, `sendmail -I` is same as running `newaliases`

- **NOTE** that mail setup this way, the mail will be sent to target alias instead of the source alias, the source alias will **NOT** receive the email
- We can also setup 'fake' account, which does not have real user account backing up, and use these 'fake' accounts to behave as distribution list, e.g.
  - `urgent: user1,user2,user3`
