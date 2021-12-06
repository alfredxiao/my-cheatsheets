# `shebang`
* `#!/bin/sh`

# To Read input
```
echo -n "What is your name? "
read NAME
```

# Arguments to your script
* `$0` for name of current script
* `$1` for first argument, `$2` for second, and so on
* `$@` represents arguments array (excluding name of current script)

Return value from script
* `exit 138`

Test conditions
* `[ condition ]` is same as `test condition`
* `test -f /mydir/myfile` file exists
* `test ! -f /mydir/myfile` file does not exists
* `test -d /mydir` dir exists
* `test -x 'apath/acmd'` a cmd is executable by the current user
* `test 1 -eq 1` numeric equality
* `test ! 1 -eq 1` numeric inequality
* `test 1 -ne 1` same as above
* `test 'a' = 'a' ` string equality
* `test 'a' != 'a' ` string inequality
* `test 1 -eq 1 -o 2 -eq 2`  OR (as `-o`)
* `test 1 -eq 1 -a 2 -eq 2`  AND (as `-a`)

`if`
* Forms
```
if [...]; then
else
fi
```

```
if [...]; then
elif
else
fi
```

`case`
```
case "$VAR" in
abc|def)
  echo "ABC or DEF"
  ;;
ghi)
  echo "GHI"
  ;;
*)
  echo "SOMETHING ELSE"
esac
```

`seq`
```
for num in `seq 1 10`
do
  touch file$num
done
```

Background Job
* `dosomething &`
  * output: `[1] 107`
    * `[1]`: job number (unique in this shell)
    * `107`: process id (unique in this OS)
* `jobs`: lists all jobs in current shell
* `fg %1` to bring job 1 to foreground
* To `PAUSE`/`suspend` a process: type `Ctrl-Z`
* `bg %1` to bring the `PAUSED` job to running in the Background
* To kill job number 1: `kill %1`

Stream Redirecting
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
