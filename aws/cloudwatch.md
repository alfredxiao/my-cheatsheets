# Display selected fields (from logs already in JSON format)
```
filter @message like /myField/
| filter myField1, myField2, myField3
| sort myField1 desc
| limit 200
```