* `set` gives both environment and shell variables.
* `set` is a shell builtin
* `env` only gives environment variables, which are passed along to child processes.
* `env` is a standalone program
* `unset` to clear a variable value
  * `unset -f myfunc` to unset a function

# `set` used in scripts
* `set -x` to enable all commands being printed
* `set +x` to disable it
