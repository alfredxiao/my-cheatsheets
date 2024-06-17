
# SSL connect, SASL
kafka-topics --command-config client.config --bootstrap-server localhost:9092 --list

# Content of client.config
security.protocol: SASL_SSL
sasl.mechanism: PLAIN
sasl.jaas.config: org.apache.kafka.common.security.plain.PlainLoginModule required username='XX' password='YY';

# of course, replace XX and YY
