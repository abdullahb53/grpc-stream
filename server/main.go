// ABDULLAH BIYIK / abdullahb5355@gmail.com
// KARADENIZ TECHNICAL UNIVERSITY COMPUTER ENGINEERING STUDENT.
package main

import (
	"fmt"
	genpb "grpc_stream/genpb/protos"
	"io"
	random "math/rand"
	"strconv"
	"time"

	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type myEventServer struct {
	genpb.UnimplementedStreamingPracticesServiceServer
}

var (
	lis net.Listener
)

func main() {
	var err error
	if lis, err = net.Listen("tcp", "localhost:8002"); err != nil {
		log.Fatalf("failed listener error, %v", err)
	}

	// opts: Options.
	// eventServer: 'MyEventServer' implementation.
	// grpcServer: in server side what we will serve.
	var opts []grpc.ServerOption
	eventServer := new(myEventServer)
	grpcServer := grpc.NewServer(opts...)

	genpb.RegisterStreamingPracticesServiceServer(grpcServer, eventServer)

	// gRPC Server Reflection provides information about publicly-accessible gRPC services on a server,
	// and assists clients at runtime to construct RPC requests and responses without precompiled service information.
	//
	// for more information you can go there: https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	//
	// Serving grpc server.
	fmt.Println("grpc serving at", lis.Addr().String())
	if err_serving := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serving error, %v", err_serving)
	}

}

func (s *myEventServer) ServerSideStreamFunc(c *genpb.BasicRequest, stream genpb.StreamingPracticesService_ServerSideStreamFuncServer) error {
	//
	// I want to send some stream-data to client.
	// i: iterator. until 10.
	//
	i := -1
	for {
		time.Sleep(2 * time.Second)
		i++
		if i > 10 {
			break
		}

		// Generate random data for streaming.
		_ = random.Int()
		strRandomVal := strconv.Itoa(i + 3456)

		//
		// Sending data.
		//
		err := stream.Send(&genpb.BasicResponse{
			HashCode: strRandomVal,
		})
		if err != nil {
			fmt.Println("stream sender error.", err)
			return err
		}
		fmt.Println("[SERVER] [SRVSTREAM] [SENDING STREAM DATA TO CLIENT]: ", strRandomVal)

	}
	return nil
}

func (s *myEventServer) ClientSideStreamFunc(stream genpb.StreamingPracticesService_ClientSideStreamFuncServer) error {

	for {
		val, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("[Server]EOF Detect. ")
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println("[SERVER] [CLISTREAM] [GETTING STREAM DATA FROM CLIENT]: ", val)

	}

}
func (s *myEventServer) BiDirectionalStreamFunc(stream genpb.StreamingPracticesService_BiDirectionalStreamFuncServer) error {

	i := -1
	for {
		time.Sleep(2 * time.Second)
		i++
		if i > 10 {
			break
		}

		// Generate random data for streaming.
		_ = random.Int()
		strRandomVal := strconv.Itoa(i + 7106223)

		//
		// Sending data.
		//
		err := stream.Send(&genpb.BasicResponse{
			HashCode: strRandomVal,
		})
		if err != nil {
			fmt.Println("stream sender error.", err)
			return err
		}
		fmt.Println("[SERVER] [BIDIRECT] [SENDING STREAM DATA TO CLIENT]: ", strRandomVal)

	}

	go func() {
		for {
			val, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("[Server]EOF Detect. ")
				break
			}

			fmt.Println("[SERVER] [BIDIRECT] [GETTING STREAM DATA FROM CLIENT]: ", val)
		}
	}()

	return nil
}
