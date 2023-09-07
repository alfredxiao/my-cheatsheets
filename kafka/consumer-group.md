# View all consumer groups (and active consumers)
`kafka-consumer-groups --bootstrap-server $BOOTSTRAP_SERVER --all-groups -describe`

# View single consumer group (and Lag)
`kafka-consumer-groups --bootstrap-server $BOOTSTRAP_SERVER --group <GROUP_ID> --describe `

# Reset Offset 
```
kafka-consumer-groups \
    --bootstrap-server $BOOTSTRAP_SERVER \
    --reset-offsets \
    --to-earliest \
    --execute \
    --group G1 \
    --topic t2
```
Dry run instead of execute: `--dry-run`
To Different offsets:
- `--to-current`                            Reset offsets to current offset.
- `--to-datetime <String: datetime>`        Reset offsets to offset from datetime.
  - Format: `YYYY-MM-DDTHH:mm:SS.sss`
- `--to-earliest`                           Reset offsets to earliest offset.
- `--to-latest`                             Reset offsets to latest offset.
- `--to-offset <Long: offset>`              Reset offsets to a specific offset.
- `--shift-by <Long: number-of-offsets>`    Reset offsets shifting current offset by 'n', where 'n' can be positive or negative.

# Delete Offsets
## Delete offset for single Group/Topic combination
```
kafka-consumer-groups \
    --bootstrap-server $BOOTSTRAP_SERVER \
    --delete-offsets \
    --execute \
    --group GROUP_ID \
    --topic TOPIC_NAME
```
## Delete offset for entire Group
```
kafka-consumer-groups \
    --bootstrap-server $BOOTSTRAP_SERVER \
    --delete \
    --execute \
    --group GROUP_ID
```

