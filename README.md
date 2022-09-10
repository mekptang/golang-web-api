
# Build image & Run container

```
$ docker build -t test-server .

$ docker run --rm -p 80:80 test-server
```

# API Usage

## Add key-value pair
```
curl -X POST \
  http://localhost/add \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"key":"asdf",
	"value":"some other value"
}'
```

## List all records
```
curl -X GET \
  http://localhost/list \
  -H 'cache-control: no-cache'
```