// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: service.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdServiceClient is the client API for AdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdServiceClient interface {
	CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*AdResponse, error)
	ChangeAdStatus(ctx context.Context, in *ChangeAdStatusRequest, opts ...grpc.CallOption) (*AdResponse, error)
	UpdateAd(ctx context.Context, in *UpdateAdRequest, opts ...grpc.CallOption) (*AdResponse, error)
	ListAds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAdResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAd(ctx context.Context, in *DeleteAdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type adServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdServiceClient(cc grpc.ClientConnInterface) AdServiceClient {
	return &adServiceClient{cc}
}

func (c *adServiceClient) CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, "/ad.AdService/CreateAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) ChangeAdStatus(ctx context.Context, in *ChangeAdStatusRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, "/ad.AdService/ChangeAdStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) UpdateAd(ctx context.Context, in *UpdateAdRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, "/ad.AdService/UpdateAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) ListAds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAdResponse, error) {
	out := new(ListAdResponse)
	err := c.cc.Invoke(ctx, "/ad.AdService/ListAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/ad.AdService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/ad.AdService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ad.AdService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) DeleteAd(ctx context.Context, in *DeleteAdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ad.AdService/DeleteAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceServer is the server API for AdService service.
// All implementations must embed UnimplementedAdServiceServer
// for forward compatibility
type AdServiceServer interface {
	CreateAd(context.Context, *CreateAdRequest) (*AdResponse, error)
	ChangeAdStatus(context.Context, *ChangeAdStatusRequest) (*AdResponse, error)
	UpdateAd(context.Context, *UpdateAdRequest) (*AdResponse, error)
	ListAds(context.Context, *emptypb.Empty) (*ListAdResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	DeleteAd(context.Context, *DeleteAdRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAdServiceServer()
}

// UnimplementedAdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdServiceServer struct {
}

func (UnimplementedAdServiceServer) CreateAd(context.Context, *CreateAdRequest) (*AdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAd not implemented")
}
func (UnimplementedAdServiceServer) ChangeAdStatus(context.Context, *ChangeAdStatusRequest) (*AdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAdStatus not implemented")
}
func (UnimplementedAdServiceServer) UpdateAd(context.Context, *UpdateAdRequest) (*AdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAd not implemented")
}
func (UnimplementedAdServiceServer) ListAds(context.Context, *emptypb.Empty) (*ListAdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAds not implemented")
}
func (UnimplementedAdServiceServer) CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAdServiceServer) GetUser(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAdServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAdServiceServer) DeleteAd(context.Context, *DeleteAdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAd not implemented")
}
func (UnimplementedAdServiceServer) mustEmbedUnimplementedAdServiceServer() {}

// UnsafeAdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdServiceServer will
// result in compilation errors.
type UnsafeAdServiceServer interface {
	mustEmbedUnimplementedAdServiceServer()
}

func RegisterAdServiceServer(s grpc.ServiceRegistrar, srv AdServiceServer) {
	s.RegisterService(&AdService_ServiceDesc, srv)
}

func _AdService_CreateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).CreateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/CreateAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).CreateAd(ctx, req.(*CreateAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_ChangeAdStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAdStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).ChangeAdStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/ChangeAdStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).ChangeAdStatus(ctx, req.(*ChangeAdStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_UpdateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).UpdateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/UpdateAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).UpdateAd(ctx, req.(*UpdateAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_ListAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).ListAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/ListAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).ListAds(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_DeleteAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).DeleteAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AdService/DeleteAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).DeleteAd(ctx, req.(*DeleteAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdService_ServiceDesc is the grpc.ServiceDesc for AdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ad.AdService",
	HandlerType: (*AdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAd",
			Handler:    _AdService_CreateAd_Handler,
		},
		{
			MethodName: "ChangeAdStatus",
			Handler:    _AdService_ChangeAdStatus_Handler,
		},
		{
			MethodName: "UpdateAd",
			Handler:    _AdService_UpdateAd_Handler,
		},
		{
			MethodName: "ListAds",
			Handler:    _AdService_ListAds_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AdService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AdService_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AdService_DeleteUser_Handler,
		},
		{
			MethodName: "DeleteAd",
			Handler:    _AdService_DeleteAd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
