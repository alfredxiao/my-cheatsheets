* `1` for STDOUT, `2` for STDERR, `0` for STDIN
* `>` by default, redirects STDOUT
  * to modify, e.g. `2>` redirects STDERR
  * Example: `ls /fake 2> stderr.txt`
* `&>` redirects both STDOUT and STDERR
  * e.g. `ls no-such-file.txt &> combine.txt` sends both stdout and stderr to `combined.txt`
* `>&` adds to
  * e.g. `a-command > combined.txt 2>&1` where `2>&1` means send `2` to where `1` is
* To redirect stdout and stderr to different files:
  * `ls /fake /etc/ppp > example.txt 2> error.txt`
  * the order of which stream goes first does not matter
* `<` to send a file to a command as STDIN
  * e.g. `tr 'a-z' 'A-Z' < example.txt`
