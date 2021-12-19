# Found out who has text status open (can receive message)
- `who -T`
  - `+` means a user tty can receive message
- `mesg`
  - `is y` means you current terminal can receive message
- `tty` returns current terminal

# Send message
- `wall "Message"`
- `wall` Enter, then input text lines, then `Ctrl-D` to finish
