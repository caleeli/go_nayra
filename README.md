
## Create dynamodb tables

```
AWS_ACCESS_KEY_ID="DUMMYIDEXAMPLE"
AWS_SECRET_ACCESS_KEY="DUMMYEXAMPLEKEY"
AWS_ENDPOINT_URL=http://localhost:8000
```

processes
```
aws dynamodb create-table --table-name processes --endpoint-url $AWS_ENDPOINT_URL --region us-west-2 --attribute-definitions "AttributeName=Id,AttributeType=S" --key-schema "AttributeName=Id,KeyType=HASH" --provisioned-throughput "ReadCapacityUnits=1,WriteCapacityUnits=1"
```

requests
```
aws dynamodb create-table --table-name requests --endpoint-url $AWS_ENDPOINT_URL --region us-west-2 --attribute-definitions "AttributeName=Id,AttributeType=S" --key-schema "AttributeName=Id,KeyType=HASH" --provisioned-throughput "ReadCapacityUnits=1,WriteCapacityUnits=1"
```