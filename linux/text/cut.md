
# cut by a delimiter
* `cut -d: -f1,5-7 mypasswd`
  * `-d:` specifies delimiter as `:`
  * `-f` specifies what field(s) to display, in this case, `1` and `5,6,7`

# cut by character position
* `cut -c1,49-`
  * `-c` specifies positions characters are at for displaying
  * this case, we display the first, then 49 onwards

# Single character
* Note that delimiter must be a single character
