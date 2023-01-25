// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: my.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamingPracticesServiceClient is the client API for StreamingPracticesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingPracticesServiceClient interface {
	ServerSideStreamFunc(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (StreamingPracticesService_ServerSideStreamFuncClient, error)
	ClientSideStreamFunc(ctx context.Context, opts ...grpc.CallOption) (StreamingPracticesService_ClientSideStreamFuncClient, error)
	BiDirectionalStreamFunc(ctx context.Context, opts ...grpc.CallOption) (StreamingPracticesService_BiDirectionalStreamFuncClient, error)
}

type streamingPracticesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingPracticesServiceClient(cc grpc.ClientConnInterface) StreamingPracticesServiceClient {
	return &streamingPracticesServiceClient{cc}
}

func (c *streamingPracticesServiceClient) ServerSideStreamFunc(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (StreamingPracticesService_ServerSideStreamFuncClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingPracticesService_ServiceDesc.Streams[0], "/StreamingPracticesService/ServerSideStreamFunc", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingPracticesServiceServerSideStreamFuncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingPracticesService_ServerSideStreamFuncClient interface {
	Recv() (*BasicResponse, error)
	grpc.ClientStream
}

type streamingPracticesServiceServerSideStreamFuncClient struct {
	grpc.ClientStream
}

func (x *streamingPracticesServiceServerSideStreamFuncClient) Recv() (*BasicResponse, error) {
	m := new(BasicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingPracticesServiceClient) ClientSideStreamFunc(ctx context.Context, opts ...grpc.CallOption) (StreamingPracticesService_ClientSideStreamFuncClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingPracticesService_ServiceDesc.Streams[1], "/StreamingPracticesService/ClientSideStreamFunc", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingPracticesServiceClientSideStreamFuncClient{stream}
	return x, nil
}

type StreamingPracticesService_ClientSideStreamFuncClient interface {
	Send(*BasicRequest) error
	CloseAndRecv() (*BasicResponse, error)
	grpc.ClientStream
}

type streamingPracticesServiceClientSideStreamFuncClient struct {
	grpc.ClientStream
}

func (x *streamingPracticesServiceClientSideStreamFuncClient) Send(m *BasicRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingPracticesServiceClientSideStreamFuncClient) CloseAndRecv() (*BasicResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BasicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingPracticesServiceClient) BiDirectionalStreamFunc(ctx context.Context, opts ...grpc.CallOption) (StreamingPracticesService_BiDirectionalStreamFuncClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingPracticesService_ServiceDesc.Streams[2], "/StreamingPracticesService/BiDirectionalStreamFunc", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingPracticesServiceBiDirectionalStreamFuncClient{stream}
	return x, nil
}

type StreamingPracticesService_BiDirectionalStreamFuncClient interface {
	Send(*BasicRequest) error
	Recv() (*BasicResponse, error)
	grpc.ClientStream
}

type streamingPracticesServiceBiDirectionalStreamFuncClient struct {
	grpc.ClientStream
}

func (x *streamingPracticesServiceBiDirectionalStreamFuncClient) Send(m *BasicRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingPracticesServiceBiDirectionalStreamFuncClient) Recv() (*BasicResponse, error) {
	m := new(BasicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingPracticesServiceServer is the server API for StreamingPracticesService service.
// All implementations must embed UnimplementedStreamingPracticesServiceServer
// for forward compatibility
type StreamingPracticesServiceServer interface {
	ServerSideStreamFunc(*BasicRequest, StreamingPracticesService_ServerSideStreamFuncServer) error
	ClientSideStreamFunc(StreamingPracticesService_ClientSideStreamFuncServer) error
	BiDirectionalStreamFunc(StreamingPracticesService_BiDirectionalStreamFuncServer) error
	mustEmbedUnimplementedStreamingPracticesServiceServer()
}

// UnimplementedStreamingPracticesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingPracticesServiceServer struct {
}

func (UnimplementedStreamingPracticesServiceServer) ServerSideStreamFunc(*BasicRequest, StreamingPracticesService_ServerSideStreamFuncServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStreamFunc not implemented")
}
func (UnimplementedStreamingPracticesServiceServer) ClientSideStreamFunc(StreamingPracticesService_ClientSideStreamFuncServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSideStreamFunc not implemented")
}
func (UnimplementedStreamingPracticesServiceServer) BiDirectionalStreamFunc(StreamingPracticesService_BiDirectionalStreamFuncServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalStreamFunc not implemented")
}
func (UnimplementedStreamingPracticesServiceServer) mustEmbedUnimplementedStreamingPracticesServiceServer() {
}

// UnsafeStreamingPracticesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingPracticesServiceServer will
// result in compilation errors.
type UnsafeStreamingPracticesServiceServer interface {
	mustEmbedUnimplementedStreamingPracticesServiceServer()
}

func RegisterStreamingPracticesServiceServer(s grpc.ServiceRegistrar, srv StreamingPracticesServiceServer) {
	s.RegisterService(&StreamingPracticesService_ServiceDesc, srv)
}

func _StreamingPracticesService_ServerSideStreamFunc_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BasicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingPracticesServiceServer).ServerSideStreamFunc(m, &streamingPracticesServiceServerSideStreamFuncServer{stream})
}

type StreamingPracticesService_ServerSideStreamFuncServer interface {
	Send(*BasicResponse) error
	grpc.ServerStream
}

type streamingPracticesServiceServerSideStreamFuncServer struct {
	grpc.ServerStream
}

func (x *streamingPracticesServiceServerSideStreamFuncServer) Send(m *BasicResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamingPracticesService_ClientSideStreamFunc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingPracticesServiceServer).ClientSideStreamFunc(&streamingPracticesServiceClientSideStreamFuncServer{stream})
}

type StreamingPracticesService_ClientSideStreamFuncServer interface {
	SendAndClose(*BasicResponse) error
	Recv() (*BasicRequest, error)
	grpc.ServerStream
}

type streamingPracticesServiceClientSideStreamFuncServer struct {
	grpc.ServerStream
}

func (x *streamingPracticesServiceClientSideStreamFuncServer) SendAndClose(m *BasicResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingPracticesServiceClientSideStreamFuncServer) Recv() (*BasicRequest, error) {
	m := new(BasicRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamingPracticesService_BiDirectionalStreamFunc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingPracticesServiceServer).BiDirectionalStreamFunc(&streamingPracticesServiceBiDirectionalStreamFuncServer{stream})
}

type StreamingPracticesService_BiDirectionalStreamFuncServer interface {
	Send(*BasicResponse) error
	Recv() (*BasicRequest, error)
	grpc.ServerStream
}

type streamingPracticesServiceBiDirectionalStreamFuncServer struct {
	grpc.ServerStream
}

func (x *streamingPracticesServiceBiDirectionalStreamFuncServer) Send(m *BasicResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingPracticesServiceBiDirectionalStreamFuncServer) Recv() (*BasicRequest, error) {
	m := new(BasicRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingPracticesService_ServiceDesc is the grpc.ServiceDesc for StreamingPracticesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingPracticesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StreamingPracticesService",
	HandlerType: (*StreamingPracticesServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStreamFunc",
			Handler:       _StreamingPracticesService_ServerSideStreamFunc_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSideStreamFunc",
			Handler:       _StreamingPracticesService_ClientSideStreamFunc_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BiDirectionalStreamFunc",
			Handler:       _StreamingPracticesService_BiDirectionalStreamFunc_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "my.proto",
}