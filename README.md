## 2 microservises (golang, rest api, mongodb, docker-compose)

### NEED TO DO [task link](https://github.com/giffone/task_for_go_dev_one/blob/main/TODO.md)

### Run

in command line
```console
make run
```

### Stop

in command line
```console
make stop
```

### 1.service - generate salt

request
```
http://localhost:8001/generate-salt
```
response
```json
{
	"salt": "293960b3d957"
}
```

### 2.service - user

request create user
```
curl --request POST \
  --url http://localhost:8002/create-user \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "aaa@mail.ru", 
	"password": "12345"
}'
```
response
```json
{
	"status": "ok",
	"message": "user created"
}
```

request get user
```
http://localhost:8002/get-user/aaa@mail.ru
```
response
```json
{
	"id": "63ee69b935dd1f2cb699ca0d",
	"email": "aaa@mail.ru",
	"password": "d8206c92cbf434c985c496fe3b6e684c",
	"salt": "825a1dc4321c"
}
```
or
```json
{
	"status": "not found",
	"message": "mongo: no documents in result"
}
```