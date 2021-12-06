# basic flow
* `somecommand | sed 'editcommand' | anothercommand`
  * meaning `sed` is not interactive but mostly scripted
  * e.g. `cat file1 | sed "2 s/old/new/"`

# replace all occurences
* `sed 's/UNIX/Linux/' myfile`
* same as `sed 's/UNIX/Linux/g' myfile`

# delete lines containting some text
* `sed '/sometext/d' myfile`

# delete lines NOT containing some text
* `sed '/sometext/!d' myfile`

# search text
* `sed -n /sometext/p myfile`
  * `-n` to disable default display
  * `/p` for print (meaning display)

# multiple scripts
* `sed -e '/windows/d' -e '/java/!d'` to delete all lines containing windows, then delete lines not containing java
* `sed -e 's/cake/donuts/ ; s/like/love/' cake.txt`

# delete lines
* `sed '1d'` deletes the first line
* `sed '3,$d'` deletes line 3 all the way to the end
* `sed '2!d'` deletes all lines except the second line
* `sed '$d'` deletes the last line

# To apply a sed script
* `sed -f myscript`

# To edit a file in place
* `sed -i '1d' myfile` deletes first line from myfile
