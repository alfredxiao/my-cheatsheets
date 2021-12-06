# meaning
* build and execute command lines from standard input
* literally, `xargs` = execute arguments

# multiple arguments
* `echo "dir1 dir2 dir3" | xargs mkdir`

# dealing with spaces
* `echo "'file 1' 'file2' 'file 3'" | xargs -d "\n" rm`

# on files from `find`
* `find /tmp -mtime +14 | xargs rm`
  * find files older then 14 days, remove each

# Displays commands to be executed
* `echo 'one two three' | xargs -t rm`

# Prompt before executing
* `echo 'one two three' | xargs -p touch`

# Multiple commands
* `echo 'one two three' | xargs -I % sh -c 'echo %; mkdir %'`

# vs. `find -exec`
* `find ./foo -type f -name "*.txt" -exec rm {} \;`
* `find ./foo -type f -name "*.txt" | xargs rm`
* In general `xargs` is much faster
