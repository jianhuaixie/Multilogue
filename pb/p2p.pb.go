// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: p2p.proto

package protocols_p2p

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// designed to be shared between all app protocols
type MessageData struct {
	// shared between all requests
	ClientVersion        string   `protobuf:"bytes,1,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Gossip               bool     `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	NodePubKey           []byte   `protobuf:"bytes,6,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Sign                 []byte   `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageData) Reset()         { *m = MessageData{} }
func (m *MessageData) String() string { return proto.CompactTextString(m) }
func (*MessageData) ProtoMessage()    {}
func (*MessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{0}
}
func (m *MessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageData.Unmarshal(m, b)
}
func (m *MessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageData.Marshal(b, m, deterministic)
}
func (m *MessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageData.Merge(m, src)
}
func (m *MessageData) XXX_Size() int {
	return xxx_messageInfo_MessageData.Size(m)
}
func (m *MessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageData.DiscardUnknown(m)
}

var xxx_messageInfo_MessageData proto.InternalMessageInfo

func (m *MessageData) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *MessageData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MessageData) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *MessageData) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MessageData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *MessageData) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

//// Multilogue protocol
type ProtocolMessage struct {
	MessageData          *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProtocolMessage) Reset()         { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()    {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{1}
}
func (m *ProtocolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolMessage.Unmarshal(m, b)
}
func (m *ProtocolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolMessage.Marshal(b, m, deterministic)
}
func (m *ProtocolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolMessage.Merge(m, src)
}
func (m *ProtocolMessage) XXX_Size() int {
	return xxx_messageInfo_ProtocolMessage.Size(m)
}
func (m *ProtocolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolMessage proto.InternalMessageInfo

func (m *ProtocolMessage) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

type ClientData struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientData) Reset()         { *m = ClientData{} }
func (m *ClientData) String() string { return proto.CompactTextString(m) }
func (*ClientData) ProtoMessage()    {}
func (*ClientData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{2}
}
func (m *ClientData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientData.Unmarshal(m, b)
}
func (m *ClientData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientData.Marshal(b, m, deterministic)
}
func (m *ClientData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientData.Merge(m, src)
}
func (m *ClientData) XXX_Size() int {
	return xxx_messageInfo_ClientData.Size(m)
}
func (m *ClientData) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientData.DiscardUnknown(m)
}

var xxx_messageInfo_ClientData proto.InternalMessageInfo

func (m *ClientData) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *ClientData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type HostData struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
	PeerId               string   `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostData) Reset()         { *m = HostData{} }
func (m *HostData) String() string { return proto.CompactTextString(m) }
func (*HostData) ProtoMessage()    {}
func (*HostData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{3}
}
func (m *HostData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostData.Unmarshal(m, b)
}
func (m *HostData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostData.Marshal(b, m, deterministic)
}
func (m *HostData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostData.Merge(m, src)
}
func (m *HostData) XXX_Size() int {
	return xxx_messageInfo_HostData.Size(m)
}
func (m *HostData) XXX_DiscardUnknown() {
	xxx_messageInfo_HostData.DiscardUnknown(m)
}

var xxx_messageInfo_HostData proto.InternalMessageInfo

func (m *HostData) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *HostData) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

// a protocol define a set of reuqest and responses
type ClientSendMessage struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	Message              string      `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClientSendMessage) Reset()         { *m = ClientSendMessage{} }
func (m *ClientSendMessage) String() string { return proto.CompactTextString(m) }
func (*ClientSendMessage) ProtoMessage()    {}
func (*ClientSendMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{4}
}
func (m *ClientSendMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSendMessage.Unmarshal(m, b)
}
func (m *ClientSendMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSendMessage.Marshal(b, m, deterministic)
}
func (m *ClientSendMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSendMessage.Merge(m, src)
}
func (m *ClientSendMessage) XXX_Size() int {
	return xxx_messageInfo_ClientSendMessage.Size(m)
}
func (m *ClientSendMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSendMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSendMessage proto.InternalMessageInfo

func (m *ClientSendMessage) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *ClientSendMessage) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *ClientSendMessage) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

func (m *ClientSendMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ClientTransmissionStart struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClientTransmissionStart) Reset()         { *m = ClientTransmissionStart{} }
func (m *ClientTransmissionStart) String() string { return proto.CompactTextString(m) }
func (*ClientTransmissionStart) ProtoMessage()    {}
func (*ClientTransmissionStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{5}
}
func (m *ClientTransmissionStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientTransmissionStart.Unmarshal(m, b)
}
func (m *ClientTransmissionStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientTransmissionStart.Marshal(b, m, deterministic)
}
func (m *ClientTransmissionStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientTransmissionStart.Merge(m, src)
}
func (m *ClientTransmissionStart) XXX_Size() int {
	return xxx_messageInfo_ClientTransmissionStart.Size(m)
}
func (m *ClientTransmissionStart) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientTransmissionStart.DiscardUnknown(m)
}

var xxx_messageInfo_ClientTransmissionStart proto.InternalMessageInfo

func (m *ClientTransmissionStart) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *ClientTransmissionStart) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *ClientTransmissionStart) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

type ClientTransmissionEnd struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClientTransmissionEnd) Reset()         { *m = ClientTransmissionEnd{} }
func (m *ClientTransmissionEnd) String() string { return proto.CompactTextString(m) }
func (*ClientTransmissionEnd) ProtoMessage()    {}
func (*ClientTransmissionEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{6}
}
func (m *ClientTransmissionEnd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientTransmissionEnd.Unmarshal(m, b)
}
func (m *ClientTransmissionEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientTransmissionEnd.Marshal(b, m, deterministic)
}
func (m *ClientTransmissionEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientTransmissionEnd.Merge(m, src)
}
func (m *ClientTransmissionEnd) XXX_Size() int {
	return xxx_messageInfo_ClientTransmissionEnd.Size(m)
}
func (m *ClientTransmissionEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientTransmissionEnd.DiscardUnknown(m)
}

var xxx_messageInfo_ClientTransmissionEnd proto.InternalMessageInfo

func (m *ClientTransmissionEnd) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *ClientTransmissionEnd) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *ClientTransmissionEnd) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

type HostAcceptTransmission struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HostAcceptTransmission) Reset()         { *m = HostAcceptTransmission{} }
func (m *HostAcceptTransmission) String() string { return proto.CompactTextString(m) }
func (*HostAcceptTransmission) ProtoMessage()    {}
func (*HostAcceptTransmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{7}
}
func (m *HostAcceptTransmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostAcceptTransmission.Unmarshal(m, b)
}
func (m *HostAcceptTransmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostAcceptTransmission.Marshal(b, m, deterministic)
}
func (m *HostAcceptTransmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostAcceptTransmission.Merge(m, src)
}
func (m *HostAcceptTransmission) XXX_Size() int {
	return xxx_messageInfo_HostAcceptTransmission.Size(m)
}
func (m *HostAcceptTransmission) XXX_DiscardUnknown() {
	xxx_messageInfo_HostAcceptTransmission.DiscardUnknown(m)
}

var xxx_messageInfo_HostAcceptTransmission proto.InternalMessageInfo

func (m *HostAcceptTransmission) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *HostAcceptTransmission) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *HostAcceptTransmission) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

type HostDenyTransmission struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	ErrorCode            int32       `protobuf:"varint,4,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HostDenyTransmission) Reset()         { *m = HostDenyTransmission{} }
func (m *HostDenyTransmission) String() string { return proto.CompactTextString(m) }
func (*HostDenyTransmission) ProtoMessage()    {}
func (*HostDenyTransmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{8}
}
func (m *HostDenyTransmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostDenyTransmission.Unmarshal(m, b)
}
func (m *HostDenyTransmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostDenyTransmission.Marshal(b, m, deterministic)
}
func (m *HostDenyTransmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostDenyTransmission.Merge(m, src)
}
func (m *HostDenyTransmission) XXX_Size() int {
	return xxx_messageInfo_HostDenyTransmission.Size(m)
}
func (m *HostDenyTransmission) XXX_DiscardUnknown() {
	xxx_messageInfo_HostDenyTransmission.DiscardUnknown(m)
}

var xxx_messageInfo_HostDenyTransmission proto.InternalMessageInfo

func (m *HostDenyTransmission) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *HostDenyTransmission) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *HostDenyTransmission) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

func (m *HostDenyTransmission) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type HostBroadcastMessage struct {
	MessageData          *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	ClientData           *ClientData  `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData    `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	Message              string       `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HostBroadcastMessage) Reset()         { *m = HostBroadcastMessage{} }
func (m *HostBroadcastMessage) String() string { return proto.CompactTextString(m) }
func (*HostBroadcastMessage) ProtoMessage()    {}
func (*HostBroadcastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{9}
}
func (m *HostBroadcastMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostBroadcastMessage.Unmarshal(m, b)
}
func (m *HostBroadcastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostBroadcastMessage.Marshal(b, m, deterministic)
}
func (m *HostBroadcastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostBroadcastMessage.Merge(m, src)
}
func (m *HostBroadcastMessage) XXX_Size() int {
	return xxx_messageInfo_HostBroadcastMessage.Size(m)
}
func (m *HostBroadcastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HostBroadcastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HostBroadcastMessage proto.InternalMessageInfo

func (m *HostBroadcastMessage) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *HostBroadcastMessage) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *HostBroadcastMessage) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

func (m *HostBroadcastMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ClientJoinChannel struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClientJoinChannel) Reset()         { *m = ClientJoinChannel{} }
func (m *ClientJoinChannel) String() string { return proto.CompactTextString(m) }
func (*ClientJoinChannel) ProtoMessage()    {}
func (*ClientJoinChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{10}
}
func (m *ClientJoinChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientJoinChannel.Unmarshal(m, b)
}
func (m *ClientJoinChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientJoinChannel.Marshal(b, m, deterministic)
}
func (m *ClientJoinChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientJoinChannel.Merge(m, src)
}
func (m *ClientJoinChannel) XXX_Size() int {
	return xxx_messageInfo_ClientJoinChannel.Size(m)
}
func (m *ClientJoinChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientJoinChannel.DiscardUnknown(m)
}

var xxx_messageInfo_ClientJoinChannel proto.InternalMessageInfo

func (m *ClientJoinChannel) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *ClientJoinChannel) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *ClientJoinChannel) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

type ClientLeaveChannel struct {
	MessageData          *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	ClientData           *ClientData  `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData    `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ClientLeaveChannel) Reset()         { *m = ClientLeaveChannel{} }
func (m *ClientLeaveChannel) String() string { return proto.CompactTextString(m) }
func (*ClientLeaveChannel) ProtoMessage()    {}
func (*ClientLeaveChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{11}
}
func (m *ClientLeaveChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientLeaveChannel.Unmarshal(m, b)
}
func (m *ClientLeaveChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientLeaveChannel.Marshal(b, m, deterministic)
}
func (m *ClientLeaveChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientLeaveChannel.Merge(m, src)
}
func (m *ClientLeaveChannel) XXX_Size() int {
	return xxx_messageInfo_ClientLeaveChannel.Size(m)
}
func (m *ClientLeaveChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientLeaveChannel.DiscardUnknown(m)
}

var xxx_messageInfo_ClientLeaveChannel proto.InternalMessageInfo

func (m *ClientLeaveChannel) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *ClientLeaveChannel) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *ClientLeaveChannel) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

type HostAcceptClient struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HostAcceptClient) Reset()         { *m = HostAcceptClient{} }
func (m *HostAcceptClient) String() string { return proto.CompactTextString(m) }
func (*HostAcceptClient) ProtoMessage()    {}
func (*HostAcceptClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{12}
}
func (m *HostAcceptClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostAcceptClient.Unmarshal(m, b)
}
func (m *HostAcceptClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostAcceptClient.Marshal(b, m, deterministic)
}
func (m *HostAcceptClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostAcceptClient.Merge(m, src)
}
func (m *HostAcceptClient) XXX_Size() int {
	return xxx_messageInfo_HostAcceptClient.Size(m)
}
func (m *HostAcceptClient) XXX_DiscardUnknown() {
	xxx_messageInfo_HostAcceptClient.DiscardUnknown(m)
}

var xxx_messageInfo_HostAcceptClient proto.InternalMessageInfo

func (m *HostAcceptClient) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *HostAcceptClient) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *HostAcceptClient) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

type HostDenyClient struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	ClientData           *ClientData `protobuf:"bytes,2,opt,name=clientData,proto3" json:"clientData,omitempty"`
	HostData             *HostData   `protobuf:"bytes,3,opt,name=hostData,proto3" json:"hostData,omitempty"`
	ErrorCode            int32       `protobuf:"varint,4,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HostDenyClient) Reset()         { *m = HostDenyClient{} }
func (m *HostDenyClient) String() string { return proto.CompactTextString(m) }
func (*HostDenyClient) ProtoMessage()    {}
func (*HostDenyClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{13}
}
func (m *HostDenyClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostDenyClient.Unmarshal(m, b)
}
func (m *HostDenyClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostDenyClient.Marshal(b, m, deterministic)
}
func (m *HostDenyClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostDenyClient.Merge(m, src)
}
func (m *HostDenyClient) XXX_Size() int {
	return xxx_messageInfo_HostDenyClient.Size(m)
}
func (m *HostDenyClient) XXX_DiscardUnknown() {
	xxx_messageInfo_HostDenyClient.DiscardUnknown(m)
}

var xxx_messageInfo_HostDenyClient proto.InternalMessageInfo

func (m *HostDenyClient) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *HostDenyClient) GetClientData() *ClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *HostDenyClient) GetHostData() *HostData {
	if m != nil {
		return m.HostData
	}
	return nil
}

func (m *HostDenyClient) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageData)(nil), "protocols.p2p.MessageData")
	proto.RegisterType((*ProtocolMessage)(nil), "protocols.p2p.ProtocolMessage")
	proto.RegisterType((*ClientData)(nil), "protocols.p2p.ClientData")
	proto.RegisterType((*HostData)(nil), "protocols.p2p.HostData")
	proto.RegisterType((*ClientSendMessage)(nil), "protocols.p2p.ClientSendMessage")
	proto.RegisterType((*ClientTransmissionStart)(nil), "protocols.p2p.ClientTransmissionStart")
	proto.RegisterType((*ClientTransmissionEnd)(nil), "protocols.p2p.ClientTransmissionEnd")
	proto.RegisterType((*HostAcceptTransmission)(nil), "protocols.p2p.HostAcceptTransmission")
	proto.RegisterType((*HostDenyTransmission)(nil), "protocols.p2p.HostDenyTransmission")
	proto.RegisterType((*HostBroadcastMessage)(nil), "protocols.p2p.HostBroadcastMessage")
	proto.RegisterType((*ClientJoinChannel)(nil), "protocols.p2p.ClientJoinChannel")
	proto.RegisterType((*ClientLeaveChannel)(nil), "protocols.p2p.ClientLeaveChannel")
	proto.RegisterType((*HostAcceptClient)(nil), "protocols.p2p.HostAcceptClient")
	proto.RegisterType((*HostDenyClient)(nil), "protocols.p2p.HostDenyClient")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor_e7fdddb109e6467a) }

var fileDescriptor_e7fdddb109e6467a = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe5, 0x6e, 0xeb, 0x9a, 0x17, 0x36, 0xc0, 0x82, 0x2e, 0x4c, 0x13, 0xaa, 0x22, 0x0e,
	0x3d, 0xf5, 0x90, 0x9d, 0x90, 0x38, 0x0c, 0x0a, 0x12, 0xe3, 0x87, 0x98, 0x3c, 0xc4, 0xdd, 0x8b,
	0x9f, 0x3a, 0x4b, 0x8d, 0x6d, 0xd9, 0x19, 0xd2, 0xfe, 0x3b, 0x40, 0x02, 0xfe, 0x08, 0x10, 0x07,
	0xfe, 0x12, 0x14, 0xc7, 0x4d, 0xd2, 0x22, 0x71, 0xe2, 0x92, 0x9d, 0xea, 0xf7, 0xed, 0x7b, 0xcf,
	0xf9, 0x3c, 0xdb, 0x5f, 0x88, 0x4c, 0x66, 0x66, 0xc6, 0xea, 0x52, 0xd3, 0x3d, 0xff, 0x93, 0xeb,
	0xa5, 0x9b, 0x99, 0xcc, 0xa4, 0x5f, 0x09, 0xc4, 0x6f, 0xd1, 0x39, 0xbe, 0xc0, 0xe7, 0xbc, 0xe4,
	0xf4, 0x11, 0xec, 0xe5, 0x4b, 0x89, 0xaa, 0xfc, 0x80, 0xd6, 0x49, 0xad, 0x12, 0x32, 0x21, 0xd3,
	0x88, 0xad, 0x8b, 0xf4, 0x08, 0xa2, 0x52, 0x16, 0xe8, 0x4a, 0x5e, 0x98, 0x64, 0x30, 0x21, 0xd3,
	0x2d, 0xd6, 0x0a, 0x74, 0x1f, 0x06, 0x52, 0x24, 0x5b, 0xbe, 0x70, 0x20, 0x05, 0x1d, 0xc3, 0x70,
	0xa1, 0x9d, 0x93, 0x26, 0xd9, 0x9e, 0x90, 0xe9, 0x88, 0x85, 0xa8, 0xd2, 0x95, 0x16, 0x78, 0x2a,
	0x92, 0x1d, 0x9f, 0x1b, 0x22, 0xfa, 0x10, 0xa0, 0x5a, 0x9d, 0x5d, 0x5d, 0xbc, 0xc6, 0xeb, 0x64,
	0x38, 0x21, 0xd3, 0x5b, 0xac, 0xa3, 0x50, 0x0a, 0xdb, 0x4e, 0x2e, 0x54, 0xb2, 0xeb, 0xff, 0xf1,
	0xeb, 0xf4, 0x1d, 0xdc, 0x3e, 0x0b, 0x60, 0x01, 0x87, 0x3e, 0x81, 0xb8, 0x68, 0xc9, 0x3c, 0x48,
	0x9c, 0x1d, 0xce, 0xd6, 0xf8, 0x67, 0x1d, 0x76, 0xd6, 0x4d, 0x4f, 0x4f, 0x00, 0xe6, 0x9e, 0xd9,
	0x8f, 0x65, 0x0c, 0x43, 0x83, 0x68, 0x4f, 0x45, 0x98, 0x47, 0x88, 0xe8, 0x21, 0x8c, 0xae, 0x1c,
	0x5a, 0xc5, 0x0b, 0xf4, 0x73, 0x88, 0x58, 0x13, 0xa7, 0x27, 0x30, 0x7a, 0xa9, 0x5d, 0x5d, 0x7f,
	0x04, 0x51, 0x7e, 0xc9, 0x95, 0xc2, 0x65, 0xd3, 0xa2, 0x15, 0x3a, 0xdd, 0x07, 0xdd, 0xee, 0xe9,
	0x0f, 0x02, 0x77, 0xeb, 0x8f, 0x38, 0x47, 0x25, 0xfe, 0x0b, 0x17, 0x7d, 0x0c, 0x90, 0x37, 0x5c,
	0x7e, 0xbf, 0x38, 0x7b, 0xb0, 0x51, 0xdc, 0x82, 0xb3, 0x4e, 0x32, 0x3d, 0x86, 0xd1, 0x65, 0x00,
	0xf2, 0xa7, 0x1b, 0x67, 0x07, 0x1b, 0x85, 0x2b, 0x5e, 0xd6, 0x24, 0xd2, 0x04, 0x76, 0xc3, 0xf6,
	0xfe, 0xf4, 0x23, 0xb6, 0x0a, 0xd3, 0xef, 0x04, 0x0e, 0xea, 0x9d, 0xde, 0x5b, 0xae, 0x5c, 0x21,
	0x5d, 0x75, 0xb7, 0xce, 0x4b, 0x6e, 0xcb, 0x7e, 0x31, 0x56, 0x8f, 0xe8, 0xfe, 0xdf, 0x24, 0x2f,
	0x94, 0xe8, 0x19, 0xc7, 0x37, 0x02, 0xe3, 0x4a, 0x7e, 0x9a, 0xe7, 0x68, 0xd6, 0x58, 0x7a, 0x06,
	0xf2, 0x9b, 0xc0, 0x3d, 0x2f, 0xa3, 0xba, 0xee, 0x2f, 0x46, 0xe5, 0x1a, 0x68, 0xad, 0xb6, 0x73,
	0x2d, 0xea, 0xd7, 0xb3, 0xc3, 0x5a, 0x21, 0xfd, 0x15, 0x20, 0x9f, 0x59, 0xcd, 0x45, 0xce, 0x5d,
	0x79, 0xd3, 0x0c, 0xe2, 0x73, 0x63, 0x7f, 0xaf, 0xb4, 0x54, 0xf3, 0xda, 0x2e, 0x7b, 0x76, 0x13,
	0xbf, 0x10, 0xa0, 0x75, 0xbf, 0x37, 0xc8, 0x3f, 0x62, 0x3f, 0x21, 0x3e, 0x11, 0xb8, 0xd3, 0xfa,
	0x42, 0xdd, 0xb9, 0x67, 0x08, 0x3f, 0x09, 0xec, 0xaf, 0x1c, 0xa1, 0x8f, 0x00, 0xff, 0xf6, 0x82,
	0x8b, 0xa1, 0xaf, 0x3f, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x36, 0x56, 0x12, 0x6f, 0xe9, 0x09,
	0x00, 0x00,
}
