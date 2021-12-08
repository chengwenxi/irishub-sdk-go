// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coinswap/query.proto

package coinswap

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/irisnet/irishub-sdk-go/types"
	query "github.com/irisnet/irishub-sdk-go/types/query"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryLiquidityPoolRequest is request type for the Query/LiquidityPool RPC
// method
type QueryLiquidityPoolRequest struct {
	LptDenom string `protobuf:"bytes,1,opt,name=lpt_denom,json=lptDenom,proto3" json:"lpt_denom,omitempty"`
}

func (m *QueryLiquidityPoolRequest) Reset()         { *m = QueryLiquidityPoolRequest{} }
func (m *QueryLiquidityPoolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidityPoolRequest) ProtoMessage()    {}
func (*QueryLiquidityPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cabf8423404f12f, []int{0}
}
func (m *QueryLiquidityPoolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidityPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidityPoolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidityPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidityPoolRequest.Merge(m, src)
}
func (m *QueryLiquidityPoolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidityPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidityPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidityPoolRequest proto.InternalMessageInfo

func (m *QueryLiquidityPoolRequest) GetLptDenom() string {
	if m != nil {
		return m.LptDenom
	}
	return ""
}

// QueryLiquidityPoolResponse is response type for the Query/LiquidityPool RPC
// method
type QueryLiquidityPoolResponse struct {
	Pool PoolInfo `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool"`
}

func (m *QueryLiquidityPoolResponse) Reset()         { *m = QueryLiquidityPoolResponse{} }
func (m *QueryLiquidityPoolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidityPoolResponse) ProtoMessage()    {}
func (*QueryLiquidityPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cabf8423404f12f, []int{1}
}
func (m *QueryLiquidityPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidityPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidityPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidityPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidityPoolResponse.Merge(m, src)
}
func (m *QueryLiquidityPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidityPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidityPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidityPoolResponse proto.InternalMessageInfo

func (m *QueryLiquidityPoolResponse) GetPool() PoolInfo {
	if m != nil {
		return m.Pool
	}
	return PoolInfo{}
}

// QueryLiquidityPoolsRequest is request type for the Query/LiquidityPools RPC
// method
type QueryLiquidityPoolsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryLiquidityPoolsRequest) Reset()         { *m = QueryLiquidityPoolsRequest{} }
func (m *QueryLiquidityPoolsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidityPoolsRequest) ProtoMessage()    {}
func (*QueryLiquidityPoolsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cabf8423404f12f, []int{2}
}
func (m *QueryLiquidityPoolsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidityPoolsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidityPoolsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidityPoolsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidityPoolsRequest.Merge(m, src)
}
func (m *QueryLiquidityPoolsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidityPoolsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidityPoolsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidityPoolsRequest proto.InternalMessageInfo

func (m *QueryLiquidityPoolsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryLiquidityPoolsResponse is response type for the Query/LiquidityPools RPC
// method
type QueryLiquidityPoolsResponse struct {
	Pools      []PoolInfo          `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryLiquidityPoolsResponse) Reset()         { *m = QueryLiquidityPoolsResponse{} }
func (m *QueryLiquidityPoolsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidityPoolsResponse) ProtoMessage()    {}
func (*QueryLiquidityPoolsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cabf8423404f12f, []int{3}
}
func (m *QueryLiquidityPoolsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidityPoolsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidityPoolsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidityPoolsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidityPoolsResponse.Merge(m, src)
}
func (m *QueryLiquidityPoolsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidityPoolsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidityPoolsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidityPoolsResponse proto.InternalMessageInfo

func (m *QueryLiquidityPoolsResponse) GetPools() []PoolInfo {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *QueryLiquidityPoolsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type PoolInfo struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// escrow account for deposit tokens
	EscrowAddress string `protobuf:"bytes,2,opt,name=escrow_address,json=escrowAddress,proto3" json:"escrow_address,omitempty"`
	// main token balance
	Standard types.Coin `protobuf:"bytes,3,opt,name=standard,proto3" json:"standard"`
	// counterparty token balance
	Token types.Coin `protobuf:"bytes,4,opt,name=token,proto3" json:"token"`
	// liquidity token balance
	Lpt types.Coin `protobuf:"bytes,5,opt,name=lpt,proto3" json:"lpt"`
	// liquidity pool fee
	Fee string `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *PoolInfo) Reset()         { *m = PoolInfo{} }
func (m *PoolInfo) String() string { return proto.CompactTextString(m) }
func (*PoolInfo) ProtoMessage()    {}
func (*PoolInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cabf8423404f12f, []int{4}
}
func (m *PoolInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolInfo.Merge(m, src)
}
func (m *PoolInfo) XXX_Size() int {
	return m.Size()
}
func (m *PoolInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PoolInfo proto.InternalMessageInfo

func (m *PoolInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PoolInfo) GetEscrowAddress() string {
	if m != nil {
		return m.EscrowAddress
	}
	return ""
}

func (m *PoolInfo) GetStandard() types.Coin {
	if m != nil {
		return m.Standard
	}
	return types.Coin{}
}

func (m *PoolInfo) GetToken() types.Coin {
	if m != nil {
		return m.Token
	}
	return types.Coin{}
}

func (m *PoolInfo) GetLpt() types.Coin {
	if m != nil {
		return m.Lpt
	}
	return types.Coin{}
}

func (m *PoolInfo) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryLiquidityPoolRequest)(nil), "irismod.coinswap.QueryLiquidityPoolRequest")
	proto.RegisterType((*QueryLiquidityPoolResponse)(nil), "irismod.coinswap.QueryLiquidityPoolResponse")
	proto.RegisterType((*QueryLiquidityPoolsRequest)(nil), "irismod.coinswap.QueryLiquidityPoolsRequest")
	proto.RegisterType((*QueryLiquidityPoolsResponse)(nil), "irismod.coinswap.QueryLiquidityPoolsResponse")
	proto.RegisterType((*PoolInfo)(nil), "irismod.coinswap.PoolInfo")
}

func init() { proto.RegisterFile("coinswap/query.proto", fileDescriptor_2cabf8423404f12f) }

var fileDescriptor_2cabf8423404f12f = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0xf9, 0x51, 0x92, 0x91, 0x86, 0x32, 0x14, 0xdc, 0x6c, 0x65, 0x5b, 0x22, 0x55,
	0xb1, 0x66, 0x87, 0xc4, 0x1f, 0x08, 0x9e, 0xac, 0xa2, 0x08, 0x3d, 0xd4, 0x3d, 0x7a, 0x29, 0x93,
	0xcc, 0x74, 0x3b, 0x74, 0x33, 0x6f, 0xb3, 0x33, 0x6b, 0x29, 0xe2, 0xc5, 0x3f, 0x40, 0x04, 0x0f,
	0xde, 0xfc, 0x23, 0xfc, 0x2b, 0x7a, 0x2c, 0x78, 0xf1, 0x24, 0x92, 0xf8, 0x57, 0x78, 0x92, 0x9d,
	0xdd, 0xc4, 0xae, 0xa6, 0xb8, 0xa7, 0x4c, 0x66, 0xbe, 0xdf, 0xf7, 0x3e, 0x6f, 0xde, 0x9b, 0x45,
	0xeb, 0x23, 0x10, 0x52, 0x9d, 0xd0, 0x88, 0x4c, 0x12, 0x1e, 0x9f, 0x7a, 0x51, 0x0c, 0x1a, 0xf0,
	0x9a, 0x88, 0x85, 0x1a, 0x03, 0xf3, 0xe6, 0xa7, 0x8e, 0x3b, 0x02, 0x35, 0x06, 0x45, 0x86, 0x54,
	0x71, 0xf2, 0xba, 0x3f, 0xe4, 0x9a, 0xf6, 0x49, 0x7a, 0x9a, 0x39, 0x9c, 0xf5, 0x00, 0x02, 0x30,
	0x4b, 0x92, 0xae, 0xf2, 0xdd, 0x6b, 0x01, 0x40, 0x10, 0x72, 0x42, 0x23, 0x41, 0xa8, 0x94, 0xa0,
	0xa9, 0x16, 0x20, 0x55, 0x7e, 0x7a, 0xfb, 0x62, 0x4c, 0x93, 0x7e, 0x11, 0x39, 0xa2, 0x81, 0x90,
	0x46, 0x9c, 0x69, 0xbb, 0x0f, 0x51, 0xe7, 0x65, 0xaa, 0xd8, 0x13, 0x93, 0x44, 0x30, 0xa1, 0x4f,
	0xf7, 0x01, 0x42, 0x9f, 0x4f, 0x12, 0xae, 0x34, 0xde, 0x40, 0xad, 0x30, 0xd2, 0x07, 0x8c, 0x4b,
	0x18, 0xdb, 0xd6, 0x96, 0x75, 0xab, 0xe5, 0x37, 0xc3, 0x48, 0x3f, 0x4d, 0xff, 0x77, 0x7d, 0xe4,
	0x2c, 0x73, 0xaa, 0x08, 0xa4, 0xe2, 0xf8, 0x1e, 0xaa, 0x47, 0x00, 0xa1, 0x71, 0x5d, 0x19, 0x38,
	0xde, 0xdf, 0x85, 0x7b, 0xa9, 0xfa, 0x85, 0x3c, 0x84, 0xdd, 0xfa, 0xd9, 0xf7, 0xcd, 0x8a, 0x6f,
	0xd4, 0x5d, 0xb6, 0x2c, 0xa6, 0x9a, 0xe3, 0x3c, 0x43, 0xe8, 0x0f, 0x7f, 0x1e, 0xf9, 0x86, 0x97,
	0x15, 0xeb, 0xa5, 0xc5, 0x7a, 0xd9, 0x5d, 0xe7, 0xc5, 0x7a, 0xfb, 0x34, 0xe0, 0xb9, 0xd7, 0xbf,
	0xe0, 0xec, 0x7e, 0xb6, 0xd0, 0xc6, 0xd2, 0x34, 0x39, 0xfb, 0x03, 0xd4, 0x48, 0x69, 0x94, 0x6d,
	0x6d, 0xd5, 0x4a, 0xc1, 0x67, 0x72, 0xfc, 0xbc, 0xc0, 0x57, 0x35, 0x7c, 0x37, 0xff, 0xcb, 0x97,
	0x25, 0x2d, 0x00, 0xfe, 0xb2, 0x50, 0x73, 0x9e, 0x02, 0xb7, 0x51, 0x55, 0xb0, 0xfc, 0xf6, 0xab,
	0x82, 0xe1, 0x6d, 0xd4, 0xe6, 0x6a, 0x14, 0xc3, 0xc9, 0x01, 0x65, 0x2c, 0xe6, 0x4a, 0x99, 0x4c,
	0x2d, 0x7f, 0x35, 0xdb, 0x7d, 0x9c, 0x6d, 0xe2, 0x47, 0xa8, 0xa9, 0x34, 0x95, 0x8c, 0xc6, 0xcc,
	0xae, 0x19, 0x94, 0x4e, 0x01, 0x65, 0x0e, 0xf1, 0x04, 0x84, 0xcc, 0xcb, 0x58, 0x18, 0xf0, 0x7d,
	0xd4, 0xd0, 0x70, 0xcc, 0xa5, 0x5d, 0x2f, 0xe7, 0xcc, 0xd4, 0xb8, 0x8f, 0x6a, 0x61, 0xa4, 0xed,
	0x46, 0x39, 0x53, 0xaa, 0xc5, 0x6b, 0xa8, 0x76, 0xc8, 0xb9, 0xbd, 0x62, 0x4a, 0x48, 0x97, 0x83,
	0x2f, 0x55, 0xd4, 0x30, 0xdd, 0xc1, 0x9f, 0x2c, 0xb4, 0x5a, 0x68, 0x11, 0xde, 0xf9, 0xb7, 0x15,
	0x97, 0x4e, 0xaf, 0x73, 0xa7, 0x9c, 0x38, 0xbb, 0xff, 0xee, 0xce, 0xbb, 0xaf, 0x3f, 0x3f, 0x56,
	0xb7, 0xf1, 0x75, 0x92, 0xbb, 0xc8, 0xe2, 0x05, 0x9b, 0xee, 0x92, 0x37, 0x8b, 0xa7, 0xf0, 0x16,
	0xbf, 0xb7, 0x50, 0xbb, 0x38, 0x3c, 0xb8, 0x54, 0xb6, 0xf9, 0x28, 0x3b, 0xbd, 0x92, 0xea, 0x1c,
	0x6e, 0xd3, 0xc0, 0x75, 0xf0, 0xd5, 0x4b, 0xe0, 0x76, 0xf7, 0xce, 0xa6, 0xae, 0x75, 0x3e, 0x75,
	0xad, 0x1f, 0x53, 0xd7, 0xfa, 0x30, 0x73, 0x2b, 0xe7, 0x33, 0xb7, 0xf2, 0x6d, 0xe6, 0x56, 0x5e,
	0x0d, 0x02, 0xa1, 0x8f, 0x92, 0xa1, 0x37, 0x82, 0xb1, 0x31, 0x4b, 0xae, 0xcd, 0xef, 0x51, 0x32,
	0xec, 0x29, 0x76, 0xdc, 0x0b, 0x80, 0x8c, 0x81, 0x25, 0x21, 0x57, 0x8b, 0x98, 0xc3, 0x15, 0xf3,
	0x6d, 0xb8, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x9d, 0xad, 0xf8, 0xc5, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// LiquidityPool returns the liquidity pool for the provided
	// lpt_denom
	LiquidityPool(ctx context.Context, in *QueryLiquidityPoolRequest, opts ...grpc.CallOption) (*QueryLiquidityPoolResponse, error)
	// LiquidityPools returns all the liquidity pools available
	LiquidityPools(ctx context.Context, in *QueryLiquidityPoolsRequest, opts ...grpc.CallOption) (*QueryLiquidityPoolsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) LiquidityPool(ctx context.Context, in *QueryLiquidityPoolRequest, opts ...grpc.CallOption) (*QueryLiquidityPoolResponse, error) {
	out := new(QueryLiquidityPoolResponse)
	err := c.cc.Invoke(ctx, "/irismod.coinswap.Query/LiquidityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidityPools(ctx context.Context, in *QueryLiquidityPoolsRequest, opts ...grpc.CallOption) (*QueryLiquidityPoolsResponse, error) {
	out := new(QueryLiquidityPoolsResponse)
	err := c.cc.Invoke(ctx, "/irismod.coinswap.Query/LiquidityPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// LiquidityPool returns the liquidity pool for the provided
	// lpt_denom
	LiquidityPool(context.Context, *QueryLiquidityPoolRequest) (*QueryLiquidityPoolResponse, error)
	// LiquidityPools returns all the liquidity pools available
	LiquidityPools(context.Context, *QueryLiquidityPoolsRequest) (*QueryLiquidityPoolsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) LiquidityPool(ctx context.Context, req *QueryLiquidityPoolRequest) (*QueryLiquidityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidityPool not implemented")
}
func (*UnimplementedQueryServer) LiquidityPools(ctx context.Context, req *QueryLiquidityPoolsRequest) (*QueryLiquidityPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidityPools not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_LiquidityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.coinswap.Query/LiquidityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidityPool(ctx, req.(*QueryLiquidityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidityPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidityPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidityPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.coinswap.Query/LiquidityPools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidityPools(ctx, req.(*QueryLiquidityPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.coinswap.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LiquidityPool",
			Handler:    _Query_LiquidityPool_Handler,
		},
		{
			MethodName: "LiquidityPools",
			Handler:    _Query_LiquidityPools_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coinswap/query.proto",
}

func (m *QueryLiquidityPoolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidityPoolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidityPoolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LptDenom) > 0 {
		i -= len(m.LptDenom)
		copy(dAtA[i:], m.LptDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.LptDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLiquidityPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidityPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidityPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryLiquidityPoolsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidityPoolsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidityPoolsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLiquidityPoolsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidityPoolsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidityPoolsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PoolInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.Lpt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Standard.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.EscrowAddress) > 0 {
		i -= len(m.EscrowAddress)
		copy(dAtA[i:], m.EscrowAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EscrowAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryLiquidityPoolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LptDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLiquidityPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pool.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryLiquidityPoolsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLiquidityPoolsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PoolInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.EscrowAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.Standard.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.Token.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.Lpt.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryLiquidityPoolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidityPoolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidityPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LptDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LptDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLiquidityPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidityPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidityPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLiquidityPoolsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidityPoolsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidityPoolsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLiquidityPoolsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidityPoolsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidityPoolsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, PoolInfo{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Standard", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Standard.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lpt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Lpt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
