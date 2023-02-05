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
	streamForServ, err := client.ServerSideStreamFunc(ctx, &genpb.BasicRequest{
		HashCode: "[I'm one request from CLIENT.]",
	})
	if err != nil {
		log.Println("resp error from client:", err, "\nresponse:", streamForServ)
	}

	//
	// Server-side streaming data.
	//
	go func() {
		for {
			// Getting streaming data.
			val, err := streamForServ.Recv()
			if err == io.EOF {
				fmt.Println("EOF Detect.")
				break
			}

			if err != nil {
				break
			}

			fmt.Println("[CLIENT] [GETTING STREAM DATA FROM SERVER]: ", val)
		}
	}()

	// We have a client-side stream now.
	streamForCli, err := client.ClientSideStreamFunc(ctx)
	if err != nil {
		log.Println("resp error from client:", err, "\nresponse:", streamForCli)
	}

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

	// We have a client-side stream now.
	streamForBiDirect, err := client.BiDirectionalStreamFunc(ctx)
	if err != nil {
		log.Println("resp error from bi-direct:", err, "\nresponse:", streamForBiDirect)
	}

	//
	// Bidirect-side streaming data.
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
			fmt.Println("[BIDIRECT] [SENDING STREAM DATA TO SERVER]: ", strRandomVal)

		}
	}()

	go func() {
		for {
			// Getting streaming data.
			val, err := streamForBiDirect.Recv()
			if err == io.EOF {
				fmt.Println("EOF Detect.")
				break
			}

			if err != nil {
				break
			}

			fmt.Println("[BIDIRECT] [GETTING STREAM DATA FROM SERVER]: ", val)
		}
	}()

	// We'll have to wait to see the data.
	c := make(chan bool)
	<-c

}
