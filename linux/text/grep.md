# gives count only
* `grep -c ABC myfile.txt`

# multiple files
* `grep ABC file1 file2`

# Multiple patterns
- `grep -e 'abc' -e 'x.z'`

# highlight found
* `grep --color ABC myfile.txt`

# To give exact line number in original file
* `grep -n ABC myfile.txt`

# To negate the matching
* `grep -v ABC myfile.txt` lines not containing ABC

# To ignore case
* `grep -i abc myfile.txt`

# Whole word only
* `grep -w abc myfile.txt`

# Using grokking
- `grep [AJ] myfile` find either `A` or `J`

# using regex
* `grep '[^0-9]' myfile.txt` for standard regex
* `grep -E 'e+' myfile.txt` for extended regex
  * Extended Regex: allows the use of `?` `+` and `|`
* `grep -Ev '^#|^$' afile` removes empty or comment lines

# Don't use regex
- `grep -F "US$" myfile`

# Recursive directory
* `grep -R abc mydir`
  * same as `-r`

# Load patterns from a file
* `grep -f patterns.txt myfile`

# quote mode (for scripting):
* `grep -q abc myfile.txt`
  * scripts use exit status code to find out if matches are found or not

# Context Lines
- `grep -B 5 abc myfile.txt` print 5 lines **Before** the matching line
- `grep -A 5 abc myfile.txt` print 5 lines **After** the matching line
- `grep -C 2 abc myfile.txt` print 2 lines **Both** before and after the matching line

# grep family
* `egrep` is `grep -E`
* `fgrep` is `grep -F` meaning there is no regrex
* `rgrep` is `grep -r`
* `pgrep` is for finding process ids
