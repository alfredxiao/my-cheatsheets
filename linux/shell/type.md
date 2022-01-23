use `type` to determine if a command is
* alias
* shell builtin
* regular program

* `type mycmd` to tell type of a command (is it a shell built-in, alias, external command file, etc.)
* `type -a mycmd` to list all (aliases and path to the command, for example)

# brief mode
- `type -t ls` --> `alias`
