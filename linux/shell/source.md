With `source` itself being a shell builtin command. And `source` and the `.` operator being synonyms.
A dot in that context means to "source" the contents of that file into the current shell.

e.g. `source test.sh` or `. test.sh` would just like having the scripts running in current shell, where the content of `test.sh` does not need shebang, nor does the file need `x` permission
