generate:
	mkdir -p ./pkg/servicespb
	protoc --proto_path=api/v1/proto \
				--go_out=pkg/servicespb --go_opt=paths=source_relative \
				--go-grpc_out=pkg/servicespb --go-grpc_opt=paths=source_relative \
				services.proto

run provider:
	go run .provider/cmd/client/main.go &
	go run .provider/cmd/server/main.go

run receiver:
	go run .receiver/cmd/client/main.go
	go run .receiver/cmd/server/main.go

