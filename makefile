.PHONY: proto

proto:
	@protoc -I internal/api/ --go_out=plugins=grpc:internal/api internal/api/audio.proto

build-services:
	@docker build -t jonnypillar/somniloquy-services -f ./build/package/Dockerfile .

run-services:
	@go run cmd/services/main.go