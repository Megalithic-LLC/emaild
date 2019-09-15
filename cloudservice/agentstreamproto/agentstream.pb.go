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
	//	*ClientMessage_EmailcdnGetAccountsRequest
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

type ClientMessage_EmailcdnGetAccountsRequest struct {
	EmailcdnGetAccountsRequest *EmailcdnGetAccountsRequest `protobuf:"bytes,6,opt,name=emailcdn_get_accounts_request,json=emailcdnGetAccountsRequest,proto3,oneof"`
}

func (*ClientMessage_AckResponse) isClientMessage_MessageType() {}

func (*ClientMessage_NoopRequest) isClientMessage_MessageType() {}

func (*ClientMessage_StartupRequest) isClientMessage_MessageType() {}

func (*ClientMessage_ErrorResponse) isClientMessage_MessageType() {}

func (*ClientMessage_EmailcdnGetAccountsRequest) isClientMessage_MessageType() {}

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

func (m *ClientMessage) GetEmailcdnGetAccountsRequest() *EmailcdnGetAccountsRequest {
	if x, ok := m.GetMessageType().(*ClientMessage_EmailcdnGetAccountsRequest); ok {
		return x.EmailcdnGetAccountsRequest
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
		(*ClientMessage_EmailcdnGetAccountsRequest)(nil),
	}
}

type ServerMessage struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to MessageType:
	//	*ServerMessage_AckResponse
	//	*ServerMessage_NoopRequest
	//	*ServerMessage_ErrorResponse
	//	*ServerMessage_ClaimRequest
	//	*ServerMessage_ConfigChangedResponse
	//	*ServerMessage_EmailcdnTestAccountCredentialsRequest
	//	*ServerMessage_EmailcdnGetAccountsResponse
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

type ServerMessage_ConfigChangedResponse struct {
	ConfigChangedResponse *ConfigChangedResponse `protobuf:"bytes,6,opt,name=config_changed_response,json=configChangedResponse,proto3,oneof"`
}

type ServerMessage_EmailcdnTestAccountCredentialsRequest struct {
	EmailcdnTestAccountCredentialsRequest *EmailcdnTestAccountCredentialsRequest `protobuf:"bytes,7,opt,name=emailcdn_test_account_credentials_request,json=emailcdnTestAccountCredentialsRequest,proto3,oneof"`
}

type ServerMessage_EmailcdnGetAccountsResponse struct {
	EmailcdnGetAccountsResponse *EmailcdnGetAccountsResponse `protobuf:"bytes,8,opt,name=emailcdn_get_accounts_response,json=emailcdnGetAccountsResponse,proto3,oneof"`
}

func (*ServerMessage_AckResponse) isServerMessage_MessageType() {}

func (*ServerMessage_NoopRequest) isServerMessage_MessageType() {}

func (*ServerMessage_ErrorResponse) isServerMessage_MessageType() {}

func (*ServerMessage_ClaimRequest) isServerMessage_MessageType() {}

func (*ServerMessage_ConfigChangedResponse) isServerMessage_MessageType() {}

func (*ServerMessage_EmailcdnTestAccountCredentialsRequest) isServerMessage_MessageType() {}

func (*ServerMessage_EmailcdnGetAccountsResponse) isServerMessage_MessageType() {}

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

func (m *ServerMessage) GetConfigChangedResponse() *ConfigChangedResponse {
	if x, ok := m.GetMessageType().(*ServerMessage_ConfigChangedResponse); ok {
		return x.ConfigChangedResponse
	}
	return nil
}

func (m *ServerMessage) GetEmailcdnTestAccountCredentialsRequest() *EmailcdnTestAccountCredentialsRequest {
	if x, ok := m.GetMessageType().(*ServerMessage_EmailcdnTestAccountCredentialsRequest); ok {
		return x.EmailcdnTestAccountCredentialsRequest
	}
	return nil
}

func (m *ServerMessage) GetEmailcdnGetAccountsResponse() *EmailcdnGetAccountsResponse {
	if x, ok := m.GetMessageType().(*ServerMessage_EmailcdnGetAccountsResponse); ok {
		return x.EmailcdnGetAccountsResponse
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
		(*ServerMessage_ConfigChangedResponse)(nil),
		(*ServerMessage_EmailcdnTestAccountCredentialsRequest)(nil),
		(*ServerMessage_EmailcdnGetAccountsResponse)(nil),
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

type ConfigChangedResponse struct {
	HashesByTable        map[string][]byte `protobuf:"bytes,1,rep,name=hashes_by_table,json=hashesByTable,proto3" json:"hashes_by_table,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConfigChangedResponse) Reset()         { *m = ConfigChangedResponse{} }
func (m *ConfigChangedResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigChangedResponse) ProtoMessage()    {}
func (*ConfigChangedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{4}
}

func (m *ConfigChangedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigChangedResponse.Unmarshal(m, b)
}
func (m *ConfigChangedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigChangedResponse.Marshal(b, m, deterministic)
}
func (m *ConfigChangedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigChangedResponse.Merge(m, src)
}
func (m *ConfigChangedResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigChangedResponse.Size(m)
}
func (m *ConfigChangedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigChangedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigChangedResponse proto.InternalMessageInfo

func (m *ConfigChangedResponse) GetHashesByTable() map[string][]byte {
	if m != nil {
		return m.HashesByTable
	}
	return nil
}

type EmailcdnAccount struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Provider             string   `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	ImapHost             string   `protobuf:"bytes,5,opt,name=imap_host,json=imapHost,proto3" json:"imap_host,omitempty"`
	ImapPort             uint32   `protobuf:"varint,6,opt,name=imap_port,json=imapPort,proto3" json:"imap_port,omitempty"`
	ImapUsername         string   `protobuf:"bytes,7,opt,name=imap_username,json=imapUsername,proto3" json:"imap_username,omitempty"`
	ImapPassword         string   `protobuf:"bytes,8,opt,name=imap_password,json=imapPassword,proto3" json:"imap_password,omitempty"`
	SmtpHost             string   `protobuf:"bytes,9,opt,name=smtp_host,json=smtpHost,proto3" json:"smtp_host,omitempty"`
	SmtpPort             uint32   `protobuf:"varint,10,opt,name=smtp_port,json=smtpPort,proto3" json:"smtp_port,omitempty"`
	SmtpUsername         string   `protobuf:"bytes,11,opt,name=smtp_username,json=smtpUsername,proto3" json:"smtp_username,omitempty"`
	SmtpPassword         string   `protobuf:"bytes,12,opt,name=smtp_password,json=smtpPassword,proto3" json:"smtp_password,omitempty"`
	SslRequired          bool     `protobuf:"varint,13,opt,name=ssl_required,json=sslRequired,proto3" json:"ssl_required,omitempty"`
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

func (m *EmailcdnAccount) GetName() string {
	if m != nil {
		return m.Name
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

type EmailcdnGetAccountsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailcdnGetAccountsRequest) Reset()         { *m = EmailcdnGetAccountsRequest{} }
func (m *EmailcdnGetAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*EmailcdnGetAccountsRequest) ProtoMessage()    {}
func (*EmailcdnGetAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{6}
}

func (m *EmailcdnGetAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailcdnGetAccountsRequest.Unmarshal(m, b)
}
func (m *EmailcdnGetAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailcdnGetAccountsRequest.Marshal(b, m, deterministic)
}
func (m *EmailcdnGetAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailcdnGetAccountsRequest.Merge(m, src)
}
func (m *EmailcdnGetAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_EmailcdnGetAccountsRequest.Size(m)
}
func (m *EmailcdnGetAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailcdnGetAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailcdnGetAccountsRequest proto.InternalMessageInfo

type EmailcdnGetAccountsResponse struct {
	Accounts             []*EmailcdnAccount `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EmailcdnGetAccountsResponse) Reset()         { *m = EmailcdnGetAccountsResponse{} }
func (m *EmailcdnGetAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*EmailcdnGetAccountsResponse) ProtoMessage()    {}
func (*EmailcdnGetAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ad2f4d378dabd, []int{7}
}

func (m *EmailcdnGetAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailcdnGetAccountsResponse.Unmarshal(m, b)
}
func (m *EmailcdnGetAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailcdnGetAccountsResponse.Marshal(b, m, deterministic)
}
func (m *EmailcdnGetAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailcdnGetAccountsResponse.Merge(m, src)
}
func (m *EmailcdnGetAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_EmailcdnGetAccountsResponse.Size(m)
}
func (m *EmailcdnGetAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailcdnGetAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmailcdnGetAccountsResponse proto.InternalMessageInfo

func (m *EmailcdnGetAccountsResponse) GetAccounts() []*EmailcdnAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
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
	return fileDescriptor_652ad2f4d378dabd, []int{8}
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
	return fileDescriptor_652ad2f4d378dabd, []int{9}
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
	return fileDescriptor_652ad2f4d378dabd, []int{10}
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
	return fileDescriptor_652ad2f4d378dabd, []int{11}
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
	proto.RegisterType((*ConfigChangedResponse)(nil), "agentstreamproto.ConfigChangedResponse")
	proto.RegisterMapType((map[string][]byte)(nil), "agentstreamproto.ConfigChangedResponse.HashesByTableEntry")
	proto.RegisterType((*EmailcdnAccount)(nil), "agentstreamproto.EmailcdnAccount")
	proto.RegisterType((*EmailcdnGetAccountsRequest)(nil), "agentstreamproto.EmailcdnGetAccountsRequest")
	proto.RegisterType((*EmailcdnGetAccountsResponse)(nil), "agentstreamproto.EmailcdnGetAccountsResponse")
	proto.RegisterType((*EmailcdnTestAccountCredentialsRequest)(nil), "agentstreamproto.EmailcdnTestAccountCredentialsRequest")
	proto.RegisterType((*ErrorResponse)(nil), "agentstreamproto.ErrorResponse")
	proto.RegisterType((*NoopRequest)(nil), "agentstreamproto.NoopRequest")
	proto.RegisterType((*StartupRequest)(nil), "agentstreamproto.StartupRequest")
}

func init() { proto.RegisterFile("agentstream.proto", fileDescriptor_652ad2f4d378dabd) }

var fileDescriptor_652ad2f4d378dabd = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x5d, 0x6f, 0xeb, 0x44,
	0x10, 0xc5, 0x49, 0xda, 0xc6, 0x13, 0x3b, 0x2d, 0x2b, 0x2a, 0xa2, 0x94, 0x96, 0xd4, 0x50, 0x11,
	0x24, 0xc8, 0x43, 0x79, 0x00, 0x15, 0x21, 0xd1, 0x46, 0x15, 0x91, 0x10, 0xa8, 0x72, 0xcb, 0x1b,
	0x92, 0xb5, 0xb1, 0x87, 0xc4, 0x8a, 0xe3, 0x75, 0x77, 0x37, 0x45, 0xf9, 0x1b, 0xf0, 0x63, 0x78,
	0xe2, 0x81, 0x7f, 0x86, 0x76, 0xd7, 0x5f, 0xbd, 0x49, 0xae, 0xa2, 0xfb, 0x74, 0xdf, 0x32, 0x67,
	0x66, 0xce, 0x1e, 0xef, 0x9c, 0xd9, 0xc0, 0x87, 0x74, 0x86, 0xa9, 0x14, 0x92, 0x23, 0x5d, 0x8e,
	0x32, 0xce, 0x24, 0x23, 0x27, 0x35, 0x48, 0x23, 0xde, 0x7f, 0x4d, 0x70, 0xc7, 0x49, 0x8c, 0xa9,
	0xfc, 0x05, 0x85, 0xa0, 0x33, 0x24, 0x5d, 0x68, 0xc4, 0x51, 0xcf, 0x1a, 0x58, 0xc3, 0x96, 0xdf,
	0x88, 0x23, 0x72, 0x07, 0x0e, 0x0d, 0x17, 0x01, 0x47, 0x91, 0xb1, 0x54, 0x60, 0xaf, 0x31, 0xb0,
	0x86, 0x9d, 0xeb, 0xf3, 0xd1, 0x9b, 0x54, 0xa3, 0xdb, 0x70, 0xe1, 0xe7, 0x45, 0x93, 0x0f, 0xfc,
	0x0e, 0xad, 0x42, 0xc5, 0x91, 0x32, 0x96, 0x05, 0x1c, 0x9f, 0x57, 0x28, 0x64, 0xaf, 0xb9, 0x8b,
	0xe3, 0x57, 0xc6, 0x32, 0xdf, 0x14, 0x29, 0x8e, 0xb4, 0x0a, 0xc9, 0xcf, 0x70, 0x2c, 0x24, 0xe5,
	0x72, 0x55, 0xd1, 0xb4, 0x34, 0xcd, 0x60, 0x93, 0xe6, 0xd1, 0x14, 0x56, 0x4c, 0x5d, 0xf1, 0x0a,
	0x21, 0x13, 0xe8, 0x22, 0xe7, 0x8c, 0x57, 0x9f, 0x75, 0xa0, 0xb9, 0x3e, 0xdd, 0xe4, 0xba, 0x57,
	0x75, 0xb5, 0x0f, 0x73, 0xb1, 0x0e, 0x90, 0x67, 0x38, 0xc7, 0x25, 0x8d, 0x93, 0x30, 0x4a, 0x83,
	0x19, 0xca, 0x80, 0x86, 0x21, 0x5b, 0xa5, 0x52, 0x94, 0x22, 0x0f, 0x35, 0xf1, 0x57, 0x5b, 0x88,
	0xf3, 0xb6, 0x9f, 0x50, 0xde, 0xe6, 0x4d, 0x95, 0xe0, 0x3e, 0xee, 0xcc, 0xde, 0x75, 0xc1, 0x59,
	0x9a, 0x61, 0x05, 0x72, 0x9d, 0xa1, 0xf7, 0xcf, 0x01, 0xb8, 0x8f, 0xc8, 0x5f, 0x90, 0xbf, 0xef,
	0x33, 0xdc, 0xbc, 0xf6, 0xd6, 0x3b, 0x5e, 0xfb, 0x3d, 0xb8, 0x61, 0x42, 0xe3, 0x65, 0x29, 0xc7,
	0xcc, 0xef, 0x62, 0x93, 0x68, 0xac, 0xca, 0x2a, 0x3d, 0x4e, 0x58, 0x8b, 0x09, 0x85, 0x8f, 0x43,
	0x96, 0xfe, 0x11, 0xcf, 0x82, 0x70, 0x4e, 0xd3, 0x19, 0x46, 0x95, 0x32, 0x33, 0xb7, 0x2f, 0xb6,
	0x10, 0xea, 0x86, 0xb1, 0xa9, 0xaf, 0x29, 0x3c, 0x0d, 0xb7, 0x25, 0xc8, 0xdf, 0x16, 0x7c, 0x59,
	0x3a, 0x44, 0xa2, 0x28, 0x2d, 0x12, 0x84, 0x1c, 0x23, 0x4c, 0x65, 0x4c, 0x93, 0xca, 0x2d, 0x47,
	0xfa, 0xd4, 0x6f, 0x77, 0xbb, 0xe5, 0x09, 0x45, 0x61, 0x88, 0x71, 0xd5, 0x5f, 0x7d, 0xdf, 0x15,
	0xee, 0x53, 0x48, 0x24, 0x5c, 0xec, 0xb2, 0x6d, 0xfe, 0xfd, 0x6d, 0xad, 0xe4, 0xeb, 0x3d, 0x7d,
	0x5b, 0xde, 0xc2, 0x19, 0xee, 0x4e, 0x6f, 0x38, 0xd7, 0x85, 0x4e, 0xcd, 0x71, 0xde, 0xe7, 0xe0,
	0xd4, 0xa7, 0x45, 0x3e, 0x82, 0x03, 0xc9, 0x16, 0x98, 0x6a, 0x27, 0xdb, 0xbe, 0x09, 0xbc, 0x7f,
	0x2d, 0x38, 0xdd, 0x3a, 0x03, 0x32, 0x85, 0xe3, 0x39, 0x15, 0x73, 0x14, 0xc1, 0x74, 0x1d, 0x48,
	0x3a, 0x4d, 0xb0, 0x67, 0x0d, 0x9a, 0xc3, 0xce, 0xf5, 0xcd, 0x9e, 0x53, 0x1c, 0x4d, 0x74, 0xfb,
	0xdd, 0xfa, 0x49, 0x35, 0xdf, 0xa7, 0x92, 0xaf, 0x7d, 0x77, 0x5e, 0xc7, 0xfa, 0x3f, 0x02, 0xd9,
	0x2c, 0x22, 0x27, 0xd0, 0x5c, 0xe0, 0x3a, 0xd7, 0xa9, 0x7e, 0x2a, 0xed, 0x2f, 0x34, 0x59, 0x99,
	0x5d, 0x73, 0x7c, 0x13, 0xdc, 0x34, 0xbe, 0xb3, 0xbc, 0xbf, 0x9a, 0x70, 0x5c, 0xdc, 0x61, 0x7e,
	0x43, 0xb5, 0x85, 0xb5, 0xf5, 0xc2, 0x12, 0x68, 0xa5, 0x74, 0x69, 0x9a, 0x6d, 0x5f, 0xff, 0x26,
	0x7d, 0x68, 0x67, 0x9c, 0xbd, 0xc4, 0x11, 0x72, 0xbd, 0x7c, 0xb6, 0x5f, 0xc6, 0xea, 0x34, 0x7d,
	0xef, 0x7a, 0x9f, 0x6c, 0xdf, 0x04, 0xe4, 0x0c, 0xec, 0x78, 0x49, 0xb3, 0x60, 0xce, 0xf2, 0x05,
	0xb1, 0xfd, 0xb6, 0x02, 0x26, 0x4c, 0xc8, 0x32, 0x99, 0x31, 0x6e, 0x1e, 0x29, 0xd7, 0x24, 0x1f,
	0x18, 0x97, 0xe4, 0x33, 0x70, 0x75, 0x72, 0x25, 0x90, 0x6b, 0x21, 0x47, 0xba, 0xdb, 0x51, 0xe0,
	0x6f, 0x39, 0x56, 0x16, 0x65, 0x54, 0x88, 0x3f, 0x19, 0x8f, 0xb4, 0x65, 0xf2, 0xa2, 0x87, 0x1c,
	0x53, 0xc7, 0x88, 0xa5, 0xcc, 0x35, 0xd8, 0x46, 0x83, 0x02, 0x0a, 0x0d, 0x3a, 0xa9, 0x35, 0x80,
	0xd1, 0xa0, 0x80, 0x42, 0x83, 0x4e, 0x96, 0x1a, 0x3a, 0x86, 0x5e, 0x81, 0x75, 0x0d, 0x86, 0xa1,
	0xd0, 0xe0, 0x54, 0x45, 0xa5, 0x86, 0x4b, 0x70, 0x84, 0x48, 0xf4, 0x8e, 0xc5, 0x1c, 0xa3, 0x9e,
	0x3b, 0xb0, 0x86, 0x6d, 0xbf, 0x23, 0x44, 0xe2, 0xe7, 0x90, 0xf7, 0x09, 0xf4, 0x77, 0xbf, 0xc7,
	0xde, 0xef, 0x70, 0xf6, 0x16, 0xd7, 0x93, 0x1f, 0xa0, 0x5d, 0xec, 0x4f, 0x6e, 0xb8, 0xcb, 0xdd,
	0x6b, 0x93, 0x77, 0xfb, 0x65, 0x8b, 0x17, 0xc1, 0xd5, 0x5e, 0xdb, 0x4d, 0xbe, 0x87, 0xa3, 0xbc,
	0x49, 0x5b, 0x65, 0xaf, 0x63, 0x8a, 0x0e, 0xef, 0x0a, 0xdc, 0x57, 0x6f, 0xaa, 0xf6, 0x8c, 0x02,
	0x8a, 0xed, 0xd2, 0x81, 0x5a, 0xc9, 0xda, 0x03, 0xee, 0x9d, 0x40, 0xf7, 0xf5, 0x9f, 0xe9, 0xf4,
	0x50, 0x9f, 0xf3, 0xcd, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0x7c, 0xff, 0x2b, 0x5f, 0x08,
	0x00, 0x00,
}
