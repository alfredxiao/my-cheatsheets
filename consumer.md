# launch a Kafka consumer with specified deserializer
```
kafka-console-consumer --bootstrap-server localhost:9092 \
    --topic color-output \
    --from-beginning \
    --formatter kafka.tools.DefaultMessageFormatter \
    --property print.key=true \
    --property print.value=true \
    --property key.deserializer=org.apache.kafka.common.serialization.StringDeserializer \
    --property value.deserializer=org.apache.kafka.common.serialization.LongDeserializer
```

# launch the streams application
`kafka-console-consumer --bootstrap-server kafka1:9092 --topic color-output --from-beginning`

# List the consumer groups known to Kafka?

`kafka-consumer-groups --zookeeper localhost:2181 --list (old api)`
`kafka-consumer-groups --new-consumer --bootstrap-server localhost:9092 --list (new api)`

# View the details of a consumer group?
`kafka-consumer-groups --zookeeper localhost:2181 --describe --group <group name>`
