`/etc/inittab`
Format: identifier:run-level:action:process

`id:3:initdefault:`
- id means ‘initdefault’
- 3 means level 3
- ‘initdefault’ is predefined action, meaning defaults to run into level 3
- no process required for this action

`si::sysinit:/etc/rc.d/rc.sysint`
- si means ‘sysinit’
- run level is empty since action is applied globally and not specific to any run-level
- ‘sysinit’ action means run the process specified
- /etc/rc.d/rc.sysint : the process to run as part of ‘sysinit’ action

`L2:2:wait:/etc/rc.d/rc 2`
L2: level two stuff
2: level is 2
wait: Start the specified process once run level 2 is entered, and init will wait for its termination
