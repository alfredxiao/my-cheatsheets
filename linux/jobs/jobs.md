# Send a job to background
* `mycmd &`
* another approach
  1. `mycmd`
  2. `Ctrl-Z`
  3. `bg %n` n is number from above `Ctrl-Z` output

# List background jobs
* `jobs`
* `jobs -l` shows PID of jobs

# Bring a job to foreground
* `fg %n` n is number from `jobs` output

# Stop a background job
* `kill %n` 
