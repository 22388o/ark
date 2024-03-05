// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package arkv1

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

// ArkServiceClient is the client API for ArkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArkServiceClient interface {
	RegisterPayment(ctx context.Context, in *RegisterPaymentRequest, opts ...grpc.CallOption) (*RegisterPaymentResponse, error)
	ClaimPayment(ctx context.Context, in *ClaimPaymentRequest, opts ...grpc.CallOption) (*ClaimPaymentResponse, error)
	FinalizePayment(ctx context.Context, in *FinalizePaymentRequest, opts ...grpc.CallOption) (*FinalizePaymentResponse, error)
	GetRound(ctx context.Context, in *GetRoundRequest, opts ...grpc.CallOption) (*GetRoundResponse, error)
	GetEventStream(ctx context.Context, in *GetEventStreamRequest, opts ...grpc.CallOption) (ArkService_GetEventStreamClient, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	ListVtxos(ctx context.Context, in *ListVtxosRequest, opts ...grpc.CallOption) (*ListVtxosResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	Onboard(ctx context.Context, in *OnboardRequest, opts ...grpc.CallOption) (*OnboardResponse, error)
}

type arkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArkServiceClient(cc grpc.ClientConnInterface) ArkServiceClient {
	return &arkServiceClient{cc}
}

func (c *arkServiceClient) RegisterPayment(ctx context.Context, in *RegisterPaymentRequest, opts ...grpc.CallOption) (*RegisterPaymentResponse, error) {
	out := new(RegisterPaymentResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/RegisterPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arkServiceClient) ClaimPayment(ctx context.Context, in *ClaimPaymentRequest, opts ...grpc.CallOption) (*ClaimPaymentResponse, error) {
	out := new(ClaimPaymentResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/ClaimPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arkServiceClient) FinalizePayment(ctx context.Context, in *FinalizePaymentRequest, opts ...grpc.CallOption) (*FinalizePaymentResponse, error) {
	out := new(FinalizePaymentResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/FinalizePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arkServiceClient) GetRound(ctx context.Context, in *GetRoundRequest, opts ...grpc.CallOption) (*GetRoundResponse, error) {
	out := new(GetRoundResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/GetRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arkServiceClient) GetEventStream(ctx context.Context, in *GetEventStreamRequest, opts ...grpc.CallOption) (ArkService_GetEventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArkService_ServiceDesc.Streams[0], "/ark.v1.ArkService/GetEventStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &arkServiceGetEventStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArkService_GetEventStreamClient interface {
	Recv() (*GetEventStreamResponse, error)
	grpc.ClientStream
}

type arkServiceGetEventStreamClient struct {
	grpc.ClientStream
}

func (x *arkServiceGetEventStreamClient) Recv() (*GetEventStreamResponse, error) {
	m := new(GetEventStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arkServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arkServiceClient) ListVtxos(ctx context.Context, in *ListVtxosRequest, opts ...grpc.CallOption) (*ListVtxosResponse, error) {
	out := new(ListVtxosResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/ListVtxos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arkServiceClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arkServiceClient) Onboard(ctx context.Context, in *OnboardRequest, opts ...grpc.CallOption) (*OnboardResponse, error) {
	out := new(OnboardResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ArkService/Onboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArkServiceServer is the server API for ArkService service.
// All implementations should embed UnimplementedArkServiceServer
// for forward compatibility
type ArkServiceServer interface {
	RegisterPayment(context.Context, *RegisterPaymentRequest) (*RegisterPaymentResponse, error)
	ClaimPayment(context.Context, *ClaimPaymentRequest) (*ClaimPaymentResponse, error)
	FinalizePayment(context.Context, *FinalizePaymentRequest) (*FinalizePaymentResponse, error)
	GetRound(context.Context, *GetRoundRequest) (*GetRoundResponse, error)
	GetEventStream(*GetEventStreamRequest, ArkService_GetEventStreamServer) error
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	ListVtxos(context.Context, *ListVtxosRequest) (*ListVtxosResponse, error)
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	Onboard(context.Context, *OnboardRequest) (*OnboardResponse, error)
}

// UnimplementedArkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedArkServiceServer struct {
}

func (UnimplementedArkServiceServer) RegisterPayment(context.Context, *RegisterPaymentRequest) (*RegisterPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPayment not implemented")
}
func (UnimplementedArkServiceServer) ClaimPayment(context.Context, *ClaimPaymentRequest) (*ClaimPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimPayment not implemented")
}
func (UnimplementedArkServiceServer) FinalizePayment(context.Context, *FinalizePaymentRequest) (*FinalizePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizePayment not implemented")
}
func (UnimplementedArkServiceServer) GetRound(context.Context, *GetRoundRequest) (*GetRoundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRound not implemented")
}
func (UnimplementedArkServiceServer) GetEventStream(*GetEventStreamRequest, ArkService_GetEventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEventStream not implemented")
}
func (UnimplementedArkServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedArkServiceServer) ListVtxos(context.Context, *ListVtxosRequest) (*ListVtxosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVtxos not implemented")
}
func (UnimplementedArkServiceServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedArkServiceServer) Onboard(context.Context, *OnboardRequest) (*OnboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Onboard not implemented")
}

// UnsafeArkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArkServiceServer will
// result in compilation errors.
type UnsafeArkServiceServer interface {
	mustEmbedUnimplementedArkServiceServer()
}

func RegisterArkServiceServer(s grpc.ServiceRegistrar, srv ArkServiceServer) {
	s.RegisterService(&ArkService_ServiceDesc, srv)
}

func _ArkService_RegisterPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).RegisterPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/RegisterPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).RegisterPayment(ctx, req.(*RegisterPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArkService_ClaimPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).ClaimPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/ClaimPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).ClaimPayment(ctx, req.(*ClaimPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArkService_FinalizePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).FinalizePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/FinalizePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).FinalizePayment(ctx, req.(*FinalizePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArkService_GetRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).GetRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/GetRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).GetRound(ctx, req.(*GetRoundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArkService_GetEventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArkServiceServer).GetEventStream(m, &arkServiceGetEventStreamServer{stream})
}

type ArkService_GetEventStreamServer interface {
	Send(*GetEventStreamResponse) error
	grpc.ServerStream
}

type arkServiceGetEventStreamServer struct {
	grpc.ServerStream
}

func (x *arkServiceGetEventStreamServer) Send(m *GetEventStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ArkService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArkService_ListVtxos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVtxosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).ListVtxos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/ListVtxos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).ListVtxos(ctx, req.(*ListVtxosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArkService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArkService_Onboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArkServiceServer).Onboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ArkService/Onboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArkServiceServer).Onboard(ctx, req.(*OnboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArkService_ServiceDesc is the grpc.ServiceDesc for ArkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ark.v1.ArkService",
	HandlerType: (*ArkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPayment",
			Handler:    _ArkService_RegisterPayment_Handler,
		},
		{
			MethodName: "ClaimPayment",
			Handler:    _ArkService_ClaimPayment_Handler,
		},
		{
			MethodName: "FinalizePayment",
			Handler:    _ArkService_FinalizePayment_Handler,
		},
		{
			MethodName: "GetRound",
			Handler:    _ArkService_GetRound_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ArkService_Ping_Handler,
		},
		{
			MethodName: "ListVtxos",
			Handler:    _ArkService_ListVtxos_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _ArkService_GetInfo_Handler,
		},
		{
			MethodName: "Onboard",
			Handler:    _ArkService_Onboard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEventStream",
			Handler:       _ArkService_GetEventStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ark/v1/service.proto",
}