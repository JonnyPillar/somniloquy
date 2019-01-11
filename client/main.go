package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jonnypillar/somniloquy/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	resposne, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "Foo"})
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	fmt.Println("Response from server: ", resposne)

	us := api.NewUploadServiceClient(conn)

	var (
		writing = true
		buf     []byte
		n       int
		dat     *os.File
		status  *api.UploadStatus
	)

	dat, err = os.Open("./client/test_audio/test.mp3")
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
	defer dat.Close()

	buf = make([]byte, (1 << 12))

	stream, err := us.Upload(context.Background())
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	for writing {
		// put as many bytes as `chunkSize` into the
		// buf array.
		n, err = dat.Read(buf)
		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}

			err = errors.Wrapf(err,
				"errored while copying from file to buf")
			return
		}

		err = stream.Send(&api.Chunk{
			Content: buf[:n],
		})
		if err != nil {
			err = errors.Wrapf(err,
				"failed to send chunk via stream")
			return
		}
	}

	status, err = stream.CloseAndRecv()
	if err != nil {
		err = errors.Wrapf(err,
			"failed to receive upstream status response")
		return
	}

	if status.Code != api.UploadStatusCode_Ok {
		err = errors.Errorf(
			"upload failed - msg: %s",
			status.Message)
		return
	}

	fmt.Println("Response from server: ", status)
}
