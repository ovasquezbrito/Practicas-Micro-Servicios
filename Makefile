gen: 
	protoc --proto_path=proto ./proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm ./pb/*.go

server:
	go run cmd/server/main.go --port 8080

client:
	go run cmd/client/main.go --address 0.0.0.0:8080

test:
	go test -cover -race ./...

.PHONY: gen clean server client test

install: 
	go get -u github.com/jinzhu/copier
	go get -u github.com/google/uuid
	github.com/stretchr/testify/require