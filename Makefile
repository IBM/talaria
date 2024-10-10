include local.env
export $(shell sed 's/=.*//' local.env)

run: run_broker
	
run_broker: 
	go run .

run_test:
	go run ./_examples/client/confluent/main.go

test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./.../... ; go tool cover -html=coverage.out
