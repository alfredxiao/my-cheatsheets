- by default a Pod is not deleted unless `spec.ttlSecondsAfterFinished` is specified
  - but the Pod is in a Success or Completed state where container is not really running inside
- `spec.activeDeadlineSeconds` how long the job can be active. after that if it is still running, will be terminated by K8s
  - no matter how many pods are created, restarted
  - Note that a Job's `.spec.activeDeadlineSeconds` takes precedence over its `.spec.backoffLimit`. Therefore, a Job that is retrying one or more failed Pods will not deploy additional Pods once it reaches the time limit specified by `activeDeadlineSeconds`, even if the `backoffLimit` is not yet reached.

# imperative
- `kubectl create job counter --image=nginx -- /bin/sh -c 'echo "DONE"'`

# via YAML
```
piVersion: batch/v1
kind: Job
metadata:
  name: counter
spec:
  template:
    spec:
      containers:
      - name: counter
        image: nginx
        command:
        - /bin/sh
        - -c
        - counter=0; while [ $counter -lt 3 ]; do counter=$((counter+1)); \
          echo "$counter"; sleep 3; done;
      restartPolicy: Never
```

# `spec.parallelism`
- controls how many pods are working/running in parallel

# `spec.completions`
- Specifies the desired number of successfully finished pods the job should be run with
  - Setting to `nil` means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.
  - Setting to `1` means that parallelism is limited to `1` and the success of that pod signals the success of the job.

# `spec.backoffLimit`
- controls number of retries a Job attempts to retry if a command executed in a Pod fails (non-zero exit code)
- e.g. `backoffLimit=6` (defaults to 6) would mean the max possible run iteration for a Job is 7 (one initial + 6 retries)

# `spec.template.spec.restartPolicy`
- for a Job it can only be `Never` or `OnFailure`
  - when set to `OnFailure`, **NO** new Pod is created for the failed instance, and you would observe the **same** Pod name in each attempt where the container inside the Pod is a new run
  - if set to `Never`, and while command execution exit code is non-zero, the Pod **will** be replaced with new Pod, hence you would observe a **new** Pod named

Create job from a cronjob (ignoring schedule)
`kubectl create job pinghost --from=cronjob/pinghost`
