`ed`
* `ed -p: myfile` - use `:` as command prompt
* to print current line number (`ed` points to last line right after command starts)
  * `:.=` prints the line number
  * `:.` or `:.p` prints the line
* to print multiple lines
  * `:1,$p` prints line from line `1` to `$` which stands for last line
* to goto specific line
  * `:3` goes to line 3
* to search a line containing some text
  * `:/cdrom` search for `cdrom`, it goes to first hit searched from current line on
  * You can keep searching by retyping the same command
* delete current line
  * `:d`
* To replace some text with another set of text (within current line)
  * `:s/oldtext/newtext`
* To insert a line (before) current line
  ```
  :i
  (type new line content here) ENTER
  .  (type a single period)
  :
  ```
  * after this line inserted, all number numbers after this new line will be increased
* to insert a line (after) current line
  * `:a`
* To save changes and quite
  * `:wq`
* Quit without saving
  * `:Q`
