# meaning
* build and execute command lines from standard input
* literally, `xargs` = execute arguments

# multiple arguments
* `echo "dir1 dir2 dir3" | xargs mkdir`
  - this yields `mkdir dir1 dir2 dir3` (rather than `mkdir dir1` and `mkdir dir2` etc.)
- `echo "a b c" | tr -d "\n" | xargs -d " " -t touch`
  - this yields `touch a b c`
- `echo "a b c" | xargs -d "\n" touch`
  - this yeilds `touch "a b c"`

# dealing with spaces
* `echo "'file 1' 'file2' 'file 3'" | xargs -d "\n" rm`

# on files from `find`
* `find /tmp -mtime +14 | xargs rm`
  * find files older then 14 days, remove each

# Displays commands to be executed
* `echo 'one two three' | xargs -t rm`

# Prompt before executing
* `echo 'one two three' | xargs -p touch`

# Use variable/placeholder
* `echo 'one two three' | xargs -I % sh -c 'echo %; mkdir %'`
- `grep -l “junk” *.txt | xargs -I {}  mv {} test/bak/`
- `[command] | xargs -I @  echo "'@'"`

# NULL padded input
- `find . -name "*.txt" -print0 | xargs -0 touch`
  - `-print0` prints output to be NULL separated rather than space separated
  - this is useful if input (filenames in this case) has spaces inside them

# Limit number of items in each execution
- `echo "a b c d e" | xargs -n 3 touch`
  - yields `touch a b c` and `touch d e`
- `[command] | xargs -n 1 -P [p] sleep` can execute in parallel

# Suppress if input is empty
- `echo "    " | xargs -r -p touch` will not do anything

# vs. `find -exec`
* `find ./foo -type f -name "*.txt" -exec rm {} \;`
* `find ./foo -type f -name "*.txt" | xargs rm`
* In general `xargs` is much faster

- `find . -empty -type f | xargs rm -f`
    - `find … -exec` handles files one at a time, which is slower
- `find . -name “*.sh” | xargs ls -al > scripts.txt`

- `cut -d, -f1 file.csv | xargs touch`
