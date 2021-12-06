* cpio takes the list of files from the standard input while creating an archive, and sends the output to the standard output.

# Create archive
- `ls infile* | cpio -ov > out.cpio`
  - `-o` for output
  - `-v` for verbose

# Extract
- `cpio -idv < out.cpio`
  - `-i` for extract
  - `-d` for making directories if needed
- `cpio -idv -F out.cpio`
- `cpio -idv -F out.tar`
- `cpio -itvI out.cpio`

# Extract without preserving original fullpath
- `cpio -iv --no-absolute-filenames -I out.cpio`

# View list of files of an achieve
- `cpio -it -F out.tar`
- `cpio -itvI out.cpio`
