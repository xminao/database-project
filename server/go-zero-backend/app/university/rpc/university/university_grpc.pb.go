// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: university.proto

package university

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

// UniversityRpcClient is the client API for UniversityRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UniversityRpcClient interface {
	GetStudentInfo(ctx context.Context, in *GetStudentInfoReq, opts ...grpc.CallOption) (*GetStudentInfoResp, error)
}

type universityRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUniversityRpcClient(cc grpc.ClientConnInterface) UniversityRpcClient {
	return &universityRpcClient{cc}
}

func (c *universityRpcClient) GetStudentInfo(ctx context.Context, in *GetStudentInfoReq, opts ...grpc.CallOption) (*GetStudentInfoResp, error) {
	out := new(GetStudentInfoResp)
	err := c.cc.Invoke(ctx, "/university.UniversityRpc/GetStudentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UniversityRpcServer is the server API for UniversityRpc service.
// All implementations must embed UnimplementedUniversityRpcServer
// for forward compatibility
type UniversityRpcServer interface {
	GetStudentInfo(context.Context, *GetStudentInfoReq) (*GetStudentInfoResp, error)
	mustEmbedUnimplementedUniversityRpcServer()
}

// UnimplementedUniversityRpcServer must be embedded to have forward compatible implementations.
type UnimplementedUniversityRpcServer struct {
}

func (UnimplementedUniversityRpcServer) GetStudentInfo(context.Context, *GetStudentInfoReq) (*GetStudentInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentInfo not implemented")
}
func (UnimplementedUniversityRpcServer) mustEmbedUnimplementedUniversityRpcServer() {}

// UnsafeUniversityRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UniversityRpcServer will
// result in compilation errors.
type UnsafeUniversityRpcServer interface {
	mustEmbedUnimplementedUniversityRpcServer()
}

func RegisterUniversityRpcServer(s grpc.ServiceRegistrar, srv UniversityRpcServer) {
	s.RegisterService(&UniversityRpc_ServiceDesc, srv)
}

func _UniversityRpc_GetStudentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityRpcServer).GetStudentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/university.UniversityRpc/GetStudentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityRpcServer).GetStudentInfo(ctx, req.(*GetStudentInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UniversityRpc_ServiceDesc is the grpc.ServiceDesc for UniversityRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UniversityRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "university.UniversityRpc",
	HandlerType: (*UniversityRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudentInfo",
			Handler:    _UniversityRpc_GetStudentInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "university.proto",
}
