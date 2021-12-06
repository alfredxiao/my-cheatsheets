# `od`
* display binary file content in octal and other formats

# Character
* `od -c myfile` print each character

# Octal
* `od myfile`
* `od -t o1 myfile` print octal code by each byte (otherwise per two bytes)

# Decimal
* `od -t d1 myfile`

# Hex
* `-t x1` to print hex code by each byte (otherwise, defaults to each two bytes)
```
echo "ABC" | od -ct x1
0000000   A   B   C  \n
         41  42  43  0a
0000004
```
