// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pulse.proto

package pulse_go_proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This message allows the DCP to determine if each agent is alive or dead.
type Msg struct {
	AgentId              *AgentId `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Frequency            int32    `protobuf:"varint,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	AgentVersion         string   `protobuf:"bytes,3,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
	AgentLogsDir         string   `protobuf:"bytes,4,opt,name=agent_logs_dir,json=agentLogsDir,proto3" json:"agent_logs_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c067e3d82b299225, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetAgentId() *AgentId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *Msg) GetFrequency() int32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *Msg) GetAgentVersion() string {
	if m != nil {
		return m.AgentVersion
	}
	return ""
}

func (m *Msg) GetAgentLogsDir() string {
	if m != nil {
		return m.AgentLogsDir
	}
	return ""
}

// This message stores a unique identifier for each agent.
// The DCP can use this to separate each agent and monitor future behaviors.
type AgentId struct {
	HostName             string   `protobuf:"bytes,1,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	ProcessId            string   `protobuf:"bytes,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentId) Reset()         { *m = AgentId{} }
func (m *AgentId) String() string { return proto.CompactTextString(m) }
func (*AgentId) ProtoMessage()    {}
func (*AgentId) Descriptor() ([]byte, []int) {
	return fileDescriptor_c067e3d82b299225, []int{1}
}

func (m *AgentId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentId.Unmarshal(m, b)
}
func (m *AgentId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentId.Marshal(b, m, deterministic)
}
func (m *AgentId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentId.Merge(m, src)
}
func (m *AgentId) XXX_Size() int {
	return xxx_messageInfo_AgentId.Size(m)
}
func (m *AgentId) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentId.DiscardUnknown(m)
}

var xxx_messageInfo_AgentId proto.InternalMessageInfo

func (m *AgentId) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *AgentId) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "cloud_ingest_pulse.Msg")
	proto.RegisterType((*AgentId)(nil), "cloud_ingest_pulse.AgentId")
}

func init() { proto.RegisterFile("pulse.proto", fileDescriptor_c067e3d82b299225) }

var fileDescriptor_c067e3d82b299225 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0xa6, 0x6e, 0x79, 0xd3, 0x21, 0x39, 0x15, 0xa6, 0x50, 0xa6, 0x87, 0x9e, 0x7a,
	0x50, 0xf0, 0xae, 0xe8, 0x61, 0xa0, 0x1e, 0x7a, 0xf0, 0xe0, 0x25, 0xd4, 0xe6, 0x19, 0x03, 0x5d,
	0x5e, 0xcd, 0x6b, 0x05, 0xbf, 0x90, 0x9f, 0x53, 0x9a, 0x0c, 0x14, 0x76, 0x4a, 0xf8, 0xfd, 0x7f,
	0x24, 0xef, 0xff, 0x60, 0xd1, 0x0d, 0x2d, 0x63, 0xd9, 0x05, 0xea, 0x49, 0xa9, 0xa6, 0xa5, 0xc1,
	0x68, 0xe7, 0x2d, 0x72, 0xaf, 0x63, 0xb2, 0xfe, 0x11, 0x30, 0x7d, 0x62, 0xab, 0x6e, 0x60, 0x5e,
	0x5b, 0xf4, 0xbd, 0x76, 0x26, 0x13, 0xb9, 0x28, 0x16, 0x57, 0xab, 0x72, 0x5f, 0x2f, 0x6f, 0x47,
	0x67, 0x63, 0xaa, 0x59, 0x9d, 0x2e, 0xea, 0x0c, 0xe4, 0x7b, 0xc0, 0xcf, 0x01, 0x7d, 0xf3, 0x9d,
	0x4d, 0x72, 0x51, 0x1c, 0x56, 0x7f, 0x40, 0x5d, 0xc0, 0x49, 0x7a, 0xf5, 0x0b, 0x03, 0x3b, 0xf2,
	0xd9, 0x34, 0x17, 0x85, 0xac, 0x8e, 0x23, 0x7c, 0x49, 0x4c, 0x5d, 0xc2, 0x32, 0x49, 0x2d, 0x59,
	0xd6, 0xc6, 0x85, 0xec, 0xe0, 0x9f, 0xf5, 0x48, 0x96, 0xef, 0x5d, 0x58, 0x3f, 0xc0, 0x6c, 0xf7,
	0xb9, 0x5a, 0x81, 0xfc, 0x20, 0xee, 0xb5, 0xaf, 0xb7, 0x18, 0x87, 0x95, 0xd5, 0x7c, 0x04, 0xcf,
	0xf5, 0x16, 0xd5, 0x39, 0x40, 0x17, 0xa8, 0x41, 0xe6, 0xb1, 0xca, 0x24, 0xa6, 0x72, 0x47, 0x36,
	0xe6, 0xee, 0xf4, 0x75, 0x19, 0x9b, 0x68, 0x4b, 0x3a, 0x6e, 0xe5, 0xed, 0x28, 0x1e, 0xd7, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xbd, 0xd9, 0x4e, 0x2b, 0x01, 0x00, 0x00,
}