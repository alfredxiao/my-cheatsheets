# Basic string consumer 
```
kafka-console-consumer --bootstrap-server $BOOTSTRAP_SERVER --topic TOPIC_NAME
```

# Print more metadata
```
kafka-console-consumer \
    --bootstrap-server $BOOTSTRAP_SERVER \
    --property print.key=true \
    --property print.value=true \
    --property key.separator='|' \
    --property print.timestamp=true \
    --property print.offset=true \
    --property print.partition=true \
    --property print.headers=true \
    --property key.separator='|' \
    --property headers.separator=',' \
    --property null.literal=NULL \
    --topic TOPIC_NAME 
```

# Enable mTLS (see mtls_client.properties)
```
kafka-console-consumer \
    --bootstrap-server $BOOTSTRAP_SERVER \
    --consumer.config PATH_TO_mtls_client.properties \
    --topic TOPIC_NAME 
```

# Exit when ready 
```
--max-messages 1
```

# Basic Avro consumer 
```
kafka-avro-console-consumer \
    --bootstrap-server $BOOTSTRAP_SERVER \
	--property schema.registry.url=$SCHEMA_REGISTRY_URL \
	--property print.key=true \
	--property print.value=true \
	--property key.separator='|' \
	--topic TOPIC_NAME
```

# Avro consumer with non-Avro key
```
kafka-avro-console-consumer \
    --bootstrap-server $BOOTSTRAP_SERVER \
	--consumer.config /app/confluent/envs/$ENV/client.properties \
	--property schema.registry.url=$SCHEMA_REGISTRY_URL \
	--property print.key=true \
	--property print.value=true \
	--property key.separator='|' \
    --key-deserializer org.apache.kafka.common.serialization.StringDeserializer \
	--value-deserializer io.confluent.kafka.serializers.KafkaAvroDeserializer \
	--topic TOPIC_NAME
```

# Avro consumer with mTLS
```
kafka-avro-console-consumer \
    --bootstrap-server $BOOTSTRAP_SERVER \
	--consumer.config PATH_TO_mtls_client.properties \
	--property schema.registry.url=$SCHEMA_REGISTRY_URL \
    --property basic.auth.credentials.source=USER_INFO \
    --property basic.auth.user.info=USERNAME:PASSWORD \
	--property schema.registry.ssl.keystore.location=$KAFKA_KEYSTORE_LOCATION \
	--property schema.registry.ssl.keystore.password=$KAFKA_JKS_PASSWORD \
	--property schema.registry.ssl.truststore.location=$KAFKA_TRUSTSTORE_LOCATION \
	--property schema.registry.ssl.truststore.password=$KAFKA_JKS_PASSWORD \
	--property print.key=true \
	--property print.value=true \
	--property key.separator='|' \
	--value-deserializer io.confluent.kafka.serializers.KafkaAvroDeserializer \
    --key-deserializer org.apache.kafka.common.serialization.StringDeserializer \
	--topic TOPIC_NAME

```

# launch the streams application
`kafka-console-consumer --bootstrap-server kafka1:9092 --topic color-output --from-beginning`

# View read only committed messages
`--consumer-property "isolation.level"="read_uncommitted"`

# From Specified Offset/Partition
`--partition N --offset M`