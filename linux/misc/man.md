`man` page
* if a command/file/etc. has multiple man pages, e.g. passwd is a command but also a file that has its own and different man page
  * `man -f passwd` lists all man pages
    * this is the same as `whatis passwd`
  * `man 5 passwd` list the man page for `passwd` at section/category `5`
    * same as `man -s 5 passwd` or with `-S`
* search for a command or its short description
  * `man -k xxx`
  * this is the same as `apropos` on most linux distributions

`apropos`
* search for a command or its short description
* same as `man -k xxx`
