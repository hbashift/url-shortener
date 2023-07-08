proto:
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=shortener proto/*.proto

compile:
	env GOOS=linux GOARCH=amd64 go build -o  bin/server cmd/server/main.go
	env GOOS=linux GOARCH=amd64 go build -o  bin/api cmd/api/main.go

postgres:
	env GOOS=linux GOARCH=amd64 go build -o  bin/server cmd/server/main.go
	env GOOS=linux GOARCH=amd64 go build -o  bin/api cmd/api/main.go
	docker compose -f docker-compose-postgres.yaml up --build

redis:
	env GOOS=linux GOARCH=amd64 go build -o  bin/server cmd/server/main.go
	env GOOS=linux GOARCH=amd64 go build -o  bin/api cmd/api/main.go
	docker compose -f docker-compose-redis.yaml up --build