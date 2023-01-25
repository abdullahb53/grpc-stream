// ABDULLAH BIYIK / abdullahb5355@gmail.com
// KARADENIZ TECHNICAL UNIVERSITY COMPUTER ENGINEERING STUDENT.
package main

import (
	"fmt"
	genpb "grpc_stream/genpb/protos"
	random "math/rand"
	"strconv"

	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	i := -1
	for {
		i++
		if i > 10 {
			break
		}

		randomVal := random.Int()
		strRandomVal := strconv.Itoa(randomVal)

		err := stream.Send(&genpb.BasicResponse{
			HashCode: strRandomVal,
		})
		if err != nil {
			fmt.Println("stream sender error.", err)
			return err
		}
		fmt.Println("[i] i send this value to client..", strRandomVal)

	}
	return nil
}

func (s *myEventServer) ClientSideStreamFunc(stream genpb.StreamingPracticesService_ClientSideStreamFuncServer) error {

	return status.Errorf(codes.Unimplemented, "method ClientSideStreamFunc not implemented")
}
func (s *myEventServer) BiDirectionalStreamFunc(stream genpb.StreamingPracticesService_BiDirectionalStreamFuncServer) error {

	return status.Errorf(codes.Unimplemented, "method BiDirectionalStreamFunc not implemented")
}