# Types of Shells (in terms of login)
## interactive login shell
- most normal shell, can find out this by `echo $0`, where `-` in `-bash` indicates it is a login shell
- `login` means `.profile` stuff will be loaded
## interactive non-login shell
- running `bash` from a already running shell starts an interactive but non-login shell
- `non-login` means `.profile` stuff will not be loaded
## non-interactive login shell
- `bash --login myscript.sh`
- `--login` tells the shell that it is an **already logged in** shell, please load `/.profile` and other stuff to setup environment
## non-interactive non-login shell
- `bash myscript.sh` run within an interactive shell will start a new shell
- in the new shell, no user interaction, no loading of `.profile` stuff

# Login/Init procedure
## Order of init files for Login Shell
1. `/etc/profile` and `/etc/profile.d/`
2. `$HOME/.bash_profile`
3. `$HOME/.bash_login`
4. `$HOME/.profile`
- Most linux distribution use only one file from 2/3/4
## For Non-Login Shell
- It **will** run `$HOME/.bashrc` but not the `profile` init scripts
  - then it runs `/etc/bashrc`
- You can provide an init script to run
  - `BASH_ENV=/path/to/myscript bash mycommand`
  - where `mycommand` would see changes run by `myscript`
