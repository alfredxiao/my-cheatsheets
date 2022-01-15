# Read into a variable
- `read myvar`

# Prompt
- `read -p "Please enter your name:" name`

# Multiple
- `read -p "Enter fullname:" first last`
- when enters `Alfred Xiao`, `$first` becomes `Alfred`, `$last` becomes `Xiao`

# Read into REPLY
- `read`
- input is then stored in `$REPLY`

# Specify a timeout
- `read -t 10 -p "Enter your name (in 10 seconds):"`
- when timed out, read raises non-zero return code, which can be used in bash script's conditional logics like `if`

# Specify how many characters
- `read -n1 -p "Please enter Y/N:"`
- once one character is entered, it proceed forward, no need to press ENTER key

# Read Silently (e.g. password)
- `read -s "please enter password:" password`
