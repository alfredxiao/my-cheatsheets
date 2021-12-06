# replace (one to one)
* `cat afile | tr 'abc' 'xyz'` would replace all `a` with `x`, `b` with `y`, etc.
* `tr 'a-z' 'A-Z'` basically convert to uppercase

# squeeze
* e.g. `tr -s ' '` convert consecutive spaces into one space
* e.g. `tr -s ' ' '_'` convert consecutive spaces into one underscore

# change casing
* `tr '[:lower:]' '[:upper:]'` convert from lower to upper case

# remove
* `echo "my phone is 123-456-7890" | tr -d [:digit:]` remove all digits

# Windows text to Linux text file
* `tr -d '\015' < filename.dos > filename.linux`
