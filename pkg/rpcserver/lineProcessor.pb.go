// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lineProcessor.proto

package rpcserver

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

type Request struct {
	Sports               []string `protobuf:"bytes,1,rep,name=sports,proto3" json:"sports,omitempty"`
	TimeUpd              string   `protobuf:"bytes,2,opt,name=timeUpd,proto3" json:"timeUpd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7fd3e0762619a8e, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSports() []string {
	if m != nil {
		return m.Sports
	}
	return nil
}

func (m *Request) GetTimeUpd() string {
	if m != nil {
		return m.TimeUpd
	}
	return ""
}

type Response struct {
	Line                 map[string]float32 `protobuf:"bytes,1,rep,name=line,proto3" json:"line,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7fd3e0762619a8e, []int{1}
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

func (m *Response) GetLine() map[string]float32 {
	if m != nil {
		return m.Line
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "rpcserver.Request")
	proto.RegisterType((*Response)(nil), "rpcserver.Response")
	proto.RegisterMapType((map[string]float32)(nil), "rpcserver.Response.LineEntry")
}

func init() {
	proto.RegisterFile("lineProcessor.proto", fileDescriptor_f7fd3e0762619a8e)
}

var fileDescriptor_f7fd3e0762619a8e = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xc9, 0x56, 0x5b, 0x77, 0x44, 0x90, 0xa9, 0x94, 0x50, 0x10, 0x4a, 0x4f, 0x7b, 0x0a,
	0x5a, 0x0f, 0x8a, 0x9e, 0x7b, 0x2b, 0x28, 0x29, 0x7a, 0x77, 0xd7, 0x39, 0x04, 0x6b, 0x12, 0x67,
	0xb2, 0x0b, 0xfd, 0xf7, 0xb2, 0xd1, 0x2e, 0x0a, 0xde, 0xe6, 0x0d, 0x6f, 0xf8, 0xde, 0x1b, 0x98,
	0xee, 0x9c, 0xa7, 0x27, 0x0e, 0x0d, 0x89, 0x04, 0x36, 0x91, 0x43, 0x0a, 0x58, 0x72, 0x6c, 0x84,
	0xb8, 0x23, 0x5e, 0x3e, 0xc0, 0xc4, 0xd2, 0x67, 0x4b, 0x92, 0x70, 0x06, 0x63, 0x89, 0x81, 0x93,
	0x68, 0xb5, 0x18, 0x55, 0xa5, 0xfd, 0x51, 0xa8, 0x61, 0x92, 0xdc, 0x07, 0x3d, 0xc7, 0x37, 0x5d,
	0x2c, 0x54, 0x55, 0xda, 0x83, 0x5c, 0x76, 0x70, 0x62, 0x49, 0x62, 0xf0, 0x42, 0x78, 0x0d, 0x47,
	0x3d, 0x2a, 0xdf, 0x9e, 0xae, 0x2e, 0xcd, 0x80, 0x30, 0x07, 0x8b, 0xd9, 0x38, 0x4f, 0x6b, 0x9f,
	0x78, 0x6f, 0xb3, 0x75, 0x7e, 0x0b, 0xe5, 0xb0, 0xc2, 0x73, 0x18, 0xbd, 0xd3, 0x5e, 0xab, 0x4c,
	0xe8, 0x47, 0xbc, 0x80, 0xe3, 0xee, 0x75, 0xd7, 0x52, 0xa6, 0x16, 0xf6, 0x5b, 0xdc, 0x17, 0x77,
	0x6a, 0xf5, 0x02, 0x67, 0x9b, 0xdf, 0xb5, 0x70, 0x0d, 0xb3, 0x6d, 0x5b, 0x4b, 0xc3, 0xae, 0xa6,
	0x47, 0xbf, 0xcd, 0xb9, 0x7b, 0x87, 0x20, 0xfe, 0x09, 0x92, 0x8b, 0xce, 0xa7, 0xff, 0x84, 0xab,
	0xd4, 0x95, 0xaa, 0xc7, 0xf9, 0x3d, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xe2, 0xb7,
	0xce, 0x35, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LineProcessorClient is the client API for LineProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LineProcessorClient interface {
	SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (LineProcessor_SubscribeOnSportsLinesClient, error)
}

type lineProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewLineProcessorClient(cc grpc.ClientConnInterface) LineProcessorClient {
	return &lineProcessorClient{cc}
}

func (c *lineProcessorClient) SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (LineProcessor_SubscribeOnSportsLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LineProcessor_serviceDesc.Streams[0], "/rpcserver.LineProcessor/SubscribeOnSportsLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &lineProcessorSubscribeOnSportsLinesClient{stream}
	return x, nil
}

type LineProcessor_SubscribeOnSportsLinesClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type lineProcessorSubscribeOnSportsLinesClient struct {
	grpc.ClientStream
}

func (x *lineProcessorSubscribeOnSportsLinesClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lineProcessorSubscribeOnSportsLinesClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LineProcessorServer is the server API for LineProcessor service.
type LineProcessorServer interface {
	SubscribeOnSportsLines(LineProcessor_SubscribeOnSportsLinesServer) error
}

// UnimplementedLineProcessorServer can be embedded to have forward compatible implementations.
type UnimplementedLineProcessorServer struct {
}

func (*UnimplementedLineProcessorServer) SubscribeOnSportsLines(srv LineProcessor_SubscribeOnSportsLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOnSportsLines not implemented")
}

func RegisterLineProcessorServer(s *grpc.Server, srv LineProcessorServer) {
	s.RegisterService(&_LineProcessor_serviceDesc, srv)
}

func _LineProcessor_SubscribeOnSportsLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LineProcessorServer).SubscribeOnSportsLines(&lineProcessorSubscribeOnSportsLinesServer{stream})
}

type LineProcessor_SubscribeOnSportsLinesServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type lineProcessorSubscribeOnSportsLinesServer struct {
	grpc.ServerStream
}

func (x *lineProcessorSubscribeOnSportsLinesServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lineProcessorSubscribeOnSportsLinesServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LineProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcserver.LineProcessor",
	HandlerType: (*LineProcessorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOnSportsLines",
			Handler:       _LineProcessor_SubscribeOnSportsLines_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "lineProcessor.proto",
}
