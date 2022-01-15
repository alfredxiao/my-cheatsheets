# Send a job to background
* `mycmd &`
  * output: `[1] 107`
    * `[1]`: job number (unique in this shell)
    * `107`: process id (unique in this OS)
* another approach
  1. `mycmd`
  2. `Ctrl-Z` which stopps the job
  3. `bg %n` where n is number from above `Ctrl-Z` output

# List background jobs
* `jobs`
* `jobs -l` shows PID of jobs

# Bring a job to foreground
* `fg %n` n is number from `jobs` output

# Stop a background job
* `kill %n`

# Show PID only
- `jobs -p`

# Show Running jobs only
- `jobs -r`

# Show stopped jobs only
- `jobs -s`

# Default Job and Second Default Job
- in the output of `jobs` command
- `+` means default job
- `-` means second default job
- It would be the job referenced by any job control commands if a job number wasn’t specified in the command line.
- The job with the minus sign is the job that would become the default job when the current default job finishes processing.
- There will only be one job with the plus sign and one job with the minus sign at any time, no matter how many jobs are running in the shell.
