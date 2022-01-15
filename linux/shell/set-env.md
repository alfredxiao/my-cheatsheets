* `set` gives both environment and shell variables.
* `set` is a shell builtin
* `env` is a standalone program
* `env` only gives environment variables, which are passed along to child processes.
* `unset` to clear a variable value
  * `unset -f myfunc` to unset a function

# `set` used in scripts
* `set -x` to enable all commands being printed
* `set +x` to disable it

- It’s extremely **important** that there are **no spaces** between the environment variable name, the equal sign, and the value.
  - `var2=value2`
