* procfs or "/proc" is a special filesystem under Linux that is used to present process information and kernel processes. Although "/proc" is still used widely, much of the information found on systems running with a kernel level of 2.6 and above have been moved to another pseudo filesystem called "sysfs" which is generally mounted under "/sys". "/proc" is stored in memory, unlike other filesystems, which are stored on disk

# features
* is NOT a directory at the root of the file system
* IS a pseudo filesystem
* is NOT persistent across boots
* DOES contain primarily read-only zero-byte files and directories
* DOES have man page

# Notable Entries (for each process)
* `cmdline`: Contains the command that originally started the process.
* `cmd`: a symbolic link to the current working directory of the process.
* `environ` : This file contains variable information relating to the process.
* `exe` : a symbolic link pointing to the original executable file
* `fd` : Is a directory that contains symbolic links for each open file descriptor:
* `fdinfo` : A directory that contains files that relate to each open file descriptor found in "/fd".
* `maps` : A file containing information about mapped files and blocks.
* `mem` : A binary file that represents the processes virtual memory.
* `root` : A symbolic link to the root path as seen by the process.
* `status`: A file that contains information relating to the process regarding its current run state and memory usage.
* `task` : This directory contains links to any tasks that have been started by this process.

# Notable Directories
* `/proc/fb` : List of available frame buffers.
* `/proc/cpuinfo` : Lists information about your systems CPU - Family, vendor information, Model number, Core Speeds and CPU flag information.
* `/proc/devices` : Contains a list of character and block devices.
* `/proc/diskstats` : Lists information relating to Logical Disk Devices.
* `/proc/filesystems` : List of filesystems that are supported by the kernel.
* `/proc/interrupts` : Interrupt information can be found here.
* `/proc/iomem` : Contains a map of the systems memory for each physical device.
* `/proc/ioports` : Contains a list of currently registered port regions used for input or output communication with a device.
* `/proc/irq` : This directory contains directories that correspond to IRQs present on your system. SMP affinity information may be modified here.
* `/proc/meminfo` : Contains kernel memory information.
* `/proc/modules` : Contains currently loaded modules within the kernel. "lsmod" command obtains its information from here.
  * we can run `modinfo modulexyz` to view details of a module
* `/proc/mounts` : Contains information regarding mounts. Filesystems in use and what mount options are in use are also listed.
* `/proc/net` : Network stack information.
* `/proc/partitions` : A list of the device numbers, their size and their /dev names which the kernel has identified as a partition.
* `/proc/slabinfo` : Contains kernel slab statistics.
* `/proc/swaps` : List of active swap partitions and their size.
* `/proc/sys` : Dynamically configurable kernel options can be found here.
* `/proc/uptime` : The amount of time in seconds the kernel has been running since boot time and spent in idle mode.
* `/proc/version` : Contains kernel information, version number, gcc version number used for building the kernel.
* `/proc/cmdline` command line to start kernel image, and parameters

# IRQ
* Interrupt requests (IRQs) allow hardware devices to indicate when they have data to send to the CPU. The PnP system must assign each hardware device installed on the system a unique IRQ address.
* View current IRQs in use by `cat /proc/interrupts`

# IO Ports
* The system I/O ports are locations in memory where the CPU can send data to and receive data from the hardware device. As with IRQs, the system must assign each device a unique I/O port.
* View port assignments by `cat /proc/ioports`

# DMA
* Using I/O ports to send data to the CPU can be somewhat slow. To speed things up, many devices use direct memory access (DMA) channels. DMA channels do what the name implies—they send data from a hardware device directly to memory on the system, without having to wait for the CPU. The CPU can then read those memory locations to access the data when it’s ready.
* As with I/O ports, each hardware device that uses DMA must be assigned a unique channel number.
* View current channel assignment by `cat /proc/dma`
