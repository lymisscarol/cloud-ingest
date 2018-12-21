// Code generated by protoc-gen-go. DO NOT EDIT.
// source: listfile.proto

package listfile_go_proto

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

// List File Entry specification.
type ListFileEntry struct {
	// Types that are valid to be assigned to Entry:
	//	*ListFileEntry_FileInfo
	//	*ListFileEntry_DirectoryInfo
	Entry                isListFileEntry_Entry `protobuf_oneof:"entry"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListFileEntry) Reset()         { *m = ListFileEntry{} }
func (m *ListFileEntry) String() string { return proto.CompactTextString(m) }
func (*ListFileEntry) ProtoMessage()    {}
func (*ListFileEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_944e22c88393983d, []int{0}
}

func (m *ListFileEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFileEntry.Unmarshal(m, b)
}
func (m *ListFileEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFileEntry.Marshal(b, m, deterministic)
}
func (m *ListFileEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFileEntry.Merge(m, src)
}
func (m *ListFileEntry) XXX_Size() int {
	return xxx_messageInfo_ListFileEntry.Size(m)
}
func (m *ListFileEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFileEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ListFileEntry proto.InternalMessageInfo

type isListFileEntry_Entry interface {
	isListFileEntry_Entry()
}

type ListFileEntry_FileInfo struct {
	FileInfo *FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3,oneof"`
}

type ListFileEntry_DirectoryInfo struct {
	DirectoryInfo *DirectoryInfo `protobuf:"bytes,2,opt,name=directory_info,json=directoryInfo,proto3,oneof"`
}

func (*ListFileEntry_FileInfo) isListFileEntry_Entry() {}

func (*ListFileEntry_DirectoryInfo) isListFileEntry_Entry() {}

func (m *ListFileEntry) GetEntry() isListFileEntry_Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *ListFileEntry) GetFileInfo() *FileInfo {
	if x, ok := m.GetEntry().(*ListFileEntry_FileInfo); ok {
		return x.FileInfo
	}
	return nil
}

func (m *ListFileEntry) GetDirectoryInfo() *DirectoryInfo {
	if x, ok := m.GetEntry().(*ListFileEntry_DirectoryInfo); ok {
		return x.DirectoryInfo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListFileEntry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListFileEntry_FileInfo)(nil),
		(*ListFileEntry_DirectoryInfo)(nil),
	}
}

// Represents a single file’s metadata.
type FileInfo struct {
	// Full path of the file in the format used by the local OS.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Last modified time of the file in seconds since the epoch.
	LastModifiedTime int64 `protobuf:"varint,2,opt,name=last_modified_time,json=lastModifiedTime,proto3" json:"last_modified_time,omitempty"`
	// The size of the file in bytes.
	Size                 int64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_944e22c88393983d, []int{1}
}

func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileInfo) GetLastModifiedTime() int64 {
	if m != nil {
		return m.LastModifiedTime
	}
	return 0
}

func (m *FileInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// Represents a single directory's metadata.
type DirectoryInfo struct {
	// The full path of the directory in the format used by the local OS.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectoryInfo) Reset()         { *m = DirectoryInfo{} }
func (m *DirectoryInfo) String() string { return proto.CompactTextString(m) }
func (*DirectoryInfo) ProtoMessage()    {}
func (*DirectoryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_944e22c88393983d, []int{2}
}

func (m *DirectoryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectoryInfo.Unmarshal(m, b)
}
func (m *DirectoryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectoryInfo.Marshal(b, m, deterministic)
}
func (m *DirectoryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryInfo.Merge(m, src)
}
func (m *DirectoryInfo) XXX_Size() int {
	return xxx_messageInfo_DirectoryInfo.Size(m)
}
func (m *DirectoryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryInfo proto.InternalMessageInfo

func (m *DirectoryInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*ListFileEntry)(nil), "cloud_ingest_listfile.ListFileEntry")
	proto.RegisterType((*FileInfo)(nil), "cloud_ingest_listfile.FileInfo")
	proto.RegisterType((*DirectoryInfo)(nil), "cloud_ingest_listfile.DirectoryInfo")
}

func init() { proto.RegisterFile("listfile.proto", fileDescriptor_944e22c88393983d) }

var fileDescriptor_944e22c88393983d = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xc9, 0x2c, 0x2e,
	0x49, 0xcb, 0xcc, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0xce, 0xc9, 0x2f,
	0x4d, 0x89, 0xcf, 0xcc, 0x4b, 0x4f, 0x2d, 0x2e, 0x89, 0x87, 0x49, 0x2a, 0x2d, 0x67, 0xe4, 0xe2,
	0xf5, 0xc9, 0x2c, 0x2e, 0x71, 0xcb, 0xcc, 0x49, 0x75, 0xcd, 0x2b, 0x29, 0xaa, 0x14, 0xb2, 0xe3,
	0xe2, 0x04, 0xc9, 0xc4, 0x67, 0xe6, 0xa5, 0xe5, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xc9,
	0xeb, 0x61, 0xd5, 0xac, 0x07, 0xd2, 0xe4, 0x99, 0x97, 0x96, 0xef, 0xc1, 0x10, 0xc4, 0x91, 0x06,
	0x65, 0x0b, 0xf9, 0x72, 0xf1, 0xa5, 0x64, 0x16, 0xa5, 0x26, 0x97, 0xe4, 0x17, 0x55, 0x42, 0x0c,
	0x61, 0x02, 0x1b, 0xa2, 0x82, 0xc3, 0x10, 0x17, 0x98, 0x62, 0xa8, 0x49, 0xbc, 0x29, 0xc8, 0x02,
	0x4e, 0xec, 0x5c, 0xac, 0xa9, 0x20, 0x77, 0x29, 0x25, 0x70, 0x71, 0xc0, 0xec, 0x13, 0x12, 0xe2,
	0x62, 0x29, 0x48, 0x2c, 0xc9, 0x00, 0x3b, 0x8f, 0x33, 0x08, 0xcc, 0x16, 0xd2, 0xe1, 0x12, 0xca,
	0x49, 0x2c, 0x2e, 0x89, 0xcf, 0xcd, 0x4f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4d, 0x89, 0x2f, 0xc9, 0xcc,
	0x4d, 0x05, 0xdb, 0xcd, 0x1c, 0x24, 0x00, 0x92, 0xf1, 0x85, 0x4a, 0x84, 0x64, 0xe6, 0xa6, 0x82,
	0x4c, 0x28, 0xce, 0xac, 0x4a, 0x95, 0x60, 0x06, 0xcb, 0x83, 0xd9, 0x4a, 0xca, 0x5c, 0xbc, 0x28,
	0x8e, 0xc1, 0x66, 0x8d, 0x93, 0x70, 0x94, 0x20, 0xcc, 0xe9, 0xf1, 0xe9, 0xf9, 0xf1, 0xe0, 0xc0,
	0x4d, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xf8, 0xaa, 0x68, 0x75,
	0x01, 0x00, 0x00,
}
