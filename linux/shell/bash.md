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


# `if`
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

# `case`
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

# `for`
## enumeration
```
for color in "Red" "Green"
do
  echo "Color is: $color"
done
```
## `seq`
```
for num in `seq 1 10`
do
  touch file$num
done
```
## range
```
for i in {1..5}
do
  echo "ID-$i"
done
```
## iterate over command output
```
for file in $(ls | sort)
do
  echo "File is: $file"
done
```

# `while`
```
while [ condition ] ; do
  commands
done
```

# `until`
```
until [ CONDITION ]
do
  COMMANDS
done
```

# Functions
## Forms
- with `function` keyword
```
function fname {
  commands
}
```
- without `function` keyword
```
fname() {
  commands
}
```
## Use a function
- `fname` and **NOT** `fname()`
## Return a value and check it
```
fname() {
  return 123
}

fname
echo "Return value is $?"

# alternatively
output=$(fname)
```

Stream Redirecting


# Arithmetic
- `((result=n*n))`
- ``
