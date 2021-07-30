# A producer with key value separator
bin/kafka-console-producer.sh --broker-list kafka1:9092 --topic color-input --property "parse.key=true" --property "key.separator=,"
