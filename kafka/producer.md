# Basic String producer
```
kafka-console-producer \
  --bootstrap-server localhost:9092 \
  --topic string_input \
  --property "parse.key=true" \
  --property "key.separator=|"

Example input:
KEY|VALUE
```

# Parameterise Input
```
# produce.sh
echo $1 | kafka-console-producer \
  --bootstrap-server localhost:9092 \
  --topic $2 \
  --property "parse.key=true" \
  --property "key.separator=|"

Example command usage:
produce.sh "KEY|VALUE" my_topic
```

# Produce Headers
```
kafka-console-producer \
  --bootstrap-server localhost:9092 \
  --topic TOPIC_NAME \
  --property "parse.key=true" \
  --property "key.separator=|" \
  --property parse.headers=true \
  --property headers.delimiter='^' \
  --property headers.separator=, \
  --property headers.key.separator=:

Example input:
h1:v1,h2:v2^KEY|VALUE
```

# Enable mTLS (see mtls_client.properties)
```
kafka-console-producer \
--bootstrap-server $BOOTSTRAP_SERVER \
--producer.config "PATH_TO_mtls_client.properties" \
--property key.serializer=org.apache.kafka.common.serialization.StringSerializer \
--property value.serializer=org.apache.kafka.common.serialization.StringSerializer \
--property parse.key=true \
--property key.separator='|' \
--topic TOPIC_NAME
```

# Basic Avro Producer
## string key, avro value
```
kafka-avro-console-producer \
  --bootstrap-server $BOOTSTRAP_SERVER \
  --property schema.registry.url=$SCHEMA_REGISTRY_URL \
  --property value.schema="$(< PATH_TO_schema.avsc)" \
  --property key.serializer=org.apache.kafka.common.serialization.StringSerializer \
  --property parse.key=true \
  --property key.separator='|' \
  --topic TOPIC_NAME

Example input:
KEY|{"field1":{"string":"test001"},"field2":null,"field3":1688529120000,"decimalField":"a\t","longField":{"long":500}}
```
# both key and value are avro
```
kafka-avro-console-producer \
  --bootstrap-server $BOOTSTRAP_SERVER \
  --property schema.registry.url=$SCHEMA_REGISTRY_URL \
  --property key.schema="$(< PATH_TO_KEY_schema.avsc)" \
  --property value.schema="$(< PATH_TO_VALUE_schema.avsc)" \
  --property parse.key=true \
  --property key.separator='|' \
  --topic TOPIC_NAME
```


# Avro Producer with mTLS (mTLS with both Broker and Schema Registry)
```
kafka-avro-console-producer \
  --bootstrap-server $BOOTSTRAP_SERVER \
  --producer.config client_mtls.config \
  --property schema.registry.url=$SCHEMA_REGISTRY_URL \
  --property schema.registry.ssl.keystore.location=$KAFKA_KEYSTORE_LOCATION \
  --property schema.registry.ssl.keystore.password=$KAFKA_JKS_PASSWORD \
  --property schema.registry.ssl.truststore.location=$KAFKA_TRUSTSTORE_LOCATION \
  --property schema.registry.ssl.truststore.password=$KAFKA_JKS_PASSWORD \
  --property value.schema="$(< SCHEMA.avsc)" \
  --property key.serializer=org.apache.kafka.common.serialization.StringSerializer \
  --property parse.key=true \
  --property key.separator='|' \
  --topic TOPIC_NAME
```

# Avro Producer with SASL (USERNAME/PASSWORD with Schema Registry)
```
kafka-avro-console-producer \
  --bootstrap-server $BOOTSTRAP_SERVER \
  --producer.config client_sasl.config \
  --property schema.registry.url=$SCHEMA_REGISTRY_URL \
  --property basic.auth.credentials.source=USER_INFO \
  --property basic.auth.user.info=$SR_USERNAME:$SR_PASSWORD \
  --property value.schema="$(< SCHEMA.avsc)" \
  --property key.serializer=org.apache.kafka.common.serialization.StringSerializer \
  --property parse.key=true \
  --property key.separator='|' \
  --topic TOPIC_NAME
```

# Producing null
```
kafka-console-producer \
  --bootstrap-server localhost:9092 \
  --topic TOPIC_NAME \
  --property "parse.key=true" \
  --property "key.separator=|" \
  --property "null.marker=~"

Example input:
~|VALUE
```
This works for both key, value, and headers
