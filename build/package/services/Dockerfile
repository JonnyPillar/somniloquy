FROM golang:alpine AS builder

COPY . /go/src/github.com/jonnypillar/somniloquy/
WORKDIR /go/src/github.com/jonnypillar/somniloquy/cmd/services/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o /go/bin/somniloquy-services

FROM scratch
COPY --from=builder /go/bin/somniloquy-services /go/bin/somniloquy-services
EXPOSE 7777

CMD ["/go/bin/somniloquy-services"]