# Cron Job
- `kubectl create cronjob cj1 --schedule="* * * * *" --image=nginx -- /bin/sh -c 'echo "Current date: $(date)"`
- Pods are not deleted
- `kubectl get cronjobs`
- create from yaml
```
apiVersion: batch/v1
kind: CronJob
metadata:
  name: current-date
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: current-date
            image: nginx
            args:
            - /bin/sh
            - -c
            - 'echo "Current date: `date`"'
          restartPolicy: OnFailure
```
- In order to reconfigure the job retention history limits, set new values for the attributes `spec.successfulJobsHistoryLimit` (which defaults to 3) and `spec.failedJobsHistoryLimit` (which defaults to 1)

# `spec.concurrencyPolicy`
indicates how to treat a new Job when a previous job is still running
- `"Allow"` allows CronJobs to run concurrently.
- `"Forbid"` forbids concurrent runs, skipping next run if previous hasn't finished yet.
- `"Replace"` cancels currently running job and replaces it with a new one.

# `spec.suspend`
 is useful if you want to temporarily suspend a specific CronJob, without deleting it.
