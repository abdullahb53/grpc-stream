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
	// Dial options initialize.
	target := "localhost:8002"
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	// Wait for server. 'WithBlock'.
	opts = append(opts, grpc.WithBlock())

	//
	// Connect to server.
	//
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("dial connection error, %v", err)
	}
	fmt.Println("[CLIENT]Connected to server..")
	// defer conn.Close()

	//
	// We are CLIENT NOW..!
	//
	client := genpb.NewStreamingPracticesServiceClient(conn)

	ctx := context.Background()
	//
	// We have a stream now.
	//
	stream, err := client.ServerSideStreamFunc(ctx, &genpb.BasicRequest{
		HashCode: "[I'm request from CLIENT.]",
	})
	if err != nil {
		log.Println("resp error from client:", err, "\nresponse:", stream)
	}

	go func() {
		for {
			// Getting streaming data.
			val, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("EOF Detect.")
				break
			}
			fmt.Println(val)
		}
	}()

	// We'll have to wait to see the data.
	c := make(chan bool)
	<-c

}
