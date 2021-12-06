This is for managing the library cache. it is automatically run when we use package manager, but for any applications you are developing yourself, you’ll have to manually run the `ldconfig` command.

# Process for developing new library
1. testing:
  * export `LD_LIBRARY_PATH` to include a path to your library file
2. installation:
  * add your library file to `/usr/lib*` directory tree
  * create a library config file within `/etc/ld.so.conf.d/` and points to library file location
3. config
  * use `ldconfig` to update library cache

# Listing file cache
* `ldconfig -v`
