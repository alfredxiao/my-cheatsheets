# Basics
* `[ condition ]` is same as `test condition`
* `test 1 -eq 1` numeric equality
* `test ! 1 -eq 1` numeric inequality
* `test 1 -ne 1` numeric inequality
* `test 'a' = 'a' ` string equality
* `test 'a' != 'a' ` string inequality
* `test 1 -eq 1 -o 2 -eq 2`  OR (as `-o`)
* `test 1 -eq 1 -a 2 -eq 2`  AND (as `-a`)

# Numeric Comparison
- `-eq` equals
- `-ne` not equals
- `-ge` >=
- `-gt` >
- `-le` <=
- `-lt` <
## When inside `(())`
- `if ((x>y)) && ((y>0))`
- `((z=(x>y)?100:200))`

# String Comparison
- `=` equals **Note** there must be a space before and after `=`
- `!=` not equals
- `-n str1` str1 has non zero length
- `-z str1` str1 has zero length
## only supported in `[[`
- `[[ 'Unix System' == *'Unix'* ]]` test substring
- `[[ string1 =~ regex ]]` test matching regex
  - e.g. `[[ 'Linux System' =~ .*Linux.* ]]`

# File Test
- `test -e /path/to/file` if file exists (actually file or directory)
* `test -f /mydir/myfile` file exists and is regular file
* `test -d /mydir` dir exists
* `test ! -f /mydir/myfile` file does not exists
* `test -x /path/to/file` a cmd is executable by the current user
- `test -h /path` is symbolic link
- `test -p /path` if pipe exists
- `test -r /path` if file readable
- `test -s /path` file exists and not empty
- `test -t FD` if FD is opened on a terminal
- `test -O /path` if file effectively owned by you
- `test -G /path` if file effectively owned by your group
- `test file1 -nt file2` if file1 is newer than file2
- `test file1 -ot file2` if file1 is older than file2

# `[` vs `[[`
## Quote strings when using `[`
- `[ $myname = $yourname ]` WRONG!
- `[ "$myname" = "$yourname" ]` correct
## When using `<` and `>`
- either `[[ $name1 > $name2 ]]` but Quoting is usually going to give you the behavior that you want, so make it a habit
- or `[ "$name1" \> "$name2" ]`
- `>` and `<` have special meaning to bash, so they should be escaped

# `&&` and `||`
- only supported in `[[`
- `[[ EXPR && EXPR ]]` and `[[ EXPR || EXPR ]]`
- or `[ EXPR -a EXPR ]` and `[ EXPR -o EXPR ]`
