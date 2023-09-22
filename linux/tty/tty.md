type `tty` command to see which terminal you are in
`/dev/pts/0` meaning this is a pesudo terminal

# list logged in terminals
- `w`
```
09:40:28 up 4 min,  3 users,  load average: 0.09, 0.09, 0.04
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
alfred   tty1     -                09:35    4:44   0.04s  0.02s -bash
alfred   tty7     :0               09:36    4:58   0.82s  0.14s xfce4-session
alfred   tty6     -                09:36    3:47   0.05s  0.02s -bash

```
- `who`
```
alfred   tty1         2022-02-06 09:35
alfred   tty7         2022-02-06 09:36 (:0)
alfred   tty6         2022-02-06 09:36
```
- where you can see `:0` indicates graphical session

`sudo chvt 7` to switch to tty7
