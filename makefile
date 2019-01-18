.PHONY: install proto build-services run-services run-client test

install:
	@dep ensure -v

proto:
	@protoc -I internal/api/ --go_out=plugins=grpc:internal/api internal/api/audio.proto

build-services:
	@docker build -t jonnypillar/somniloquy-services -f ./build/package/services/Dockerfile .

run-services:
	@go run cmd/services/main.go

run-client:
	@go run cmd/client/main.go

test:
	@go test ./... -cover