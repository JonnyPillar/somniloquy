.PHONY: install proto build-services run-services run-client test

install:
	@dep ensure -v

proto:
	@protoc -I internal/api/ --go_out=plugins=grpc:internal/api internal/api/audio.proto

build-services:
	@docker build -t jonnypillar/somniloquy-services -f ./build/package/services/Dockerfile .

push-services:
	@docker push jonnypillar/somniloquy-services

run-record-service:
	@go run cmd/services/record/main.go

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