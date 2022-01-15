# Config Location
* `/etc/systemd/journald.conf`

# Config entires
* `Storage=`
  * `none`
  * `volatile` saved in RAM and in `/run/log/journal`
  * `persistent` saved in `/var/log/journal` persistently
  * `auto` (default) like persistent, ONLY if `/var/log/journal` folder exists, otherwise fallback to `volatile`
* `Compress=`
  * `yes`(default) or `no`
* `ForwardToSyslog=`
  * `yes`(default) or `no`
* `SystemMaxUse=, SystemKeepFree=, SystemMaxFileSize=, SystemMaxFiles=, RuntimeMaxUse=, RuntimeKeepFree=, RuntimeMaxFileSize=, RuntimeMaxFiles=`
  - Enforce size limits on the journal files stored. The options prefixed with `System` apply to the journal files when stored on a persistent file system, more specifically `/var/log/journal`. The options prefixed with `Runtime` apply to the journal files when stored on a volatile in-memory file system, more specifically `/run/log/journal`.
  - `SystemMaxUse=` and `RuntimeMaxUse=` control how much disk space the journal may use up at most.
  - `SystemKeepFree=` and `RuntimeKeepFree=` control how much disk space `systemd-journald` shall leave free for other uses. `systemd-journald` will respect both limits and use the smaller of the two values.
  - `SystemMaxFileSize=` and `RuntimeMaxFileSize=` control how large individual journal files may grow at most.
  - `SystemMaxFiles=` and `RuntimeMaxFiles=` control how many individual journal files to keep at most.
* `MaxFileSec=`
  * `n month/week/day` Typically this is not needed if a size limitation is employed. To turn this feature off, set the number to 0 with no time unit. (Default is 1month.)
- `MaxRetentionSec=`
  - The maximum time to store journal entries. This controls whether journal files containing entries older than the specified time span are deleted.
- `SplitMode=`
  - Controls whether to split up journal files per user, either `uid` or `none`. Split journal files are primarily useful for **access control**
  - In either case, the system’s active journal file is named `system.journal`, with user active journal files (if used) named `user-UID.journal`.
  - These journal files are rotated automatically when a certain size or time is reached, depending on the directives set in the `journal.conf` file. After the files are rotated, they are renamed and considered archived. The archived journal filenames start with either system or user-UID, contain an `@` followed by several letters and numbers, and end in a `.journal` file extension.
