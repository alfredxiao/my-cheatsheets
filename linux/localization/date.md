# Display time and timezone
- `date`

# Date yyyy-MM-dd
- `date +%F`

# show epoch
- `date +%s`

# Date MM/dd/YY
- `date +%D`
- or `date +%m/%d/%y`

# Other formatters
- `%Y` is year in four digit
- `%a` Mon/Tue/...
- `%A` Monday/Tuesday/...

# Date in UTC
- `date -u`

# Convert epoch seconds
- `date --date='@2147483647'`

# Show date in a different timezone
- `TZ='America/Los_Angeles' date`

# Set Date time (for current session)
- `date MMDDhhmm`
- `date MMDDhhmm[[CC]YY][.ss]`
- `date +%Y%m%d -s "20081128"` set date according to provided format
- `date +%T -s "10:13:13"` set time
