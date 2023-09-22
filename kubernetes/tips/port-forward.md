`kubectl port-forward service/mongo 28015:27017`
or
`kubectl port-forward deployment/mongo 28015:27017`

then run Mongo client
`mongosh --port 28015`
