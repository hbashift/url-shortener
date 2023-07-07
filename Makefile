proto:
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=shortener proto/*.proto

build:
	go build -o  bin/server.exe cmd/server/main.go
	go build -o bin/api.exe cmd/api/main.go
	go build -o bin/client.exe cmd/client/main.go

docker:
	env GOOS=linux GOARCH=amd64 go build -o  bin/server cmd/server/main.go
	docker build -f ./DockerFile -t url-shortener --progress=plain .

compose:
	docker-compose up
.PHONY: proto build