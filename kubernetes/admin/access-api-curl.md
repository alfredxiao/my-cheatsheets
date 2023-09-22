service account token located at
- `/var/run/secrets/kubernetes.io/serviceaccount/token`

kubectl create role config-reader --verb=list,get --resource=configmaps
kubectl create rolebinding config-reader-sa1 --role=config-reader --serviceaccount=default:sa1

export TOKEN=${cat /var/run/secrets/kubernetes.io/serviceaccount/token}
curl -k  https://kubernetes.default/api/v1/namespaces/default/configmaps/c1 -H "Authorization: Bearer ${TOKEN}"
