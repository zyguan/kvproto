// Code generated by protoc-gen-go.
// source: raft_serverpb.proto
// DO NOT EDIT!

/*
Package raft_serverpb is a generated protocol buffer package.

It is generated from these files:
	raft_serverpb.proto

It has these top-level messages:
	Message
	RaftMessage
	RaftTruncatedState
	KeyValue
	RaftSnapshotData
	StoreIdent
*/
package raft_serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import raftpb "github.com/pingcap/kvproto/pkg/raftpb"
import metapb "github.com/pingcap/kvproto/pkg/metapb"
import raft_cmdpb "github.com/pingcap/kvproto/pkg/raft_cmdpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MessageType int32

const (
	MessageType_None        MessageType = 0
	MessageType_Command     MessageType = 1
	MessageType_CommandResp MessageType = 2
	MessageType_Raft        MessageType = 3
)

var MessageType_name = map[int32]string{
	0: "None",
	1: "Command",
	2: "CommandResp",
	3: "Raft",
}
var MessageType_value = map[string]int32{
	"None":        0,
	"Command":     1,
	"CommandResp": 2,
	"Raft":        3,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Message holds all messages communicating with raft server.
type Message struct {
	MsgType          *MessageType                    `protobuf:"varint,1,opt,name=msg_type,enum=raft_serverpb.MessageType" json:"msg_type,omitempty"`
	CmdReq           *raft_cmdpb.RaftCommandRequest  `protobuf:"bytes,2,opt,name=cmd_req" json:"cmd_req,omitempty"`
	CmdResp          *raft_cmdpb.RaftCommandResponse `protobuf:"bytes,3,opt,name=cmd_resp" json:"cmd_resp,omitempty"`
	Raft             *RaftMessage                    `protobuf:"bytes,4,opt,name=raft" json:"raft,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetMsgType() MessageType {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return MessageType_None
}

func (m *Message) GetCmdReq() *raft_cmdpb.RaftCommandRequest {
	if m != nil {
		return m.CmdReq
	}
	return nil
}

func (m *Message) GetCmdResp() *raft_cmdpb.RaftCommandResponse {
	if m != nil {
		return m.CmdResp
	}
	return nil
}

func (m *Message) GetRaft() *RaftMessage {
	if m != nil {
		return m.Raft
	}
	return nil
}

type RaftMessage struct {
	RegionId         *uint64             `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	FromPeer         *metapb.Peer        `protobuf:"bytes,2,opt,name=from_peer" json:"from_peer,omitempty"`
	ToPeer           *metapb.Peer        `protobuf:"bytes,3,opt,name=to_peer" json:"to_peer,omitempty"`
	Message          *raftpb.Message     `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	RegionEpoch      *metapb.RegionEpoch `protobuf:"bytes,5,opt,name=region_epoch" json:"region_epoch,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RaftMessage) Reset()                    { *m = RaftMessage{} }
func (m *RaftMessage) String() string            { return proto.CompactTextString(m) }
func (*RaftMessage) ProtoMessage()               {}
func (*RaftMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RaftMessage) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *RaftMessage) GetFromPeer() *metapb.Peer {
	if m != nil {
		return m.FromPeer
	}
	return nil
}

func (m *RaftMessage) GetToPeer() *metapb.Peer {
	if m != nil {
		return m.ToPeer
	}
	return nil
}

func (m *RaftMessage) GetMessage() *raftpb.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RaftMessage) GetRegionEpoch() *metapb.RegionEpoch {
	if m != nil {
		return m.RegionEpoch
	}
	return nil
}

type RaftTruncatedState struct {
	Index            *uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Term             *uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RaftTruncatedState) Reset()                    { *m = RaftTruncatedState{} }
func (m *RaftTruncatedState) String() string            { return proto.CompactTextString(m) }
func (*RaftTruncatedState) ProtoMessage()               {}
func (*RaftTruncatedState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RaftTruncatedState) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *RaftTruncatedState) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

type KeyValue struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RaftSnapshotData struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	Data             []*KeyValue    `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RaftSnapshotData) Reset()                    { *m = RaftSnapshotData{} }
func (m *RaftSnapshotData) String() string            { return proto.CompactTextString(m) }
func (*RaftSnapshotData) ProtoMessage()               {}
func (*RaftSnapshotData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RaftSnapshotData) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *RaftSnapshotData) GetData() []*KeyValue {
	if m != nil {
		return m.Data
	}
	return nil
}

type StoreIdent struct {
	ClusterId        *uint64 `protobuf:"varint,1,opt,name=cluster_id" json:"cluster_id,omitempty"`
	NodeId           *uint64 `protobuf:"varint,2,opt,name=node_id" json:"node_id,omitempty"`
	StoreId          *uint64 `protobuf:"varint,3,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StoreIdent) Reset()                    { *m = StoreIdent{} }
func (m *StoreIdent) String() string            { return proto.CompactTextString(m) }
func (*StoreIdent) ProtoMessage()               {}
func (*StoreIdent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StoreIdent) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

func (m *StoreIdent) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *StoreIdent) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "raft_serverpb.Message")
	proto.RegisterType((*RaftMessage)(nil), "raft_serverpb.RaftMessage")
	proto.RegisterType((*RaftTruncatedState)(nil), "raft_serverpb.RaftTruncatedState")
	proto.RegisterType((*KeyValue)(nil), "raft_serverpb.KeyValue")
	proto.RegisterType((*RaftSnapshotData)(nil), "raft_serverpb.RaftSnapshotData")
	proto.RegisterType((*StoreIdent)(nil), "raft_serverpb.StoreIdent")
	proto.RegisterEnum("raft_serverpb.MessageType", MessageType_name, MessageType_value)
}

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x4b, 0x46, 0xb2, 0x93, 0x6c, 0x0b, 0x67, 0x17, 0x44, 0x95, 0xd8, 0xa6, 0x48, 0xa0,
	0x81, 0x50, 0xd0, 0xfa, 0x02, 0x5c, 0x0c, 0x2e, 0x10, 0x02, 0xa1, 0x76, 0x42, 0xe2, 0x2a, 0x32,
	0xcd, 0x59, 0x57, 0xb1, 0xd8, 0x9e, 0xed, 0x4e, 0xf4, 0xa1, 0x78, 0x0c, 0xde, 0x0b, 0xdb, 0x75,
	0xa0, 0x9d, 0xd8, 0x4d, 0x14, 0x7f, 0x7f, 0xe7, 0xf3, 0x91, 0xe1, 0x48, 0xb1, 0x2b, 0xd3, 0x6a,
	0x52, 0x77, 0xa4, 0xe4, 0xf7, 0x46, 0x2a, 0x61, 0x04, 0xee, 0x6f, 0x81, 0xa3, 0xc2, 0x1d, 0x07,
	0x72, 0x54, 0xf4, 0x64, 0xd8, 0xdf, 0x53, 0xe9, 0xa5, 0xb3, 0xbe, 0x1b, 0x90, 0xfa, 0x77, 0x04,
	0xe9, 0x27, 0xd2, 0x9a, 0xcd, 0x09, 0x5f, 0x43, 0xd6, 0xeb, 0x79, 0x6b, 0x56, 0x92, 0xaa, 0xe8,
	0x34, 0x3a, 0x3b, 0x18, 0x8f, 0x9a, 0xed, 0x81, 0x41, 0x79, 0x69, 0x15, 0xf8, 0x06, 0x52, 0x1b,
	0xd4, 0x2a, 0xba, 0xad, 0x76, 0xac, 0x38, 0x1f, 0x1f, 0x37, 0x1b, 0xe9, 0x13, 0xfb, 0x7b, 0x21,
	0xfa, 0x9e, 0xf1, 0x6e, 0x42, 0xb7, 0x4b, 0xd2, 0x06, 0xcf, 0x21, 0x5b, 0x1b, 0xb4, 0xac, 0x62,
	0xef, 0x38, 0x79, 0xd0, 0xa1, 0xa5, 0xe0, 0x9a, 0xf0, 0x0c, 0x12, 0xa7, 0xa8, 0x12, 0x2f, 0xbf,
	0xdf, 0xc6, 0x39, 0x42, 0xa3, 0xfa, 0x57, 0x04, 0xf9, 0xc6, 0x19, 0x9f, 0xc0, 0x9e, 0xa2, 0xf9,
	0x42, 0xf0, 0x76, 0xd1, 0xf9, 0xcb, 0x24, 0x78, 0x02, 0x7b, 0x57, 0x4a, 0xf4, 0xad, 0x24, 0x52,
	0xa1, 0x72, 0xd1, 0x84, 0xf5, 0x7c, 0xb1, 0x18, 0x3e, 0x83, 0xd4, 0x88, 0x35, 0x1d, 0xff, 0x87,
	0x3e, 0x85, 0xb4, 0x5f, 0xa7, 0x87, 0x3e, 0x87, 0x4d, 0x58, 0xf5, 0x30, 0xf4, 0x25, 0x14, 0x61,
	0x28, 0x49, 0x31, 0xbb, 0xae, 0x76, 0xbd, 0xec, 0x68, 0x48, 0x99, 0x78, 0xee, 0xbd, 0xa3, 0xea,
	0x73, 0x40, 0x57, 0xf7, 0x52, 0x2d, 0xf9, 0x8c, 0x19, 0xea, 0xa6, 0xc6, 0x7e, 0x71, 0x1f, 0x76,
	0x17, 0xbc, 0xa3, 0x9f, 0xa1, 0x71, 0x01, 0x89, 0x21, 0xd5, 0xfb, 0xb2, 0x49, 0xfd, 0x02, 0xb2,
	0x8f, 0xb4, 0xfa, 0xca, 0x6e, 0x96, 0x84, 0x39, 0xc4, 0x3f, 0x68, 0xe5, 0x65, 0x85, 0x73, 0xdd,
	0x39, 0xd4, 0xeb, 0x8a, 0xfa, 0x1b, 0x94, 0x2e, 0x7a, 0xca, 0x99, 0xd4, 0xd7, 0xc2, 0xbc, 0x63,
	0x86, 0xe1, 0x31, 0x3c, 0x5e, 0x37, 0xf3, 0x96, 0x7c, 0x7c, 0xb0, 0xdd, 0x09, 0x9f, 0x43, 0xd2,
	0x59, 0x9d, 0x4d, 0x88, 0x2d, 0xfb, 0xf4, 0xde, 0xa2, 0x87, 0xb1, 0xf5, 0x05, 0xc0, 0xd4, 0x08,
	0x45, 0x1f, 0x3a, 0xe2, 0x06, 0x11, 0x60, 0x76, 0xb3, 0xd4, 0xb6, 0xe2, 0xbf, 0x25, 0x1f, 0x42,
	0xca, 0x45, 0x47, 0x0e, 0xf0, 0xad, 0xb1, 0x84, 0x4c, 0x3b, 0x8b, 0x43, 0xdc, 0x56, 0x93, 0x57,
	0x6f, 0x21, 0xdf, 0x7c, 0x47, 0x19, 0x24, 0x9f, 0x05, 0xa7, 0xf2, 0x91, 0xbd, 0x54, 0x1a, 0x1e,
	0x40, 0x19, 0xd9, 0xa0, 0x7c, 0xe3, 0x35, 0x94, 0x3b, 0x4e, 0xe7, 0xae, 0x55, 0xc6, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xe7, 0x4a, 0xa9, 0x06, 0x03, 0x00, 0x00,
}