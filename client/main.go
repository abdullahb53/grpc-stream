package main

import (
	"context"
	"fmt"
	genpb "grpc_stream/genpb/protos"
	"io"
	"log"
	random "math/rand"
	"strconv"
	"time"

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

	// We have a server-side stream now.
	stream, err := client.ServerSideStreamFunc(ctx, &genpb.BasicRequest{
		HashCode: "[I'm one request from CLIENT.]",
	})
	if err != nil {
		log.Println("resp error from client:", err, "\nresponse:", stream)
	}

	// We have a client-side stream now.
	streamForCli, err := client.ClientSideStreamFunc(ctx)
	if err != nil {
		log.Println("resp error from client:", err, "\nresponse:", streamForCli)
	}

	//
	// Server-side streaming data.
	//
	go func() {
		for {
			// Getting streaming data.
			val, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("EOF Detect.")
				break
			}
			fmt.Println("[CLIENT] [GETTING STREAM DATA FROM SERVER]: ", val)
		}
	}()

	//
	// Client-side streaming data.
	//
	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			if i > 10 {
				break
			}
			i++

			_ = random.Int()
			strRandomVal := strconv.Itoa(i + 12)
			streamForCli.Send(&genpb.BasicRequest{
				HashCode: strRandomVal,
			})
			fmt.Println("[Client] [SENDING STREAM DATA TO SERVER]: ", strRandomVal)

		}
	}()

	// We'll have to wait to see the data.
	c := make(chan bool)
	<-c

}
