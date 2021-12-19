First, you should determine if your system’s CPU has these hardware extensions available. You can research this via the /proc/cpuinfo file’s flag information. Type
`grep ^flags /proc/cpuinfo` to view the various enabled features of your server’s CPU. If enabled, you should see one of the following:

- For Intel CPUs: `vmx`
- For AMD CPUs: `svm`
