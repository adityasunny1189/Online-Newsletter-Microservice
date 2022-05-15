// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/admin.proto

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Status, error)
	CreatePlan(ctx context.Context, in *PlanInfo, opts ...grpc.CallOption) (*ID, error)
	DeletePlan(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error)
	CreateNews(ctx context.Context, in *NewsInfo, opts ...grpc.CallOption) (*ID, error)
	GetPlans(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AdminService_GetPlansClient, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AdminService_GetUsersClient, error)
	GetSubscriptions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AdminService_GetSubscriptionsClient, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/admin.AdminService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreatePlan(ctx context.Context, in *PlanInfo, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/admin.AdminService/CreatePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeletePlan(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/admin.AdminService/DeletePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateNews(ctx context.Context, in *NewsInfo, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/admin.AdminService/CreateNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetPlans(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AdminService_GetPlansClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdminService_ServiceDesc.Streams[0], "/admin.AdminService/GetPlans", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminServiceGetPlansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdminService_GetPlansClient interface {
	Recv() (*PlanInfo, error)
	grpc.ClientStream
}

type adminServiceGetPlansClient struct {
	grpc.ClientStream
}

func (x *adminServiceGetPlansClient) Recv() (*PlanInfo, error) {
	m := new(PlanInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminServiceClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AdminService_GetUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdminService_ServiceDesc.Streams[1], "/admin.AdminService/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminServiceGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdminService_GetUsersClient interface {
	Recv() (*UserInfo, error)
	grpc.ClientStream
}

type adminServiceGetUsersClient struct {
	grpc.ClientStream
}

func (x *adminServiceGetUsersClient) Recv() (*UserInfo, error) {
	m := new(UserInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminServiceClient) GetSubscriptions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AdminService_GetSubscriptionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdminService_ServiceDesc.Streams[2], "/admin.AdminService/GetSubscriptions", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminServiceGetSubscriptionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdminService_GetSubscriptionsClient interface {
	Recv() (*SubscriptionInfo, error)
	grpc.ClientStream
}

type adminServiceGetSubscriptionsClient struct {
	grpc.ClientStream
}

func (x *adminServiceGetSubscriptionsClient) Recv() (*SubscriptionInfo, error) {
	m := new(SubscriptionInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	Login(context.Context, *LoginRequest) (*Status, error)
	CreatePlan(context.Context, *PlanInfo) (*ID, error)
	DeletePlan(context.Context, *ID) (*Status, error)
	CreateNews(context.Context, *NewsInfo) (*ID, error)
	GetPlans(*Empty, AdminService_GetPlansServer) error
	GetUsers(*Empty, AdminService_GetUsersServer) error
	GetSubscriptions(*Empty, AdminService_GetSubscriptionsServer) error
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) Login(context.Context, *LoginRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAdminServiceServer) CreatePlan(context.Context, *PlanInfo) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlan not implemented")
}
func (UnimplementedAdminServiceServer) DeletePlan(context.Context, *ID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlan not implemented")
}
func (UnimplementedAdminServiceServer) CreateNews(context.Context, *NewsInfo) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNews not implemented")
}
func (UnimplementedAdminServiceServer) GetPlans(*Empty, AdminService_GetPlansServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPlans not implemented")
}
func (UnimplementedAdminServiceServer) GetUsers(*Empty, AdminService_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedAdminServiceServer) GetSubscriptions(*Empty, AdminService_GetSubscriptionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/CreatePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreatePlan(ctx, req.(*PlanInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeletePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeletePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/DeletePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeletePlan(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/CreateNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateNews(ctx, req.(*NewsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetPlans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServiceServer).GetPlans(m, &adminServiceGetPlansServer{stream})
}

type AdminService_GetPlansServer interface {
	Send(*PlanInfo) error
	grpc.ServerStream
}

type adminServiceGetPlansServer struct {
	grpc.ServerStream
}

func (x *adminServiceGetPlansServer) Send(m *PlanInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _AdminService_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServiceServer).GetUsers(m, &adminServiceGetUsersServer{stream})
}

type AdminService_GetUsersServer interface {
	Send(*UserInfo) error
	grpc.ServerStream
}

type adminServiceGetUsersServer struct {
	grpc.ServerStream
}

func (x *adminServiceGetUsersServer) Send(m *UserInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _AdminService_GetSubscriptions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServiceServer).GetSubscriptions(m, &adminServiceGetSubscriptionsServer{stream})
}

type AdminService_GetSubscriptionsServer interface {
	Send(*SubscriptionInfo) error
	grpc.ServerStream
}

type adminServiceGetSubscriptionsServer struct {
	grpc.ServerStream
}

func (x *adminServiceGetSubscriptionsServer) Send(m *SubscriptionInfo) error {
	return x.ServerStream.SendMsg(m)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AdminService_Login_Handler,
		},
		{
			MethodName: "CreatePlan",
			Handler:    _AdminService_CreatePlan_Handler,
		},
		{
			MethodName: "DeletePlan",
			Handler:    _AdminService_DeletePlan_Handler,
		},
		{
			MethodName: "CreateNews",
			Handler:    _AdminService_CreateNews_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPlans",
			Handler:       _AdminService_GetPlans_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetUsers",
			Handler:       _AdminService_GetUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSubscriptions",
			Handler:       _AdminService_GetSubscriptions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/admin.proto",
}