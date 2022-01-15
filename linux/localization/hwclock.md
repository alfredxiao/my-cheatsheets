# Basics
- Displays or sets the time as kept on the internal BIOS or UEFI clock on the workstation or server
- There are two clock concepts
  - RTC (Real Time Clock), as kept by `hwclock`
  - Linux system time, as kept by `date` command

# Write hwclock with Linux system time
- `hwclock -w` same as `--systohc`

# Write Linux system time with hwclock time
- `hwclock -s` same as `--hctosys`
