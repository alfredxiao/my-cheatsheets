# 1 `SIGHUP`
- The **SIGHUP** (“hang-up”) signal is used to report that the user's terminal is disconnected, perhaps because a network or telephone connection was broken.

# 15 `SIGTERM`
- The **SIGTERM** signal is a generic signal used to cause program termination. Unlike `SIGKILL`, this signal can be blocked, handled, and ignored. It is the normal way to politely ask a program to terminate.

# 9 `SIGKILL`
- The **SIGKILL** signal is used to cause immediate program termination. It **cannot** be handled or ignored, and is therefore always fatal. It is also not possible to block this signal.
- This signal is usually generated only by explicit request. Since it cannot be handled, you should generate it only as a last resort, after first trying a less drastic method such as C-c or SIGTERM. If a process does not respond to any other termination signals, sending it a SIGKILL signal will almost always cause it to go away.
- In fact, if `SIGKILL` fails to terminate a process, that by itself constitutes an operating system bug which you should report.
- A process's exit code due to `SIGKILL` can be different from a `SIGTERM` signal

# 3 `SIGQUIT`
- The **SIGQUIT** signal is similar to `SIGINT`, except that it’s controlled by a different key—the QUIT character, usually `Ctrl-\` and produces a core dump when it terminates the process, just like a program error signal. You can think of this as a program error condition “detected” by the user.
  - JVM prints thread dump when `kill -3 PID`

# 2 `SIGINT`
- The **SIGINT** (“program interrupt”) signal is sent when the user types the INTR character (normally `Ctrl-c`).

# 20 `SIGTSTP`
- This is **SIGTSTP**. Essentially the same as `SIGSTOP`. This is the signal sent when the user hits
`Ctrl+Z` on the terminal.
- a suspend signal which immediately places a process in suspension until (for example) a `fg` command is issued for the same process which will bring it back to the foreground.

# 19 `SIGSTOP`
- **SIGSTOP** is a signal sent programmatically (eg: `kill -STOP pid` )
- `SIGSTOP` cannot be ignored. `SIGTSTP` might be.

# 18 `SIGCONT`
- `kill -SIGCONT pid` resumes processes suspended by `SIGSTOP` or `SIGTSTP`

# Exit code
- when using `CTRL+c` at the command line the exit code will always be `130`, `137` when killed with `kill -9`, and `143` when `kill -15` was used.
