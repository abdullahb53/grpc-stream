package main

import (
	"context"
	"fmt"
	genpb "grpc_stream/genpb/protos"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	target := "localhost:8002"
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("dial connection error, %v", err)
	}
	fmt.Println("[CLIENT]Connected to server..")
	// defer conn.Close()

	client := genpb.NewStreamingPracticesServiceClient(conn)

	ctx := context.Background()
	stream, err := client.ServerSideStreamFunc(ctx, &genpb.BasicRequest{
		HashCode: "[I'm request from CLIENT.]",
	})
	if err != nil {
		log.Println("resp error from client:", err, "\nresponse:", stream)
	}

	go func() {
		for {
			val, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("EOF Detect.")
				break
			}
			fmt.Println(val)
		}
	}()

	c := make(chan bool)
	<-c

}
