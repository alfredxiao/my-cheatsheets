# Basics
- `trap` allows you to catch signals and execute code when they occur.
- There is a bash-generated psuedo-signal named EXIT
- It does not really trap those signals that cannot be handled, like `SIGKILL`

# Usage
```
function cleanup() {
  echo "this is run when you hit Ctrl-C"
}

trap cleanup SIGINT

sleep 900
```

#  Trick
```
# Run something important, no Ctrl-C allowed.
trap "" SIGINT
important_command

# Less important stuff from here on out, Ctrl-C allowed.
trap - SIGINT
not_so_important_command
```
