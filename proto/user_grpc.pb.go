// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/user.proto

package proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*CreatedUser, error)
	GreetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_GreetUserClient, error)
	GreetMany(ctx context.Context, opts ...grpc.CallOption) (UserService_GreetManyClient, error)
	GreetManyWithMultipleGreet(ctx context.Context, opts ...grpc.CallOption) (UserService_GreetManyWithMultipleGreetClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*CreatedUser, error) {
	out := new(CreatedUser)
	err := c.cc.Invoke(ctx, "/greet_service.UserService/createUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GreetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_GreetUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/greet_service.UserService/greetUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGreetUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GreetUserClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type userServiceGreetUserClient struct {
	grpc.ClientStream
}

func (x *userServiceGreetUserClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GreetMany(ctx context.Context, opts ...grpc.CallOption) (UserService_GreetManyClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/greet_service.UserService/greetMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGreetManyClient{stream}
	return x, nil
}

type UserService_GreetManyClient interface {
	Send(*User) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type userServiceGreetManyClient struct {
	grpc.ClientStream
}

func (x *userServiceGreetManyClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGreetManyClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GreetManyWithMultipleGreet(ctx context.Context, opts ...grpc.CallOption) (UserService_GreetManyWithMultipleGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[2], "/greet_service.UserService/greetManyWithMultipleGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGreetManyWithMultipleGreetClient{stream}
	return x, nil
}

type UserService_GreetManyWithMultipleGreetClient interface {
	Send(*User) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type userServiceGreetManyWithMultipleGreetClient struct {
	grpc.ClientStream
}

func (x *userServiceGreetManyWithMultipleGreetClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGreetManyWithMultipleGreetClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*CreatedUser, error)
	GreetUser(*User, UserService_GreetUserServer) error
	GreetMany(UserService_GreetManyServer) error
	GreetManyWithMultipleGreet(UserService_GreetManyWithMultipleGreetServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *User) (*CreatedUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GreetUser(*User, UserService_GreetUserServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetUser not implemented")
}
func (UnimplementedUserServiceServer) GreetMany(UserService_GreetManyServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetMany not implemented")
}
func (UnimplementedUserServiceServer) GreetManyWithMultipleGreet(UserService_GreetManyWithMultipleGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyWithMultipleGreet not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet_service.UserService/createUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GreetUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GreetUser(m, &userServiceGreetUserServer{stream})
}

type UserService_GreetUserServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type userServiceGreetUserServer struct {
	grpc.ServerStream
}

func (x *userServiceGreetUserServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_GreetMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GreetMany(&userServiceGreetManyServer{stream})
}

type UserService_GreetManyServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*User, error)
	grpc.ServerStream
}

type userServiceGreetManyServer struct {
	grpc.ServerStream
}

func (x *userServiceGreetManyServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGreetManyServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_GreetManyWithMultipleGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GreetManyWithMultipleGreet(&userServiceGreetManyWithMultipleGreetServer{stream})
}

type UserService_GreetManyWithMultipleGreetServer interface {
	Send(*GreetResponse) error
	Recv() (*User, error)
	grpc.ServerStream
}

type userServiceGreetManyWithMultipleGreetServer struct {
	grpc.ServerStream
}

func (x *userServiceGreetManyWithMultipleGreetServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGreetManyWithMultipleGreetServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "greetUser",
			Handler:       _UserService_GreetUser_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "greetMany",
			Handler:       _UserService_GreetMany_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "greetManyWithMultipleGreet",
			Handler:       _UserService_GreetManyWithMultipleGreet_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/user.proto",
}
