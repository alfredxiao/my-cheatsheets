# Basics
- changes label of an ext2/3/4 file system

# Display current label (or volume name)
- `e2label /dev/sda1`

# Set label
- `e2label /dev/sda1 MYDATA001`
- or `tune2fs -L MYDATA001 /dev/sda1`

# Remove a label
- `e2label /dev/sda1 ""`
- or `tune2fs -L "" /dev/sda1`
