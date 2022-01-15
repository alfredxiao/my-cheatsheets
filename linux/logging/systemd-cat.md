similar to `logger`,
# create entries manually
-`echo "my message" | systemd-cat`

# Specifiy a severity
- `echo 'info' | systemd-cat -p info`
- `echo 'hello' | systemd-cat -p warning`
- `echo 'hello' | systemd-cat -p emerg`

# Specify a tag
- `echo 'hello' | systemd-cat -t someapp -p emerg`
