# launch a Kafka consumer
bin/kafka-console-consumer.sh --bootstrap-server kafka1:9092 \
    --topic color-output \
    --from-beginning \
    --formatter kafka.tools.DefaultMessageFormatter \
    --property print.key=true \
    --property print.value=true \
    --property key.deserializer=org.apache.kafka.common.serialization.StringDeserializer \
    --property value.deserializer=org.apache.kafka.common.serialization.LongDeserializer

# launch the streams application
bin/kafka-console-consumer.sh --bootstrap-server kafka1:9092 --topic color-output --from-beginning

# List the consumer groups known to Kafka

`bin/kafka-consumer-groups.sh --zookeeper localhost:2181 --list (old api)`

`bin/kafka-consumer-groups.sh --new-consumer --bootstrap-server localhost:9092 --list (new api)`

# View the details of a consumer group

`bin/kafka-consumer-groups.sh --zookeeper localhost:2181 --describe --group <group name>`
