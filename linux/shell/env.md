# `printenv`
- it prints environment variables

# `env`
- it by default does same thing as `printenv`
## run program with modified environment
- `env -i newcommand` run with empty environment
- `env -u VAR123 newcommand` run with `VAR123` environment variable removed

# Common variables
- `USER` – your current username.
- `SHELL` – the path to the current command shell (for example, /bin/bash).
- `PWD` – the current working directory.
- `HOSTNAME` – the hostname of the computer.
- `HOME` – your home directory.
- `MAIL` – the location of the user’s mail spool. Usually /var/spool/mail/USER.
- `LANG` – your current language.
- `TZ` – your time zone.
- `PS1` – the default prompt in bash.
- `TERM` – the current terminal type (for example, xterm).
- `DISPLAY` – the display used by X. This variable is usually set to :0.0, which means the first display on the current computer.
- `HISTFILESIZE` – the maximum number of lines contained in the history file.
- `EDITOR` – the user’s preferred text editor.
- `MANPATH` – the list of directories to search for manual pages.
- `OSTYPE` – the type of operating system.
- `_` - last executed command

# To make environment variable to all users in the system
- `vi /etc/environment`
  - add a line `MY_GLOBAL_VAR="My Awesome Value"`
  - will be available to new login sessions
- `vi /etc/bash.bashrc`
- `vi /etc/profile.d/myfile`
