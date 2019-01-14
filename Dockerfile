FROM golang:alpine AS builder

COPY . /go/src/github.com/jonnypillar/somniloquy/
WORKDIR /go/src/github.com/jonnypillar/somniloquy/server

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o /go/bin/hello

FROM scratch
COPY --from=builder /go/bin/hello /go/bin/hello
EXPOSE 7777

CMD ["/go/bin/hello"]