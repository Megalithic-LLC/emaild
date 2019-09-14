// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agentstream.proto

package agentstreamproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ClientMessage struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to MessageType:
	//	*ClientMessage_AckResponse
	//	*ClientMessage_NoopRequest
	//	*ClientMessage_StartupRequest
	//	*ClientMessage_ErrorResponse
	MessageType          isClientMessage_MessageType `protobuf_oneof:"message_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ClientMessage) Reset()         { *m = ClientMessage{} }
func (m *ClientMessage) String() string { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()    {}
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{0}
}

func (m *ClientMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientMessage.Unmarshal(m, b)
}
func (m *ClientMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientMessage.Marshal(b, m, deterministic)
}
func (m *ClientMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientMessage.Merge(m, src)
}
func (m *ClientMessage) XXX_Size() int {
	return xxx_messageInfo_ClientMessage.Size(m)
}
func (m *ClientMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClientMessage proto.InternalMessageInfo

func (m *ClientMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isClientMessage_MessageType interface {
	isClientMessage_MessageType()
}

type ClientMessage_AckResponse struct {
	AckResponse *AckResponse `protobuf:"bytes,2,opt,name=ack_response,json=ackResponse,proto3,oneof"`
}

type ClientMessage_NoopRequest struct {
	NoopRequest *NoopRequest `protobuf:"bytes,3,opt,name=noop_request,json=noopRequest,proto3,oneof"`
}

type ClientMessage_StartupRequest struct {
	StartupRequest *StartupRequest `protobuf:"bytes,4,opt,name=startup_request,json=startupRequest,proto3,oneof"`
}

type ClientMessage_ErrorResponse struct {
	ErrorResponse *ErrorResponse `protobuf:"bytes,5,opt,name=error_response,json=errorResponse,proto3,oneof"`
}

func (*ClientMessage_AckResponse) isClientMessage_MessageType() {}

func (*ClientMessage_NoopRequest) isClientMessage_MessageType() {}

func (*ClientMessage_StartupRequest) isClientMessage_MessageType() {}

func (*ClientMessage_ErrorResponse) isClientMessage_MessageType() {}

func (m *ClientMessage) GetMessageType() isClientMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *ClientMessage) GetAckResponse() *AckResponse {
	if x, ok := m.GetMessageType().(*ClientMessage_AckResponse); ok {
		return x.AckResponse
	}
	return nil
}

func (m *ClientMessage) GetNoopRequest() *NoopRequest {
	if x, ok := m.GetMessageType().(*ClientMessage_NoopRequest); ok {
		return x.NoopRequest
	}
	return nil
}

func (m *ClientMessage) GetStartupRequest() *StartupRequest {
	if x, ok := m.GetMessageType().(*ClientMessage_StartupRequest); ok {
		return x.StartupRequest
	}
	return nil
}

func (m *ClientMessage) GetErrorResponse() *ErrorResponse {
	if x, ok := m.GetMessageType().(*ClientMessage_ErrorResponse); ok {
		return x.ErrorResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClientMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClientMessage_AckResponse)(nil),
		(*ClientMessage_NoopRequest)(nil),
		(*ClientMessage_StartupRequest)(nil),
		(*ClientMessage_ErrorResponse)(nil),
	}
}

type ServerMessage struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to MessageType:
	//	*ServerMessage_AckResponse
	//	*ServerMessage_NoopRequest
	//	*ServerMessage_ErrorResponse
	//	*ServerMessage_ClaimRequest
	//	*ServerMessage_ConfigChangedRequest
	//	*ServerMessage_EmailcdnTestAccountCredentialsRequest
	MessageType          isServerMessage_MessageType `protobuf_oneof:"message_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ServerMessage) Reset()         { *m = ServerMessage{} }
func (m *ServerMessage) String() string { return proto.CompactTextString(m) }
func (*ServerMessage) ProtoMessage()    {}
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{1}
}

func (m *ServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMessage.Unmarshal(m, b)
}
func (m *ServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMessage.Marshal(b, m, deterministic)
}
func (m *ServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMessage.Merge(m, src)
}
func (m *ServerMessage) XXX_Size() int {
	return xxx_messageInfo_ServerMessage.Size(m)
}
func (m *ServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMessage proto.InternalMessageInfo

func (m *ServerMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isServerMessage_MessageType interface {
	isServerMessage_MessageType()
}

type ServerMessage_AckResponse struct {
	AckResponse *AckResponse `protobuf:"bytes,2,opt,name=ack_response,json=ackResponse,proto3,oneof"`
}

type ServerMessage_NoopRequest struct {
	NoopRequest *NoopRequest `protobuf:"bytes,3,opt,name=noop_request,json=noopRequest,proto3,oneof"`
}

type ServerMessage_ErrorResponse struct {
	ErrorResponse *ErrorResponse `protobuf:"bytes,4,opt,name=error_response,json=errorResponse,proto3,oneof"`
}

type ServerMessage_ClaimRequest struct {
	ClaimRequest *ClaimRequest `protobuf:"bytes,5,opt,name=claim_request,json=claimRequest,proto3,oneof"`
}

type ServerMessage_ConfigChangedRequest struct {
	ConfigChangedRequest *ConfigChangedRequest `protobuf:"bytes,6,opt,name=config_changed_request,json=configChangedRequest,proto3,oneof"`
}

type ServerMessage_EmailcdnTestAccountCredentialsRequest struct {
	EmailcdnTestAccountCredentialsRequest *EmailcdnTestAccountCredentialsRequest `protobuf:"bytes,7,opt,name=emailcdn_test_account_credentials_request,json=emailcdnTestAccountCredentialsRequest,proto3,oneof"`
}

func (*ServerMessage_AckResponse) isServerMessage_MessageType() {}

func (*ServerMessage_NoopRequest) isServerMessage_MessageType() {}

func (*ServerMessage_ErrorResponse) isServerMessage_MessageType() {}

func (*ServerMessage_ClaimRequest) isServerMessage_MessageType() {}

func (*ServerMessage_ConfigChangedRequest) isServerMessage_MessageType() {}

func (*ServerMessage_EmailcdnTestAccountCredentialsRequest) isServerMessage_MessageType() {}

func (m *ServerMessage) GetMessageType() isServerMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *ServerMessage) GetAckResponse() *AckResponse {
	if x, ok := m.GetMessageType().(*ServerMessage_AckResponse); ok {
		return x.AckResponse
	}
	return nil
}

func (m *ServerMessage) GetNoopRequest() *NoopRequest {
	if x, ok := m.GetMessageType().(*ServerMessage_NoopRequest); ok {
		return x.NoopRequest
	}
	return nil
}

func (m *ServerMessage) GetErrorResponse() *ErrorResponse {
	if x, ok := m.GetMessageType().(*ServerMessage_ErrorResponse); ok {
		return x.ErrorResponse
	}
	return nil
}

func (m *ServerMessage) GetClaimRequest() *ClaimRequest {
	if x, ok := m.GetMessageType().(*ServerMessage_ClaimRequest); ok {
		return x.ClaimRequest
	}
	return nil
}

func (m *ServerMessage) GetConfigChangedRequest() *ConfigChangedRequest {
	if x, ok := m.GetMessageType().(*ServerMessage_ConfigChangedRequest); ok {
		return x.ConfigChangedRequest
	}
	return nil
}

func (m *ServerMessage) GetEmailcdnTestAccountCredentialsRequest() *EmailcdnTestAccountCredentialsRequest {
	if x, ok := m.GetMessageType().(*ServerMessage_EmailcdnTestAccountCredentialsRequest); ok {
		return x.EmailcdnTestAccountCredentialsRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerMessage_AckResponse)(nil),
		(*ServerMessage_NoopRequest)(nil),
		(*ServerMessage_ErrorResponse)(nil),
		(*ServerMessage_ClaimRequest)(nil),
		(*ServerMessage_ConfigChangedRequest)(nil),
		(*ServerMessage_EmailcdnTestAccountCredentialsRequest)(nil),
	}
}

type AckResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckResponse) Reset()         { *m = AckResponse{} }
func (m *AckResponse) String() string { return proto.CompactTextString(m) }
func (*AckResponse) ProtoMessage()    {}
func (*AckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{2}
}

func (m *AckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckResponse.Unmarshal(m, b)
}
func (m *AckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckResponse.Marshal(b, m, deterministic)
}
func (m *AckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckResponse.Merge(m, src)
}
func (m *AckResponse) XXX_Size() int {
	return xxx_messageInfo_AckResponse.Size(m)
}
func (m *AckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AckResponse proto.InternalMessageInfo

type ClaimRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClaimRequest) Reset()         { *m = ClaimRequest{} }
func (m *ClaimRequest) String() string { return proto.CompactTextString(m) }
func (*ClaimRequest) ProtoMessage()    {}
func (*ClaimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{3}
}

func (m *ClaimRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimRequest.Unmarshal(m, b)
}
func (m *ClaimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimRequest.Marshal(b, m, deterministic)
}
func (m *ClaimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRequest.Merge(m, src)
}
func (m *ClaimRequest) XXX_Size() int {
	return xxx_messageInfo_ClaimRequest.Size(m)
}
func (m *ClaimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRequest proto.InternalMessageInfo

func (m *ClaimRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ConfigChangedRequest struct {
	HashesByTable        map[string][]byte `protobuf:"bytes,1,rep,name=hashes_by_table,json=hashesByTable,proto3" json:"hashes_by_table,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConfigChangedRequest) Reset()         { *m = ConfigChangedRequest{} }
func (m *ConfigChangedRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigChangedRequest) ProtoMessage()    {}
func (*ConfigChangedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{4}
}

func (m *ConfigChangedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigChangedRequest.Unmarshal(m, b)
}
func (m *ConfigChangedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigChangedRequest.Marshal(b, m, deterministic)
}
func (m *ConfigChangedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigChangedRequest.Merge(m, src)
}
func (m *ConfigChangedRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigChangedRequest.Size(m)
}
func (m *ConfigChangedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigChangedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigChangedRequest proto.InternalMessageInfo

func (m *ConfigChangedRequest) GetHashesByTable() map[string][]byte {
	if m != nil {
		return m.HashesByTable
	}
	return nil
}

type EmailcdnAccount struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Provider             string   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	ImapHost             string   `protobuf:"bytes,4,opt,name=imap_host,json=imapHost,proto3" json:"imap_host,omitempty"`
	ImapPort             uint32   `protobuf:"varint,5,opt,name=imap_port,json=imapPort,proto3" json:"imap_port,omitempty"`
	ImapUsername         string   `protobuf:"bytes,6,opt,name=imap_username,json=imapUsername,proto3" json:"imap_username,omitempty"`
	ImapPassword         string   `protobuf:"bytes,7,opt,name=imap_password,json=imapPassword,proto3" json:"imap_password,omitempty"`
	SmtpHost             string   `protobuf:"bytes,8,opt,name=smtp_host,json=smtpHost,proto3" json:"smtp_host,omitempty"`
	SmtpPort             uint32   `protobuf:"varint,9,opt,name=smtp_port,json=smtpPort,proto3" json:"smtp_port,omitempty"`
	SmtpUsername         string   `protobuf:"bytes,10,opt,name=smtp_username,json=smtpUsername,proto3" json:"smtp_username,omitempty"`
	SmtpPassword         string   `protobuf:"bytes,11,opt,name=smtp_password,json=smtpPassword,proto3" json:"smtp_password,omitempty"`
	SslRequired          bool     `protobuf:"varint,12,opt,name=ssl_required,json=sslRequired,proto3" json:"ssl_required,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailcdnAccount) Reset()         { *m = EmailcdnAccount{} }
func (m *EmailcdnAccount) String() string { return proto.CompactTextString(m) }
func (*EmailcdnAccount) ProtoMessage()    {}
func (*EmailcdnAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{5}
}

func (m *EmailcdnAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailcdnAccount.Unmarshal(m, b)
}
func (m *EmailcdnAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailcdnAccount.Marshal(b, m, deterministic)
}
func (m *EmailcdnAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailcdnAccount.Merge(m, src)
}
func (m *EmailcdnAccount) XXX_Size() int {
	return xxx_messageInfo_EmailcdnAccount.Size(m)
}
func (m *EmailcdnAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailcdnAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EmailcdnAccount proto.InternalMessageInfo

func (m *EmailcdnAccount) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EmailcdnAccount) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *EmailcdnAccount) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailcdnAccount) GetImapHost() string {
	if m != nil {
		return m.ImapHost
	}
	return ""
}

func (m *EmailcdnAccount) GetImapPort() uint32 {
	if m != nil {
		return m.ImapPort
	}
	return 0
}

func (m *EmailcdnAccount) GetImapUsername() string {
	if m != nil {
		return m.ImapUsername
	}
	return ""
}

func (m *EmailcdnAccount) GetImapPassword() string {
	if m != nil {
		return m.ImapPassword
	}
	return ""
}

func (m *EmailcdnAccount) GetSmtpHost() string {
	if m != nil {
		return m.SmtpHost
	}
	return ""
}

func (m *EmailcdnAccount) GetSmtpPort() uint32 {
	if m != nil {
		return m.SmtpPort
	}
	return 0
}

func (m *EmailcdnAccount) GetSmtpUsername() string {
	if m != nil {
		return m.SmtpUsername
	}
	return ""
}

func (m *EmailcdnAccount) GetSmtpPassword() string {
	if m != nil {
		return m.SmtpPassword
	}
	return ""
}

func (m *EmailcdnAccount) GetSslRequired() bool {
	if m != nil {
		return m.SslRequired
	}
	return false
}

type EmailcdnTestAccountCredentialsRequest struct {
	Account              *EmailcdnAccount `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EmailcdnTestAccountCredentialsRequest) Reset()         { *m = EmailcdnTestAccountCredentialsRequest{} }
func (m *EmailcdnTestAccountCredentialsRequest) String() string { return proto.CompactTextString(m) }
func (*EmailcdnTestAccountCredentialsRequest) ProtoMessage()    {}
func (*EmailcdnTestAccountCredentialsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{6}
}

func (m *EmailcdnTestAccountCredentialsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailcdnTestAccountCredentialsRequest.Unmarshal(m, b)
}
func (m *EmailcdnTestAccountCredentialsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailcdnTestAccountCredentialsRequest.Marshal(b, m, deterministic)
}
func (m *EmailcdnTestAccountCredentialsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailcdnTestAccountCredentialsRequest.Merge(m, src)
}
func (m *EmailcdnTestAccountCredentialsRequest) XXX_Size() int {
	return xxx_messageInfo_EmailcdnTestAccountCredentialsRequest.Size(m)
}
func (m *EmailcdnTestAccountCredentialsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailcdnTestAccountCredentialsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailcdnTestAccountCredentialsRequest proto.InternalMessageInfo

func (m *EmailcdnTestAccountCredentialsRequest) GetAccount() *EmailcdnAccount {
	if m != nil {
		return m.Account
	}
	return nil
}

type ErrorResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{7}
}

func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(m, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type NoopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoopRequest) Reset()         { *m = NoopRequest{} }
func (m *NoopRequest) String() string { return proto.CompactTextString(m) }
func (*NoopRequest) ProtoMessage()    {}
func (*NoopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{8}
}

func (m *NoopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoopRequest.Unmarshal(m, b)
}
func (m *NoopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoopRequest.Marshal(b, m, deterministic)
}
func (m *NoopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoopRequest.Merge(m, src)
}
func (m *NoopRequest) XXX_Size() int {
	return xxx_messageInfo_NoopRequest.Size(m)
}
func (m *NoopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NoopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NoopRequest proto.InternalMessageInfo

type StartupRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartupRequest) Reset()         { *m = StartupRequest{} }
func (m *StartupRequest) String() string { return proto.CompactTextString(m) }
func (*StartupRequest) ProtoMessage()    {}
func (*StartupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{9}
}

func (m *StartupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartupRequest.Unmarshal(m, b)
}
func (m *StartupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartupRequest.Marshal(b, m, deterministic)
}
func (m *StartupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartupRequest.Merge(m, src)
}
func (m *StartupRequest) XXX_Size() int {
	return xxx_messageInfo_StartupRequest.Size(m)
}
func (m *StartupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartupRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientMessage)(nil), "agentstreamproto.ClientMessage")
	proto.RegisterType((*ServerMessage)(nil), "agentstreamproto.ServerMessage")
	proto.RegisterType((*AckResponse)(nil), "agentstreamproto.AckResponse")
	proto.RegisterType((*ClaimRequest)(nil), "agentstreamproto.ClaimRequest")
	proto.RegisterType((*ConfigChangedRequest)(nil), "agentstreamproto.ConfigChangedRequest")
	proto.RegisterMapType((map[string][]byte)(nil), "agentstreamproto.ConfigChangedRequest.HashesByTableEntry")
	proto.RegisterType((*EmailcdnAccount)(nil), "agentstreamproto.EmailcdnAccount")
	proto.RegisterType((*EmailcdnTestAccountCredentialsRequest)(nil), "agentstreamproto.EmailcdnTestAccountCredentialsRequest")
	proto.RegisterType((*ErrorResponse)(nil), "agentstreamproto.ErrorResponse")
	proto.RegisterType((*NoopRequest)(nil), "agentstreamproto.NoopRequest")
	proto.RegisterType((*StartupRequest)(nil), "agentstreamproto.StartupRequest")
}

func init() { proto.RegisterFile("agentstream.proto", fileDescriptor_652ad2f4d378dabd) }

var fileDescriptor_652ad2f4d378dabd = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x6f, 0xdb, 0x38,
	0x10, 0x5d, 0x39, 0x4e, 0x62, 0x8d, 0x25, 0x27, 0x4b, 0x04, 0x0b, 0x23, 0x8b, 0xdd, 0x3a, 0x6a,
	0x53, 0xb8, 0x17, 0x1f, 0xd2, 0x43, 0xbf, 0x2e, 0x4d, 0x8c, 0x00, 0x06, 0x8a, 0x16, 0x01, 0x93,
	0x5e, 0x2b, 0x30, 0xd2, 0xd4, 0x16, 0x2c, 0x89, 0x2a, 0x49, 0xa7, 0xf0, 0xef, 0xe8, 0xcf, 0x29,
	0xd0, 0x9f, 0xd0, 0x9f, 0x54, 0x14, 0x24, 0x65, 0x49, 0xa9, 0x1d, 0x20, 0xe8, 0xa9, 0x37, 0xcd,
	0xe3, 0xe3, 0xe3, 0xd3, 0xe3, 0x0c, 0xe1, 0x6f, 0x36, 0xc5, 0x5c, 0x49, 0x25, 0x90, 0x65, 0xa3,
	0x42, 0x70, 0xc5, 0xc9, 0x7e, 0x03, 0x32, 0x48, 0xf0, 0xbd, 0x05, 0xfe, 0x38, 0x4d, 0x30, 0x57,
	0x6f, 0x51, 0x4a, 0x36, 0x45, 0xd2, 0x83, 0x56, 0x12, 0xf7, 0x9d, 0x81, 0x33, 0x6c, 0xd3, 0x56,
	0x12, 0x93, 0x33, 0xf0, 0x58, 0x34, 0x0f, 0x05, 0xca, 0x82, 0xe7, 0x12, 0xfb, 0xad, 0x81, 0x33,
	0xec, 0x9e, 0xfc, 0x37, 0xfa, 0x55, 0x6a, 0x74, 0x1a, 0xcd, 0x69, 0x49, 0x9a, 0xfc, 0x45, 0xbb,
	0xac, 0x2e, 0xb5, 0x46, 0xce, 0x79, 0x11, 0x0a, 0xfc, 0xb4, 0x40, 0xa9, 0xfa, 0x5b, 0x77, 0x69,
	0xbc, 0xe3, 0xbc, 0xa0, 0x96, 0xa4, 0x35, 0xf2, 0xba, 0x24, 0x6f, 0x60, 0x4f, 0x2a, 0x26, 0xd4,
	0xa2, 0x96, 0x69, 0x1b, 0x99, 0xc1, 0xba, 0xcc, 0xa5, 0x25, 0xd6, 0x4a, 0x3d, 0x79, 0x0b, 0x21,
	0x13, 0xe8, 0xa1, 0x10, 0x5c, 0xd4, 0xbf, 0xb5, 0x6d, 0xb4, 0x1e, 0xac, 0x6b, 0x9d, 0x6b, 0x5e,
	0xe3, 0xc7, 0x7c, 0x6c, 0x02, 0x67, 0x3d, 0xf0, 0x32, 0x9b, 0x5c, 0xa8, 0x96, 0x05, 0x06, 0xdf,
	0xda, 0xe0, 0x5f, 0xa2, 0xb8, 0x41, 0xf1, 0xa7, 0x07, 0xba, 0x9e, 0x41, 0xfb, 0xf7, 0x32, 0x20,
	0xe7, 0xe0, 0x47, 0x29, 0x4b, 0xb2, 0xca, 0x8e, 0x0d, 0xf3, 0xff, 0x75, 0xa1, 0xb1, 0xa6, 0xd5,
	0x7e, 0xbc, 0xa8, 0x51, 0x93, 0x0f, 0xf0, 0x4f, 0xc4, 0xf3, 0x8f, 0xc9, 0x34, 0x8c, 0x66, 0x2c,
	0x9f, 0x62, 0x5c, 0xe9, 0xed, 0x18, 0xbd, 0xc7, 0x1b, 0xf4, 0x0c, 0x7f, 0x6c, 0xe9, 0xb5, 0xee,
	0x41, 0xb4, 0x01, 0x27, 0x5f, 0x1c, 0x78, 0x82, 0x19, 0x4b, 0xd2, 0x28, 0xce, 0x43, 0x85, 0x52,
	0x85, 0x2c, 0x8a, 0xf8, 0x22, 0x57, 0x61, 0x24, 0x30, 0xc6, 0x5c, 0x25, 0x2c, 0x95, 0xd5, 0x99,
	0xbb, 0xe6, 0xcc, 0x67, 0x1b, 0xc2, 0x28, 0x25, 0xae, 0x50, 0xaa, 0x53, 0x2b, 0x30, 0xae, 0xf7,
	0xd7, 0x26, 0x8e, 0xf1, 0x3e, 0xc4, 0xb5, 0x06, 0xf2, 0xa1, 0xdb, 0xb8, 0xf8, 0xe0, 0x11, 0x78,
	0xcd, 0xd0, 0xc8, 0x01, 0x6c, 0x2b, 0x3e, 0xc7, 0xdc, 0x34, 0x94, 0x4b, 0x6d, 0x11, 0x7c, 0x75,
	0xe0, 0x60, 0x53, 0x16, 0x84, 0xc1, 0xde, 0x8c, 0xc9, 0x19, 0xca, 0xf0, 0x7a, 0x19, 0x2a, 0x76,
	0x9d, 0x62, 0xdf, 0x19, 0x6c, 0x0d, 0xbb, 0x27, 0x2f, 0xee, 0x17, 0xe6, 0x68, 0x62, 0x76, 0x9f,
	0x2d, 0xaf, 0xf4, 0xde, 0xf3, 0x5c, 0x89, 0x25, 0xf5, 0x67, 0x4d, 0xec, 0xf0, 0x35, 0x90, 0x75,
	0x12, 0xd9, 0x87, 0xad, 0x39, 0x2e, 0x4b, 0x97, 0xfa, 0x53, 0x3b, 0xbf, 0x61, 0xe9, 0xc2, 0x36,
	0xbc, 0x47, 0x6d, 0xf1, 0xb2, 0xf5, 0xdc, 0x09, 0x7e, 0xb4, 0x60, 0x6f, 0x95, 0x6a, 0x19, 0x54,
	0x63, 0x6a, 0x5c, 0x33, 0x35, 0x87, 0xd0, 0x29, 0x04, 0xbf, 0x49, 0x62, 0x14, 0x46, 0xc0, 0xa5,
	0x55, 0xad, 0x95, 0x4d, 0xd6, 0x66, 0x0c, 0x5c, 0x6a, 0x0b, 0xf2, 0x2f, 0xb8, 0x49, 0xc6, 0x8a,
	0x70, 0xc6, 0xcb, 0xa7, 0xc2, 0xa5, 0x1d, 0x0d, 0x4c, 0xb8, 0x54, 0xd5, 0x62, 0xc1, 0x85, 0x6d,
	0x57, 0xdf, 0x2e, 0x5e, 0x70, 0xa1, 0xc8, 0x43, 0xf0, 0xcd, 0xe2, 0x42, 0xa2, 0xc8, 0x59, 0x86,
	0xa6, 0xff, 0x5c, 0xea, 0x69, 0xf0, 0x7d, 0x89, 0x55, 0xa4, 0x82, 0x49, 0xf9, 0x99, 0x8b, 0xd8,
	0x34, 0x4c, 0x49, 0xba, 0x28, 0x31, 0x7d, 0x8c, 0xcc, 0x54, 0xe9, 0xa1, 0x63, 0x3d, 0x68, 0x60,
	0xe5, 0xc1, 0x2c, 0x1a, 0x0f, 0xae, 0xf5, 0xa0, 0x81, 0x95, 0x07, 0xb3, 0x58, 0x79, 0x00, 0x2b,
	0xaf, 0xc1, 0xa6, 0x07, 0xab, 0xb0, 0xf2, 0xd0, 0xad, 0x49, 0x95, 0x87, 0x23, 0xf0, 0xa4, 0x4c,
	0x4d, 0x5f, 0x27, 0x02, 0xe3, 0xbe, 0x37, 0x70, 0x86, 0x1d, 0xda, 0x95, 0x32, 0xa5, 0x25, 0x14,
	0xc4, 0x70, 0x7c, 0xaf, 0xae, 0x26, 0xaf, 0x60, 0xb7, 0x9c, 0x19, 0x73, 0x35, 0xdd, 0x93, 0xa3,
	0xbb, 0xe7, 0xa3, 0x54, 0xa1, 0xab, 0x1d, 0xc1, 0x31, 0xf8, 0xb7, 0x1e, 0x12, 0x73, 0x6f, 0x1a,
	0x58, 0xf5, 0xb2, 0x29, 0xf4, 0x00, 0x34, 0x5e, 0xad, 0x60, 0x1f, 0x7a, 0xb7, 0x9f, 0xf3, 0xeb,
	0x1d, 0x73, 0xce, 0xd3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0x4e, 0x01, 0xa9, 0xe1, 0x06,
	0x00, 0x00,
}