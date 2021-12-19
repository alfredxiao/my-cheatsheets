# Concepts
- `halt`: Stops all processes and shuts down the CPU.
- `poweroff`: Stops all processes, shuts down the CPU, and sends signals to the hardware
to power down.
- `reboot`: Stops all processes, shuts down the CPU, and then restarts the system.
- `shutdown`: Stops all processes, shuts down the CPU, and sends signals to the hardware to power down.

# Commands
## Common parameters
- `--halt`: Makes the command behave like halt.
- `-p`, `--poweroff`: Makes the command behave like poweroff.
- `--reboot`: Makes the command behave like reboot.
- `halt`
- `poweroff`
- `shutdown`
  - `shutdown -P` poweroff
  - `shutdown -H` halt
  - `shutdown -r` reboot
  - `shutdown -P now`
  - `shutdown -P 20:00` shutdown at 20:00
  - `shutdown -P +2` shutdown in 2 minutes
  - `shutdown -P +5 "Shutdown in 5 mins, please backup and exit!"`
  - `shutdown -k` do not shutdown, just write wall message
- `reboot`
