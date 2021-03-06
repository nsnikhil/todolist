// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type AddRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{1}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddRequest) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *AddRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type GetRequest struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

type RemoveRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{3}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RemoveRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type UpdateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status               bool     `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateRequest) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *UpdateRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ApiResponse struct {
	Data                 *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiResponse) Reset()         { *m = ApiResponse{} }
func (m *ApiResponse) String() string { return proto.CompactTextString(m) }
func (*ApiResponse) ProtoMessage()    {}
func (*ApiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{5}
}

func (m *ApiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiResponse.Unmarshal(m, b)
}
func (m *ApiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiResponse.Marshal(b, m, deterministic)
}
func (m *ApiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiResponse.Merge(m, src)
}
func (m *ApiResponse) XXX_Size() int {
	return xxx_messageInfo_ApiResponse.Size(m)
}
func (m *ApiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiResponse proto.InternalMessageInfo

func (m *ApiResponse) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "proto.ping_request")
	proto.RegisterType((*AddRequest)(nil), "proto.add_request")
	proto.RegisterType((*GetRequest)(nil), "proto.get_request")
	proto.RegisterType((*RemoveRequest)(nil), "proto.remove_request")
	proto.RegisterType((*UpdateRequest)(nil), "proto.update_request")
	proto.RegisterType((*ApiResponse)(nil), "proto.api_response")
}

func init() { proto.RegisterFile("proto/api.proto", fileDescriptor_ecf0878b123623e2) }

var fileDescriptor_ecf0878b123623e2 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xc9, 0x47, 0x8b, 0x9d, 0x94, 0x28, 0x63, 0x95, 0x58, 0x10, 0x42, 0x4e, 0x39, 0xa5,
	0x52, 0x3d, 0x78, 0xf5, 0x15, 0xf2, 0x02, 0x65, 0xeb, 0xae, 0x61, 0xa1, 0x66, 0xb7, 0xd9, 0x89,
	0xd0, 0xb3, 0x3e, 0xb8, 0x74, 0xb7, 0x89, 0x29, 0xb6, 0x3d, 0x65, 0x3e, 0x7e, 0xc3, 0xe6, 0xcf,
	0x0f, 0xae, 0x75, 0xa3, 0x48, 0x2d, 0x98, 0x96, 0x85, 0xad, 0x70, 0x64, 0x3f, 0xf3, 0x87, 0x4a,
	0xa9, 0x6a, 0x23, 0x16, 0xb6, 0x5b, 0xb7, 0x1f, 0x0b, 0x56, 0xef, 0x1c, 0x91, 0xc5, 0x30, 0xd5,
	0xb2, 0xae, 0x56, 0x8d, 0xd8, 0xb6, 0xc2, 0x50, 0xb6, 0x85, 0x88, 0x71, 0xde, 0xb5, 0x38, 0x83,
	0x11, 0x49, 0xda, 0x88, 0xc4, 0x4b, 0xbd, 0x7c, 0x52, 0xba, 0x06, 0x53, 0x88, 0xb8, 0x30, 0xef,
	0x8d, 0xd4, 0x24, 0x55, 0x9d, 0xf8, 0x76, 0x37, 0x1c, 0xe1, 0x3d, 0x8c, 0x0d, 0x31, 0x6a, 0x4d,
	0x12, 0xa4, 0x5e, 0x7e, 0x55, 0x1e, 0x3a, 0x44, 0x08, 0x89, 0x55, 0x26, 0x09, 0xd3, 0x20, 0x9f,
	0x94, 0xb6, 0xce, 0x1e, 0x21, 0xaa, 0x04, 0xf5, 0x4f, 0xc6, 0xe0, 0x4b, 0x9e, 0x78, 0x16, 0xf0,
	0x25, 0xcf, 0x96, 0x10, 0x37, 0xe2, 0x53, 0x7d, 0x89, 0x7f, 0x84, 0xe7, 0x08, 0xbc, 0x81, 0x40,
	0x72, 0x93, 0xf8, 0xf6, 0x64, 0x5f, 0x66, 0x3f, 0x1e, 0xc4, 0xad, 0xe6, 0x8c, 0xce, 0x1f, 0xf5,
	0xc9, 0xfc, 0x0b, 0xc9, 0x82, 0x4b, 0xc9, 0xc2, 0x93, 0xc9, 0x46, 0x83, 0x64, 0xaf, 0x30, 0x65,
	0x5a, 0xae, 0x1a, 0x61, 0xb4, 0xaa, 0x8d, 0xc0, 0x1c, 0x42, 0xce, 0x88, 0xd9, 0xbf, 0x88, 0x96,
	0xb3, 0xc2, 0x69, 0x29, 0x3a, 0x2d, 0xc5, 0x5b, 0xbd, 0x2b, 0x2d, 0xb1, 0xfc, 0xf6, 0x21, 0x60,
	0x5a, 0xe2, 0x13, 0x84, 0x7b, 0x3d, 0x78, 0xeb, 0xa0, 0x62, 0xe8, 0x6a, 0xde, 0x0d, 0x8f, 0xde,
	0x28, 0x20, 0x60, 0x9c, 0x23, 0x76, 0xbb, 0x3f, 0x99, 0x67, 0xf9, 0x4a, 0x50, 0xcf, 0x0f, 0x4c,
	0x9c, 0xe6, 0x5f, 0x60, 0xec, 0x74, 0xe0, 0xdd, 0x61, 0x7d, 0x6c, 0xe7, 0xec, 0x95, 0xf3, 0xd1,
	0x5f, 0x1d, 0xeb, 0x39, 0x79, 0xb5, 0x1e, 0xdb, 0xd9, 0xf3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x82, 0xab, 0xae, 0x7a, 0xd8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*ApiResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*ApiResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ApiResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*ApiResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ApiResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/proto.api/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/proto.api/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/proto.api/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/proto.api/remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/proto.api/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	Ping(context.Context, *PingRequest) (*ApiResponse, error)
	Add(context.Context, *AddRequest) (*ApiResponse, error)
	Get(context.Context, *GetRequest) (*ApiResponse, error)
	Remove(context.Context, *RemoveRequest) (*ApiResponse, error)
	Update(context.Context, *UpdateRequest) (*ApiResponse, error)
}

// UnimplementedApiServer can be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (*UnimplementedApiServer) Ping(ctx context.Context, req *PingRequest) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedApiServer) Add(ctx context.Context, req *AddRequest) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedApiServer) Get(ctx context.Context, req *GetRequest) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedApiServer) Remove(ctx context.Context, req *RemoveRequest) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedApiServer) Update(ctx context.Context, req *UpdateRequest) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ping",
			Handler:    _Api_Ping_Handler,
		},
		{
			MethodName: "add",
			Handler:    _Api_Add_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Api_Get_Handler,
		},
		{
			MethodName: "remove",
			Handler:    _Api_Remove_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Api_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
