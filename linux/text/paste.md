# concatenate lines on each file
* `paste f1 f2 f3` concatenate each line in f1 to f3

# using delimiter
* `paste -d : f1 f2`

# serial mode
* `paste -s f1 f2` would join all lines in `f1` into one line, then next line for lines from `f2`, etc.
* `paste -s -d ',' f1 f2` use comma as delimiter rather than the default tab
