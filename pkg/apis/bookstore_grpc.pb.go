// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apis

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

// BookstoreServiceClient is the client API for BookstoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookstoreServiceClient interface {
	GetBookInfoList(ctx context.Context, in *GetBookInfoListRequest, opts ...grpc.CallOption) (*GetBookInfoListResponse, error)
}

type bookstoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookstoreServiceClient(cc grpc.ClientConnInterface) BookstoreServiceClient {
	return &bookstoreServiceClient{cc}
}

func (c *bookstoreServiceClient) GetBookInfoList(ctx context.Context, in *GetBookInfoListRequest, opts ...grpc.CallOption) (*GetBookInfoListResponse, error) {
	out := new(GetBookInfoListResponse)
	err := c.cc.Invoke(ctx, "/BookstoreService/GetBookInfoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookstoreServiceServer is the server API for BookstoreService service.
// All implementations must embed UnimplementedBookstoreServiceServer
// for forward compatibility
type BookstoreServiceServer interface {
	GetBookInfoList(context.Context, *GetBookInfoListRequest) (*GetBookInfoListResponse, error)
	mustEmbedUnimplementedBookstoreServiceServer()
}

// UnimplementedBookstoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookstoreServiceServer struct {
}

func (UnimplementedBookstoreServiceServer) GetBookInfoList(context.Context, *GetBookInfoListRequest) (*GetBookInfoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfoList not implemented")
}
func (UnimplementedBookstoreServiceServer) mustEmbedUnimplementedBookstoreServiceServer() {}

// UnsafeBookstoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookstoreServiceServer will
// result in compilation errors.
type UnsafeBookstoreServiceServer interface {
	mustEmbedUnimplementedBookstoreServiceServer()
}

func RegisterBookstoreServiceServer(s grpc.ServiceRegistrar, srv BookstoreServiceServer) {
	s.RegisterService(&BookstoreService_ServiceDesc, srv)
}

func _BookstoreService_GetBookInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookInfoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServiceServer).GetBookInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookstoreService/GetBookInfoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServiceServer).GetBookInfoList(ctx, req.(*GetBookInfoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookstoreService_ServiceDesc is the grpc.ServiceDesc for BookstoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookstoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookstoreService",
	HandlerType: (*BookstoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfoList",
			Handler:    _BookstoreService_GetBookInfoList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookstore.proto",
}