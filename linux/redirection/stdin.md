# std input from file
* `uniq < myfile`

# Note the difference
- `wc test.txt`         —> empty stdin, `wc` loads content from file
- `wc < test.txt`      —> stdin from file content
- `cat test.txt | wc` —> input of ‘wc’ is from stdout of ‘cat’
