# Basics
- two ways to change system localization
- one by changing environment variables manually
- the other is by `localectl`

# Set Var
- Instead of having to change all of the `LC_` environment variables individually, the `LANG` environment variable controls all of them at one place: `export LANG=en_GB.UTF-8`
- Some Linux systems require that you also set the `LC_ALL` environment variable, so it’s usually a good idea to set that along with the `LANG` environment variable.
- This method changes the localization for your current login session. If you need to permanently change the localization, you’ll need to add the export command to the `.bashrc` file in your `$HOME` folder so that it runs each time you log in.

# Current Value
`localectl`

# List installed locales
`localectl list-locales`

# Set locale
`localectl set-locale LANG=en_GB.utf8`
