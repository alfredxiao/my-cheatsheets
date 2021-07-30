clean a topic

# List topics
bin/kafka-topics.sh --bootstrap-server kafka1:9092 --list

# Create a Topic
bin/kafka-topics.sh --bootstrap-server kafka1:9092 --create --topic color-input --partitions 1 --replication-factor 1

# With compaction enabled
--config cleanup.policy=compact

# Describe a topic
kafka-topics --bootstrap-server localhost:9092 --describe --topic titles

# Alter topic config
kafka-configs --bootstrap-server localhost:9092 --alter --entity-type topics --entity-name titles --add-config segment.ms=1000,min.cleanable.dirty.ratio=0.01,delete.retention.ms=100

# Create a topic that is instantly compacted
kafka-topics --bootstrap-server localhost:9092 --create --topic quickcompact --partitions 1 --replication-factor 1 --config cleanup.policy=compact --config segment.ms=1000 --config min.cleanable.dirty.ratio=0.01 --config delete.retention.ms=100

#Get the earliest offset still in a topic
`bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic mytopic --time -2`

#Get the latest offset still in a topic
`bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic mytopic --time -1`

# Get the consumer offsets for a topic
`bin/kafka-consumer-offset-checker.sh --zookeeper=localhost:2181 --topic=mytopic --group=my_consumer_group`

# Purge a topic
```
bin/kafka-topics.sh --zookeeper localhost:2181 --alter --topic mytopic --config retention.ms=1000
... wait a minute ...
bin/kafka-topics.sh --zookeeper localhost:2181 --alter --topic mytopic --delete-config retention.ms
```

# Read from __consumer_offsets
Add the following property to config/consumer.properties: exclude.internal.topics=false
```
bin/kafka-console-consumer.sh --consumer.config config/consumer.properties --from-beginning --topic __consumer_offsets --zookeeper localhost:2181 --formatter "kafka.coordinator.GroupMetadataManager\$OffsetsMessageFormatter"
```
