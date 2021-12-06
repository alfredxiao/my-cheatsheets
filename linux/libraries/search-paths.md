When a program is using a shared function, the system will search for the function’s library file in a specific order; looking in directories stored within the

1. `LD_LIBRARY_PATH` environment variable
2. Program's `PATH` environment variable
3. `/etc/ld.so.conf.d/` folder
4. `/etc/ld.so.conf` file
5. `/lib*/` and `/usr/lib*/` folders

Be aware that the order of #3 and #4 may be flip-flopped on your system.

It is important to know that the `/lib*/` folders, such as `/lib/` and `/lib64/`, are for libraries needed by system utilities that reside in the `/bin/` and `/sbin/` directories, whereas the `/usr/lib*/` folders, such as `/usr/lib/` and `/usr/lib64/`, are for libraries needed by additional software, such as database utilities like MariaDB.
