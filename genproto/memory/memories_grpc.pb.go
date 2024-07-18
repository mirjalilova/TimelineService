// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: memories.proto

package memory

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MemoryService_Create_FullMethodName              = "/memory.MemoryService/Create"
	MemoryService_Get_FullMethodName                 = "/memory.MemoryService/Get"
	MemoryService_GetHistoricalMemory_FullMethodName = "/memory.MemoryService/GetHistoricalMemory"
	MemoryService_GetByTagMemory_FullMethodName      = "/memory.MemoryService/GetByTagMemory"
	MemoryService_GetAll_FullMethodName              = "/memory.MemoryService/GetAll"
	MemoryService_Update_FullMethodName              = "/memory.MemoryService/Update"
	MemoryService_Delete_FullMethodName              = "/memory.MemoryService/Delete"
)

// MemoryServiceClient is the client API for MemoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoryServiceClient interface {
	Create(ctx context.Context, in *MemoryCreate, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*MemoryRes, error)
	GetHistoricalMemory(ctx context.Context, in *GetByUser, opts ...grpc.CallOption) (*GetAllRes, error)
	GetByTagMemory(ctx context.Context, in *GetByTag, opts ...grpc.CallOption) (*GetAllRes, error)
	GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllRes, error)
	Update(ctx context.Context, in *MemoryUpdate, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Void, error)
}

type memoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoryServiceClient(cc grpc.ClientConnInterface) MemoryServiceClient {
	return &memoryServiceClient{cc}
}

func (c *memoryServiceClient) Create(ctx context.Context, in *MemoryCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, MemoryService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) Get(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*MemoryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemoryRes)
	err := c.cc.Invoke(ctx, MemoryService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) GetHistoricalMemory(ctx context.Context, in *GetByUser, opts ...grpc.CallOption) (*GetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllRes)
	err := c.cc.Invoke(ctx, MemoryService_GetHistoricalMemory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) GetByTagMemory(ctx context.Context, in *GetByTag, opts ...grpc.CallOption) (*GetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllRes)
	err := c.cc.Invoke(ctx, MemoryService_GetByTagMemory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllRes)
	err := c.cc.Invoke(ctx, MemoryService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) Update(ctx context.Context, in *MemoryUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, MemoryService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) Delete(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, MemoryService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoryServiceServer is the server API for MemoryService service.
// All implementations must embed UnimplementedMemoryServiceServer
// for forward compatibility
type MemoryServiceServer interface {
	Create(context.Context, *MemoryCreate) (*Void, error)
	Get(context.Context, *GetById) (*MemoryRes, error)
	GetHistoricalMemory(context.Context, *GetByUser) (*GetAllRes, error)
	GetByTagMemory(context.Context, *GetByTag) (*GetAllRes, error)
	GetAll(context.Context, *GetAllReq) (*GetAllRes, error)
	Update(context.Context, *MemoryUpdate) (*Void, error)
	Delete(context.Context, *GetById) (*Void, error)
	mustEmbedUnimplementedMemoryServiceServer()
}

// UnimplementedMemoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemoryServiceServer struct {
}

func (UnimplementedMemoryServiceServer) Create(context.Context, *MemoryCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMemoryServiceServer) Get(context.Context, *GetById) (*MemoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMemoryServiceServer) GetHistoricalMemory(context.Context, *GetByUser) (*GetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoricalMemory not implemented")
}
func (UnimplementedMemoryServiceServer) GetByTagMemory(context.Context, *GetByTag) (*GetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByTagMemory not implemented")
}
func (UnimplementedMemoryServiceServer) GetAll(context.Context, *GetAllReq) (*GetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMemoryServiceServer) Update(context.Context, *MemoryUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMemoryServiceServer) Delete(context.Context, *GetById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMemoryServiceServer) mustEmbedUnimplementedMemoryServiceServer() {}

// UnsafeMemoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoryServiceServer will
// result in compilation errors.
type UnsafeMemoryServiceServer interface {
	mustEmbedUnimplementedMemoryServiceServer()
}

func RegisterMemoryServiceServer(s grpc.ServiceRegistrar, srv MemoryServiceServer) {
	s.RegisterService(&MemoryService_ServiceDesc, srv)
}

func _MemoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemoryCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoryService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).Create(ctx, req.(*MemoryCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoryService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).Get(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_GetHistoricalMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).GetHistoricalMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoryService_GetHistoricalMemory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).GetHistoricalMemory(ctx, req.(*GetByUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_GetByTagMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).GetByTagMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoryService_GetByTagMemory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).GetByTagMemory(ctx, req.(*GetByTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoryService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).GetAll(ctx, req.(*GetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemoryUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoryService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).Update(ctx, req.(*MemoryUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoryService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).Delete(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

// MemoryService_ServiceDesc is the grpc.ServiceDesc for MemoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memory.MemoryService",
	HandlerType: (*MemoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MemoryService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MemoryService_Get_Handler,
		},
		{
			MethodName: "GetHistoricalMemory",
			Handler:    _MemoryService_GetHistoricalMemory_Handler,
		},
		{
			MethodName: "GetByTagMemory",
			Handler:    _MemoryService_GetByTagMemory_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MemoryService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MemoryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MemoryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "memories.proto",
}