applies to `less`, `more`, `vi`, `man` pages

* next page: `f`, `SPACE`
* previous page: `b`





`tr`

`comm`
* for comparing two sorted files
* output has three columns,
  * first column: lines exist only in file1
  * second: lines exist only in file2
  * third: shared lines in file1 and file2
* to not display one ore more columns
  * `comm -23 file1 file2` to not display second or third column

`expand`
* converts `TAB` into blank keys, default to 8 blanks?

`unexpand`
* converts spaces into `TAB`

`fold`
* wraps each line into specified width, e.g. 80 (which is default)
* `fold -w 80 myfile`
* similar to `fmt` command



`xxd`
* `xxd file1` display hex code for a file

`paste`
