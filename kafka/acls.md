# List
```
kafka-acls \
  --bootstrap-server $BOOTSTRAP_SERVER \
  -list
```
## When mTLS is enabled
```
kafka-acls \
  --bootstrap-server $BOOTSTRAP_SERVER \
  --command-config PATH_TO_mtls_client.properties \
  -list
```
