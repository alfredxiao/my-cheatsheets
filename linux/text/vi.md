# How to quit
* quit and not save `:q!`
* save and quit : `:qw`  `Zz` `:x` `:x!`
* `w` save but not quit
* `ESC` to exit editing mode
* `:e!` to discard changes and reload file

# Move cursor
* `w` forward one word
* `b` backward one word
* `e` forward to end of word
* `0` or `^` begin of line
* `$` end of line
* `gg` first line of the file
* `G` last line of the file
* `5G` move to line 5
* `Ctrl-f` page down
* `Ctrl-b` page up

# Search text
* `/` forward search
* `?` backward search
* `n` next match
* `N` previous match

# Delete
* `dd` delete one line
* `5dd` delete 5 lines
* `d$` or `D` delete from cursor to end of line
* `dw` delete a word
* `d3w` delete 3 words
* `x` delete one character
* `3x` delete 3 characters

# Modify
* `i` to start insert mode
* `I` to start insert mode at beginning of line
* `a` to start insert mode after cursor
* `A` to start insert mode at end of line
* `o` to start insert mode by appending new line
* `O` to start insert mode by prepending new line
* `j` join two lines
* `3j` join three lines
* `p` paste after cursor
* `r` to replace current character
* `~` to alter case (from upper to lower, or the other way around)
* `5~` to alter case for next 5 characters

# Find and Replace
* `:%s/abc/newabc/g` to replace `abc` with `newabc`

# Copy
* `yy` copy a line
* `yw` copy a word
* `y3w` copy 3 words

# copy a file's content
* `:r file1`

# Undo
* `u` undo once
* `2u` undo twice

# Shell
* `:! ls /` execute `ls /`, display result, come back to editor, do not change buffer
* `:r! ls /` execute `ls /`, copy result to current editor buffer

# Tutorial
* `vimtutor`
