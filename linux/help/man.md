man sections
Section # 1 : User command (executable programs or shell commands)
Section # 2 : System calls (functions provided by the kernel)
Section # 3 : Library calls (functions within program libraries)
Section # 4 : Special files (usually found in /dev)
Section # 5 : File formats and conventions eg /etc/passwd
Section # 6 : Games
Section # 7 : Miscellaneous (including macro packages and conventions),
Section # 8 : System administration commands (usually only for root)
Section # 9 : Kernel routines [Non standard]

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
