# `?`
* means exactly **ONE** character

# `[]`
* `ls a[123]b` would match `a1b` `a2b` and `a3b`
* `ls a[1-3]b` would match the same as above
* `ls a[^123]b` would match anything like `a?b` but middle char is not one of `123`
  * `[^123]` is the same as `[!123]`

# Using '{}'
- `ls a{1,2,3}b` would match the same as above
- or `ls a{1..3}b`
- or `ls a[1-3]b`
- `mkdir projects/{data,software,document}` make three directories

# Expanding
* Running `xxx *.txt` in the shell ends up running `xxx a.txt b.txt c.txt`
  * Because it is the shell that interprets `*`, not `xxx` command
  * It does NOT matter what command `xxx` is, be it `ls`, `java`, same
* To pass `*` character to command, it has to be quoted
