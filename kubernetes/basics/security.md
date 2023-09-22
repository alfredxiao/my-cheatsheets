`kubectl create clusterrolebinding cluster-admin --clusterrole=cluster-admin --user=user1 --user=user2 --group=group1`
kubectl -n practice create serviceaccount pvviewer-practice
kubectl create clusterrole pvviewer-role-practice --resource=pv --verb=list
kubectl create clusterrolebinding pvviewer-role-binding-practice --clusterrole=pvviewer-role-practice --serviceaccount=practice:pvviewer-practice
