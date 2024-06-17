# Default mode `auto.register.schemas=true` and `use.latest.version=false`
1. producer send first message
`POST /subjects/transactions-value/versions HTTP/1.1`
`Content-Type: application/vnd.schemaregistry.v1+json`
```
{"schema":"{\"type\":\"record\",\"name\":\"Transaction\",\"namespace\":\"demo.model\",\"fields\":[{\"name\":\"amount1\",\"type\":\"long\"},{\"name\":\"amount2\",\"type\":\"long\"}]}"}
```
response
```
200 OK
{"id":1}
```
2. Same producer send again, same thing repeats
## Lessons Learned
1. When `auto.register.schemas` is `true`, producer (really it is the avro serializer) registers the schema **every time**
2. Upon success registration, Schema Registry responds with 200 and and `id` number **every time**
3. This is what `schemaRegistry.register(subject, schema)` does, this method defined in interface `SchemaRegistryClient`, real class `CachedSchemaRegistryClient`

# Set `auto.register.schemas=false` while `use.latest.version=false`
1. with previous schema already registered
`POST /subjects/transactions-value?deleted=false HTTP/1.1`
same schema payload as step 1
response
```
200 OK
{"subject":"transactions-value","version":1,"id":1,"schema":"{\"type\":\"record\",\"name\":\"Transaction\",\"namespace\":\"demo.model\",\"fields\":[{\"name\":\"amount1\",\"type\":\"long\"},{\"name\":\"amount2\",\"type\":\"long\"}]}"}
```
2. Same producer send again, same thing repeats
## Lessons Learned
1. Even not registering new schema version, the producer ask for the id given the schema it is trying to use;
2. This is what `schemaRegistry.getId(subject, schema)` does
3. If request payload in the POST is NOT an already registered schema version, an 404 "Schema not found" error will be returned from Schema Registry and serialization throws exception

# Set `auto.register.schemas=false` and `use.latest.version=true`
1. `GET /subjects/transactions-value/versions/latest HTTP/1.1`
`{"subject":"transactions-value","version":1,"id":1,"schema":"{\"type\":\"record\",\"name\":\"Transaction\",\"namespace\":\"demo.model\",\"fields\":[{\"name\":\"amount\",\"type\":\"long\"}]}"}`
2. `POST /subjects/transactions-value?deleted=false HTTP/1.1`
```
{"schema":"{\"type\":\"record\",\"name\":\"Transaction\",\"namespace\":\"demo.model\",\"fields\":[{\"name\":\"amount\",\"type\":\"long\"}]}"}
```
`{"subject":"transactions-value","version":1,"id":1,"schema":"{\"type\":\"record\",\"name\":\"Transaction\",\"namespace\":\"demo.model\",\"fields\":[{\"name\":\"amount\",\"type\":\"long\"}]}"}`
## Lessons Learned
1. It uses `schemaRegistry.getLatestSchemaMetadata(subject)` to GET latest schema version;
2. It uses `ParsedSchema.isBackwardCompatible(ParsedSchema previousSchema)` (implemented in `AvroSchema` class) to check compatibility, this is a pure client side check, not performed by Schema Registry
  - Validate that the 'latest' schema can be used to read data written with `previousSchema` - which is the schema the producer was compiled with in its `SpecificRecord`
3. If step 2 outcome is unsuccessful, serialization throws exception
  - This can happen if the local schema is missing a mandatory field which is present in 'latestVersion'
4. If step 2 outcome is successful, it does a `schemaRegistry.getId(subject, schema)` (where `schema` is the latestVersion looked up from Schema Registry) call (HTTP POST though) to retrieve an integer id used to write into bytes representing the avro record.


# Field order in avro definition
## Those fields under `fields` attribute
Manually do a POST with request similar to step 3, but swap order of two fields to have `amount2` go before `amount1`, like
```
{"schema":"{\"type\":\"record\",\"name\":\"Transaction\",\"namespace\":\"demo.model\",\"fields\":[{\"name\":\"amount2\",\"type\":\"long\"},{\"name\":\"amount1\",\"type\":\"long\"}]}"}
```
response
```
404 Not Found
{
    "error_code": 40403,
    "message": "Schema not found"
}
```
This means ordering of them **Matters**
## Those attrbibutes (excluding things under `fields`)
Manually do a POST with request similar to step 3, but swap order of two attributes to have `namespace` go before `name`, plus extra blanks, like
```
{"schema":"{\"type\":\"record\" ,  \"namespace\":\"demo.model\",  \"name\": \"Transaction\",\"fields\":[{\"name\":\"amount2\",\"type\":\"long\"},{\"name\":\"amount1\",\"type\":\"long\"}]}"}
```
response same as in step 3
Meaning ordering of the does **NOT** matter
