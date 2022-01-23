- CUPS provides a common interface for working with any type of printer on your Linux system.
- It accepts print jobs using the PostScript document format and sends them to printers using a print queue system.
- The CUPS software uses the `Ghostscript` program to convert the `PostScript` document into a format understood by the different printers. The Ghostscript program requires different drivers for the different printer types to know how to convert the document to make it printable on that type of printer. This is done using configuration files and drivers. Config files for them are stored under `/etc/cups`

# Jobs Management
- `cancel` cancel a print request
- `cupsaccept CUPS-PDF` accept jobs sent to a destination
- `cupsreject CUPS-PDF` reject new jobs for a destination
- `cupsenable CUPS-PDF` start a printer
- `cupsdisable CUPS-PDF` : stops a printer (but still accept jobs)

# Legacy BSD style utilities
- `lpc` Start, stop, or pause the print queue
  - `lpc status` : check status of printers
- `lpq`: Display the print queue status, along with any print jobs waiting in the queue
  - `lpq` : view jobs in queue
  - `lpq -P EPSON_ET_3750_Series`
- `lpr`: Submit files for printing
  - `lpr afile` : send a file to printer
  - `lpr -P EPSON_ET_3750_Series test.txt` with -P to specify which printer
- `lprm`: Remove a specific print job from the print queue
  - `lprm nnn` : remove job from queue identified by a number seen from output of `lpq`

# What printers are installed
- `lpstat -s` : check what printers are installed
- `lpinfo -v` : check what types of printer connections are available
