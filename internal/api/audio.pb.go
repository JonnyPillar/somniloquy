// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audio.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fae679ade197b92f, []int{0}
}

type UploadRecordRequest struct {
	Content              []int32  `protobuf:"varint,1,rep,packed,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRecordRequest) Reset()         { *m = UploadRecordRequest{} }
func (m *UploadRecordRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRecordRequest) ProtoMessage()    {}
func (*UploadRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fae679ade197b92f, []int{0}
}

func (m *UploadRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRecordRequest.Unmarshal(m, b)
}
func (m *UploadRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRecordRequest.Marshal(b, m, deterministic)
}
func (m *UploadRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRecordRequest.Merge(m, src)
}
func (m *UploadRecordRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRecordRequest.Size(m)
}
func (m *UploadRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRecordRequest proto.InternalMessageInfo

func (m *UploadRecordRequest) GetContent() []int32 {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=api.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fae679ade197b92f, []int{1}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("api.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*UploadRecordRequest)(nil), "api.UploadRecordRequest")
	proto.RegisterType((*UploadStatus)(nil), "api.UploadStatus")
}

func init() { proto.RegisterFile("audio.proto", fileDescriptor_fae679ade197b92f) }

var fileDescriptor_fae679ade197b92f = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x9b, 0xad, 0x6e, 0x71, 0xaa, 0x12, 0x47, 0x84, 0xe0, 0x69, 0xd9, 0x53, 0xf4, 0xb0,
	0x42, 0x7b, 0xf2, 0x5c, 0xf0, 0x20, 0x88, 0x90, 0xa5, 0x3f, 0x20, 0x36, 0x83, 0x84, 0x96, 0x4c,
	0xdc, 0x64, 0xf5, 0xef, 0x4b, 0xbb, 0x14, 0x56, 0x7b, 0x7c, 0xbc, 0xf7, 0x66, 0xbe, 0x19, 0x98,
	0xdb, 0xde, 0x79, 0x6e, 0x62, 0xc7, 0x99, 0x71, 0x6a, 0xa3, 0xaf, 0x9f, 0xe0, 0x76, 0x1d, 0x77,
	0x6c, 0x9d, 0xa1, 0x0d, 0x77, 0xce, 0xd0, 0x57, 0x4f, 0x29, 0xa3, 0x82, 0xd9, 0x8a, 0x43, 0xa6,
	0x90, 0x95, 0xa8, 0xa6, 0xfa, 0xdc, 0x1c, 0x65, 0xdd, 0xc2, 0xe5, 0x50, 0x68, 0xb3, 0xcd, 0x7d,
	0xda, 0x27, 0xdf, 0x28, 0x25, 0xfb, 0x49, 0x4a, 0x54, 0x42, 0x5f, 0x98, 0xa3, 0xc4, 0x07, 0x38,
	0x5b, 0xb1, 0x23, 0x55, 0x54, 0x42, 0x5f, 0x2f, 0xee, 0x1a, 0x1b, 0x7d, 0x33, 0xae, 0xee, 0x4d,
	0x73, 0x88, 0x3c, 0x2e, 0x41, 0xfe, 0x77, 0x70, 0x0e, 0xb3, 0x75, 0xd8, 0x06, 0xfe, 0x09, 0x72,
	0x82, 0x25, 0x14, 0xef, 0x5b, 0x29, 0x10, 0xa0, 0x7c, 0xb1, 0x7e, 0x47, 0x4e, 0x16, 0x8b, 0x57,
	0xb8, 0x1a, 0xa0, 0x5b, 0xea, 0xbe, 0xfd, 0x86, 0xf0, 0x19, 0xca, 0x61, 0x0a, 0xaa, 0xd1, 0xb2,
	0x3f, 0x87, 0xdd, 0xdf, 0x9c, 0x60, 0xd4, 0x13, 0x2d, 0x3e, 0xca, 0xc3, 0x4b, 0x96, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0x2a, 0xa6, 0x20, 0x21, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (RecordService_UploadClient, error)
}

type recordServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecordServiceClient(cc *grpc.ClientConn) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (RecordService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecordService_serviceDesc.Streams[0], "/api.RecordService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordServiceUploadClient{stream}
	return x, nil
}

type RecordService_UploadClient interface {
	Send(*UploadRecordRequest) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type recordServiceUploadClient struct {
	grpc.ClientStream
}

func (x *recordServiceUploadClient) Send(m *UploadRecordRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recordServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecordServiceServer is the server API for RecordService service.
type RecordServiceServer interface {
	Upload(RecordService_UploadServer) error
}

func RegisterRecordServiceServer(s *grpc.Server, srv RecordServiceServer) {
	s.RegisterService(&_RecordService_serviceDesc, srv)
}

func _RecordService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecordServiceServer).Upload(&recordServiceUploadServer{stream})
}

type RecordService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*UploadRecordRequest, error)
	grpc.ServerStream
}

type recordServiceUploadServer struct {
	grpc.ServerStream
}

func (x *recordServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recordServiceUploadServer) Recv() (*UploadRecordRequest, error) {
	m := new(UploadRecordRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RecordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _RecordService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "audio.proto",
}
