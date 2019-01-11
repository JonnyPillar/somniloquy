FROM golang:alpine AS builder

# RUN apk update && apk add --no-cache git
COPY . /go/src/github.com/jonnypillar/somniloquy/
WORKDIR /go/src/github.com/jonnypillar/somniloquy/server

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o /go/bin/hello
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/hello

FROM scratch
COPY --from=builder /go/bin/hello /go/bin/hello
EXPOSE 7777

ENTRYPOINT ["/go/bin/hello"]