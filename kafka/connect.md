# Cluster
## View Connectors
### non SSL
`curl $CONNECT_URL/connectors | jq`
### Include connector status
`curl $CONNECT_URL/connectors?expand=status | jq`
### SSL
`curl -k $CONNECT_URL/connectors | jq`
### mTLS
```
curl --cacert path_to_caroot.pem \
     --key parth_to_private_key.pem \
     --cert path_to_public_cert.pem \
     $CONNECT_URL/connectors | jq
```

## View Available PlugIns
`curl $CONNECT_URL/connector-plugins/`
## Update Debug Level
```
curl 
    -X PUT \
    -H "Content-Type: application/json" \
    $CONNECT_URL/admin/loggers/io.confluent.connect.s3 -d '{"level":"TRACE"}' | jq
```


# Connector
## View Info
`curl $CONNECT_URL/connectors/<CONNECTOR_NAME>`
`curl $CONNECT_URL/connectors/<CONNECTOR_NAME>/config`
`curl $CONNECT_URL/connectors/<CONNECTOR_NAME>/status`
## Restart a connector
`curl -X POST $CONNECT_URL/connectors/<CONNECTOR_NAME>/restart`
### including tasks
`curl -X POST $CONNECT_URL/connectors/<CONNECTOR_NAME>/restart?includeTasks=true`
### including tasks, but only restart failed ones
`curl -X POST $CONNECT_URL/connectors/<CONNECTOR_NAME>/restart?includeTasks=true&onlyFailed=true`
## Pause
`curl -X PUT $CONNECT_URL/connectors/<CONNECTOR_NAME>/pause`
## Resume 
`curl -X PUT $CONNECT_URL/connectors/<CONNECTOR_NAME>/resume`
## Delete a Connector
`curl -X DELETE $CONNECT_URL/connectors/<CONNECTOR_NAME>/`
## Create New Connector
`curl -X POST -H "Content-Type: application/json" $CONNECT_URL/connectors`
```
{
    "name": "hdfs-sink-connector",
    "config": {
        "connector.class": "io.confluent.connect.hdfs.HdfsSinkConnector",
        "tasks.max": "10",
        "topics": "test-topic",
        "hdfs.url": "hdfs://fakehost:9000",
        "hadoop.conf.dir": "/opt/hadoop/conf",
        "hadoop.home": "/opt/hadoop",
        "flush.size": "100",
        "rotate.interval.ms": "1000"
    }
}
```
## Create or Update a Connector 
`curl -X PUT $CONNECT_URL/connectors/CONNECTOR_NAME/config`
```
{
    "connector.class": "io.confluent.connect.hdfs.HdfsSinkConnector",
    "tasks.max": "10",
    "topics": "test-topic",
    "hdfs.url": "hdfs://fakehost:9000",
    "hadoop.conf.dir": "/opt/hadoop/conf",
    "hadoop.home": "/opt/hadoop",
    "flush.size": "100",
    "rotate.interval.ms": "1000"
}
```


# Tasks
## List tasks
`curl $CONNECT_URL/connectors/<CONNECTOR_NAME>/tasks`
## Status of a single task 
`curl $CONNECT_URL/connectors/<CONNECTOR_NAME>/tasks/TASK_ID/status`
## Restart a task 
`curl -X POST $CONNECT_URL/connectors/<CONNECTOR_NAME>/tasks/TASK_ID/restart`

# Offsets
## View Connect Offset
`kafka-console-consumer --bootstrap-server $BOOTSTRAP_SERVER --topic connect-offsets`
### The topic name
can be different from above in actual cluster, can be confirmed by looking at config or logs of connect worker
`offset.storage.topic=connect-offsets`

## Reset Connect Offset
Once the key and value format is identified, enter new values below and write to the offset topic.
### Write null marker
```
echo "KEY|~" | kafka-console-producer \
  --bootstrap-server localhost:9092 \
  --topic $2 \
  --property "parse.key=true" \
  --property "key.separator=|" \
  --property "null.marker=~" 
```
### Write specified value
```
echo "KEY|VALUE" | kafka-console-producer \
  --bootstrap-server localhost:9092 \
  --topic $2 \
  --property "parse.key=true" \
  --property "key.separator=|"
```
