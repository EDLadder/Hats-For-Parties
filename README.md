# Hats-For-Parties

```
docker-compose up
```

1. parties-service run on localhost:8081
2. party-web run on localhost:8082


| Method | Url | Body |
| ------ | ------ | ------ |
| POST | http://localhost:8081/party/start | { "hats": 10, "name": "Test 1"} |
| PATCH | http://localhost:8081/party/stop/{party_id} | no |
| GET | http://localhost:8081/party | no |
| GET | http://localhost:8081/hat | no |
