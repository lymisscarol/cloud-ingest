// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package task_go_proto

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

// Specifies the task operation that a task performs.
type Type int32

const (
	Type_UNSET_TYPE   Type = 0
	Type_LIST         Type = 1
	Type_PROCESS_LIST Type = 2
	Type_COPY         Type = 3
)

var Type_name = map[int32]string{
	0: "UNSET_TYPE",
	1: "LIST",
	2: "PROCESS_LIST",
	3: "COPY",
}

var Type_value = map[string]int32{
	"UNSET_TYPE":   0,
	"LIST":         1,
	"PROCESS_LIST": 2,
	"COPY":         3,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

// Specifies the current status of the cloud ingest task.
type Status int32

const (
	Status_UNSET_STATUS Status = 0
	Status_QUEUED       Status = 1
	Status_FAILED       Status = 2
	Status_SUCCESS      Status = 3
)

var Status_name = map[int32]string{
	0: "UNSET_STATUS",
	1: "QUEUED",
	2: "FAILED",
	3: "SUCCESS",
}

var Status_value = map[string]int32{
	"UNSET_STATUS": 0,
	"QUEUED":       1,
	"FAILED":       2,
	"SUCCESS":      3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

// Specifies the classes of failures that a task can have.
type FailureType int32

const (
	FailureType_UNSET_FAILURE_TYPE     FailureType = 0
	FailureType_UNKNOWN_FAILURE        FailureType = 1
	FailureType_FILE_MODIFIED_FAILURE  FailureType = 2
	FailureType_HASH_MISMATCH_FAILURE  FailureType = 3
	FailureType_PRECONDITION_FAILURE   FailureType = 4
	FailureType_FILE_NOT_FOUND_FAILURE FailureType = 5
	FailureType_PERMISSION_FAILURE     FailureType = 6
	FailureType_NOT_ACTIVE_JOBRUN      FailureType = 7
)

var FailureType_name = map[int32]string{
	0: "UNSET_FAILURE_TYPE",
	1: "UNKNOWN_FAILURE",
	2: "FILE_MODIFIED_FAILURE",
	3: "HASH_MISMATCH_FAILURE",
	4: "PRECONDITION_FAILURE",
	5: "FILE_NOT_FOUND_FAILURE",
	6: "PERMISSION_FAILURE",
	7: "NOT_ACTIVE_JOBRUN",
}

var FailureType_value = map[string]int32{
	"UNSET_FAILURE_TYPE":     0,
	"UNKNOWN_FAILURE":        1,
	"FILE_MODIFIED_FAILURE":  2,
	"HASH_MISMATCH_FAILURE":  3,
	"PRECONDITION_FAILURE":   4,
	"FILE_NOT_FOUND_FAILURE": 5,
	"PERMISSION_FAILURE":     6,
	"NOT_ACTIVE_JOBRUN":      7,
}

func (x FailureType) String() string {
	return proto.EnumName(FailureType_name, int32(x))
}

func (FailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

// Contains information about a task. A task is a unit of work, one of:
// 1) listing the contents of a single directory
// 2) processing a list file
// 3) copying a single file
// Tasks might be incremental and require multiple request-response round trips
// to complete.
type Spec struct {
	// Types that are valid to be assigned to Spec:
	//	*Spec_ListSpec
	//	*Spec_ProcessListSpec
	//	*Spec_CopySpec
	Spec                 isSpec_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec.Unmarshal(m, b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return xxx_messageInfo_Spec.Size(m)
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

type isSpec_Spec interface {
	isSpec_Spec()
}

type Spec_ListSpec struct {
	ListSpec *ListSpec `protobuf:"bytes,1,opt,name=list_spec,json=listSpec,proto3,oneof"`
}

type Spec_ProcessListSpec struct {
	ProcessListSpec *ProcessListSpec `protobuf:"bytes,2,opt,name=process_list_spec,json=processListSpec,proto3,oneof"`
}

type Spec_CopySpec struct {
	CopySpec *CopySpec `protobuf:"bytes,3,opt,name=copy_spec,json=copySpec,proto3,oneof"`
}

func (*Spec_ListSpec) isSpec_Spec() {}

func (*Spec_ProcessListSpec) isSpec_Spec() {}

func (*Spec_CopySpec) isSpec_Spec() {}

func (m *Spec) GetSpec() isSpec_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Spec) GetListSpec() *ListSpec {
	if x, ok := m.GetSpec().(*Spec_ListSpec); ok {
		return x.ListSpec
	}
	return nil
}

func (m *Spec) GetProcessListSpec() *ProcessListSpec {
	if x, ok := m.GetSpec().(*Spec_ProcessListSpec); ok {
		return x.ProcessListSpec
	}
	return nil
}

func (m *Spec) GetCopySpec() *CopySpec {
	if x, ok := m.GetSpec().(*Spec_CopySpec); ok {
		return x.CopySpec
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Spec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Spec_OneofMarshaler, _Spec_OneofUnmarshaler, _Spec_OneofSizer, []interface{}{
		(*Spec_ListSpec)(nil),
		(*Spec_ProcessListSpec)(nil),
		(*Spec_CopySpec)(nil),
	}
}

func _Spec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Spec)
	// spec
	switch x := m.Spec.(type) {
	case *Spec_ListSpec:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListSpec); err != nil {
			return err
		}
	case *Spec_ProcessListSpec:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProcessListSpec); err != nil {
			return err
		}
	case *Spec_CopySpec:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CopySpec); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Spec.Spec has unexpected type %T", x)
	}
	return nil
}

func _Spec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Spec)
	switch tag {
	case 1: // spec.list_spec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListSpec)
		err := b.DecodeMessage(msg)
		m.Spec = &Spec_ListSpec{msg}
		return true, err
	case 2: // spec.process_list_spec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProcessListSpec)
		err := b.DecodeMessage(msg)
		m.Spec = &Spec_ProcessListSpec{msg}
		return true, err
	case 3: // spec.copy_spec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CopySpec)
		err := b.DecodeMessage(msg)
		m.Spec = &Spec_CopySpec{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Spec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Spec)
	// spec
	switch x := m.Spec.(type) {
	case *Spec_ListSpec:
		s := proto.Size(x.ListSpec)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Spec_ProcessListSpec:
		s := proto.Size(x.ProcessListSpec)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Spec_CopySpec:
		s := proto.Size(x.CopySpec)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Contains the information about a list task. A list task is responsible for
// listing the contents of a directory.
type ListSpec struct {
	DstListResultBucket   string   `protobuf:"bytes,1,opt,name=dst_list_result_bucket,json=dstListResultBucket,proto3" json:"dst_list_result_bucket,omitempty"`
	DstListResultObject   string   `protobuf:"bytes,2,opt,name=dst_list_result_object,json=dstListResultObject,proto3" json:"dst_list_result_object,omitempty"`
	SrcDirectory          string   `protobuf:"bytes,3,opt,name=src_directory,json=srcDirectory,proto3" json:"src_directory,omitempty"`
	ExpectedGenerationNum int64    `protobuf:"varint,4,opt,name=expected_generation_num,json=expectedGenerationNum,proto3" json:"expected_generation_num,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ListSpec) Reset()         { *m = ListSpec{} }
func (m *ListSpec) String() string { return proto.CompactTextString(m) }
func (*ListSpec) ProtoMessage()    {}
func (*ListSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *ListSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSpec.Unmarshal(m, b)
}
func (m *ListSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSpec.Marshal(b, m, deterministic)
}
func (m *ListSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSpec.Merge(m, src)
}
func (m *ListSpec) XXX_Size() int {
	return xxx_messageInfo_ListSpec.Size(m)
}
func (m *ListSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ListSpec proto.InternalMessageInfo

func (m *ListSpec) GetDstListResultBucket() string {
	if m != nil {
		return m.DstListResultBucket
	}
	return ""
}

func (m *ListSpec) GetDstListResultObject() string {
	if m != nil {
		return m.DstListResultObject
	}
	return ""
}

func (m *ListSpec) GetSrcDirectory() string {
	if m != nil {
		return m.SrcDirectory
	}
	return ""
}

func (m *ListSpec) GetExpectedGenerationNum() int64 {
	if m != nil {
		return m.ExpectedGenerationNum
	}
	return 0
}

// Contains the information about a process list task. A process list task is
// responsible for processing the list file produced by a list task.
type ProcessListSpec struct {
	DstListResultBucket  string   `protobuf:"bytes,1,opt,name=dst_list_result_bucket,json=dstListResultBucket,proto3" json:"dst_list_result_bucket,omitempty"`
	DstListResultObject  string   `protobuf:"bytes,2,opt,name=dst_list_result_object,json=dstListResultObject,proto3" json:"dst_list_result_object,omitempty"`
	SrcDirectory         string   `protobuf:"bytes,3,opt,name=src_directory,json=srcDirectory,proto3" json:"src_directory,omitempty"`
	ByteOffset           int64    `protobuf:"varint,4,opt,name=byte_offset,json=byteOffset,proto3" json:"byte_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessListSpec) Reset()         { *m = ProcessListSpec{} }
func (m *ProcessListSpec) String() string { return proto.CompactTextString(m) }
func (*ProcessListSpec) ProtoMessage()    {}
func (*ProcessListSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *ProcessListSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessListSpec.Unmarshal(m, b)
}
func (m *ProcessListSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessListSpec.Marshal(b, m, deterministic)
}
func (m *ProcessListSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessListSpec.Merge(m, src)
}
func (m *ProcessListSpec) XXX_Size() int {
	return xxx_messageInfo_ProcessListSpec.Size(m)
}
func (m *ProcessListSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessListSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessListSpec proto.InternalMessageInfo

func (m *ProcessListSpec) GetDstListResultBucket() string {
	if m != nil {
		return m.DstListResultBucket
	}
	return ""
}

func (m *ProcessListSpec) GetDstListResultObject() string {
	if m != nil {
		return m.DstListResultObject
	}
	return ""
}

func (m *ProcessListSpec) GetSrcDirectory() string {
	if m != nil {
		return m.SrcDirectory
	}
	return ""
}

func (m *ProcessListSpec) GetByteOffset() int64 {
	if m != nil {
		return m.ByteOffset
	}
	return 0
}

// Contains the information about a copy task. A copy task is responsible for
// copying a single file.
type CopySpec struct {
	SrcFile               string `protobuf:"bytes,1,opt,name=src_file,json=srcFile,proto3" json:"src_file,omitempty"`
	DstBucket             string `protobuf:"bytes,2,opt,name=dst_bucket,json=dstBucket,proto3" json:"dst_bucket,omitempty"`
	DstObject             string `protobuf:"bytes,3,opt,name=dst_object,json=dstObject,proto3" json:"dst_object,omitempty"`
	ExpectedGenerationNum int64  `protobuf:"varint,4,opt,name=expected_generation_num,json=expectedGenerationNum,proto3" json:"expected_generation_num,omitempty"`
	// Fields for bandwidth management.
	Bandwidth int64 `protobuf:"varint,5,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Fields only for managing resumable copies.
	FileBytes            int64    `protobuf:"varint,6,opt,name=file_bytes,json=fileBytes,proto3" json:"file_bytes,omitempty"`
	FileMTime            int64    `protobuf:"varint,7,opt,name=file_m_time,json=fileMTime,proto3" json:"file_m_time,omitempty"`
	BytesCopied          int64    `protobuf:"varint,8,opt,name=bytes_copied,json=bytesCopied,proto3" json:"bytes_copied,omitempty"`
	Crc32C               uint32   `protobuf:"varint,9,opt,name=crc32c,proto3" json:"crc32c,omitempty"`
	BytesToCopy          int64    `protobuf:"varint,10,opt,name=bytes_to_copy,json=bytesToCopy,proto3" json:"bytes_to_copy,omitempty"`
	ResumableUploadId    string   `protobuf:"bytes,11,opt,name=resumable_upload_id,json=resumableUploadId,proto3" json:"resumable_upload_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopySpec) Reset()         { *m = CopySpec{} }
func (m *CopySpec) String() string { return proto.CompactTextString(m) }
func (*CopySpec) ProtoMessage()    {}
func (*CopySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{3}
}

func (m *CopySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopySpec.Unmarshal(m, b)
}
func (m *CopySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopySpec.Marshal(b, m, deterministic)
}
func (m *CopySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopySpec.Merge(m, src)
}
func (m *CopySpec) XXX_Size() int {
	return xxx_messageInfo_CopySpec.Size(m)
}
func (m *CopySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CopySpec.DiscardUnknown(m)
}

var xxx_messageInfo_CopySpec proto.InternalMessageInfo

func (m *CopySpec) GetSrcFile() string {
	if m != nil {
		return m.SrcFile
	}
	return ""
}

func (m *CopySpec) GetDstBucket() string {
	if m != nil {
		return m.DstBucket
	}
	return ""
}

func (m *CopySpec) GetDstObject() string {
	if m != nil {
		return m.DstObject
	}
	return ""
}

func (m *CopySpec) GetExpectedGenerationNum() int64 {
	if m != nil {
		return m.ExpectedGenerationNum
	}
	return 0
}

func (m *CopySpec) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *CopySpec) GetFileBytes() int64 {
	if m != nil {
		return m.FileBytes
	}
	return 0
}

func (m *CopySpec) GetFileMTime() int64 {
	if m != nil {
		return m.FileMTime
	}
	return 0
}

func (m *CopySpec) GetBytesCopied() int64 {
	if m != nil {
		return m.BytesCopied
	}
	return 0
}

func (m *CopySpec) GetCrc32C() uint32 {
	if m != nil {
		return m.Crc32C
	}
	return 0
}

func (m *CopySpec) GetBytesToCopy() int64 {
	if m != nil {
		return m.BytesToCopy
	}
	return 0
}

func (m *CopySpec) GetResumableUploadId() string {
	if m != nil {
		return m.ResumableUploadId
	}
	return ""
}

// Contains the message sent from the DCP to an Agent to issue a task request.
type TaskReqMsg struct {
	TaskRelRsrcName      string   `protobuf:"bytes,1,opt,name=task_rel_rsrc_name,json=taskRelRsrcName,proto3" json:"task_rel_rsrc_name,omitempty"`
	JobrunRelRsrcName    string   `protobuf:"bytes,3,opt,name=jobrun_rel_rsrc_name,json=jobrunRelRsrcName,proto3" json:"jobrun_rel_rsrc_name,omitempty"`
	Spec                 *Spec    `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskReqMsg) Reset()         { *m = TaskReqMsg{} }
func (m *TaskReqMsg) String() string { return proto.CompactTextString(m) }
func (*TaskReqMsg) ProtoMessage()    {}
func (*TaskReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{4}
}

func (m *TaskReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskReqMsg.Unmarshal(m, b)
}
func (m *TaskReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskReqMsg.Marshal(b, m, deterministic)
}
func (m *TaskReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskReqMsg.Merge(m, src)
}
func (m *TaskReqMsg) XXX_Size() int {
	return xxx_messageInfo_TaskReqMsg.Size(m)
}
func (m *TaskReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TaskReqMsg proto.InternalMessageInfo

func (m *TaskReqMsg) GetTaskRelRsrcName() string {
	if m != nil {
		return m.TaskRelRsrcName
	}
	return ""
}

func (m *TaskReqMsg) GetJobrunRelRsrcName() string {
	if m != nil {
		return m.JobrunRelRsrcName
	}
	return ""
}

func (m *TaskReqMsg) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Contains the message sent from the Agent to the DCP in response to a task
// request.
type TaskRespMsg struct {
	TaskRelRsrcName      string      `protobuf:"bytes,1,opt,name=task_rel_rsrc_name,json=taskRelRsrcName,proto3" json:"task_rel_rsrc_name,omitempty"`
	Status               string      `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	FailureType          FailureType `protobuf:"varint,3,opt,name=failure_type,json=failureType,proto3,enum=cloud_ingest_task.FailureType" json:"failure_type,omitempty"`
	FailureMessage       string      `protobuf:"bytes,4,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	Log                  *Log        `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
	ReqSpec              *Spec       `protobuf:"bytes,6,opt,name=req_spec,json=reqSpec,proto3" json:"req_spec,omitempty"`
	RespSpec             *Spec       `protobuf:"bytes,7,opt,name=resp_spec,json=respSpec,proto3" json:"resp_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaskRespMsg) Reset()         { *m = TaskRespMsg{} }
func (m *TaskRespMsg) String() string { return proto.CompactTextString(m) }
func (*TaskRespMsg) ProtoMessage()    {}
func (*TaskRespMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{5}
}

func (m *TaskRespMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRespMsg.Unmarshal(m, b)
}
func (m *TaskRespMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRespMsg.Marshal(b, m, deterministic)
}
func (m *TaskRespMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRespMsg.Merge(m, src)
}
func (m *TaskRespMsg) XXX_Size() int {
	return xxx_messageInfo_TaskRespMsg.Size(m)
}
func (m *TaskRespMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRespMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRespMsg proto.InternalMessageInfo

func (m *TaskRespMsg) GetTaskRelRsrcName() string {
	if m != nil {
		return m.TaskRelRsrcName
	}
	return ""
}

func (m *TaskRespMsg) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskRespMsg) GetFailureType() FailureType {
	if m != nil {
		return m.FailureType
	}
	return FailureType_UNSET_FAILURE_TYPE
}

func (m *TaskRespMsg) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *TaskRespMsg) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *TaskRespMsg) GetReqSpec() *Spec {
	if m != nil {
		return m.ReqSpec
	}
	return nil
}

func (m *TaskRespMsg) GetRespSpec() *Spec {
	if m != nil {
		return m.RespSpec
	}
	return nil
}

// Contains log information for a task. This message is suitable for the "Log"
// field in the LogEntries Spanner queue. Note that this info is eventually
// dumped into the user's GCS bucket.
type Log struct {
	// Types that are valid to be assigned to Log:
	//	*Log_ListLog
	//	*Log_ProcessListLog
	//	*Log_CopyLog
	Log                  isLog_Log `protobuf_oneof:"log"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{6}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

type isLog_Log interface {
	isLog_Log()
}

type Log_ListLog struct {
	ListLog *ListLog `protobuf:"bytes,1,opt,name=list_log,json=listLog,proto3,oneof"`
}

type Log_ProcessListLog struct {
	ProcessListLog *ProcessListLog `protobuf:"bytes,2,opt,name=process_list_log,json=processListLog,proto3,oneof"`
}

type Log_CopyLog struct {
	CopyLog *CopyLog `protobuf:"bytes,3,opt,name=copy_log,json=copyLog,proto3,oneof"`
}

func (*Log_ListLog) isLog_Log() {}

func (*Log_ProcessListLog) isLog_Log() {}

func (*Log_CopyLog) isLog_Log() {}

func (m *Log) GetLog() isLog_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Log) GetListLog() *ListLog {
	if x, ok := m.GetLog().(*Log_ListLog); ok {
		return x.ListLog
	}
	return nil
}

func (m *Log) GetProcessListLog() *ProcessListLog {
	if x, ok := m.GetLog().(*Log_ProcessListLog); ok {
		return x.ProcessListLog
	}
	return nil
}

func (m *Log) GetCopyLog() *CopyLog {
	if x, ok := m.GetLog().(*Log_CopyLog); ok {
		return x.CopyLog
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Log) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Log_OneofMarshaler, _Log_OneofUnmarshaler, _Log_OneofSizer, []interface{}{
		(*Log_ListLog)(nil),
		(*Log_ProcessListLog)(nil),
		(*Log_CopyLog)(nil),
	}
}

func _Log_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Log)
	// log
	switch x := m.Log.(type) {
	case *Log_ListLog:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListLog); err != nil {
			return err
		}
	case *Log_ProcessListLog:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProcessListLog); err != nil {
			return err
		}
	case *Log_CopyLog:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CopyLog); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Log.Log has unexpected type %T", x)
	}
	return nil
}

func _Log_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Log)
	switch tag {
	case 1: // log.list_log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListLog)
		err := b.DecodeMessage(msg)
		m.Log = &Log_ListLog{msg}
		return true, err
	case 2: // log.process_list_log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProcessListLog)
		err := b.DecodeMessage(msg)
		m.Log = &Log_ProcessListLog{msg}
		return true, err
	case 3: // log.copy_log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CopyLog)
		err := b.DecodeMessage(msg)
		m.Log = &Log_CopyLog{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Log_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Log)
	// log
	switch x := m.Log.(type) {
	case *Log_ListLog:
		s := proto.Size(x.ListLog)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Log_ProcessListLog:
		s := proto.Size(x.ProcessListLog)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Log_CopyLog:
		s := proto.Size(x.CopyLog)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Contains log fields for a List task.
type ListLog struct {
	FilesFound           int64    `protobuf:"varint,1,opt,name=files_found,json=filesFound,proto3" json:"files_found,omitempty"`
	BytesFound           int64    `protobuf:"varint,2,opt,name=bytes_found,json=bytesFound,proto3" json:"bytes_found,omitempty"`
	DirsFound            int64    `protobuf:"varint,3,opt,name=dirs_found,json=dirsFound,proto3" json:"dirs_found,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLog) Reset()         { *m = ListLog{} }
func (m *ListLog) String() string { return proto.CompactTextString(m) }
func (*ListLog) ProtoMessage()    {}
func (*ListLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{7}
}

func (m *ListLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLog.Unmarshal(m, b)
}
func (m *ListLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLog.Marshal(b, m, deterministic)
}
func (m *ListLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLog.Merge(m, src)
}
func (m *ListLog) XXX_Size() int {
	return xxx_messageInfo_ListLog.Size(m)
}
func (m *ListLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLog.DiscardUnknown(m)
}

var xxx_messageInfo_ListLog proto.InternalMessageInfo

func (m *ListLog) GetFilesFound() int64 {
	if m != nil {
		return m.FilesFound
	}
	return 0
}

func (m *ListLog) GetBytesFound() int64 {
	if m != nil {
		return m.BytesFound
	}
	return 0
}

func (m *ListLog) GetDirsFound() int64 {
	if m != nil {
		return m.DirsFound
	}
	return 0
}

// Contains log fields for a ProcessList task.
type ProcessListLog struct {
	EntriesProcessed     int64    `protobuf:"varint,1,opt,name=entries_processed,json=entriesProcessed,proto3" json:"entries_processed,omitempty"`
	StartingOffset       int64    `protobuf:"varint,2,opt,name=starting_offset,json=startingOffset,proto3" json:"starting_offset,omitempty"`
	EndingOffset         int64    `protobuf:"varint,3,opt,name=ending_offset,json=endingOffset,proto3" json:"ending_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessListLog) Reset()         { *m = ProcessListLog{} }
func (m *ProcessListLog) String() string { return proto.CompactTextString(m) }
func (*ProcessListLog) ProtoMessage()    {}
func (*ProcessListLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{8}
}

func (m *ProcessListLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessListLog.Unmarshal(m, b)
}
func (m *ProcessListLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessListLog.Marshal(b, m, deterministic)
}
func (m *ProcessListLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessListLog.Merge(m, src)
}
func (m *ProcessListLog) XXX_Size() int {
	return xxx_messageInfo_ProcessListLog.Size(m)
}
func (m *ProcessListLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessListLog.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessListLog proto.InternalMessageInfo

func (m *ProcessListLog) GetEntriesProcessed() int64 {
	if m != nil {
		return m.EntriesProcessed
	}
	return 0
}

func (m *ProcessListLog) GetStartingOffset() int64 {
	if m != nil {
		return m.StartingOffset
	}
	return 0
}

func (m *ProcessListLog) GetEndingOffset() int64 {
	if m != nil {
		return m.EndingOffset
	}
	return 0
}

// Contains log fields for a Copy task.
type CopyLog struct {
	SrcFile              string   `protobuf:"bytes,1,opt,name=src_file,json=srcFile,proto3" json:"src_file,omitempty"`
	SrcBytes             int64    `protobuf:"varint,2,opt,name=src_bytes,json=srcBytes,proto3" json:"src_bytes,omitempty"`
	SrcMTime             int64    `protobuf:"varint,3,opt,name=src_m_time,json=srcMTime,proto3" json:"src_m_time,omitempty"`
	SrcCrc32C            uint32   `protobuf:"varint,4,opt,name=src_crc32c,json=srcCrc32c,proto3" json:"src_crc32c,omitempty"`
	DstFile              string   `protobuf:"bytes,5,opt,name=dst_file,json=dstFile,proto3" json:"dst_file,omitempty"`
	DstBytes             int64    `protobuf:"varint,6,opt,name=dst_bytes,json=dstBytes,proto3" json:"dst_bytes,omitempty"`
	DstMTime             int64    `protobuf:"varint,7,opt,name=dst_m_time,json=dstMTime,proto3" json:"dst_m_time,omitempty"`
	DstCrc32C            uint32   `protobuf:"varint,8,opt,name=dst_crc32c,json=dstCrc32c,proto3" json:"dst_crc32c,omitempty"`
	BytesCopied          int64    `protobuf:"varint,9,opt,name=bytes_copied,json=bytesCopied,proto3" json:"bytes_copied,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopyLog) Reset()         { *m = CopyLog{} }
func (m *CopyLog) String() string { return proto.CompactTextString(m) }
func (*CopyLog) ProtoMessage()    {}
func (*CopyLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{9}
}

func (m *CopyLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyLog.Unmarshal(m, b)
}
func (m *CopyLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyLog.Marshal(b, m, deterministic)
}
func (m *CopyLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyLog.Merge(m, src)
}
func (m *CopyLog) XXX_Size() int {
	return xxx_messageInfo_CopyLog.Size(m)
}
func (m *CopyLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyLog.DiscardUnknown(m)
}

var xxx_messageInfo_CopyLog proto.InternalMessageInfo

func (m *CopyLog) GetSrcFile() string {
	if m != nil {
		return m.SrcFile
	}
	return ""
}

func (m *CopyLog) GetSrcBytes() int64 {
	if m != nil {
		return m.SrcBytes
	}
	return 0
}

func (m *CopyLog) GetSrcMTime() int64 {
	if m != nil {
		return m.SrcMTime
	}
	return 0
}

func (m *CopyLog) GetSrcCrc32C() uint32 {
	if m != nil {
		return m.SrcCrc32C
	}
	return 0
}

func (m *CopyLog) GetDstFile() string {
	if m != nil {
		return m.DstFile
	}
	return ""
}

func (m *CopyLog) GetDstBytes() int64 {
	if m != nil {
		return m.DstBytes
	}
	return 0
}

func (m *CopyLog) GetDstMTime() int64 {
	if m != nil {
		return m.DstMTime
	}
	return 0
}

func (m *CopyLog) GetDstCrc32C() uint32 {
	if m != nil {
		return m.DstCrc32C
	}
	return 0
}

func (m *CopyLog) GetBytesCopied() int64 {
	if m != nil {
		return m.BytesCopied
	}
	return 0
}

func init() {
	proto.RegisterEnum("cloud_ingest_task.Type", Type_name, Type_value)
	proto.RegisterEnum("cloud_ingest_task.Status", Status_name, Status_value)
	proto.RegisterEnum("cloud_ingest_task.FailureType", FailureType_name, FailureType_value)
	proto.RegisterType((*Spec)(nil), "cloud_ingest_task.Spec")
	proto.RegisterType((*ListSpec)(nil), "cloud_ingest_task.ListSpec")
	proto.RegisterType((*ProcessListSpec)(nil), "cloud_ingest_task.ProcessListSpec")
	proto.RegisterType((*CopySpec)(nil), "cloud_ingest_task.CopySpec")
	proto.RegisterType((*TaskReqMsg)(nil), "cloud_ingest_task.TaskReqMsg")
	proto.RegisterType((*TaskRespMsg)(nil), "cloud_ingest_task.TaskRespMsg")
	proto.RegisterType((*Log)(nil), "cloud_ingest_task.Log")
	proto.RegisterType((*ListLog)(nil), "cloud_ingest_task.ListLog")
	proto.RegisterType((*ProcessListLog)(nil), "cloud_ingest_task.ProcessListLog")
	proto.RegisterType((*CopyLog)(nil), "cloud_ingest_task.CopyLog")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 1165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x5d, 0x8f, 0xdb, 0x44,
	0x17, 0x5e, 0xc7, 0xd9, 0x7c, 0x9c, 0xec, 0x6e, 0x9c, 0x69, 0x9b, 0xa6, 0xdf, 0x6d, 0xde, 0x8b,
	0x77, 0xd5, 0x4a, 0x8b, 0xb4, 0x45, 0x20, 0x21, 0x24, 0x94, 0x0f, 0x87, 0x0d, 0xe4, 0x0b, 0xdb,
	0x01, 0x95, 0x9b, 0x91, 0x63, 0x4f, 0x82, 0x5b, 0x27, 0x76, 0x3d, 0x13, 0x41, 0x7e, 0x01, 0x17,
	0xdc, 0xc3, 0xef, 0xe1, 0x07, 0x70, 0x01, 0x12, 0x12, 0x3f, 0x07, 0x9d, 0x19, 0x3b, 0xc9, 0x6e,
	0xb7, 0x2b, 0xc1, 0x15, 0x57, 0x19, 0x9f, 0xe7, 0x9c, 0xe7, 0x3c, 0x67, 0x3e, 0x1e, 0x05, 0x40,
	0xb8, 0xfc, 0xcd, 0x59, 0x9c, 0x44, 0x22, 0x22, 0x35, 0x2f, 0x8c, 0xd6, 0x3e, 0x0d, 0x56, 0x0b,
	0xc6, 0x05, 0x45, 0xa0, 0xf9, 0x97, 0x06, 0x79, 0x3b, 0x66, 0x1e, 0xf9, 0x04, 0xca, 0x61, 0xc0,
	0x05, 0xe5, 0x31, 0xf3, 0x1a, 0xda, 0x53, 0xed, 0xb4, 0x72, 0xfe, 0xe0, 0xec, 0x9d, 0xfc, 0xb3,
	0x41, 0xc0, 0x05, 0xe6, 0x5f, 0x1c, 0x58, 0xa5, 0x30, 0x5d, 0x93, 0x09, 0xd4, 0xe2, 0x24, 0xf2,
	0x18, 0xe7, 0x74, 0xc7, 0x91, 0x93, 0x1c, 0xcd, 0x6b, 0x38, 0x26, 0x2a, 0x77, 0x8f, 0xaa, 0x1a,
	0x5f, 0x0e, 0xa1, 0x1a, 0x2f, 0x8a, 0x37, 0x8a, 0x49, 0x7f, 0xaf, 0x9a, 0x4e, 0x14, 0x6f, 0x32,
	0x35, 0x5e, 0xba, 0x6e, 0x17, 0x20, 0x8f, 0x65, 0xcd, 0xdf, 0x35, 0x28, 0x6d, 0x09, 0x5f, 0x42,
	0xdd, 0xe7, 0x42, 0xc9, 0x4b, 0x18, 0x5f, 0x87, 0x82, 0xce, 0xd6, 0xde, 0x1b, 0x26, 0xe4, 0xac,
	0x65, 0xeb, 0x96, 0xcf, 0x05, 0x26, 0x5b, 0x12, 0x6b, 0x4b, 0xe8, 0xba, 0xa2, 0x68, 0xf6, 0x9a,
	0x79, 0x42, 0x0e, 0x77, 0xb5, 0x68, 0x2c, 0x21, 0xf2, 0x3f, 0x38, 0xe6, 0x89, 0x47, 0xfd, 0x20,
	0x61, 0x9e, 0x88, 0x92, 0x8d, 0x94, 0x5f, 0xb6, 0x8e, 0x78, 0xe2, 0x75, 0xb3, 0x18, 0xf9, 0x08,
	0xee, 0xb2, 0x1f, 0x62, 0xe6, 0x09, 0xe6, 0xd3, 0x05, 0x5b, 0xb1, 0xc4, 0x15, 0x41, 0xb4, 0xa2,
	0xab, 0xf5, 0xb2, 0x91, 0x7f, 0xaa, 0x9d, 0xea, 0xd6, 0x9d, 0x0c, 0xfe, 0x7c, 0x8b, 0x8e, 0xd6,
	0xcb, 0xe6, 0xaf, 0x1a, 0x54, 0xaf, 0x6c, 0xdf, 0x7f, 0x6d, 0xb4, 0x27, 0x50, 0x99, 0x6d, 0x04,
	0xa3, 0xd1, 0x7c, 0xce, 0x99, 0x48, 0xc7, 0x01, 0x0c, 0x8d, 0x65, 0xa4, 0xf9, 0xa3, 0x0e, 0xa5,
	0xec, 0xe0, 0xc8, 0x3d, 0x28, 0x21, 0xe5, 0x3c, 0x08, 0x59, 0x2a, 0xb7, 0xc8, 0x13, 0xaf, 0x17,
	0x84, 0x8c, 0x3c, 0x02, 0x40, 0x89, 0xe9, 0x2c, 0x4a, 0x56, 0xd9, 0xe7, 0xd9, 0x04, 0x29, 0x9c,
	0xaa, 0xd6, 0xb7, 0x70, 0xaa, 0xf5, 0x5f, 0xee, 0x30, 0x79, 0x08, 0xe5, 0x99, 0xbb, 0xf2, 0xbf,
	0x0f, 0x7c, 0xf1, 0x5d, 0xe3, 0x50, 0x66, 0xee, 0x02, 0xd8, 0x14, 0xa5, 0x52, 0x1c, 0x87, 0x37,
	0x0a, 0x0a, 0xc6, 0x48, 0x1b, 0x03, 0xe4, 0x31, 0x54, 0x24, 0xbc, 0xa4, 0x22, 0x58, 0xb2, 0x46,
	0x71, 0x87, 0x0f, 0x9d, 0x60, 0xc9, 0xc8, 0x33, 0x38, 0x92, 0x95, 0xd4, 0x8b, 0xe2, 0x80, 0xf9,
	0x8d, 0x92, 0x4c, 0x90, 0xfb, 0xc5, 0x3b, 0x32, 0x44, 0xea, 0x50, 0xf0, 0x12, 0xef, 0xe5, 0xb9,
	0xd7, 0x28, 0x3f, 0xd5, 0x4e, 0x8f, 0xad, 0xf4, 0x8b, 0x34, 0xe1, 0x58, 0x95, 0x8a, 0x08, 0xab,
	0x37, 0x0d, 0xd8, 0xab, 0x75, 0x22, 0xdc, 0x50, 0x72, 0x06, 0xb7, 0xf0, 0x2c, 0x97, 0xee, 0x2c,
	0x64, 0x74, 0x1d, 0x87, 0x91, 0xeb, 0xd3, 0xc0, 0x6f, 0x54, 0xe4, 0xde, 0xd4, 0xb6, 0xd0, 0x54,
	0x22, 0x7d, 0xbf, 0xf9, 0x8b, 0x06, 0xe0, 0xb8, 0xfc, 0x8d, 0xc5, 0xde, 0x0e, 0xf9, 0x82, 0xbc,
	0x00, 0x82, 0xaf, 0x8a, 0x26, 0x2c, 0xa4, 0x09, 0x9e, 0xca, 0xca, 0x5d, 0x66, 0xa7, 0x52, 0x15,
	0x32, 0x2f, 0xb4, 0x78, 0xe2, 0x8d, 0xdc, 0x25, 0x23, 0x1f, 0xc0, 0xed, 0xd7, 0xd1, 0x2c, 0x59,
	0xaf, 0xae, 0xa4, 0xab, 0x83, 0xa8, 0x29, 0x6c, 0xbf, 0xe0, 0x85, 0x7a, 0x96, 0xa9, 0x2f, 0xdc,
	0xbd, 0xe6, 0x35, 0xe3, 0x85, 0xb0, 0xd4, 0xdb, 0xfd, 0x23, 0x07, 0x15, 0xa5, 0x8c, 0xc7, 0xff,
	0x58, 0x5a, 0x1d, 0x0a, 0x5c, 0xb8, 0x62, 0xcd, 0xd3, 0x4b, 0x93, 0x7e, 0x91, 0x16, 0x1c, 0xcd,
	0xdd, 0x20, 0x5c, 0x27, 0x8c, 0x8a, 0x4d, 0xac, 0xa4, 0x9e, 0x9c, 0x3f, 0xbe, 0x46, 0x49, 0x4f,
	0xa5, 0x39, 0x9b, 0x98, 0x59, 0x95, 0xf9, 0xee, 0x83, 0xfc, 0x1f, 0xaa, 0x19, 0xc5, 0x92, 0x71,
	0xee, 0x2e, 0x98, 0xbc, 0x4d, 0x65, 0xeb, 0x24, 0x0d, 0x0f, 0x55, 0x94, 0x9c, 0x82, 0x1e, 0x46,
	0x0b, 0x79, 0x81, 0x2a, 0xe7, 0xf5, 0xeb, 0x8c, 0x34, 0x5a, 0x58, 0x98, 0x42, 0xce, 0xa1, 0x94,
	0xb0, 0xb7, 0xca, 0xe9, 0x0a, 0x37, 0xef, 0x4d, 0x31, 0x61, 0x6f, 0xe5, 0xab, 0xf9, 0x10, 0xca,
	0x09, 0xe3, 0xb1, 0x2a, 0x2a, 0xde, 0x5c, 0x54, 0xc2, 0x4c, 0x5c, 0x35, 0x7f, 0xd3, 0x40, 0x1f,
	0x44, 0x0b, 0xf2, 0x31, 0x48, 0xeb, 0xa6, 0x28, 0x50, 0x39, 0xfd, 0xfd, 0xf7, 0x38, 0xfd, 0x20,
	0x5a, 0x5c, 0x1c, 0x58, 0xc5, 0x50, 0x2d, 0xc9, 0x10, 0x8c, 0x4b, 0x3e, 0x8f, 0x04, 0xea, 0x38,
	0x9f, 0xdd, 0x6c, 0xf3, 0x8a, 0xe7, 0x24, 0xbe, 0x14, 0x41, 0x1d, 0xd2, 0xe4, 0x91, 0x46, 0x7f,
	0xaf, 0x0e, 0xbc, 0xd9, 0xa9, 0x0e, 0x4f, 0x2d, 0xdb, 0x87, 0x72, 0x73, 0x9b, 0xaf, 0xa1, 0x98,
	0x51, 0x3d, 0x51, 0x0f, 0x8f, 0xd3, 0x79, 0xb4, 0x5e, 0xf9, 0x72, 0x2a, 0xdd, 0x92, 0x4f, 0x95,
	0xf7, 0x30, 0x92, 0xb9, 0x52, 0x96, 0x90, 0xdb, 0xb9, 0x52, 0x9a, 0x80, 0x76, 0x12, 0x24, 0x19,
	0xae, 0xab, 0x97, 0x8b, 0x11, 0x09, 0x37, 0x7f, 0xd2, 0xe0, 0xe4, 0xf2, 0x40, 0xe4, 0x05, 0xd4,
	0xd8, 0x4a, 0x24, 0x01, 0xe3, 0x34, 0x1d, 0x8c, 0x65, 0x9d, 0x8d, 0x14, 0x98, 0x64, 0x71, 0xbc,
	0x38, 0x5c, 0xb8, 0x89, 0x08, 0x56, 0x8b, 0xcc, 0x19, 0x95, 0x86, 0x93, 0x2c, 0xac, 0xdc, 0x11,
	0x3d, 0x96, 0xad, 0xfc, 0xbd, 0x34, 0x25, 0xe5, 0x48, 0x05, 0x53, 0x0b, 0xfd, 0x39, 0x07, 0xc5,
	0x74, 0x5f, 0x6e, 0x72, 0xd0, 0x07, 0x50, 0x46, 0x48, 0x99, 0x95, 0x6a, 0x87, 0xb9, 0xca, 0xab,
	0x1e, 0x02, 0x20, 0x98, 0x5a, 0x95, 0xbe, 0x45, 0x95, 0x53, 0x3d, 0x52, 0x68, 0x6a, 0x45, 0x79,
	0x69, 0x45, 0x48, 0xd6, 0x51, 0x6e, 0x74, 0x0f, 0x4a, 0x68, 0xbe, 0xb2, 0xe9, 0xa1, 0x6a, 0xea,
	0x73, 0x91, 0x35, 0x95, 0xb6, 0xbd, 0xe7, 0x90, 0x98, 0xbb, 0x6d, 0x8a, 0xe0, 0x25, 0x7f, 0x44,
	0x74, 0xdb, 0x14, 0xd1, 0xb4, 0x69, 0x49, 0x35, 0xf5, 0xb9, 0x48, 0x9b, 0x5e, 0x75, 0xcf, 0xf2,
	0x3b, 0xee, 0xf9, 0xfc, 0x53, 0xc8, 0xcb, 0x77, 0x7a, 0x02, 0x30, 0x1d, 0xd9, 0xa6, 0x43, 0x9d,
	0x57, 0x13, 0xd3, 0x38, 0x20, 0x25, 0xc8, 0x0f, 0xfa, 0xb6, 0x63, 0x68, 0xc4, 0x80, 0xa3, 0x89,
	0x35, 0xee, 0x98, 0xb6, 0x4d, 0x65, 0x24, 0x87, 0x58, 0x67, 0x3c, 0x79, 0x65, 0xe8, 0xcf, 0x3f,
	0x83, 0x82, 0xad, 0xac, 0xc2, 0x80, 0x23, 0x55, 0x6f, 0x3b, 0x2d, 0x67, 0x6a, 0x1b, 0x07, 0x04,
	0xa0, 0xf0, 0xd5, 0xd4, 0x9c, 0x9a, 0x5d, 0x43, 0xc3, 0x75, 0xaf, 0xd5, 0x1f, 0x98, 0x5d, 0x23,
	0x47, 0x2a, 0x50, 0xb4, 0xa7, 0x1d, 0xe4, 0x33, 0xf4, 0xe7, 0x7f, 0x6a, 0x50, 0xd9, 0xf3, 0x0e,
	0x52, 0x07, 0xa2, 0x68, 0x30, 0x7d, 0x6a, 0x99, 0x99, 0x9c, 0x5b, 0x50, 0x9d, 0x8e, 0xbe, 0x1c,
	0x8d, 0xbf, 0x19, 0x65, 0x88, 0xa1, 0x91, 0x7b, 0x70, 0xa7, 0xd7, 0x1f, 0x98, 0x74, 0x38, 0xee,
	0xf6, 0x7b, 0x7d, 0xb3, 0xbb, 0x85, 0x72, 0x08, 0x5d, 0xb4, 0xec, 0x0b, 0x3a, 0xec, 0xdb, 0xc3,
	0x96, 0xd3, 0xb9, 0xd8, 0x42, 0x3a, 0x69, 0xc0, 0xed, 0x89, 0x65, 0x76, 0xc6, 0xa3, 0x6e, 0xdf,
	0xe9, 0x8f, 0x77, 0x7c, 0x79, 0x72, 0x1f, 0xea, 0x92, 0x6f, 0x34, 0x76, 0x68, 0x6f, 0x3c, 0x1d,
	0xed, 0x08, 0x0f, 0x51, 0xd8, 0xc4, 0xb4, 0x86, 0x7d, 0xdb, 0xde, 0xaf, 0x29, 0x90, 0x3b, 0x50,
	0xc3, 0xf4, 0x56, 0xc7, 0xe9, 0x7f, 0x6d, 0xd2, 0x2f, 0xc6, 0x6d, 0x6b, 0x3a, 0x32, 0x8a, 0xed,
	0xea, 0xb7, 0xc7, 0xd2, 0x7e, 0x17, 0x11, 0x95, 0xff, 0x24, 0x67, 0x05, 0xf9, 0xf3, 0xf2, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x00, 0xac, 0x2c, 0x5e, 0x0a, 0x00, 0x00,
}
