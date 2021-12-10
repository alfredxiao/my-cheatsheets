# `?`
* means exactly one character

# `[]`
* `ls a[123]b` would match `a1b` `a2b` and `a3b`
* `ls a[1-3]b` would match the same as above
* `ls a[^123]b` would match anything like `a?b` but middle char is not one of `123`

# Using '{}'
- `ls a{1,2,3}b` would match the same as above
- `mkdir projects/{data,software,document}` make three directories
