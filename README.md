# Flight Path microservice

##### Build and run on local
```shell
sudo go build -o ./bin/flightPath -i ./cmd/flightPath
./bin/flightPath
```

##### run tests
```shell
go test ./...
```

##### example api
```shell
curl -X POST \
http://localhost:8080/airport \
-H 'content-type: application/json' \
-d '{"flights":[["SFO", "ATL"], ["ATL", "GSO"]]}'

curl -X POST \
http://localhost:8080/airport \
-H 'content-type: application/json' \
-d '{"flights":[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]}'

curl -X POST \
http://localhost:8080/airport \
-H 'content-type: application/json' \
-d '{"flights":[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"],["SFO", "GSO"]]}'

curl -X POST \
http://localhost:8080/airport \
-H 'content-type: application/json' \
-d '{"flights":[["JFK","SFO"],["JFK","ATL"]]}'
```