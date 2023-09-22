# What images are being run
`kubectl get pods -A -o jsonpath="{.items[*].spec.containers[*].image}" | tr -s ' ' '\n' | sort | uniq -c`
