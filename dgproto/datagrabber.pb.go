// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datagrabber.proto

package dgproto

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

type HeaderReq struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=Identifier,proto3" json:"Identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderReq) Reset()         { *m = HeaderReq{} }
func (m *HeaderReq) String() string { return proto.CompactTextString(m) }
func (*HeaderReq) ProtoMessage()    {}
func (*HeaderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c794cd1986f2f27b, []int{0}
}

func (m *HeaderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderReq.Unmarshal(m, b)
}
func (m *HeaderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderReq.Marshal(b, m, deterministic)
}
func (m *HeaderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderReq.Merge(m, src)
}
func (m *HeaderReq) XXX_Size() int {
	return xxx_messageInfo_HeaderReq.Size(m)
}
func (m *HeaderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderReq proto.InternalMessageInfo

func (m *HeaderReq) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type HeaderResp struct {
	ColHeader            string   `protobuf:"bytes,1,opt,name=ColHeader,proto3" json:"ColHeader,omitempty"`
	RowHeader            string   `protobuf:"bytes,2,opt,name=RowHeader,proto3" json:"RowHeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderResp) Reset()         { *m = HeaderResp{} }
func (m *HeaderResp) String() string { return proto.CompactTextString(m) }
func (*HeaderResp) ProtoMessage()    {}
func (*HeaderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c794cd1986f2f27b, []int{1}
}

func (m *HeaderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderResp.Unmarshal(m, b)
}
func (m *HeaderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderResp.Marshal(b, m, deterministic)
}
func (m *HeaderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderResp.Merge(m, src)
}
func (m *HeaderResp) XXX_Size() int {
	return xxx_messageInfo_HeaderResp.Size(m)
}
func (m *HeaderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderResp.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderResp proto.InternalMessageInfo

func (m *HeaderResp) GetColHeader() string {
	if m != nil {
		return m.ColHeader
	}
	return ""
}

func (m *HeaderResp) GetRowHeader() string {
	if m != nil {
		return m.RowHeader
	}
	return ""
}

type DataReq struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=Identifier,proto3" json:"Identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataReq) Reset()         { *m = DataReq{} }
func (m *DataReq) String() string { return proto.CompactTextString(m) }
func (*DataReq) ProtoMessage()    {}
func (*DataReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c794cd1986f2f27b, []int{2}
}

func (m *DataReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataReq.Unmarshal(m, b)
}
func (m *DataReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataReq.Marshal(b, m, deterministic)
}
func (m *DataReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataReq.Merge(m, src)
}
func (m *DataReq) XXX_Size() int {
	return xxx_messageInfo_DataReq.Size(m)
}
func (m *DataReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DataReq.DiscardUnknown(m)
}

var xxx_messageInfo_DataReq proto.InternalMessageInfo

func (m *DataReq) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type PointResp struct {
	X                    string   `protobuf:"bytes,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	XLabel               string   `protobuf:"bytes,4,opt,name=XLabel,proto3" json:"XLabel,omitempty"`
	YLabel               string   `protobuf:"bytes,5,opt,name=YLabel,proto3" json:"YLabel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointResp) Reset()         { *m = PointResp{} }
func (m *PointResp) String() string { return proto.CompactTextString(m) }
func (*PointResp) ProtoMessage()    {}
func (*PointResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c794cd1986f2f27b, []int{3}
}

func (m *PointResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointResp.Unmarshal(m, b)
}
func (m *PointResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointResp.Marshal(b, m, deterministic)
}
func (m *PointResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointResp.Merge(m, src)
}
func (m *PointResp) XXX_Size() int {
	return xxx_messageInfo_PointResp.Size(m)
}
func (m *PointResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PointResp.DiscardUnknown(m)
}

var xxx_messageInfo_PointResp proto.InternalMessageInfo

func (m *PointResp) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *PointResp) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *PointResp) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PointResp) GetXLabel() string {
	if m != nil {
		return m.XLabel
	}
	return ""
}

func (m *PointResp) GetYLabel() string {
	if m != nil {
		return m.YLabel
	}
	return ""
}

func init() {
	proto.RegisterType((*HeaderReq)(nil), "dgproto.HeaderReq")
	proto.RegisterType((*HeaderResp)(nil), "dgproto.HeaderResp")
	proto.RegisterType((*DataReq)(nil), "dgproto.DataReq")
	proto.RegisterType((*PointResp)(nil), "dgproto.PointResp")
}

func init() {
	proto.RegisterFile("datagrabber.proto", fileDescriptor_c794cd1986f2f27b)
}

var fileDescriptor_c794cd1986f2f27b = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x59, 0x35, 0x2d, 0x19, 0x15, 0x74, 0x14, 0x09, 0x45, 0x44, 0x72, 0x52, 0x84, 0x50,
	0x14, 0x2f, 0x5e, 0xf5, 0xd0, 0x82, 0x07, 0x89, 0x1e, 0x92, 0xe3, 0xc6, 0x8c, 0x65, 0x21, 0x66,
	0xd3, 0xed, 0xaa, 0x78, 0xf3, 0xa3, 0x4b, 0x76, 0xb6, 0x1b, 0xc1, 0x4b, 0x6f, 0xfb, 0x7e, 0xef,
	0xed, 0xfc, 0x83, 0xc3, 0x5a, 0x5a, 0xb9, 0x30, 0xb2, 0xaa, 0xc8, 0x64, 0x9d, 0xd1, 0x56, 0xe3,
	0xb8, 0x5e, 0xb8, 0x47, 0x7a, 0x05, 0xf1, 0x8c, 0x64, 0x4d, 0x26, 0xa7, 0x25, 0x9e, 0x01, 0xcc,
	0x6b, 0x6a, 0xad, 0x7a, 0x53, 0x64, 0x12, 0x71, 0x2e, 0x2e, 0xe2, 0xfc, 0x0f, 0x49, 0x67, 0x00,
	0xeb, 0xf0, 0xaa, 0xc3, 0x53, 0x88, 0xef, 0x75, 0xc3, 0xc0, 0x87, 0x07, 0xd0, 0xbb, 0xb9, 0xfe,
	0xf2, 0xee, 0x16, 0xbb, 0x01, 0xa4, 0x97, 0x30, 0x7e, 0x90, 0x56, 0x6e, 0xd2, 0xf4, 0x1d, 0xe2,
	0x27, 0xad, 0x5a, 0xeb, 0x7a, 0xee, 0x81, 0x28, 0x7c, 0x46, 0x14, 0xbd, 0x2a, 0x5d, 0xed, 0x28,
	0x17, 0x25, 0x1e, 0x43, 0xf4, 0xa2, 0x6c, 0x43, 0xc9, 0xb6, 0xf3, 0x59, 0xe0, 0x09, 0x8c, 0x8a,
	0x47, 0x59, 0x51, 0x93, 0xec, 0x38, 0xec, 0x55, 0xcf, 0x4b, 0xe6, 0x11, 0x73, 0x56, 0xd7, 0x3f,
	0x02, 0x76, 0xfb, 0xd1, 0x9e, 0xc9, 0x7c, 0xaa, 0x57, 0xc2, 0x3b, 0xd8, 0xe7, 0x99, 0xe7, 0xed,
	0xf2, 0x43, 0x99, 0x6f, 0xc4, 0xcc, 0xdf, 0x2e, 0x0b, 0x87, 0x9b, 0x1c, 0xfd, 0x63, 0xab, 0x6e,
	0x2a, 0xf0, 0x96, 0x4b, 0xad, 0x7f, 0x1e, 0x84, 0x94, 0xdf, 0x7d, 0x32, 0xd4, 0x0a, 0x2b, 0x4e,
	0x45, 0x35, 0x72, 0xe8, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x79, 0x02, 0x7f, 0xb8, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataServiceClient interface {
	HeaderInquiry(ctx context.Context, in *HeaderReq, opts ...grpc.CallOption) (DataService_HeaderInquiryClient, error)
	DataInquiry(ctx context.Context, in *DataReq, opts ...grpc.CallOption) (DataService_DataInquiryClient, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) HeaderInquiry(ctx context.Context, in *HeaderReq, opts ...grpc.CallOption) (DataService_HeaderInquiryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataService_serviceDesc.Streams[0], "/dgproto.DataService/HeaderInquiry", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceHeaderInquiryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_HeaderInquiryClient interface {
	Recv() (*HeaderResp, error)
	grpc.ClientStream
}

type dataServiceHeaderInquiryClient struct {
	grpc.ClientStream
}

func (x *dataServiceHeaderInquiryClient) Recv() (*HeaderResp, error) {
	m := new(HeaderResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) DataInquiry(ctx context.Context, in *DataReq, opts ...grpc.CallOption) (DataService_DataInquiryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataService_serviceDesc.Streams[1], "/dgproto.DataService/DataInquiry", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceDataInquiryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_DataInquiryClient interface {
	Recv() (*PointResp, error)
	grpc.ClientStream
}

type dataServiceDataInquiryClient struct {
	grpc.ClientStream
}

func (x *dataServiceDataInquiryClient) Recv() (*PointResp, error) {
	m := new(PointResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServiceServer is the server API for DataService service.
type DataServiceServer interface {
	HeaderInquiry(*HeaderReq, DataService_HeaderInquiryServer) error
	DataInquiry(*DataReq, DataService_DataInquiryServer) error
}

// UnimplementedDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (*UnimplementedDataServiceServer) HeaderInquiry(req *HeaderReq, srv DataService_HeaderInquiryServer) error {
	return status.Errorf(codes.Unimplemented, "method HeaderInquiry not implemented")
}
func (*UnimplementedDataServiceServer) DataInquiry(req *DataReq, srv DataService_DataInquiryServer) error {
	return status.Errorf(codes.Unimplemented, "method DataInquiry not implemented")
}

func RegisterDataServiceServer(s *grpc.Server, srv DataServiceServer) {
	s.RegisterService(&_DataService_serviceDesc, srv)
}

func _DataService_HeaderInquiry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HeaderReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).HeaderInquiry(m, &dataServiceHeaderInquiryServer{stream})
}

type DataService_HeaderInquiryServer interface {
	Send(*HeaderResp) error
	grpc.ServerStream
}

type dataServiceHeaderInquiryServer struct {
	grpc.ServerStream
}

func (x *dataServiceHeaderInquiryServer) Send(m *HeaderResp) error {
	return x.ServerStream.SendMsg(m)
}

func _DataService_DataInquiry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).DataInquiry(m, &dataServiceDataInquiryServer{stream})
}

type DataService_DataInquiryServer interface {
	Send(*PointResp) error
	grpc.ServerStream
}

type dataServiceDataInquiryServer struct {
	grpc.ServerStream
}

func (x *dataServiceDataInquiryServer) Send(m *PointResp) error {
	return x.ServerStream.SendMsg(m)
}

var _DataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dgproto.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HeaderInquiry",
			Handler:       _DataService_HeaderInquiry_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DataInquiry",
			Handler:       _DataService_DataInquiry_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "datagrabber.proto",
}
