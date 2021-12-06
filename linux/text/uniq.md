# default behaviour
* displays unique groups (multiple same adjacent lines considered one group)
  * `uniq myfile`

# display unique lines among entire file
* `cat myfile | sort uniq -u`

# display duplicated groups
* `uniq -d myfile`

# display duplicated lines
* `uniq -D myfile`

# ignore casing
* `uniq -i myfile`
