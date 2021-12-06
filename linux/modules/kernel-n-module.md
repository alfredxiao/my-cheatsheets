 the Linux kernel uses kernel modules, which are individual hardware driver files that can be linked into the kernel at runtime. That way, the system can link only the modules needed for the hardware present on your system

 The .ko file extension is used to identify the module object files.

 The standard location for storing module object files is in the `/lib/modules` directory. This is where the Linux module utilities (such as `insmod` and `modprobe`) look for module object library files by default.
