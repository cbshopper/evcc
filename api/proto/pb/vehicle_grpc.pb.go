// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/vehicle.proto

package pb

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

// VehicleClient is the client API for Vehicle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewReply, error)
	SoC(ctx context.Context, in *SoCRequest, opts ...grpc.CallOption) (*SoCReply, error)
}

type vehicleClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleClient(cc grpc.ClientConnInterface) VehicleClient {
	return &vehicleClient{cc}
}

func (c *vehicleClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewReply, error) {
	out := new(NewReply)
	err := c.cc.Invoke(ctx, "/Vehicle/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) SoC(ctx context.Context, in *SoCRequest, opts ...grpc.CallOption) (*SoCReply, error) {
	out := new(SoCReply)
	err := c.cc.Invoke(ctx, "/Vehicle/SoC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServer is the server API for Vehicle service.
// All implementations must embed UnimplementedVehicleServer
// for forward compatibility
type VehicleServer interface {
	New(context.Context, *NewRequest) (*NewReply, error)
	SoC(context.Context, *SoCRequest) (*SoCReply, error)
	mustEmbedUnimplementedVehicleServer()
}

// UnimplementedVehicleServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleServer struct {
}

func (UnimplementedVehicleServer) New(context.Context, *NewRequest) (*NewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (UnimplementedVehicleServer) SoC(context.Context, *SoCRequest) (*SoCReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoC not implemented")
}
func (UnimplementedVehicleServer) mustEmbedUnimplementedVehicleServer() {}

// UnsafeVehicleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServer will
// result in compilation errors.
type UnsafeVehicleServer interface {
	mustEmbedUnimplementedVehicleServer()
}

func RegisterVehicleServer(s grpc.ServiceRegistrar, srv VehicleServer) {
	s.RegisterService(&Vehicle_ServiceDesc, srv)
}

func _Vehicle_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Vehicle/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_SoC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SoCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).SoC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Vehicle/SoC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).SoC(ctx, req.(*SoCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vehicle_ServiceDesc is the grpc.ServiceDesc for Vehicle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vehicle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Vehicle",
	HandlerType: (*VehicleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Vehicle_New_Handler,
		},
		{
			MethodName: "SoC",
			Handler:    _Vehicle_SoC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vehicle.proto",
}
