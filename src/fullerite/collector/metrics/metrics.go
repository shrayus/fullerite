// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics/metrics.proto

package metrics

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MetricsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsRequest) Reset()         { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()    {}
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1af212157fe6fb, []int{0}
}

func (m *MetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsRequest.Unmarshal(m, b)
}
func (m *MetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsRequest.Marshal(b, m, deterministic)
}
func (m *MetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsRequest.Merge(m, src)
}
func (m *MetricsRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsRequest.Size(m)
}
func (m *MetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsRequest proto.InternalMessageInfo

type MetricsResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsResponse) Reset()         { *m = MetricsResponse{} }
func (m *MetricsResponse) String() string { return proto.CompactTextString(m) }
func (*MetricsResponse) ProtoMessage()    {}
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1af212157fe6fb, []int{1}
}

func (m *MetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsResponse.Unmarshal(m, b)
}
func (m *MetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsResponse.Marshal(b, m, deterministic)
}
func (m *MetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsResponse.Merge(m, src)
}
func (m *MetricsResponse) XXX_Size() int {
	return xxx_messageInfo_MetricsResponse.Size(m)
}
func (m *MetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsResponse proto.InternalMessageInfo

func (m *MetricsResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*MetricsRequest)(nil), "com.yelp.generated.queryengine.metrics.MetricsRequest")
	proto.RegisterType((*MetricsResponse)(nil), "com.yelp.generated.queryengine.metrics.MetricsResponse")
}

func init() { proto.RegisterFile("metrics/metrics.proto", fileDescriptor_cb1af212157fe6fb) }

var fileDescriptor_cb1af212157fe6fb = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0x2d, 0x29,
	0xca, 0x4c, 0x2e, 0xd6, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6a, 0xc9, 0xf9,
	0xb9, 0x7a, 0x95, 0xa9, 0x39, 0x05, 0x7a, 0xe9, 0xa9, 0x79, 0xa9, 0x45, 0x89, 0x25, 0xa9, 0x29,
	0x7a, 0x85, 0xa5, 0xa9, 0x45, 0x95, 0xa9, 0x79, 0xe9, 0x99, 0x79, 0xa9, 0x7a, 0x50, 0xd5, 0x4a,
	0x02, 0x5c, 0x7c, 0xbe, 0x10, 0x66, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x2a, 0x17,
	0x3f, 0x5c, 0xa4, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24,
	0x51, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x36, 0x6a, 0x67, 0xe4, 0x62, 0x87, 0xaa,
	0x13, 0xaa, 0x41, 0x30, 0xcd, 0xf4, 0x88, 0xb3, 0x58, 0x0f, 0xd5, 0x56, 0x29, 0x73, 0x92, 0xf5,
	0x41, 0xdc, 0xa6, 0xc4, 0xe0, 0x64, 0x12, 0x65, 0x04, 0xd2, 0x97, 0x9c, 0x5f, 0x54, 0x00, 0x32,
	0x44, 0xbf, 0x28, 0x31, 0x2f, 0x3b, 0x33, 0x2f, 0x3d, 0xbe, 0x20, 0x27, 0xb1, 0x24, 0x2d, 0xbf,
	0x28, 0x57, 0x1f, 0x6c, 0x4c, 0x3c, 0xc4, 0x1c, 0x58, 0x30, 0x25, 0xb1, 0x81, 0xc3, 0xc9, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xda, 0xaa, 0xee, 0x40, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsClient interface {
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type metricsClient struct {
	cc *grpc.ClientConn
}

func NewMetricsClient(cc *grpc.ClientConn) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/com.yelp.generated.queryengine.metrics.Metrics/Metrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServer is the server API for Metrics service.
type MetricsServer interface {
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
}

// UnimplementedMetricsServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (*UnimplementedMetricsServer) Metrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}

func RegisterMetricsServer(s *grpc.Server, srv MetricsServer) {
	s.RegisterService(&_Metrics_serviceDesc, srv)
}

func _Metrics_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.yelp.generated.queryengine.metrics.Metrics/Metrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).Metrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.yelp.generated.queryengine.metrics.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Metrics",
			Handler:    _Metrics_Metrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics/metrics.proto",
}
