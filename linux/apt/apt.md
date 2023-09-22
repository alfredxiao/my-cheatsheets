# What is installed
* `sudo apt list --installed` what is installed
* `sudo apt list --upgradable` what is upgradable
- not applicable to `apt-get`

# Similarity with `apt-get`
## Same Parameters
| apt               | apt-get            | function                            |
| ------------------|--------------------|-------------------------------------|
| apt install       | apt-get install    |                                     |
| apt remove        | apt-get remove     |                                     |
| apt purge         | atp-get purge      | remove a package with configuration |
| apt update        | apt-get update     | refreshes repository index          |
| apt remove        | apt-get remove     |                                     |
| apt upgrade       | apt-get upgrade    | upgrades all upgradeable            |
| apt autoremove    | apt-get autoremove | removes unwanted packages           |

## Differences
| apt               | apt-get              | function                           |
| ------------------|----------------------|------------------------------------|
| apt full-upgrade  | apt-get dist-upgrade | upgrades packages with auto-handling of dependencies |
| apt search        | apt-cache search     | searches for the program           |
| apt show          | apt-cache show       | show package details               |
| apt list          | dpkg-query list, apt-cache pkgnames     | list what packages are avilable    |

# Only Upgrade
- `sudo apt --only-upgrade install xyz`

# Search
- `apt search xyz`

# Show
- `apt show xyz`
