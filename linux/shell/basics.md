`~`
* `~` refers to home directory of current user
* `~user1` refers to home directory of user named `user1`

# alias
* `alias` lists all aliases
* to avoid using alias `ls`, you need to type `\ls`
  * by default `ls` is alias of `ls --color=auto`

# Variables vs Environments
* Variables are available to shell (and children shells), not processes started in the shell
* Environments are available to both shell and processes started in this shell
* `myvar=MyValue` defines a local variable
* `export myvar=MyValue` defines an environment

# Quoting
* Single quote
  * as is, no special interpretation
  * `'$ is $, $PATH is $PATH'`
* Double quote
  * Variables are substituted with their values
    * `echo D*` and `echo "D*"` are different
      * `echo D*` yields `Desktop Documents Downloads`
      * So actually **no quotes** is also a special form compared with single/double/backtick quoting
  * `"Path is $PATH"` - `$PATH` will be replaced with its value
  * If you want to escape the `$` sign, use `"Path variable is just \$PATH"` - where it is NOT substituted
  * backtick and `$()` ARE interpreted in double quotes
* Backtick
  * Replaced with command output
  * `echo "Today is `date`"` yields `Today is Wed Dec 30 21:58:08 UTC 2020`
  * Same as `echo "Today is $(date)"`

# Arithmetic
- `echo $((2*100+30))`
- `echo "Result is $((2*100+30))"`

# `;` vs `&&` when running multiple commands
* `;` runs regardless of outcome of previous command
* `&&` halts on failure, as it means logical AND

# `true` and `false`
* they are shell built-in commands
* `true` ends with code `0`
* `false` ends with code `1`

# `||`
- means logical OR, Depending on the result of the first command, the second command will either run or be skipped.
- `mkdir may-already-exists || true `
