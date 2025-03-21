// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: ai_agent.proto

package ai_agent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AIAgent_Single_FullMethodName     = "/ai_agent.AIAgent/Single"
	AIAgent_StreamChat_FullMethodName = "/ai_agent.AIAgent/StreamChat"
)

// AIAgentClient is the client API for AIAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AIAgentClient interface {
	Single(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
	StreamChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatResponse], error)
}

type aIAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewAIAgentClient(cc grpc.ClientConnInterface) AIAgentClient {
	return &aIAgentClient{cc}
}

func (c *aIAgentClient) Single(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, AIAgent_Single_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIAgentClient) StreamChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AIAgent_ServiceDesc.Streams[0], AIAgent_StreamChat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ChatRequest, ChatResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AIAgent_StreamChatClient = grpc.ServerStreamingClient[ChatResponse]

// AIAgentServer is the server API for AIAgent service.
// All implementations must embed UnimplementedAIAgentServer
// for forward compatibility.
type AIAgentServer interface {
	Single(context.Context, *ChatRequest) (*ChatResponse, error)
	StreamChat(*ChatRequest, grpc.ServerStreamingServer[ChatResponse]) error
	mustEmbedUnimplementedAIAgentServer()
}

// UnimplementedAIAgentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAIAgentServer struct{}

func (UnimplementedAIAgentServer) Single(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Single not implemented")
}
func (UnimplementedAIAgentServer) StreamChat(*ChatRequest, grpc.ServerStreamingServer[ChatResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamChat not implemented")
}
func (UnimplementedAIAgentServer) mustEmbedUnimplementedAIAgentServer() {}
func (UnimplementedAIAgentServer) testEmbeddedByValue()                 {}

// UnsafeAIAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIAgentServer will
// result in compilation errors.
type UnsafeAIAgentServer interface {
	mustEmbedUnimplementedAIAgentServer()
}

func RegisterAIAgentServer(s grpc.ServiceRegistrar, srv AIAgentServer) {
	// If the following call pancis, it indicates UnimplementedAIAgentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AIAgent_ServiceDesc, srv)
}

func _AIAgent_Single_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAgentServer).Single(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAgent_Single_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAgentServer).Single(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIAgent_StreamChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AIAgentServer).StreamChat(m, &grpc.GenericServerStream[ChatRequest, ChatResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AIAgent_StreamChatServer = grpc.ServerStreamingServer[ChatResponse]

// AIAgent_ServiceDesc is the grpc.ServiceDesc for AIAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ai_agent.AIAgent",
	HandlerType: (*AIAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Single",
			Handler:    _AIAgent_Single_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChat",
			Handler:       _AIAgent_StreamChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ai_agent.proto",
}
