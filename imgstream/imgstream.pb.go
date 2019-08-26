// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imgstream/imgstream.proto

package imgstream

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

type ImageRequest struct {
	W                    int32    `protobuf:"varint,1,opt,name=w,proto3" json:"w,omitempty"`
	H                    int32    `protobuf:"varint,2,opt,name=h,proto3" json:"h,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageRequest) Reset()         { *m = ImageRequest{} }
func (m *ImageRequest) String() string { return proto.CompactTextString(m) }
func (*ImageRequest) ProtoMessage()    {}
func (*ImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a616566742048269, []int{0}
}

func (m *ImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageRequest.Unmarshal(m, b)
}
func (m *ImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageRequest.Marshal(b, m, deterministic)
}
func (m *ImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageRequest.Merge(m, src)
}
func (m *ImageRequest) XXX_Size() int {
	return xxx_messageInfo_ImageRequest.Size(m)
}
func (m *ImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImageRequest proto.InternalMessageInfo

func (m *ImageRequest) GetW() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *ImageRequest) GetH() int32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *ImageRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ImageStream struct {
	W                    int32    `protobuf:"varint,1,opt,name=w,proto3" json:"w,omitempty"`
	H                    int32    `protobuf:"varint,2,opt,name=h,proto3" json:"h,omitempty"`
	Image                []byte   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageStream) Reset()         { *m = ImageStream{} }
func (m *ImageStream) String() string { return proto.CompactTextString(m) }
func (*ImageStream) ProtoMessage()    {}
func (*ImageStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_a616566742048269, []int{1}
}

func (m *ImageStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageStream.Unmarshal(m, b)
}
func (m *ImageStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageStream.Marshal(b, m, deterministic)
}
func (m *ImageStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageStream.Merge(m, src)
}
func (m *ImageStream) XXX_Size() int {
	return xxx_messageInfo_ImageStream.Size(m)
}
func (m *ImageStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageStream.DiscardUnknown(m)
}

var xxx_messageInfo_ImageStream proto.InternalMessageInfo

func (m *ImageStream) GetW() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *ImageStream) GetH() int32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *ImageStream) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type CloseStreamRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseStreamRequest) Reset()         { *m = CloseStreamRequest{} }
func (m *CloseStreamRequest) String() string { return proto.CompactTextString(m) }
func (*CloseStreamRequest) ProtoMessage()    {}
func (*CloseStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a616566742048269, []int{2}
}

func (m *CloseStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseStreamRequest.Unmarshal(m, b)
}
func (m *CloseStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseStreamRequest.Marshal(b, m, deterministic)
}
func (m *CloseStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseStreamRequest.Merge(m, src)
}
func (m *CloseStreamRequest) XXX_Size() int {
	return xxx_messageInfo_CloseStreamRequest.Size(m)
}
func (m *CloseStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseStreamRequest proto.InternalMessageInfo

func (m *CloseStreamRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CloseStreamResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseStreamResponse) Reset()         { *m = CloseStreamResponse{} }
func (m *CloseStreamResponse) String() string { return proto.CompactTextString(m) }
func (*CloseStreamResponse) ProtoMessage()    {}
func (*CloseStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a616566742048269, []int{3}
}

func (m *CloseStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseStreamResponse.Unmarshal(m, b)
}
func (m *CloseStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseStreamResponse.Marshal(b, m, deterministic)
}
func (m *CloseStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseStreamResponse.Merge(m, src)
}
func (m *CloseStreamResponse) XXX_Size() int {
	return xxx_messageInfo_CloseStreamResponse.Size(m)
}
func (m *CloseStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseStreamResponse proto.InternalMessageInfo

func (m *CloseStreamResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageRequest)(nil), "imgstream.ImageRequest")
	proto.RegisterType((*ImageStream)(nil), "imgstream.ImageStream")
	proto.RegisterType((*CloseStreamRequest)(nil), "imgstream.CloseStreamRequest")
	proto.RegisterType((*CloseStreamResponse)(nil), "imgstream.CloseStreamResponse")
}

func init() { proto.RegisterFile("imgstream/imgstream.proto", fileDescriptor_a616566742048269) }

var fileDescriptor_a616566742048269 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4f, 0xb3, 0x40,
	0x10, 0xc6, 0xbb, 0x6d, 0xfa, 0xe6, 0xed, 0x80, 0xa6, 0x59, 0x8d, 0x22, 0x89, 0xa6, 0xe1, 0x22,
	0xa7, 0x4a, 0xf4, 0x03, 0x68, 0xdb, 0x84, 0xa4, 0x89, 0x5e, 0xc0, 0x93, 0x37, 0x6a, 0x26, 0x74,
	0x53, 0x96, 0x45, 0x86, 0xda, 0x6f, 0xee, 0xd9, 0xf0, 0x47, 0x8a, 0x35, 0xeb, 0xc1, 0xdb, 0x3e,
	0xcc, 0xc3, 0x6f, 0x9e, 0x99, 0x0c, 0x5c, 0x08, 0x19, 0x53, 0x91, 0x63, 0x24, 0x6f, 0xda, 0xd7,
	0x34, 0xcb, 0x55, 0xa1, 0xf8, 0xa8, 0xfd, 0xe0, 0xf8, 0x60, 0x2e, 0x65, 0x14, 0x63, 0x80, 0x6f,
	0x5b, 0xa4, 0x82, 0x9b, 0xc0, 0x76, 0x16, 0x9b, 0x30, 0x77, 0x18, 0xb0, 0x5d, 0xa9, 0xd6, 0x56,
	0xbf, 0x56, 0x6b, 0x6e, 0xc3, 0xff, 0x2d, 0x61, 0x9e, 0x46, 0x12, 0xad, 0xc1, 0x84, 0xb9, 0xa3,
	0xa0, 0xd5, 0xce, 0x3d, 0x18, 0x15, 0x27, 0xac, 0xb0, 0xbf, 0x62, 0x4e, 0x61, 0x28, 0x4a, 0x6b,
	0xc5, 0x30, 0x83, 0x5a, 0x38, 0x1e, 0xf0, 0x45, 0xa2, 0xa8, 0x01, 0x7c, 0xc5, 0xe9, 0xb6, 0x64,
	0x07, 0x2d, 0xaf, 0xe1, 0xe4, 0xdb, 0x1f, 0x94, 0xa9, 0x94, 0x90, 0x8f, 0x61, 0x20, 0x29, 0x6e,
	0xdc, 0xe5, 0xf3, 0xf6, 0xa3, 0x0f, 0xe3, 0xa5, 0x8c, 0x6b, 0x5f, 0x88, 0xf9, 0xbb, 0x78, 0x45,
	0x3e, 0x07, 0x63, 0x46, 0x9b, 0x67, 0x15, 0x44, 0x94, 0x65, 0x82, 0x9f, 0x4f, 0xf7, 0x4b, 0xea,
	0x2e, 0xc4, 0x3e, 0x3b, 0x2c, 0xd4, 0x18, 0xa7, 0xe7, 0x31, 0xfe, 0x08, 0x46, 0x27, 0x01, 0xbf,
	0xec, 0x58, 0x7f, 0xce, 0x62, 0x5f, 0xe9, 0xca, 0x4d, 0x70, 0x1f, 0x8e, 0x43, 0x4c, 0x0b, 0x3f,
	0x57, 0xb2, 0x09, 0xa5, 0xe9, 0xad, 0xcf, 0xe4, 0x32, 0x8f, 0x71, 0x1f, 0x8e, 0x66, 0xb4, 0x29,
	0x31, 0x4f, 0x6a, 0x25, 0x12, 0xfc, 0xc3, 0x6c, 0x15, 0xe7, 0x01, 0x20, 0x14, 0x32, 0x4b, 0x70,
	0x11, 0x25, 0x89, 0x1e, 0xa2, 0x2b, 0x38, 0xbd, 0xb9, 0xf1, 0xb2, 0xbf, 0xb4, 0xd5, 0xbf, 0xea,
	0xf6, 0xee, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x64, 0x81, 0xd1, 0x5d, 0x98, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImgStreamServiceClient is the client API for ImgStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImgStreamServiceClient interface {
	AskToRasppi(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (ImgStreamService_AskToRasppiClient, error)
	CloseStream(ctx context.Context, in *CloseStreamRequest, opts ...grpc.CallOption) (*CloseStreamResponse, error)
	SentFromRasppi(ctx context.Context, opts ...grpc.CallOption) (ImgStreamService_SentFromRasppiClient, error)
	AskFromMobile(ctx context.Context, opts ...grpc.CallOption) (ImgStreamService_AskFromMobileClient, error)
	SimpleCall(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageRequest, error)
}

type imgStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewImgStreamServiceClient(cc *grpc.ClientConn) ImgStreamServiceClient {
	return &imgStreamServiceClient{cc}
}

func (c *imgStreamServiceClient) AskToRasppi(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (ImgStreamService_AskToRasppiClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImgStreamService_serviceDesc.Streams[0], "/imgstream.ImgStreamService/AskToRasppi", opts...)
	if err != nil {
		return nil, err
	}
	x := &imgStreamServiceAskToRasppiClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImgStreamService_AskToRasppiClient interface {
	Recv() (*ImageStream, error)
	grpc.ClientStream
}

type imgStreamServiceAskToRasppiClient struct {
	grpc.ClientStream
}

func (x *imgStreamServiceAskToRasppiClient) Recv() (*ImageStream, error) {
	m := new(ImageStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imgStreamServiceClient) CloseStream(ctx context.Context, in *CloseStreamRequest, opts ...grpc.CallOption) (*CloseStreamResponse, error) {
	out := new(CloseStreamResponse)
	err := c.cc.Invoke(ctx, "/imgstream.ImgStreamService/CloseStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imgStreamServiceClient) SentFromRasppi(ctx context.Context, opts ...grpc.CallOption) (ImgStreamService_SentFromRasppiClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImgStreamService_serviceDesc.Streams[1], "/imgstream.ImgStreamService/SentFromRasppi", opts...)
	if err != nil {
		return nil, err
	}
	x := &imgStreamServiceSentFromRasppiClient{stream}
	return x, nil
}

type ImgStreamService_SentFromRasppiClient interface {
	Send(*ImageStream) error
	Recv() (*ImageStream, error)
	grpc.ClientStream
}

type imgStreamServiceSentFromRasppiClient struct {
	grpc.ClientStream
}

func (x *imgStreamServiceSentFromRasppiClient) Send(m *ImageStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imgStreamServiceSentFromRasppiClient) Recv() (*ImageStream, error) {
	m := new(ImageStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imgStreamServiceClient) AskFromMobile(ctx context.Context, opts ...grpc.CallOption) (ImgStreamService_AskFromMobileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImgStreamService_serviceDesc.Streams[2], "/imgstream.ImgStreamService/AskFromMobile", opts...)
	if err != nil {
		return nil, err
	}
	x := &imgStreamServiceAskFromMobileClient{stream}
	return x, nil
}

type ImgStreamService_AskFromMobileClient interface {
	Send(*ImageRequest) error
	Recv() (*ImageStream, error)
	grpc.ClientStream
}

type imgStreamServiceAskFromMobileClient struct {
	grpc.ClientStream
}

func (x *imgStreamServiceAskFromMobileClient) Send(m *ImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imgStreamServiceAskFromMobileClient) Recv() (*ImageStream, error) {
	m := new(ImageStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imgStreamServiceClient) SimpleCall(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageRequest, error) {
	out := new(ImageRequest)
	err := c.cc.Invoke(ctx, "/imgstream.ImgStreamService/SimpleCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImgStreamServiceServer is the server API for ImgStreamService service.
type ImgStreamServiceServer interface {
	AskToRasppi(*ImageRequest, ImgStreamService_AskToRasppiServer) error
	CloseStream(context.Context, *CloseStreamRequest) (*CloseStreamResponse, error)
	SentFromRasppi(ImgStreamService_SentFromRasppiServer) error
	AskFromMobile(ImgStreamService_AskFromMobileServer) error
	SimpleCall(context.Context, *ImageRequest) (*ImageRequest, error)
}

// UnimplementedImgStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImgStreamServiceServer struct {
}

func (*UnimplementedImgStreamServiceServer) AskToRasppi(req *ImageRequest, srv ImgStreamService_AskToRasppiServer) error {
	return status.Errorf(codes.Unimplemented, "method AskToRasppi not implemented")
}
func (*UnimplementedImgStreamServiceServer) CloseStream(ctx context.Context, req *CloseStreamRequest) (*CloseStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseStream not implemented")
}
func (*UnimplementedImgStreamServiceServer) SentFromRasppi(srv ImgStreamService_SentFromRasppiServer) error {
	return status.Errorf(codes.Unimplemented, "method SentFromRasppi not implemented")
}
func (*UnimplementedImgStreamServiceServer) AskFromMobile(srv ImgStreamService_AskFromMobileServer) error {
	return status.Errorf(codes.Unimplemented, "method AskFromMobile not implemented")
}
func (*UnimplementedImgStreamServiceServer) SimpleCall(ctx context.Context, req *ImageRequest) (*ImageRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleCall not implemented")
}

func RegisterImgStreamServiceServer(s *grpc.Server, srv ImgStreamServiceServer) {
	s.RegisterService(&_ImgStreamService_serviceDesc, srv)
}

func _ImgStreamService_AskToRasppi_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImgStreamServiceServer).AskToRasppi(m, &imgStreamServiceAskToRasppiServer{stream})
}

type ImgStreamService_AskToRasppiServer interface {
	Send(*ImageStream) error
	grpc.ServerStream
}

type imgStreamServiceAskToRasppiServer struct {
	grpc.ServerStream
}

func (x *imgStreamServiceAskToRasppiServer) Send(m *ImageStream) error {
	return x.ServerStream.SendMsg(m)
}

func _ImgStreamService_CloseStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImgStreamServiceServer).CloseStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imgstream.ImgStreamService/CloseStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImgStreamServiceServer).CloseStream(ctx, req.(*CloseStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImgStreamService_SentFromRasppi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImgStreamServiceServer).SentFromRasppi(&imgStreamServiceSentFromRasppiServer{stream})
}

type ImgStreamService_SentFromRasppiServer interface {
	Send(*ImageStream) error
	Recv() (*ImageStream, error)
	grpc.ServerStream
}

type imgStreamServiceSentFromRasppiServer struct {
	grpc.ServerStream
}

func (x *imgStreamServiceSentFromRasppiServer) Send(m *ImageStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imgStreamServiceSentFromRasppiServer) Recv() (*ImageStream, error) {
	m := new(ImageStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ImgStreamService_AskFromMobile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImgStreamServiceServer).AskFromMobile(&imgStreamServiceAskFromMobileServer{stream})
}

type ImgStreamService_AskFromMobileServer interface {
	Send(*ImageStream) error
	Recv() (*ImageRequest, error)
	grpc.ServerStream
}

type imgStreamServiceAskFromMobileServer struct {
	grpc.ServerStream
}

func (x *imgStreamServiceAskFromMobileServer) Send(m *ImageStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imgStreamServiceAskFromMobileServer) Recv() (*ImageRequest, error) {
	m := new(ImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ImgStreamService_SimpleCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImgStreamServiceServer).SimpleCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imgstream.ImgStreamService/SimpleCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImgStreamServiceServer).SimpleCall(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImgStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imgstream.ImgStreamService",
	HandlerType: (*ImgStreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloseStream",
			Handler:    _ImgStreamService_CloseStream_Handler,
		},
		{
			MethodName: "SimpleCall",
			Handler:    _ImgStreamService_SimpleCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AskToRasppi",
			Handler:       _ImgStreamService_AskToRasppi_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SentFromRasppi",
			Handler:       _ImgStreamService_SentFromRasppi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AskFromMobile",
			Handler:       _ImgStreamService_AskFromMobile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "imgstream/imgstream.proto",
}
