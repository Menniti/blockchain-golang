# blockchain-golang


### Start
```
go run main.go
```

### Teste
```
localhost:8080

Get the blockchain
GET:
curl -X GET \
  http://localhost:8080 \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"message": 40
}'

New block to be mined
POST: 

curl -X POST \
  http://localhost:8080 \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"message": 40
}'

```
