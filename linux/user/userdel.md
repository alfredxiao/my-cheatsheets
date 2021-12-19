* deletes the account, but not home directory: `userdel auser`
* to delete account and home directory as well as mail spool: `userdel -r auser`
  * this does not delete files outside of the user home directory
