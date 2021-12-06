# `sort`

# sort as string
* `sort myfile`

# sort as number
* `sort -n myfile`

# reverse order
* `sort -r myfile`

# write output to a file
* `sort myfile -o sortedfile`

# de-duplicate lines
* `sort -u myfile`

# sort by field
* `sort -t: -k3 mypasswd`
  * `-t:` defines separator as `:`
  * `-k3` or `-k 3` defines the third field is used for sorting
  * `-n` define sorting as a number field (rather than string)
* Multi-sort
  * `sort -k2 -k3n -k4 -t, myfile` sorts by field `2` first, then field `3` as number, then field `4`
