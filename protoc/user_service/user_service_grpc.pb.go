// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: user_service.proto

package user_service

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

const (
	UserService_Register_FullMethodName         = "/userService.UserService/Register"
	UserService_Login_FullMethodName            = "/userService.UserService/Login"
	UserService_Logout_FullMethodName           = "/userService.UserService/LogoutByToken"
	UserService_SearchRole_FullMethodName       = "/userService.UserService/SearchRole"
	UserService_SearchGroup_FullMethodName      = "/userService.UserService/SearchGroup"
	UserService_SearchPermission_FullMethodName = "/userService.UserService/SearchPermission"
	UserService_Edit_FullMethodName             = "/userService.UserService/Edit"
	UserService_DelRole_FullMethodName          = "/userService.UserService/DelRole"
	UserService_DelGroup_FullMethodName         = "/userService.UserService/DelGroup"
	UserService_CreateRole_FullMethodName       = "/userService.UserService/CreateRole"
	UserService_CreateGroup_FullMethodName      = "/userService.UserService/CreateGroup"
	UserService_AddGroupMembers_FullMethodName  = "/userService.UserService/AddGroupMembers"
	UserService_ShowAllGroup_FullMethodName     = "/userService.UserService/ShowAllGroup"
	UserService_ShowAllRole_FullMethodName      = "/userService.UserService/ShowAllRole"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*Stdout, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Stdout, error)
	Logout(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error)
	SearchRole(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error)
	SearchGroup(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error)
	SearchPermission(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error)
	Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*Stdout, error)
	DelRole(ctx context.Context, in *DelRoleReq, opts ...grpc.CallOption) (*Stdout, error)
	DelGroup(ctx context.Context, in *DelGroupReq, opts ...grpc.CallOption) (*Stdout, error)
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*Stdout, error)
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*Stdout, error)
	AddGroupMembers(ctx context.Context, in *AddGroupMembersReq, opts ...grpc.CallOption) (*Stdout, error)
	ShowAllGroup(ctx context.Context, in *ShowAllGroupReq, opts ...grpc.CallOption) (*Stdout, error)
	ShowAllRole(ctx context.Context, in *ShowAllRoleReq, opts ...grpc.CallOption) (*Stdout, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SearchRole(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_SearchRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SearchGroup(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_SearchGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SearchPermission(ctx context.Context, in *NoneReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_SearchPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_Edit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DelRole(ctx context.Context, in *DelRoleReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_DelRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DelGroup(ctx context.Context, in *DelGroupReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_DelGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddGroupMembers(ctx context.Context, in *AddGroupMembersReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_AddGroupMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ShowAllGroup(ctx context.Context, in *ShowAllGroupReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_ShowAllGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ShowAllRole(ctx context.Context, in *ShowAllRoleReq, opts ...grpc.CallOption) (*Stdout, error) {
	out := new(Stdout)
	err := c.cc.Invoke(ctx, UserService_ShowAllRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *RegisterReq) (*Stdout, error)
	Login(context.Context, *LoginReq) (*Stdout, error)
	Logout(context.Context, *NoneReq) (*Stdout, error)
	SearchRole(context.Context, *NoneReq) (*Stdout, error)
	SearchGroup(context.Context, *NoneReq) (*Stdout, error)
	SearchPermission(context.Context, *NoneReq) (*Stdout, error)
	Edit(context.Context, *EditReq) (*Stdout, error)
	DelRole(context.Context, *DelRoleReq) (*Stdout, error)
	DelGroup(context.Context, *DelGroupReq) (*Stdout, error)
	CreateRole(context.Context, *CreateRoleReq) (*Stdout, error)
	CreateGroup(context.Context, *CreateGroupReq) (*Stdout, error)
	AddGroupMembers(context.Context, *AddGroupMembersReq) (*Stdout, error)
	ShowAllGroup(context.Context, *ShowAllGroupReq) (*Stdout, error)
	ShowAllRole(context.Context, *ShowAllRoleReq) (*Stdout, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Logout(context.Context, *NoneReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutByToken not implemented")
}
func (UnimplementedUserServiceServer) SearchRole(context.Context, *NoneReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRole not implemented")
}
func (UnimplementedUserServiceServer) SearchGroup(context.Context, *NoneReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGroup not implemented")
}
func (UnimplementedUserServiceServer) SearchPermission(context.Context, *NoneReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPermission not implemented")
}
func (UnimplementedUserServiceServer) Edit(context.Context, *EditReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedUserServiceServer) DelRole(context.Context, *DelRoleReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelRole not implemented")
}
func (UnimplementedUserServiceServer) DelGroup(context.Context, *DelGroupReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelGroup not implemented")
}
func (UnimplementedUserServiceServer) CreateRole(context.Context, *CreateRoleReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedUserServiceServer) CreateGroup(context.Context, *CreateGroupReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedUserServiceServer) AddGroupMembers(context.Context, *AddGroupMembersReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupMembers not implemented")
}
func (UnimplementedUserServiceServer) ShowAllGroup(context.Context, *ShowAllGroupReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllGroup not implemented")
}
func (UnimplementedUserServiceServer) ShowAllRole(context.Context, *ShowAllRoleReq) (*Stdout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllRole not implemented")
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

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*NoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SearchRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SearchRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SearchRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SearchRole(ctx, req.(*NoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SearchGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SearchGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SearchGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SearchGroup(ctx, req.(*NoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SearchPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SearchPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SearchPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SearchPermission(ctx, req.(*NoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Edit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Edit(ctx, req.(*EditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DelRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DelRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DelRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DelRole(ctx, req.(*DelRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DelGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DelGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DelGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DelGroup(ctx, req.(*DelGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateGroup(ctx, req.(*CreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddGroupMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddGroupMembers(ctx, req.(*AddGroupMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ShowAllGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAllGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ShowAllGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ShowAllGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ShowAllGroup(ctx, req.(*ShowAllGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ShowAllRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAllRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ShowAllRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ShowAllRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ShowAllRole(ctx, req.(*ShowAllRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userService.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "LogoutByToken",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "SearchRole",
			Handler:    _UserService_SearchRole_Handler,
		},
		{
			MethodName: "SearchGroup",
			Handler:    _UserService_SearchGroup_Handler,
		},
		{
			MethodName: "SearchPermission",
			Handler:    _UserService_SearchPermission_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _UserService_Edit_Handler,
		},
		{
			MethodName: "DelRole",
			Handler:    _UserService_DelRole_Handler,
		},
		{
			MethodName: "DelGroup",
			Handler:    _UserService_DelGroup_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _UserService_CreateRole_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _UserService_CreateGroup_Handler,
		},
		{
			MethodName: "AddGroupMembers",
			Handler:    _UserService_AddGroupMembers_Handler,
		},
		{
			MethodName: "ShowAllGroup",
			Handler:    _UserService_ShowAllGroup_Handler,
		},
		{
			MethodName: "ShowAllRole",
			Handler:    _UserService_ShowAllRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
