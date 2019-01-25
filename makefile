.PHONY: install proto build-services run-services run-client test

install:
	@dep ensure -v

proto:
	@protoc -I internal/api/ --go_out=plugins=grpc:internal/api internal/api/audio.proto

build-services:
	@docker build -t jonnypillar/somniloquy-services -f ./build/package/services/Dockerfile .

run-audio-service:
	@go run cmd/services/audio/main.go

run-conversion-service:
	@go run cmd/services/conversion/main.go

run-transcription-service:
	@go run cmd/services/transcription/main.go

run-client:
	@go run cmd/client/main.go

test:
	@go test ./... -cover

test-verbose:
	@go test ./... -cover -v