gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative v1/fiber.proto

run:
	@go run main.go

build:
	@go build -o main.go bin/fiber-grpc

.PHONY: gen run build