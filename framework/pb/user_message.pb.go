// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_message.proto

package pb

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

type UserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserAuthRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAuthRequest) Reset()         { *m = UserAuthRequest{} }
func (m *UserAuthRequest) String() string { return proto.CompactTextString(m) }
func (*UserAuthRequest) ProtoMessage()    {}
func (*UserAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{2}
}

func (m *UserAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthRequest.Unmarshal(m, b)
}
func (m *UserAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthRequest.Marshal(b, m, deterministic)
}
func (m *UserAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthRequest.Merge(m, src)
}
func (m *UserAuthRequest) XXX_Size() int {
	return xxx_messageInfo_UserAuthRequest.Size(m)
}
func (m *UserAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthRequest proto.InternalMessageInfo

func (m *UserAuthRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserAuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "education.code.codeedu.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "education.code.codeedu.UserResponse")
	proto.RegisterType((*UserAuthRequest)(nil), "education.code.codeedu.UserAuthRequest")
}

func init() {
	proto.RegisterFile("user_message.proto", fileDescriptor_7965fb6944d08275)
}

var fileDescriptor_7965fb6944d08275 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x6d, 0xad, 0xa2, 0xa3, 0x20, 0x0c, 0x22, 0x65, 0x4f, 0x52, 0x17, 0xf4, 0xd4, 0x83,
	0xfe, 0x02, 0xdd, 0x7f, 0xb0, 0x8b, 0x88, 0x5e, 0x24, 0xdb, 0x3e, 0xb4, 0x68, 0x93, 0x98, 0x49,
	0xf4, 0xf7, 0xf9, 0xcf, 0xa4, 0x89, 0x2e, 0x2b, 0x68, 0xd9, 0x4b, 0xc8, 0x9b, 0x3c, 0xde, 0xcb,
	0xc7, 0x10, 0x07, 0x81, 0x7b, 0xec, 0x21, 0xa2, 0x9e, 0x50, 0x5b, 0x67, 0xbc, 0xe1, 0x13, 0xb4,
	0xa1, 0x51, 0xbe, 0x33, 0xba, 0x6e, 0x4c, 0x8b, 0x78, 0xa0, 0x0d, 0xd5, 0x82, 0x0e, 0x6e, 0x05,
	0x6e, 0x8e, 0xb7, 0x00, 0xf1, 0xcc, 0x54, 0x68, 0xd5, 0xa3, 0xcc, 0x4e, 0xb3, 0x8b, 0xfd, 0x79,
	0xbc, 0xf3, 0x31, 0xed, 0xa0, 0x57, 0xdd, 0x6b, 0x99, 0xc7, 0x61, 0x12, 0x3c, 0xa1, 0x3d, 0xab,
	0x44, 0x3e, 0x8c, 0x6b, 0xcb, 0xed, 0xf8, 0xb0, 0xd2, 0xd5, 0x94, 0x0e, 0x53, 0xa8, 0x58, 0xa3,
	0x25, 0x26, 0x78, 0xf3, 0x02, 0xfd, 0x1d, 0x9b, 0x44, 0x35, 0xa3, 0xa3, 0xc1, 0x75, 0x1d, 0xfc,
	0xf3, 0x4f, 0xfd, 0xaa, 0x2a, 0xfb, 0xaf, 0x2a, 0xff, 0x5d, 0x75, 0xf9, 0x99, 0x25, 0x80, 0x05,
	0xdc, 0x7b, 0xd7, 0x80, 0xef, 0x89, 0x66, 0x0e, 0xca, 0x63, 0x18, 0xf2, 0x59, 0xfd, 0x37, 0x76,
	0xbd, 0xc6, 0x3c, 0x99, 0x8e, 0x9b, 0x12, 0x43, 0xb5, 0xc5, 0x77, 0x54, 0x0c, 0x7f, 0xe5, 0xf3,
	0x31, 0xff, 0x1a, 0xcd, 0xa6, 0xc1, 0x37, 0xc5, 0x43, 0x6e, 0x97, 0xcb, 0xdd, 0xb8, 0xa8, 0xab,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xeb, 0xf2, 0x9d, 0xbe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Auth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/education.code.codeedu.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/education.code.codeedu.UserService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
	Auth(context.Context, *UserAuthRequest) (*UserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) Auth(ctx context.Context, req *UserAuthRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/education.code.codeedu.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/education.code.codeedu.UserService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Auth(ctx, req.(*UserAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "education.code.codeedu.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _UserService_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_message.proto",
}
