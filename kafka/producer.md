# A producer with key value separator
`kafka-console-producer --broker-list localhost:9092 --topic topic1 --property "parse.key=true" --property "key.separator=,"`

# A producer with non string key/value type
