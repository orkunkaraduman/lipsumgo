// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: lipsumgo/lipsumgo.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Time(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*timestamppb.Timestamp, error)
	GetSentence(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiGetSentenceReply, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/lipsumgo.Api/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Time(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*timestamppb.Timestamp, error) {
	out := new(timestamppb.Timestamp)
	err := c.cc.Invoke(ctx, "/lipsumgo.Api/Time", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetSentence(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiGetSentenceReply, error) {
	out := new(ApiGetSentenceReply)
	err := c.cc.Invoke(ctx, "/lipsumgo.Api/GetSentence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Time(context.Context, *emptypb.Empty) (*timestamppb.Timestamp, error)
	GetSentence(context.Context, *emptypb.Empty) (*ApiGetSentenceReply, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedApiServer) Time(context.Context, *emptypb.Empty) (*timestamppb.Timestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Time not implemented")
}
func (UnimplementedApiServer) GetSentence(context.Context, *emptypb.Empty) (*ApiGetSentenceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSentence not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lipsumgo.Api/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Time_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Time(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lipsumgo.Api/Time",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Time(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetSentence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetSentence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lipsumgo.Api/GetSentence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetSentence(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lipsumgo.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Api_Ping_Handler,
		},
		{
			MethodName: "Time",
			Handler:    _Api_Time_Handler,
		},
		{
			MethodName: "GetSentence",
			Handler:    _Api_GetSentence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lipsumgo/lipsumgo.proto",
}
