# __consumer_offsets
## View Offset Info
```
kafka-console-consumer \
  --consumer-property "exclude.internal.topics"="false" \
  --from-beginning \
  --topic __consumer_offsets \
  --bootstrap-server localhost:9092 \
  --formatter "kafka.coordinator.group.GroupMetadataManager\$OffsetsMessageFormatter"
```
## View Group Info
```
kafka-console-consumer \
  --consumer-property "exclude.internal.topics"="false" \
  --from-beginning \
  --topic __consumer_offsets \
  --bootstrap-server localhost:9092 \
  --formatter "kafka.coordinator.group.GroupMetadataManager\$GroupMetadataMessageFormatter"
```
- Add the following property to `consumer.properties`: `exclude.internal.topics=false`
  `kafka-console-consumer --consumer.config consumer.properties --from-beginning --topic __consumer_offsets --bootstrap-server localhost:9092 --formatter "kafka.coordinator.group.GroupMetadataManager\$OffsetsMessageFormatter"`

Background: Since Kafka 0.10, the list of Consumer Groups are stored in the __consumer_offsets topic. That topic contains both the committed offsets and the groups metadata (group.id, members, generation, leader, ...). Groups are stored using GroupMetadataMessage messages (Offsets use OffsetsMessage).

## For older kafka
Will need to add property `exclude.internal.topics` and use a different class name
```
kafka-console-consumer \
  --consumer-property "exclude.internal.topics"="false" \
  --from-beginning \
  --topic __consumer_offsets \
  --bootstrap-server localhost:9092 \
  --formatter "kafka.coordinator.GroupMetadataManager\$OffsetsMessageFormatter"
```

# __transaction_state
```
kafka-console-consumer \
  --consumer-property "exclude.internal.topics"="false" \
  --from-beginning \
  --topic __transaction_state \
  --bootstrap-server localhost:9092 \
  --formatter "kafka.coordinator.transaction.TransactionLog\$TransactionLogMessageFormatter"
```
