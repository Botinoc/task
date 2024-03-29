// Code generated by protoc-gen-go. DO NOT EDIT.
// source: awesomeProject1/proto/laziness.proto

package laziness

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

type Laziness struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Power                int32    `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Laziness) Reset()         { *m = Laziness{} }
func (m *Laziness) String() string { return proto.CompactTextString(m) }
func (*Laziness) ProtoMessage()    {}
func (*Laziness) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed9585bf2f4eb99, []int{0}
}

func (m *Laziness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laziness.Unmarshal(m, b)
}
func (m *Laziness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laziness.Marshal(b, m, deterministic)
}
func (m *Laziness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laziness.Merge(m, src)
}
func (m *Laziness) XXX_Size() int {
	return xxx_messageInfo_Laziness.Size(m)
}
func (m *Laziness) XXX_DiscardUnknown() {
	xxx_messageInfo_Laziness.DiscardUnknown(m)
}

var xxx_messageInfo_Laziness proto.InternalMessageInfo

func (m *Laziness) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laziness) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Laziness) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

type Index struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed9585bf2f4eb99, []int{1}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
	Created              bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Laziness             *Laziness `protobuf:"bytes,2,opt,name=laziness,proto3" json:"laziness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed9585bf2f4eb99, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetLaziness() *Laziness {
	if m != nil {
		return m.Laziness
	}
	return nil
}

func init() {
	proto.RegisterType((*Laziness)(nil), "Laziness")
	proto.RegisterType((*Index)(nil), "Index")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("awesomeProject1/proto/laziness.proto", fileDescriptor_5ed9585bf2f4eb99)
}

var fileDescriptor_5ed9585bf2f4eb99 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x75, 0x57, 0xb6, 0xdd, 0x9d, 0x42, 0x85, 0x41, 0x30, 0x78, 0x90, 0x25, 0x28, 0xec, 0x29,
	0xc5, 0xfa, 0x07, 0x7a, 0x90, 0xa2, 0x07, 0x89, 0x47, 0x4f, 0x6b, 0x32, 0x87, 0x88, 0x26, 0x21,
	0x09, 0x56, 0xfd, 0x7a, 0x21, 0x35, 0xab, 0xf6, 0xf8, 0xde, 0x9b, 0x79, 0xef, 0xcd, 0xc0, 0xf9,
	0xb8, 0xa5, 0xe8, 0xde, 0xe8, 0x21, 0xb8, 0x17, 0x52, 0xe9, 0x72, 0xe5, 0x83, 0x4b, 0x6e, 0xf5,
	0x3a, 0x7e, 0x19, 0x4b, 0x31, 0x8a, 0x0c, 0xb9, 0x84, 0xf6, 0xfe, 0x87, 0xc1, 0x25, 0xd4, 0x46,
	0xb3, 0xaa, 0xaf, 0x86, 0x4e, 0xd6, 0x46, 0x63, 0x0f, 0x0b, 0x4d, 0x51, 0x05, 0xe3, 0x93, 0x71,
	0x96, 0xd5, 0x59, 0xf8, 0x4b, 0xe1, 0x31, 0x34, 0xde, 0x6d, 0x29, 0xb0, 0xc3, 0xbe, 0x1a, 0x1a,
	0xb9, 0x03, 0xfc, 0x04, 0x9a, 0x8d, 0xd5, 0xf4, 0xb1, 0x6f, 0xc8, 0xef, 0xa0, 0x95, 0x14, 0xbd,
	0xb3, 0x91, 0x90, 0xc1, 0x5c, 0x05, 0x1a, 0x13, 0xed, 0x06, 0x5a, 0x59, 0x20, 0x5e, 0x40, 0x5b,
	0x4a, 0xe6, 0xcc, 0xc5, 0xba, 0x13, 0xa5, 0xa3, 0x9c, 0xa4, 0xf5, 0x13, 0x1c, 0x15, 0xf6, 0x91,
	0xc2, 0xbb, 0x51, 0x84, 0x03, 0x2c, 0x6f, 0xb2, 0xc9, 0x74, 0xd2, 0xef, 0xe6, 0x69, 0x27, 0x4a,
	0x36, 0x3f, 0xc0, 0x33, 0x98, 0xdf, 0x52, 0xba, 0xfe, 0xdc, 0x68, 0x9c, 0x89, 0x5c, 0xf6, 0x9f,
	0xfe, 0x3c, 0xcb, 0xdf, 0xb9, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xce, 0xc9, 0xce, 0x28, 0x45,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LazinessServiceClient is the client API for LazinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LazinessServiceClient interface {
	CreateLaziness(ctx context.Context, in *Laziness, opts ...grpc.CallOption) (*Response, error)
	GetById(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Response, error)
}

type lazinessServiceClient struct {
	cc *grpc.ClientConn
}

func NewLazinessServiceClient(cc *grpc.ClientConn) LazinessServiceClient {
	return &lazinessServiceClient{cc}
}

func (c *lazinessServiceClient) CreateLaziness(ctx context.Context, in *Laziness, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/LazinessService/CreateLaziness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lazinessServiceClient) GetById(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/LazinessService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LazinessServiceServer is the server API for LazinessService service.
type LazinessServiceServer interface {
	CreateLaziness(context.Context, *Laziness) (*Response, error)
	GetById(context.Context, *Index) (*Response, error)
}

// UnimplementedLazinessServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLazinessServiceServer struct {
}

func (*UnimplementedLazinessServiceServer) CreateLaziness(ctx context.Context, req *Laziness) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaziness not implemented")
}
func (*UnimplementedLazinessServiceServer) GetById(ctx context.Context, req *Index) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}

func RegisterLazinessServiceServer(s *grpc.Server, srv LazinessServiceServer) {
	s.RegisterService(&_LazinessService_serviceDesc, srv)
}

func _LazinessService_CreateLaziness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Laziness)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LazinessServiceServer).CreateLaziness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LazinessService/CreateLaziness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LazinessServiceServer).CreateLaziness(ctx, req.(*Laziness))
	}
	return interceptor(ctx, in, info, handler)
}

func _LazinessService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Index)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LazinessServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LazinessService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LazinessServiceServer).GetById(ctx, req.(*Index))
	}
	return interceptor(ctx, in, info, handler)
}

var _LazinessService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LazinessService",
	HandlerType: (*LazinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaziness",
			Handler:    _LazinessService_CreateLaziness_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _LazinessService_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "awesomeProject1/proto/laziness.proto",
}
