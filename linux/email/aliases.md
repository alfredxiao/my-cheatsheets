locaton `/etc/aliases`

Run the `newaliases` command to update the aliases database, `/etc/aliases.db`.
- newer distribution does not seem to have this .db file (no need, right?)

The format of the alias records in the `/etc/aliases` file is
`ALIAS-NAME: RECIPIENT1[,RECIPIENT2[,...]]`

in `postfix`, `sendmail -I` is same as running `newaliases`
