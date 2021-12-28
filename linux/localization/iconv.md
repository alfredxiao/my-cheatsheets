# List available charsets
- `iconv -l`

# Convert
- `cat file1 | iconv -f UTF-8 -t ASCII`

# Write output to a file
- `cat file2 | iconv -o file2`

# From source file
- `iconv -f ISO-8859-1 -t UTF-8 -o outputl.txt source.txt`
