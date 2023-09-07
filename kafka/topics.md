# List topics
`kafka-topics --bootstrap-server localhost:9092 --list`

# Create a Topic
`kafka-topics --bootstrap-server localhost:9092 --create --topic topic1 --partitions 1 --replication-factor 1`

# Create a topic With compaction enabled
`kafka-topics --bootstrap-server localhost:9092 --create --topic topic1 --partitions 1 --replication-factor 1 --config cleanup.policy=compact`

# Create a topic with smaller segment size and retention size
`kafka-topics --bootstrap-server localhost:9092 --create --topic topic1 --partitions 1 --replication-factor 1 --config cleanup.policy=compact --config segment.ms=1000,retention.ms=10000`

# Create a topic that is instantly compacted
`kafka-topics --bootstrap-server localhost:9092 --create --topic topic1 --partitions 1 --replication-factor 1 --config cleanup.policy=compact --config segment.ms=1000 --config min.cleanable.dirty.ratio=0.01 --config delete.retention.ms=100`

# Describe a topic
`kafka-topics --bootstrap-server localhost:9092 --describe --topic topic1`

# Alter topic config
`kafka-configs --bootstrap-server localhost:9092 --alter --entity-type topics --entity-name topic1 --add-config segment.ms=1000,min.cleanable.dirty.ratio=0.01,delete.retention.ms=100`

# Get the earliest offset still in a topic
`kafka-run-class kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic topic1 --time -2`

# Get the latest offset still in a topic
`kafka-run-class kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic topic1 --time -1`

# Purge a topic
`kafka-topics --zookeeper localhost:2181 --alter --topic topic --config retention.ms=1000`
Then reset it
`kafka-topics --zookeeper localhost:2181 --alter --topic mytopic --delete-config retention.ms`
