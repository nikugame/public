// Code generated by protoc-gen-go.
// source: LoginProto.proto
// DO NOT EDIT!

/*
Package SDKProto is a generated protocol buffer package.

It is generated from these files:
	LoginProto.proto

It has these top-level messages:
	CertificateRequest
	CertificateReply
	DeviceInfo
	RegisterRequest
	AccountLoginRequest
	PhoneLoginRequest
	LoginReply
	AuthCodeRequest
	Reply
	AutoLoginRequest
	BindingPhoneRequest
	BindingAccountRequest
	ChangePassWordRequest
	WeixinLoginRequest
	HeartBeatMsg
*/
package SDKProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CertificateRequest struct {
	UUID          string `protobuf:"bytes,1,opt,name=UUID" json:"UUID,omitempty"`
	Type          int32  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	CertificateID string `protobuf:"bytes,4,opt,name=CertificateID" json:"CertificateID,omitempty"`
}

func (m *CertificateRequest) Reset()                    { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()               {}
func (*CertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CertificateRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *CertificateRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CertificateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateRequest) GetCertificateID() string {
	if m != nil {
		return m.CertificateID
	}
	return ""
}

type CertificateReply struct {
	Result  bool   `protobuf:"varint,1,opt,name=Result" json:"Result,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=Comment" json:"Comment,omitempty"`
}

func (m *CertificateReply) Reset()                    { *m = CertificateReply{} }
func (m *CertificateReply) String() string            { return proto.CompactTextString(m) }
func (*CertificateReply) ProtoMessage()               {}
func (*CertificateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *CertificateReply) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type DeviceInfo struct {
	GameID     string `protobuf:"bytes,1,opt,name=GameID" json:"GameID,omitempty"`
	Channel    string `protobuf:"bytes,2,opt,name=Channel" json:"Channel,omitempty"`
	DeviceType string `protobuf:"bytes,3,opt,name=DeviceType" json:"DeviceType,omitempty"`
	IMEI       string `protobuf:"bytes,4,opt,name=IMEI" json:"IMEI,omitempty"`
	MAC        string `protobuf:"bytes,5,opt,name=MAC" json:"MAC,omitempty"`
	IPv4       string `protobuf:"bytes,6,opt,name=IPv4" json:"IPv4,omitempty"`
	IPv6       string `protobuf:"bytes,7,opt,name=IPv6" json:"IPv6,omitempty"`
	OS         string `protobuf:"bytes,8,opt,name=OS" json:"OS,omitempty"`
	Latitude   string `protobuf:"bytes,9,opt,name=Latitude" json:"Latitude,omitempty"`
	Longitude  string `protobuf:"bytes,10,opt,name=Longitude" json:"Longitude,omitempty"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeviceInfo) GetGameID() string {
	if m != nil {
		return m.GameID
	}
	return ""
}

func (m *DeviceInfo) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *DeviceInfo) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *DeviceInfo) GetIMEI() string {
	if m != nil {
		return m.IMEI
	}
	return ""
}

func (m *DeviceInfo) GetMAC() string {
	if m != nil {
		return m.MAC
	}
	return ""
}

func (m *DeviceInfo) GetIPv4() string {
	if m != nil {
		return m.IPv4
	}
	return ""
}

func (m *DeviceInfo) GetIPv6() string {
	if m != nil {
		return m.IPv6
	}
	return ""
}

func (m *DeviceInfo) GetOS() string {
	if m != nil {
		return m.OS
	}
	return ""
}

func (m *DeviceInfo) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *DeviceInfo) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

// 注册普通账号
type RegisterRequest struct {
	UserName   string      `protobuf:"bytes,1,opt,name=UserName" json:"UserName,omitempty"`
	PassWord   string      `protobuf:"bytes,2,opt,name=PassWord" json:"PassWord,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,3,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterRequest) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *RegisterRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// 普通账号登陆
type AccountLoginRequest struct {
	UserName   string      `protobuf:"bytes,1,opt,name=UserName" json:"UserName,omitempty"`
	PassWord   string      `protobuf:"bytes,2,opt,name=PassWord" json:"PassWord,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,3,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *AccountLoginRequest) Reset()                    { *m = AccountLoginRequest{} }
func (m *AccountLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountLoginRequest) ProtoMessage()               {}
func (*AccountLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AccountLoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AccountLoginRequest) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *AccountLoginRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// 手机号码登录
type PhoneLoginRequest struct {
	Phone      string      `protobuf:"bytes,1,opt,name=Phone" json:"Phone,omitempty"`
	Code       string      `protobuf:"bytes,2,opt,name=Code" json:"Code,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,3,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *PhoneLoginRequest) Reset()                    { *m = PhoneLoginRequest{} }
func (m *PhoneLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*PhoneLoginRequest) ProtoMessage()               {}
func (*PhoneLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PhoneLoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PhoneLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PhoneLoginRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

type LoginReply struct {
	UUID         string `protobuf:"bytes,1,opt,name=UUID" json:"UUID,omitempty"`
	AutoCode     string `protobuf:"bytes,2,opt,name=AutoCode" json:"AutoCode,omitempty"`
	Token        string `protobuf:"bytes,3,opt,name=Token" json:"Token,omitempty"`
	Phone        string `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
	Certificated bool   `protobuf:"varint,5,opt,name=Certificated" json:"Certificated,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoginReply) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *LoginReply) GetAutoCode() string {
	if m != nil {
		return m.AutoCode
	}
	return ""
}

func (m *LoginReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReply) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginReply) GetCertificated() bool {
	if m != nil {
		return m.Certificated
	}
	return false
}

// 发送手机验证码
type AuthCodeRequest struct {
	Phone      string      `protobuf:"bytes,1,opt,name=Phone" json:"Phone,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,2,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *AuthCodeRequest) Reset()                    { *m = AuthCodeRequest{} }
func (m *AuthCodeRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthCodeRequest) ProtoMessage()               {}
func (*AuthCodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AuthCodeRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AuthCodeRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

type Reply struct {
	Result bool `protobuf:"varint,1,opt,name=Result" json:"Result,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Reply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// 自动登陆
type AutoLoginRequest struct {
	AutoCode   string      `protobuf:"bytes,1,opt,name=AutoCode" json:"AutoCode,omitempty"`
	UUID       string      `protobuf:"bytes,2,opt,name=UUID" json:"UUID,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,3,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *AutoLoginRequest) Reset()                    { *m = AutoLoginRequest{} }
func (m *AutoLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*AutoLoginRequest) ProtoMessage()               {}
func (*AutoLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AutoLoginRequest) GetAutoCode() string {
	if m != nil {
		return m.AutoCode
	}
	return ""
}

func (m *AutoLoginRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *AutoLoginRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// 账号绑定手机号码
type BindingPhoneRequest struct {
	Phone      string      `protobuf:"bytes,1,opt,name=Phone" json:"Phone,omitempty"`
	UUID       string      `protobuf:"bytes,2,opt,name=UUID" json:"UUID,omitempty"`
	AuthCode   string      `protobuf:"bytes,3,opt,name=AuthCode" json:"AuthCode,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,4,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *BindingPhoneRequest) Reset()                    { *m = BindingPhoneRequest{} }
func (m *BindingPhoneRequest) String() string            { return proto.CompactTextString(m) }
func (*BindingPhoneRequest) ProtoMessage()               {}
func (*BindingPhoneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BindingPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *BindingPhoneRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *BindingPhoneRequest) GetAuthCode() string {
	if m != nil {
		return m.AuthCode
	}
	return ""
}

func (m *BindingPhoneRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// 绑定普通账号
type BindingAccountRequest struct {
	AutoCode   string      `protobuf:"bytes,1,opt,name=AutoCode" json:"AutoCode,omitempty"`
	UserName   string      `protobuf:"bytes,2,opt,name=UserName" json:"UserName,omitempty"`
	PassWord   string      `protobuf:"bytes,3,opt,name=PassWord" json:"PassWord,omitempty"`
	UUID       string      `protobuf:"bytes,4,opt,name=UUID" json:"UUID,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,5,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *BindingAccountRequest) Reset()                    { *m = BindingAccountRequest{} }
func (m *BindingAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*BindingAccountRequest) ProtoMessage()               {}
func (*BindingAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BindingAccountRequest) GetAutoCode() string {
	if m != nil {
		return m.AutoCode
	}
	return ""
}

func (m *BindingAccountRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *BindingAccountRequest) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *BindingAccountRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *BindingAccountRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// 修改密码
type ChangePassWordRequest struct {
	PassWord   string      `protobuf:"bytes,1,opt,name=PassWord" json:"PassWord,omitempty"`
	UUID       string      `protobuf:"bytes,2,opt,name=UUID" json:"UUID,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,3,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *ChangePassWordRequest) Reset()                    { *m = ChangePassWordRequest{} }
func (m *ChangePassWordRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangePassWordRequest) ProtoMessage()               {}
func (*ChangePassWordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ChangePassWordRequest) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *ChangePassWordRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *ChangePassWordRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// 微信登陆
type WeixinLoginRequest struct {
	Code       string      `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,2,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *WeixinLoginRequest) Reset()                    { *m = WeixinLoginRequest{} }
func (m *WeixinLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*WeixinLoginRequest) ProtoMessage()               {}
func (*WeixinLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *WeixinLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *WeixinLoginRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// 心跳
type HeartBeatMsg struct {
	DeviceId string `protobuf:"bytes,1,opt,name=DeviceId" json:"DeviceId,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *HeartBeatMsg) Reset()                    { *m = HeartBeatMsg{} }
func (m *HeartBeatMsg) String() string            { return proto.CompactTextString(m) }
func (*HeartBeatMsg) ProtoMessage()               {}
func (*HeartBeatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *HeartBeatMsg) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *HeartBeatMsg) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "SDKProto.CertificateRequest")
	proto.RegisterType((*CertificateReply)(nil), "SDKProto.CertificateReply")
	proto.RegisterType((*DeviceInfo)(nil), "SDKProto.DeviceInfo")
	proto.RegisterType((*RegisterRequest)(nil), "SDKProto.RegisterRequest")
	proto.RegisterType((*AccountLoginRequest)(nil), "SDKProto.AccountLoginRequest")
	proto.RegisterType((*PhoneLoginRequest)(nil), "SDKProto.PhoneLoginRequest")
	proto.RegisterType((*LoginReply)(nil), "SDKProto.LoginReply")
	proto.RegisterType((*AuthCodeRequest)(nil), "SDKProto.AuthCodeRequest")
	proto.RegisterType((*Reply)(nil), "SDKProto.Reply")
	proto.RegisterType((*AutoLoginRequest)(nil), "SDKProto.AutoLoginRequest")
	proto.RegisterType((*BindingPhoneRequest)(nil), "SDKProto.BindingPhoneRequest")
	proto.RegisterType((*BindingAccountRequest)(nil), "SDKProto.BindingAccountRequest")
	proto.RegisterType((*ChangePassWordRequest)(nil), "SDKProto.ChangePassWordRequest")
	proto.RegisterType((*WeixinLoginRequest)(nil), "SDKProto.WeixinLoginRequest")
	proto.RegisterType((*HeartBeatMsg)(nil), "SDKProto.HeartBeatMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Login service

type LoginClient interface {
	Active(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Reply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*LoginReply, error)
	LoginByAccount(ctx context.Context, in *AccountLoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	LoginByPhone(ctx context.Context, in *PhoneLoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	AutoLogin(ctx context.Context, in *AutoLoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	ActhCode(ctx context.Context, in *AuthCodeRequest, opts ...grpc.CallOption) (*Reply, error)
	BindingPhone(ctx context.Context, in *BindingPhoneRequest, opts ...grpc.CallOption) (*LoginReply, error)
	ChangePassword(ctx context.Context, in *ChangePassWordRequest, opts ...grpc.CallOption) (*Reply, error)
	LoginByWeiXin(ctx context.Context, in *WeixinLoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	CertificateID(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateReply, error)
	HeartBeat(ctx context.Context, in *HeartBeatMsg, opts ...grpc.CallOption) (*Reply, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Active(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/Active", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginByAccount(ctx context.Context, in *AccountLoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/LoginByAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginByPhone(ctx context.Context, in *PhoneLoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/LoginByPhone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) AutoLogin(ctx context.Context, in *AutoLoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/AutoLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) ActhCode(ctx context.Context, in *AuthCodeRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/ActhCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) BindingPhone(ctx context.Context, in *BindingPhoneRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/BindingPhone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) ChangePassword(ctx context.Context, in *ChangePassWordRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/ChangePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginByWeiXin(ctx context.Context, in *WeixinLoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/LoginByWeiXin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) CertificateID(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateReply, error) {
	out := new(CertificateReply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/CertificateID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) HeartBeat(ctx context.Context, in *HeartBeatMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/SDKProto.Login/HeartBeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginServer interface {
	Active(context.Context, *DeviceInfo) (*Reply, error)
	Register(context.Context, *RegisterRequest) (*LoginReply, error)
	LoginByAccount(context.Context, *AccountLoginRequest) (*LoginReply, error)
	LoginByPhone(context.Context, *PhoneLoginRequest) (*LoginReply, error)
	AutoLogin(context.Context, *AutoLoginRequest) (*LoginReply, error)
	ActhCode(context.Context, *AuthCodeRequest) (*Reply, error)
	BindingPhone(context.Context, *BindingPhoneRequest) (*LoginReply, error)
	ChangePassword(context.Context, *ChangePassWordRequest) (*Reply, error)
	LoginByWeiXin(context.Context, *WeixinLoginRequest) (*LoginReply, error)
	CertificateID(context.Context, *CertificateRequest) (*CertificateReply, error)
	HeartBeat(context.Context, *HeartBeatMsg) (*Reply, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Active_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Active(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/Active",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Active(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/LoginByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginByAccount(ctx, req.(*AccountLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/LoginByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginByPhone(ctx, req.(*PhoneLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_AutoLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutoLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).AutoLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/AutoLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).AutoLogin(ctx, req.(*AutoLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_ActhCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).ActhCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/ActhCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).ActhCode(ctx, req.(*AuthCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_BindingPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindingPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).BindingPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/BindingPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).BindingPhone(ctx, req.(*BindingPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePassWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).ChangePassword(ctx, req.(*ChangePassWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginByWeiXin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeixinLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginByWeiXin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/LoginByWeiXin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginByWeiXin(ctx, req.(*WeixinLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_CertificateID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).CertificateID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/CertificateID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).CertificateID(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SDKProto.Login/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).HeartBeat(ctx, req.(*HeartBeatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SDKProto.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Active",
			Handler:    _Login_Active_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Login_Register_Handler,
		},
		{
			MethodName: "LoginByAccount",
			Handler:    _Login_LoginByAccount_Handler,
		},
		{
			MethodName: "LoginByPhone",
			Handler:    _Login_LoginByPhone_Handler,
		},
		{
			MethodName: "AutoLogin",
			Handler:    _Login_AutoLogin_Handler,
		},
		{
			MethodName: "ActhCode",
			Handler:    _Login_ActhCode_Handler,
		},
		{
			MethodName: "BindingPhone",
			Handler:    _Login_BindingPhone_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Login_ChangePassword_Handler,
		},
		{
			MethodName: "LoginByWeiXin",
			Handler:    _Login_LoginByWeiXin_Handler,
		},
		{
			MethodName: "CertificateID",
			Handler:    _Login_CertificateID_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _Login_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LoginProto.proto",
}

func init() { proto.RegisterFile("LoginProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0xc6, 0x21, 0x09, 0xce, 0xbc, 0x01, 0xf2, 0x2e, 0x1f, 0x32, 0x2e, 0x2d, 0xc8, 0xea, 0x81,
	0x53, 0xa4, 0x52, 0x84, 0x7a, 0xa9, 0xaa, 0x7c, 0x20, 0x1a, 0x01, 0x25, 0x32, 0x20, 0x7a, 0x69,
	0x25, 0xd7, 0x59, 0x8c, 0x05, 0xb1, 0x69, 0xbc, 0x49, 0xe1, 0x50, 0xf5, 0xd0, 0x6b, 0x2f, 0xfd,
	0x31, 0xfd, 0x35, 0xfd, 0x1f, 0x3d, 0x57, 0x3b, 0x5e, 0xdb, 0xeb, 0xc4, 0xa6, 0x11, 0x3d, 0xf4,
	0x12, 0xed, 0xcc, 0x4e, 0x66, 0x9e, 0x67, 0xf6, 0x99, 0x91, 0xa1, 0x76, 0xe8, 0x3b, 0xae, 0xd7,
	0x1d, 0xf8, 0xcc, 0xaf, 0xdf, 0xf0, 0x5f, 0xa2, 0x9e, 0xb4, 0x0f, 0xd0, 0x36, 0x46, 0x40, 0x5a,
	0x74, 0xc0, 0xdc, 0x0b, 0xd7, 0xb6, 0x18, 0x35, 0xe9, 0xc7, 0x21, 0x0d, 0x18, 0x21, 0x50, 0x3c,
	0x3b, 0xeb, 0xb4, 0x35, 0x65, 0x53, 0xd9, 0xaa, 0x98, 0x78, 0xe6, 0xbe, 0xd3, 0xbb, 0x1b, 0xaa,
	0x15, 0x36, 0x95, 0xad, 0x92, 0x89, 0x67, 0xee, 0x7b, 0x63, 0xf5, 0xa9, 0x36, 0x1b, 0xc6, 0xf1,
	0x33, 0x79, 0x0a, 0xf3, 0x52, 0xc6, 0x4e, 0x5b, 0x2b, 0xe2, 0x65, 0xda, 0x69, 0xb4, 0xa1, 0x96,
	0xaa, 0x7b, 0x73, 0x7d, 0x47, 0x56, 0xa1, 0x6c, 0xd2, 0x60, 0x78, 0xcd, 0xb0, 0xae, 0x6a, 0x0a,
	0x8b, 0x68, 0x30, 0xd7, 0xf2, 0xfb, 0x7d, 0xea, 0x31, 0x2c, 0x5e, 0x31, 0x23, 0xd3, 0xf8, 0xa5,
	0x00, 0xb4, 0xe9, 0xc8, 0xb5, 0x69, 0xc7, 0xbb, 0xf0, 0x79, 0x82, 0x7d, 0xab, 0x4f, 0x63, 0xe0,
	0xc2, 0xc2, 0x04, 0x97, 0x96, 0xe7, 0xd1, 0xeb, 0x38, 0x41, 0x68, 0x92, 0x27, 0xd1, 0xff, 0x91,
	0x5a, 0x48, 0x43, 0xf2, 0x70, 0x82, 0x9d, 0xa3, 0xbd, 0x8e, 0xe0, 0x80, 0x67, 0x52, 0x83, 0xd9,
	0xa3, 0x46, 0x4b, 0x2b, 0xa1, 0x8b, 0x1f, 0x31, 0xaa, 0x3b, 0xda, 0xd1, 0xca, 0x22, 0xaa, 0x3b,
	0xda, 0x11, 0xbe, 0x5d, 0x6d, 0x2e, 0xf6, 0xed, 0x92, 0x05, 0x28, 0x1c, 0x9f, 0x68, 0x2a, 0x7a,
	0x0a, 0xc7, 0x27, 0x44, 0x07, 0xf5, 0xd0, 0x62, 0x2e, 0x1b, 0xf6, 0xa8, 0x56, 0x41, 0x6f, 0x6c,
	0x93, 0x75, 0xa8, 0x1c, 0xfa, 0x9e, 0x13, 0x5e, 0x02, 0x5e, 0x26, 0x0e, 0xe3, 0x0b, 0x2c, 0x9a,
	0xd4, 0x71, 0x03, 0x46, 0x07, 0xd1, 0x9b, 0xe9, 0xa0, 0x9e, 0x05, 0x74, 0x80, 0xef, 0x11, 0xd2,
	0x8f, 0x6d, 0x7e, 0xd7, 0xb5, 0x82, 0xe0, 0xdc, 0x1f, 0xf4, 0x44, 0x07, 0x62, 0x9b, 0xec, 0xc8,
	0x2d, 0xc4, 0x16, 0xfc, 0xb7, 0xbd, 0x5c, 0x8f, 0x04, 0x52, 0x4f, 0xee, 0x4c, 0x29, 0xce, 0xf8,
	0xaa, 0xc0, 0x52, 0xc3, 0xb6, 0xfd, 0xa1, 0xc7, 0x50, 0x5d, 0xff, 0x06, 0x45, 0x00, 0xff, 0x77,
	0x2f, 0x7d, 0x8f, 0xa6, 0x20, 0x2c, 0x43, 0x09, 0x9d, 0xa2, 0x7e, 0x68, 0xf0, 0xf7, 0x68, 0xf9,
	0x3d, 0x2a, 0x0a, 0xe3, 0xf9, 0x81, 0x45, 0xbf, 0x29, 0x00, 0xa2, 0x20, 0x57, 0x6d, 0xd6, 0xac,
	0xe8, 0xa0, 0x36, 0x86, 0xcc, 0x97, 0x0a, 0xc6, 0x36, 0x87, 0x77, 0xea, 0x5f, 0x51, 0x4f, 0xa8,
	0x2d, 0x34, 0x12, 0xd0, 0x45, 0x19, 0xb4, 0x01, 0x55, 0x69, 0x4a, 0x7a, 0xa8, 0x39, 0xd5, 0x4c,
	0xf9, 0x8c, 0x77, 0xb0, 0xd8, 0x18, 0xb2, 0x4b, 0x9e, 0xfb, 0xfe, 0x0e, 0xa4, 0xd9, 0x16, 0xa6,
	0x64, 0xbb, 0x01, 0xa5, 0x7b, 0xa7, 0xd3, 0xb8, 0x85, 0x1a, 0xe7, 0x36, 0xae, 0x82, 0x98, 0xbf,
	0x32, 0xc6, 0x3f, 0xea, 0x57, 0x41, 0xea, 0xd7, 0xc3, 0x1e, 0xe2, 0xbb, 0x02, 0x4b, 0x4d, 0xd7,
	0xeb, 0xb9, 0x9e, 0x83, 0x0c, 0xff, 0x28, 0x80, 0x89, 0xba, 0x21, 0x4e, 0xec, 0x9d, 0x78, 0x8e,
	0xd8, 0x1e, 0xc3, 0x54, 0x9c, 0x12, 0xd3, 0x0f, 0x05, 0x56, 0x04, 0x26, 0x31, 0x1e, 0xd3, 0xf4,
	0x44, 0x9e, 0x9a, 0xc2, 0x3d, 0x53, 0x33, 0x3b, 0x36, 0x35, 0x11, 0xa7, 0x62, 0x6e, 0x2f, 0x4b,
	0x53, 0xe2, 0xfe, 0x0c, 0x2b, 0x7c, 0x27, 0x3a, 0x34, 0xca, 0x2d, 0xc1, 0x8e, 0xcb, 0x2b, 0x39,
	0xe5, 0xff, 0xfe, 0x29, 0xdf, 0x03, 0x39, 0xa7, 0xee, 0xad, 0xeb, 0xa5, 0x64, 0x44, 0xa0, 0x68,
	0x27, 0xed, 0xc2, 0xf3, 0x03, 0x55, 0xdc, 0x84, 0xea, 0x6b, 0x6a, 0x0d, 0x58, 0x93, 0x5a, 0xec,
	0x28, 0x70, 0x38, 0x2b, 0x71, 0x1b, 0xb3, 0x8a, 0x6c, 0x2e, 0x74, 0xde, 0xfc, 0x4e, 0xb4, 0xa4,
	0x84, 0xb5, 0xfd, 0xb3, 0x04, 0x25, 0x84, 0x47, 0x9e, 0x41, 0xb9, 0x61, 0x33, 0x77, 0x44, 0x49,
	0x66, 0x65, 0x7d, 0x31, 0xf1, 0xe2, 0xec, 0x18, 0x33, 0xe4, 0x25, 0xa8, 0xd1, 0xc2, 0x26, 0x6b,
	0xf2, 0x75, 0x6a, 0x89, 0xeb, 0x52, 0xbe, 0x64, 0xc5, 0x18, 0x33, 0x64, 0x1f, 0x16, 0xd0, 0x6e,
	0xde, 0x09, 0x55, 0x91, 0xc7, 0x49, 0x64, 0xc6, 0x1e, 0xce, 0x4d, 0xd4, 0x82, 0xaa, 0x48, 0x14,
	0x4e, 0xc5, 0xa3, 0x24, 0x6e, 0x62, 0x93, 0xe6, 0x26, 0x79, 0x05, 0x95, 0x78, 0xe4, 0x89, 0x2e,
	0x01, 0x19, 0xdb, 0x03, 0xb9, 0x09, 0x5e, 0x80, 0xda, 0xb0, 0xc5, 0x9c, 0xad, 0xa5, 0xfe, 0x2f,
	0xef, 0xb1, 0xac, 0x3e, 0xee, 0x41, 0x55, 0x1e, 0x79, 0xb9, 0x0d, 0x19, 0xab, 0x20, 0x17, 0x40,
	0x1b, 0x16, 0x12, 0xb9, 0x7f, 0xe2, 0x5a, 0xde, 0x48, 0x22, 0x33, 0x07, 0x21, 0x1b, 0xcc, 0xbc,
	0x68, 0xe6, 0x39, 0x75, 0xdf, 0xba, 0x1e, 0x59, 0x4f, 0x62, 0x26, 0xe5, 0x9c, 0x0b, 0xe6, 0x60,
	0xec, 0x8b, 0x49, 0x4e, 0x33, 0xf9, 0x71, 0xa6, 0xeb, 0x39, 0xb7, 0x61, 0xb2, 0x5d, 0xa8, 0xc4,
	0x4a, 0x27, 0xab, 0x49, 0xa8, 0x2c, 0xff, 0x0c, 0x2e, 0xcd, 0x2a, 0x80, 0xed, 0xf7, 0xeb, 0xde,
	0x95, 0x63, 0xf5, 0xe9, 0x87, 0x32, 0x7e, 0x27, 0x3e, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x82,
	0x1e, 0x23, 0xad, 0x3b, 0x0a, 0x00, 0x00,
}
