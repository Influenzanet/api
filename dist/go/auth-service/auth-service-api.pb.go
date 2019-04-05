// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth-service-api.proto

package auth_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_go "github.com/influenzanet/api/dist/go"
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

type EncodedToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncodedToken) Reset()         { *m = EncodedToken{} }
func (m *EncodedToken) String() string { return proto.CompactTextString(m) }
func (*EncodedToken) ProtoMessage()    {}
func (*EncodedToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{0}
}

func (m *EncodedToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodedToken.Unmarshal(m, b)
}
func (m *EncodedToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodedToken.Marshal(b, m, deterministic)
}
func (m *EncodedToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedToken.Merge(m, src)
}
func (m *EncodedToken) XXX_Size() int {
	return xxx_messageInfo_EncodedToken.Size(m)
}
func (m *EncodedToken) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedToken.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedToken proto.InternalMessageInfo

func (m *EncodedToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TempTokenInfo struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Expiration           int64    `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Purpose              string   `protobuf:"bytes,3,opt,name=purpose,proto3" json:"purpose,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Info                 string   `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	InstanceId           string   `protobuf:"bytes,6,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TempTokenInfo) Reset()         { *m = TempTokenInfo{} }
func (m *TempTokenInfo) String() string { return proto.CompactTextString(m) }
func (*TempTokenInfo) ProtoMessage()    {}
func (*TempTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{1}
}

func (m *TempTokenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TempTokenInfo.Unmarshal(m, b)
}
func (m *TempTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TempTokenInfo.Marshal(b, m, deterministic)
}
func (m *TempTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TempTokenInfo.Merge(m, src)
}
func (m *TempTokenInfo) XXX_Size() int {
	return xxx_messageInfo_TempTokenInfo.Size(m)
}
func (m *TempTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TempTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TempTokenInfo proto.InternalMessageInfo

func (m *TempTokenInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TempTokenInfo) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *TempTokenInfo) GetPurpose() string {
	if m != nil {
		return m.Purpose
	}
	return ""
}

func (m *TempTokenInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TempTokenInfo) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *TempTokenInfo) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

type TempTokenInfoResponse struct {
	TempToken            *TempTokenInfo `protobuf:"bytes,1,opt,name=temp_token,json=tempToken,proto3" json:"temp_token,omitempty"`
	Status               *_go.Status    `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TempTokenInfoResponse) Reset()         { *m = TempTokenInfoResponse{} }
func (m *TempTokenInfoResponse) String() string { return proto.CompactTextString(m) }
func (*TempTokenInfoResponse) ProtoMessage()    {}
func (*TempTokenInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{2}
}

func (m *TempTokenInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TempTokenInfoResponse.Unmarshal(m, b)
}
func (m *TempTokenInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TempTokenInfoResponse.Marshal(b, m, deterministic)
}
func (m *TempTokenInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TempTokenInfoResponse.Merge(m, src)
}
func (m *TempTokenInfoResponse) XXX_Size() int {
	return xxx_messageInfo_TempTokenInfoResponse.Size(m)
}
func (m *TempTokenInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TempTokenInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TempTokenInfoResponse proto.InternalMessageInfo

func (m *TempTokenInfoResponse) GetTempToken() *TempTokenInfo {
	if m != nil {
		return m.TempToken
	}
	return nil
}

func (m *TempTokenInfoResponse) GetStatus() *_go.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type TempTokenInfos struct {
	TempTokens           []*TempTokenInfo `protobuf:"bytes,1,rep,name=temp_tokens,json=tempTokens,proto3" json:"temp_tokens,omitempty"`
	Status               *_go.Status      `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TempTokenInfos) Reset()         { *m = TempTokenInfos{} }
func (m *TempTokenInfos) String() string { return proto.CompactTextString(m) }
func (*TempTokenInfos) ProtoMessage()    {}
func (*TempTokenInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{3}
}

func (m *TempTokenInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TempTokenInfos.Unmarshal(m, b)
}
func (m *TempTokenInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TempTokenInfos.Marshal(b, m, deterministic)
}
func (m *TempTokenInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TempTokenInfos.Merge(m, src)
}
func (m *TempTokenInfos) XXX_Size() int {
	return xxx_messageInfo_TempTokenInfos.Size(m)
}
func (m *TempTokenInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_TempTokenInfos.DiscardUnknown(m)
}

var xxx_messageInfo_TempTokenInfos proto.InternalMessageInfo

func (m *TempTokenInfos) GetTempTokens() []*TempTokenInfo {
	if m != nil {
		return m.TempTokens
	}
	return nil
}

func (m *TempTokenInfos) GetStatus() *_go.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type TempToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TempToken) Reset()         { *m = TempToken{} }
func (m *TempToken) String() string { return proto.CompactTextString(m) }
func (*TempToken) ProtoMessage()    {}
func (*TempToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{4}
}

func (m *TempToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TempToken.Unmarshal(m, b)
}
func (m *TempToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TempToken.Marshal(b, m, deterministic)
}
func (m *TempToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TempToken.Merge(m, src)
}
func (m *TempToken) XXX_Size() int {
	return xxx_messageInfo_TempToken.Size(m)
}
func (m *TempToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TempToken.DiscardUnknown(m)
}

var xxx_messageInfo_TempToken proto.InternalMessageInfo

func (m *TempToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TempTokenResponse struct {
	TempToken            *TempToken  `protobuf:"bytes,1,opt,name=temp_token,json=tempToken,proto3" json:"temp_token,omitempty"`
	Status               *_go.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TempTokenResponse) Reset()         { *m = TempTokenResponse{} }
func (m *TempTokenResponse) String() string { return proto.CompactTextString(m) }
func (*TempTokenResponse) ProtoMessage()    {}
func (*TempTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{5}
}

func (m *TempTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TempTokenResponse.Unmarshal(m, b)
}
func (m *TempTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TempTokenResponse.Marshal(b, m, deterministic)
}
func (m *TempTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TempTokenResponse.Merge(m, src)
}
func (m *TempTokenResponse) XXX_Size() int {
	return xxx_messageInfo_TempTokenResponse.Size(m)
}
func (m *TempTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TempTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TempTokenResponse proto.InternalMessageInfo

func (m *TempTokenResponse) GetTempToken() *TempToken {
	if m != nil {
		return m.TempToken
	}
	return nil
}

func (m *TempTokenResponse) GetStatus() *_go.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*EncodedToken)(nil), "influenzanet.auth_service_api.EncodedToken")
	proto.RegisterType((*TempTokenInfo)(nil), "influenzanet.auth_service_api.TempTokenInfo")
	proto.RegisterType((*TempTokenInfoResponse)(nil), "influenzanet.auth_service_api.TempTokenInfoResponse")
	proto.RegisterType((*TempTokenInfos)(nil), "influenzanet.auth_service_api.TempTokenInfos")
	proto.RegisterType((*TempToken)(nil), "influenzanet.auth_service_api.TempToken")
	proto.RegisterType((*TempTokenResponse)(nil), "influenzanet.auth_service_api.TempTokenResponse")
}

func init() { proto.RegisterFile("auth-service-api.proto", fileDescriptor_e8e86308594dc230) }

var fileDescriptor_e8e86308594dc230 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x95, 0xbf, 0xb4, 0xee, 0x97, 0x49, 0x9b, 0x8a, 0x55, 0x29, 0xc6, 0xa8, 0x10, 0x22, 0x2e,
	0x22, 0xd1, 0xd8, 0x55, 0x40, 0xdc, 0xb7, 0x10, 0x45, 0xe1, 0x47, 0x42, 0x49, 0xda, 0x4a, 0x08,
	0x14, 0x6d, 0xe2, 0x89, 0xb3, 0xc2, 0xd9, 0x5d, 0x79, 0xd7, 0x40, 0x78, 0x04, 0x24, 0x6e, 0x78,
	0x0f, 0x5e, 0x8a, 0x27, 0x41, 0xb6, 0xe3, 0x24, 0x96, 0x02, 0xc4, 0x48, 0xdc, 0x79, 0x76, 0xce,
	0x9c, 0x3d, 0x33, 0x7b, 0x26, 0x81, 0x63, 0x1a, 0xe9, 0x69, 0x53, 0x61, 0xf8, 0x81, 0x8d, 0xb1,
	0x49, 0x25, 0x73, 0x64, 0x28, 0xb4, 0x20, 0x27, 0x8c, 0x4f, 0x82, 0x08, 0xf9, 0x67, 0xca, 0x51,
	0x3b, 0x31, 0x68, 0xb8, 0x00, 0x0d, 0xa9, 0x64, 0x36, 0xf1, 0x03, 0x31, 0xa2, 0x41, 0x53, 0xcf,
	0x25, 0xaa, 0xb4, 0xc4, 0xbe, 0xe3, 0x0b, 0xe1, 0x07, 0xe8, 0x26, 0xd1, 0x28, 0x9a, 0xb8, 0x38,
	0x93, 0x7a, 0x9e, 0x26, 0xeb, 0x0f, 0x60, 0xbf, 0xcd, 0xc7, 0xc2, 0x43, 0x6f, 0x20, 0xde, 0x23,
	0x27, 0x47, 0xb0, 0xab, 0xe3, 0x0f, 0xcb, 0xa8, 0x19, 0x8d, 0x72, 0x2f, 0x0d, 0xea, 0xdf, 0x0d,
	0x38, 0x18, 0xe0, 0x4c, 0x26, 0x98, 0x2e, 0x9f, 0x88, 0xcd, 0x38, 0x72, 0x17, 0x00, 0x3f, 0x49,
	0x16, 0x52, 0xcd, 0x04, 0xb7, 0xfe, 0xab, 0x19, 0x8d, 0x52, 0x6f, 0xed, 0x84, 0x58, 0xb0, 0x27,
	0xa3, 0x50, 0x0a, 0x85, 0x56, 0x29, 0xa9, 0xcb, 0x42, 0x72, 0x0b, 0xf6, 0x22, 0x85, 0xe1, 0x90,
	0x79, 0xd6, 0x4e, 0x92, 0x31, 0xe3, 0xb0, 0xeb, 0x11, 0x02, 0x3b, 0x8c, 0x4f, 0x84, 0xb5, 0x9b,
	0x9c, 0x26, 0xdf, 0xe4, 0x1e, 0x54, 0x18, 0x57, 0x9a, 0xf2, 0x31, 0xc6, 0x05, 0x66, 0x92, 0x82,
	0xec, 0xa8, 0xeb, 0xd5, 0xbf, 0x19, 0x70, 0x33, 0xa7, 0xb7, 0x87, 0x4a, 0x0a, 0xae, 0x90, 0xbc,
	0x00, 0xd0, 0x38, 0x93, 0xc3, 0x95, 0xf8, 0x4a, 0xeb, 0xd4, 0xf9, 0xed, 0x50, 0x9d, 0x3c, 0x53,
	0x59, 0x67, 0x21, 0x39, 0x05, 0x53, 0x69, 0xaa, 0x23, 0x95, 0xb4, 0x5a, 0x69, 0x1d, 0xe5, 0x89,
	0xfa, 0x49, 0xae, 0xb7, 0xc0, 0xd4, 0xbf, 0x1a, 0x50, 0xcd, 0x51, 0x29, 0xf2, 0x0a, 0x2a, 0x2b,
	0x35, 0xca, 0x32, 0x6a, 0xa5, 0xc2, 0x72, 0x60, 0x29, 0x47, 0x15, 0xd4, 0x73, 0x1f, 0xca, 0x4b,
	0xaa, 0x5f, 0xbc, 0xfb, 0x17, 0x03, 0x6e, 0x2c, 0x31, 0xcb, 0x19, 0x76, 0x36, 0xcc, 0xb0, 0xb1,
	0xad, 0xe8, 0xbf, 0x9e, 0x5f, 0xeb, 0x87, 0x09, 0xd5, 0xf3, 0x48, 0x4f, 0xfb, 0x29, 0xed, 0xb9,
	0x64, 0xe4, 0x09, 0x98, 0x29, 0x88, 0x1c, 0x3b, 0xa9, 0xcb, 0x9d, 0xcc, 0xe5, 0x4e, 0x3b, 0x76,
	0xb9, 0xbd, 0x91, 0x92, 0xbc, 0x85, 0xea, 0x4b, 0xe1, 0x33, 0x7e, 0xcd, 0xf4, 0xb4, 0x3d, 0xa3,
	0x2c, 0x20, 0x27, 0x79, 0xdc, 0xa5, 0xc2, 0xf0, 0x69, 0x88, 0x1e, 0x72, 0xcd, 0x68, 0xa0, 0xec,
	0x87, 0x7f, 0x68, 0x2f, 0xb7, 0x43, 0xef, 0xe0, 0xb0, 0xcf, 0x7c, 0x1e, 0xc9, 0x7f, 0x43, 0x7f,
	0x09, 0x95, 0x2b, 0x1a, 0x30, 0x8f, 0x6a, 0x7c, 0x7e, 0x3d, 0x20, 0x45, 0x6a, 0xed, 0xdb, 0x79,
	0xf0, 0x6b, 0x1a, 0xaa, 0x8c, 0xd6, 0x83, 0xff, 0x7b, 0xc8, 0xf1, 0x63, 0x61, 0xce, 0x42, 0xe2,
	0x05, 0x1c, 0x74, 0x90, 0x63, 0x48, 0x35, 0x2e, 0x3c, 0x50, 0xc4, 0xed, 0xf6, 0xd9, 0xd6, 0x36,
	0xcb, 0xcc, 0xca, 0x61, 0xff, 0x8a, 0x06, 0xde, 0xf2, 0xbe, 0xad, 0x8d, 0x6a, 0x3f, 0x2e, 0xb4,
	0x87, 0xab, 0xfb, 0x0e, 0x3b, 0xa8, 0xd3, 0x85, 0xbc, 0x98, 0xc7, 0x0f, 0x5d, 0xb0, 0xc5, 0x66,
	0x11, 0x74, 0xf2, 0x13, 0xf2, 0x0c, 0x03, 0x2c, 0xde, 0xde, 0xc6, 0xcd, 0xb8, 0x68, 0xbd, 0x39,
	0xf3, 0x99, 0x9e, 0x46, 0x23, 0x67, 0x2c, 0x66, 0xee, 0x3a, 0xc2, 0xa5, 0x92, 0xb9, 0x1e, 0x53,
	0xda, 0xf5, 0x85, 0xbb, 0xfe, 0xef, 0x34, 0x32, 0x93, 0x9d, 0x7b, 0xf4, 0x33, 0x00, 0x00, 0xff,
	0xff, 0x2c, 0xd0, 0xd2, 0x7c, 0xb4, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceApiClient is the client API for AuthServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceApiClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*_go.Status, error)
	// Auth:
	LoginWithEmail(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error)
	SignupWithEmail(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error)
	// Token handling:
	ValidateJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*_go.ParsedToken, error)
	RenewJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*EncodedToken, error)
	// Temporary Tokens handling:
	GenerateToken(ctx context.Context, in *TempTokenInfo, opts ...grpc.CallOption) (*TempTokenResponse, error)
	ValdateToken(ctx context.Context, in *TempToken, opts ...grpc.CallOption) (*TempTokenInfoResponse, error)
	GetTokensByUser(ctx context.Context, in *TempTokenInfo, opts ...grpc.CallOption) (*TempTokenInfos, error)
	DeleteToken(ctx context.Context, in *TempToken, opts ...grpc.CallOption) (*_go.Status, error)
}

type authServiceApiClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceApiClient(cc *grpc.ClientConn) AuthServiceApiClient {
	return &authServiceApiClient{cc}
}

func (c *authServiceApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*_go.Status, error) {
	out := new(_go.Status)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) LoginWithEmail(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/LoginWithEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) SignupWithEmail(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/SignupWithEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) ValidateJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*_go.ParsedToken, error) {
	out := new(_go.ParsedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/ValidateJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) RenewJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/RenewJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) GenerateToken(ctx context.Context, in *TempTokenInfo, opts ...grpc.CallOption) (*TempTokenResponse, error) {
	out := new(TempTokenResponse)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) ValdateToken(ctx context.Context, in *TempToken, opts ...grpc.CallOption) (*TempTokenInfoResponse, error) {
	out := new(TempTokenInfoResponse)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/ValdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) GetTokensByUser(ctx context.Context, in *TempTokenInfo, opts ...grpc.CallOption) (*TempTokenInfos, error) {
	out := new(TempTokenInfos)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/GetTokensByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) DeleteToken(ctx context.Context, in *TempToken, opts ...grpc.CallOption) (*_go.Status, error) {
	out := new(_go.Status)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceApiServer is the server API for AuthServiceApi service.
type AuthServiceApiServer interface {
	Status(context.Context, *empty.Empty) (*_go.Status, error)
	// Auth:
	LoginWithEmail(context.Context, *_go.UserCredentials) (*EncodedToken, error)
	SignupWithEmail(context.Context, *_go.UserCredentials) (*EncodedToken, error)
	// Token handling:
	ValidateJWT(context.Context, *EncodedToken) (*_go.ParsedToken, error)
	RenewJWT(context.Context, *EncodedToken) (*EncodedToken, error)
	// Temporary Tokens handling:
	GenerateToken(context.Context, *TempTokenInfo) (*TempTokenResponse, error)
	ValdateToken(context.Context, *TempToken) (*TempTokenInfoResponse, error)
	GetTokensByUser(context.Context, *TempTokenInfo) (*TempTokenInfos, error)
	DeleteToken(context.Context, *TempToken) (*_go.Status, error)
}

func RegisterAuthServiceApiServer(s *grpc.Server, srv AuthServiceApiServer) {
	s.RegisterService(&_AuthServiceApi_serviceDesc, srv)
}

func _AuthServiceApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_LoginWithEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).LoginWithEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/LoginWithEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).LoginWithEmail(ctx, req.(*_go.UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_SignupWithEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).SignupWithEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/SignupWithEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).SignupWithEmail(ctx, req.(*_go.UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_ValidateJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodedToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).ValidateJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/ValidateJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).ValidateJWT(ctx, req.(*EncodedToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_RenewJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodedToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).RenewJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/RenewJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).RenewJWT(ctx, req.(*EncodedToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TempTokenInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).GenerateToken(ctx, req.(*TempTokenInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_ValdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TempToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).ValdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/ValdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).ValdateToken(ctx, req.(*TempToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_GetTokensByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TempTokenInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).GetTokensByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/GetTokensByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).GetTokensByUser(ctx, req.(*TempTokenInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TempToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).DeleteToken(ctx, req.(*TempToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthServiceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "influenzanet.auth_service_api.AuthServiceApi",
	HandlerType: (*AuthServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AuthServiceApi_Status_Handler,
		},
		{
			MethodName: "LoginWithEmail",
			Handler:    _AuthServiceApi_LoginWithEmail_Handler,
		},
		{
			MethodName: "SignupWithEmail",
			Handler:    _AuthServiceApi_SignupWithEmail_Handler,
		},
		{
			MethodName: "ValidateJWT",
			Handler:    _AuthServiceApi_ValidateJWT_Handler,
		},
		{
			MethodName: "RenewJWT",
			Handler:    _AuthServiceApi_RenewJWT_Handler,
		},
		{
			MethodName: "GenerateToken",
			Handler:    _AuthServiceApi_GenerateToken_Handler,
		},
		{
			MethodName: "ValdateToken",
			Handler:    _AuthServiceApi_ValdateToken_Handler,
		},
		{
			MethodName: "GetTokensByUser",
			Handler:    _AuthServiceApi_GetTokensByUser_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _AuthServiceApi_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-service-api.proto",
}
