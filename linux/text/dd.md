# Create file of specified size
`dd if=/dev/zero of=file.txt bs=1024 count=10240 `

`head --bytes 300K /dev/zero > file3.txt`

`tail --bytes 1G /dev/zero > file4.txt`

# Page Cache
## Advise to drop cache for whole file
`dd if=ifile iflag=nocache count=0`

## Ensure drop cache for the whole file
`dd of=ofile oflag=nocache conv=notrunc,fdatasync count=0`

## Drop cache for part of file
`dd if=ifile iflag=nocache skip=10 count=10 of=/dev/null`

## Stream data using just the read-ahead cache
`dd if=ifile of=ofile iflag=nocache oflag=nocache`
- this seems to have no cache effect on input file but output file still goes to page cache
