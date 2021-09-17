// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync.proto

package sync

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SyncRoutesRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRoutesRequest) Reset()         { *m = SyncRoutesRequest{} }
func (m *SyncRoutesRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRoutesRequest) ProtoMessage()    {}
func (*SyncRoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5273b98214de8075, []int{0}
}

func (m *SyncRoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRoutesRequest.Unmarshal(m, b)
}
func (m *SyncRoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRoutesRequest.Marshal(b, m, deterministic)
}
func (m *SyncRoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRoutesRequest.Merge(m, src)
}
func (m *SyncRoutesRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRoutesRequest.Size(m)
}
func (m *SyncRoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRoutesRequest proto.InternalMessageInfo

func (m *SyncRoutesRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SyncRoutesResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRoutesResponse) Reset()         { *m = SyncRoutesResponse{} }
func (m *SyncRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*SyncRoutesResponse) ProtoMessage()    {}
func (*SyncRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5273b98214de8075, []int{1}
}

func (m *SyncRoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRoutesResponse.Unmarshal(m, b)
}
func (m *SyncRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRoutesResponse.Marshal(b, m, deterministic)
}
func (m *SyncRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRoutesResponse.Merge(m, src)
}
func (m *SyncRoutesResponse) XXX_Size() int {
	return xxx_messageInfo_SyncRoutesResponse.Size(m)
}
func (m *SyncRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRoutesResponse proto.InternalMessageInfo

func (m *SyncRoutesResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*SyncRoutesRequest)(nil), "sync.SyncRoutesRequest")
	proto.RegisterType((*SyncRoutesResponse)(nil), "sync.SyncRoutesResponse")
}

func init() { proto.RegisterFile("sync.proto", fileDescriptor_5273b98214de8075) }

var fileDescriptor_5273b98214de8075 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xae, 0xcc, 0x4b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xd4, 0xb9, 0x04, 0x83, 0x2b,
	0xf3, 0x92, 0x83, 0xf2, 0x4b, 0x4b, 0x52, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c,
	0x25, 0x1d, 0x2e, 0x21, 0x64, 0x85, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c,
	0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x60, 0xb5, 0x1c, 0x41, 0x50, 0x9e, 0x91, 0x27, 0x17, 0x0b,
	0x48, 0xb5, 0x90, 0x23, 0x17, 0x17, 0x42, 0x97, 0x90, 0xb8, 0x1e, 0xd8, 0x7e, 0x0c, 0x0b, 0xa5,
	0x24, 0x30, 0x25, 0x20, 0x16, 0x28, 0x31, 0x38, 0xb1, 0x47, 0xb1, 0xea, 0x83, 0x64, 0x93, 0xd8,
	0xc0, 0xee, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x40, 0x06, 0xf3, 0xc5, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncClient is the client API for Sync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncClient interface {
	SyncRoutes(ctx context.Context, in *SyncRoutesRequest, opts ...grpc.CallOption) (*SyncRoutesResponse, error)
}

type syncClient struct {
	cc *grpc.ClientConn
}

func NewSyncClient(cc *grpc.ClientConn) SyncClient {
	return &syncClient{cc}
}

func (c *syncClient) SyncRoutes(ctx context.Context, in *SyncRoutesRequest, opts ...grpc.CallOption) (*SyncRoutesResponse, error) {
	out := new(SyncRoutesResponse)
	err := c.cc.Invoke(ctx, "/sync.Sync/SyncRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServer is the server API for Sync service.
type SyncServer interface {
	SyncRoutes(context.Context, *SyncRoutesRequest) (*SyncRoutesResponse, error)
}

// UnimplementedSyncServer can be embedded to have forward compatible implementations.
type UnimplementedSyncServer struct {
}

func (*UnimplementedSyncServer) SyncRoutes(ctx context.Context, req *SyncRoutesRequest) (*SyncRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRoutes not implemented")
}

func RegisterSyncServer(s *grpc.Server, srv SyncServer) {
	s.RegisterService(&_Sync_serviceDesc, srv)
}

func _Sync_SyncRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServer).SyncRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.Sync/SyncRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServer).SyncRoutes(ctx, req.(*SyncRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sync.Sync",
	HandlerType: (*SyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncRoutes",
			Handler:    _Sync_SyncRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync.proto",
}
