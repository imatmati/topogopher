// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent/messages/order.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Topography struct {
	Distances            map[string]float64 `protobuf:"bytes,1,rep,name=distances,proto3" json:"distances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Topography) Reset()         { *m = Topography{} }
func (m *Topography) String() string { return proto.CompactTextString(m) }
func (*Topography) ProtoMessage()    {}
func (*Topography) Descriptor() ([]byte, []int) {
	return fileDescriptor_85d88341b42affcc, []int{0}
}

func (m *Topography) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topography.Unmarshal(m, b)
}
func (m *Topography) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topography.Marshal(b, m, deterministic)
}
func (m *Topography) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topography.Merge(m, src)
}
func (m *Topography) XXX_Size() int {
	return xxx_messageInfo_Topography.Size(m)
}
func (m *Topography) XXX_DiscardUnknown() {
	xxx_messageInfo_Topography.DiscardUnknown(m)
}

var xxx_messageInfo_Topography proto.InternalMessageInfo

func (m *Topography) GetDistances() map[string]float64 {
	if m != nil {
		return m.Distances
	}
	return nil
}

type Locations struct {
	Location             []string `protobuf:"bytes,1,rep,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Locations) Reset()         { *m = Locations{} }
func (m *Locations) String() string { return proto.CompactTextString(m) }
func (*Locations) ProtoMessage()    {}
func (*Locations) Descriptor() ([]byte, []int) {
	return fileDescriptor_85d88341b42affcc, []int{1}
}

func (m *Locations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Locations.Unmarshal(m, b)
}
func (m *Locations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Locations.Marshal(b, m, deterministic)
}
func (m *Locations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Locations.Merge(m, src)
}
func (m *Locations) XXX_Size() int {
	return xxx_messageInfo_Locations.Size(m)
}
func (m *Locations) XXX_DiscardUnknown() {
	xxx_messageInfo_Locations.DiscardUnknown(m)
}

var xxx_messageInfo_Locations proto.InternalMessageInfo

func (m *Locations) GetLocation() []string {
	if m != nil {
		return m.Location
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_85d88341b42affcc, []int{2}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85d88341b42affcc, []int{3}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Topography)(nil), "messages.Topography")
	proto.RegisterMapType((map[string]float64)(nil), "messages.Topography.DistancesEntry")
	proto.RegisterType((*Locations)(nil), "messages.Locations")
	proto.RegisterType((*Void)(nil), "messages.Void")
	proto.RegisterType((*PingResponse)(nil), "messages.PingResponse")
}

func init() { proto.RegisterFile("agent/messages/order.proto", fileDescriptor_85d88341b42affcc) }

var fileDescriptor_85d88341b42affcc = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x97, 0x6d, 0x8e, 0xf5, 0x4d, 0x8a, 0xc4, 0x21, 0xa3, 0xa7, 0x12, 0x0f, 0xf6, 0xd4,
	0xc9, 0xbc, 0x88, 0xe8, 0x41, 0xd0, 0x9b, 0x07, 0x89, 0xe2, 0x3d, 0xb6, 0x8f, 0x1a, 0x9c, 0x79,
	0x21, 0x89, 0x83, 0x5e, 0xfc, 0x0f, 0xfc, 0x9f, 0xa5, 0xd3, 0x2e, 0x0a, 0xde, 0xf2, 0xcb, 0xfb,
	0x92, 0xef, 0xfb, 0x1e, 0x64, 0xaa, 0x41, 0x13, 0x96, 0x6f, 0xe8, 0xbd, 0x6a, 0xd0, 0x2f, 0xc9,
	0xd5, 0xe8, 0x4a, 0xeb, 0x28, 0x10, 0x9f, 0xf6, 0xb7, 0xe2, 0x93, 0x01, 0x3c, 0x92, 0xa5, 0xc6,
	0x29, 0xfb, 0xd2, 0xf2, 0x6b, 0x48, 0x6a, 0xed, 0x83, 0x32, 0x15, 0xfa, 0x05, 0xcb, 0x47, 0xc5,
	0x6c, 0x75, 0x5c, 0xf6, 0xe2, 0x32, 0x0a, 0xcb, 0x9b, 0x5e, 0x75, 0x6b, 0x82, 0x6b, 0x65, 0x7c,
	0x95, 0x5d, 0x42, 0xfa, 0x77, 0xc8, 0x0f, 0x60, 0xf4, 0x8a, 0xed, 0x82, 0xe5, 0xac, 0x48, 0x64,
	0x77, 0xe4, 0x73, 0xd8, 0xdb, 0xa8, 0xf5, 0x3b, 0x2e, 0x86, 0x39, 0x2b, 0x98, 0xfc, 0x86, 0x8b,
	0xe1, 0x39, 0x13, 0x27, 0x90, 0xdc, 0x51, 0xa5, 0x82, 0x26, 0xe3, 0x79, 0x06, 0xd3, 0xf5, 0x0f,
	0x6c, 0xc3, 0x24, 0x72, 0xc7, 0x62, 0x02, 0xe3, 0x27, 0xd2, 0xb5, 0x48, 0x61, 0xff, 0x5e, 0x9b,
	0x46, 0xa2, 0xb7, 0x64, 0x3c, 0xae, 0x3e, 0x60, 0xd6, 0xc5, 0x7c, 0x40, 0xb7, 0xd1, 0x15, 0xf2,
	0x2b, 0x48, 0x25, 0x5a, 0x72, 0xa1, 0xcf, 0xc4, 0x0f, 0x63, 0x9f, 0x9d, 0x53, 0x36, 0xff, 0xaf,
	0xa4, 0x18, 0xf0, 0x53, 0x18, 0x77, 0xbf, 0xf3, 0x34, 0xce, 0x3b, 0xd7, 0xec, 0x28, 0xf2, 0x6f,
	0x77, 0x31, 0x78, 0x9e, 0x6c, 0x37, 0x7c, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x56, 0x4d,
	0x65, 0x7f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopoServiceClient is the client API for TopoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopoServiceClient interface {
	ReportDistance(ctx context.Context, in *Locations, opts ...grpc.CallOption) (*Topography, error)
	Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PingResponse, error)
}

type topoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopoServiceClient(cc *grpc.ClientConn) TopoServiceClient {
	return &topoServiceClient{cc}
}

func (c *topoServiceClient) ReportDistance(ctx context.Context, in *Locations, opts ...grpc.CallOption) (*Topography, error) {
	out := new(Topography)
	err := c.cc.Invoke(ctx, "/messages.TopoService/ReportDistance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topoServiceClient) Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/messages.TopoService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopoServiceServer is the server API for TopoService service.
type TopoServiceServer interface {
	ReportDistance(context.Context, *Locations) (*Topography, error)
	Ping(context.Context, *Void) (*PingResponse, error)
}

func RegisterTopoServiceServer(s *grpc.Server, srv TopoServiceServer) {
	s.RegisterService(&_TopoService_serviceDesc, srv)
}

func _TopoService_ReportDistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Locations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopoServiceServer).ReportDistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.TopoService/ReportDistance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopoServiceServer).ReportDistance(ctx, req.(*Locations))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopoService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopoServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.TopoService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopoServiceServer).Ping(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.TopoService",
	HandlerType: (*TopoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportDistance",
			Handler:    _TopoService_ReportDistance_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TopoService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent/messages/order.proto",
}
