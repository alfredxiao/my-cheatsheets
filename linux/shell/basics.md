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
- or `echo $[ 2*100 + 30 ]`
- **Note** this form accepts integer only, not floats

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


# Find info about current shell
- `echo $0` -> current shell name, `-bash` means login shell
- `cat /proc/$$/cmdline`
- `ps -f` -> current shell process info
- `readlink /proc/$$/exe` -> current shell executable name
- `echo $SHELL` shell set for current user, **not necessarily** the current running shell
- `cat /etc/shells` all valid login shells currently installed
- `shopt login_shell` tells if current shell is login shell or not

# Shell options
- `echo $-` displays current shell options
- `h` - Remember the location of commands as they are looked up
- `i` - It means shell is interactive
- `m` - Job control is enabled
- `B` - The shell will perform brace expansion
- `H` - Enable ! style history substitution. This flag is on by default when the shell is interactive.
- `s` - Commands are read from the standard input device such as keyboard. This option allows the positional parameters to be set when invoking an interactive shell or when reading input through a pipe.
