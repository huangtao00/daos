// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package mgmt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StoragePrepareReq struct {
	Nvme                 *PrepareNvmeReq `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm                  *PrepareScmReq  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StoragePrepareReq) Reset()         { *m = StoragePrepareReq{} }
func (m *StoragePrepareReq) String() string { return proto.CompactTextString(m) }
func (*StoragePrepareReq) ProtoMessage()    {}
func (*StoragePrepareReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{0}
}
func (m *StoragePrepareReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoragePrepareReq.Unmarshal(m, b)
}
func (m *StoragePrepareReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoragePrepareReq.Marshal(b, m, deterministic)
}
func (dst *StoragePrepareReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoragePrepareReq.Merge(dst, src)
}
func (m *StoragePrepareReq) XXX_Size() int {
	return xxx_messageInfo_StoragePrepareReq.Size(m)
}
func (m *StoragePrepareReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StoragePrepareReq.DiscardUnknown(m)
}

var xxx_messageInfo_StoragePrepareReq proto.InternalMessageInfo

func (m *StoragePrepareReq) GetNvme() *PrepareNvmeReq {
	if m != nil {
		return m.Nvme
	}
	return nil
}

func (m *StoragePrepareReq) GetScm() *PrepareScmReq {
	if m != nil {
		return m.Scm
	}
	return nil
}

type StoragePrepareResp struct {
	Nvme                 *PrepareNvmeResp `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm                  *PrepareScmResp  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StoragePrepareResp) Reset()         { *m = StoragePrepareResp{} }
func (m *StoragePrepareResp) String() string { return proto.CompactTextString(m) }
func (*StoragePrepareResp) ProtoMessage()    {}
func (*StoragePrepareResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{1}
}
func (m *StoragePrepareResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoragePrepareResp.Unmarshal(m, b)
}
func (m *StoragePrepareResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoragePrepareResp.Marshal(b, m, deterministic)
}
func (dst *StoragePrepareResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoragePrepareResp.Merge(dst, src)
}
func (m *StoragePrepareResp) XXX_Size() int {
	return xxx_messageInfo_StoragePrepareResp.Size(m)
}
func (m *StoragePrepareResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StoragePrepareResp.DiscardUnknown(m)
}

var xxx_messageInfo_StoragePrepareResp proto.InternalMessageInfo

func (m *StoragePrepareResp) GetNvme() *PrepareNvmeResp {
	if m != nil {
		return m.Nvme
	}
	return nil
}

func (m *StoragePrepareResp) GetScm() *PrepareScmResp {
	if m != nil {
		return m.Scm
	}
	return nil
}

type StorageScanReq struct {
	Nvme                 *ScanNvmeReq `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm                  *ScanScmReq  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StorageScanReq) Reset()         { *m = StorageScanReq{} }
func (m *StorageScanReq) String() string { return proto.CompactTextString(m) }
func (*StorageScanReq) ProtoMessage()    {}
func (*StorageScanReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{2}
}
func (m *StorageScanReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageScanReq.Unmarshal(m, b)
}
func (m *StorageScanReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageScanReq.Marshal(b, m, deterministic)
}
func (dst *StorageScanReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageScanReq.Merge(dst, src)
}
func (m *StorageScanReq) XXX_Size() int {
	return xxx_messageInfo_StorageScanReq.Size(m)
}
func (m *StorageScanReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageScanReq.DiscardUnknown(m)
}

var xxx_messageInfo_StorageScanReq proto.InternalMessageInfo

func (m *StorageScanReq) GetNvme() *ScanNvmeReq {
	if m != nil {
		return m.Nvme
	}
	return nil
}

func (m *StorageScanReq) GetScm() *ScanScmReq {
	if m != nil {
		return m.Scm
	}
	return nil
}

type StorageScanResp struct {
	Nvme                 *ScanNvmeResp `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm                  *ScanScmResp  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StorageScanResp) Reset()         { *m = StorageScanResp{} }
func (m *StorageScanResp) String() string { return proto.CompactTextString(m) }
func (*StorageScanResp) ProtoMessage()    {}
func (*StorageScanResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{3}
}
func (m *StorageScanResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageScanResp.Unmarshal(m, b)
}
func (m *StorageScanResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageScanResp.Marshal(b, m, deterministic)
}
func (dst *StorageScanResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageScanResp.Merge(dst, src)
}
func (m *StorageScanResp) XXX_Size() int {
	return xxx_messageInfo_StorageScanResp.Size(m)
}
func (m *StorageScanResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageScanResp.DiscardUnknown(m)
}

var xxx_messageInfo_StorageScanResp proto.InternalMessageInfo

func (m *StorageScanResp) GetNvme() *ScanNvmeResp {
	if m != nil {
		return m.Nvme
	}
	return nil
}

func (m *StorageScanResp) GetScm() *ScanScmResp {
	if m != nil {
		return m.Scm
	}
	return nil
}

type StorageFormatReq struct {
	Nvme                 *FormatNvmeReq `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm                  *FormatScmReq  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StorageFormatReq) Reset()         { *m = StorageFormatReq{} }
func (m *StorageFormatReq) String() string { return proto.CompactTextString(m) }
func (*StorageFormatReq) ProtoMessage()    {}
func (*StorageFormatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{4}
}
func (m *StorageFormatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageFormatReq.Unmarshal(m, b)
}
func (m *StorageFormatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageFormatReq.Marshal(b, m, deterministic)
}
func (dst *StorageFormatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageFormatReq.Merge(dst, src)
}
func (m *StorageFormatReq) XXX_Size() int {
	return xxx_messageInfo_StorageFormatReq.Size(m)
}
func (m *StorageFormatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageFormatReq.DiscardUnknown(m)
}

var xxx_messageInfo_StorageFormatReq proto.InternalMessageInfo

func (m *StorageFormatReq) GetNvme() *FormatNvmeReq {
	if m != nil {
		return m.Nvme
	}
	return nil
}

func (m *StorageFormatReq) GetScm() *FormatScmReq {
	if m != nil {
		return m.Scm
	}
	return nil
}

type StorageFormatResp struct {
	Crets                []*NvmeControllerResult `protobuf:"bytes,1,rep,name=crets,proto3" json:"crets,omitempty"`
	Mrets                []*ScmMountResult       `protobuf:"bytes,2,rep,name=mrets,proto3" json:"mrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StorageFormatResp) Reset()         { *m = StorageFormatResp{} }
func (m *StorageFormatResp) String() string { return proto.CompactTextString(m) }
func (*StorageFormatResp) ProtoMessage()    {}
func (*StorageFormatResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{5}
}
func (m *StorageFormatResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageFormatResp.Unmarshal(m, b)
}
func (m *StorageFormatResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageFormatResp.Marshal(b, m, deterministic)
}
func (dst *StorageFormatResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageFormatResp.Merge(dst, src)
}
func (m *StorageFormatResp) XXX_Size() int {
	return xxx_messageInfo_StorageFormatResp.Size(m)
}
func (m *StorageFormatResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageFormatResp.DiscardUnknown(m)
}

var xxx_messageInfo_StorageFormatResp proto.InternalMessageInfo

func (m *StorageFormatResp) GetCrets() []*NvmeControllerResult {
	if m != nil {
		return m.Crets
	}
	return nil
}

func (m *StorageFormatResp) GetMrets() []*ScmMountResult {
	if m != nil {
		return m.Mrets
	}
	return nil
}

type StorageUpdateReq struct {
	Nvme                 *UpdateNvmeReq `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm                  *UpdateScmReq  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StorageUpdateReq) Reset()         { *m = StorageUpdateReq{} }
func (m *StorageUpdateReq) String() string { return proto.CompactTextString(m) }
func (*StorageUpdateReq) ProtoMessage()    {}
func (*StorageUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{6}
}
func (m *StorageUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageUpdateReq.Unmarshal(m, b)
}
func (m *StorageUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageUpdateReq.Marshal(b, m, deterministic)
}
func (dst *StorageUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageUpdateReq.Merge(dst, src)
}
func (m *StorageUpdateReq) XXX_Size() int {
	return xxx_messageInfo_StorageUpdateReq.Size(m)
}
func (m *StorageUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_StorageUpdateReq proto.InternalMessageInfo

func (m *StorageUpdateReq) GetNvme() *UpdateNvmeReq {
	if m != nil {
		return m.Nvme
	}
	return nil
}

func (m *StorageUpdateReq) GetScm() *UpdateScmReq {
	if m != nil {
		return m.Scm
	}
	return nil
}

type StorageUpdateResp struct {
	Crets                []*NvmeControllerResult `protobuf:"bytes,1,rep,name=crets,proto3" json:"crets,omitempty"`
	Mrets                []*ScmModuleResult      `protobuf:"bytes,2,rep,name=mrets,proto3" json:"mrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StorageUpdateResp) Reset()         { *m = StorageUpdateResp{} }
func (m *StorageUpdateResp) String() string { return proto.CompactTextString(m) }
func (*StorageUpdateResp) ProtoMessage()    {}
func (*StorageUpdateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{7}
}
func (m *StorageUpdateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageUpdateResp.Unmarshal(m, b)
}
func (m *StorageUpdateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageUpdateResp.Marshal(b, m, deterministic)
}
func (dst *StorageUpdateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageUpdateResp.Merge(dst, src)
}
func (m *StorageUpdateResp) XXX_Size() int {
	return xxx_messageInfo_StorageUpdateResp.Size(m)
}
func (m *StorageUpdateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageUpdateResp.DiscardUnknown(m)
}

var xxx_messageInfo_StorageUpdateResp proto.InternalMessageInfo

func (m *StorageUpdateResp) GetCrets() []*NvmeControllerResult {
	if m != nil {
		return m.Crets
	}
	return nil
}

func (m *StorageUpdateResp) GetMrets() []*ScmModuleResult {
	if m != nil {
		return m.Mrets
	}
	return nil
}

type StorageBurnInReq struct {
	Nvme                 *BurninNvmeReq `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm                  *BurninScmReq  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StorageBurnInReq) Reset()         { *m = StorageBurnInReq{} }
func (m *StorageBurnInReq) String() string { return proto.CompactTextString(m) }
func (*StorageBurnInReq) ProtoMessage()    {}
func (*StorageBurnInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{8}
}
func (m *StorageBurnInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageBurnInReq.Unmarshal(m, b)
}
func (m *StorageBurnInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageBurnInReq.Marshal(b, m, deterministic)
}
func (dst *StorageBurnInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageBurnInReq.Merge(dst, src)
}
func (m *StorageBurnInReq) XXX_Size() int {
	return xxx_messageInfo_StorageBurnInReq.Size(m)
}
func (m *StorageBurnInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageBurnInReq.DiscardUnknown(m)
}

var xxx_messageInfo_StorageBurnInReq proto.InternalMessageInfo

func (m *StorageBurnInReq) GetNvme() *BurninNvmeReq {
	if m != nil {
		return m.Nvme
	}
	return nil
}

func (m *StorageBurnInReq) GetScm() *BurninScmReq {
	if m != nil {
		return m.Scm
	}
	return nil
}

type StorageBurnInResp struct {
	Crets                []*NvmeControllerResult `protobuf:"bytes,1,rep,name=crets,proto3" json:"crets,omitempty"`
	Mrets                []*ScmMountResult       `protobuf:"bytes,2,rep,name=mrets,proto3" json:"mrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StorageBurnInResp) Reset()         { *m = StorageBurnInResp{} }
func (m *StorageBurnInResp) String() string { return proto.CompactTextString(m) }
func (*StorageBurnInResp) ProtoMessage()    {}
func (*StorageBurnInResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_a5934919a5b6df32, []int{9}
}
func (m *StorageBurnInResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageBurnInResp.Unmarshal(m, b)
}
func (m *StorageBurnInResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageBurnInResp.Marshal(b, m, deterministic)
}
func (dst *StorageBurnInResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageBurnInResp.Merge(dst, src)
}
func (m *StorageBurnInResp) XXX_Size() int {
	return xxx_messageInfo_StorageBurnInResp.Size(m)
}
func (m *StorageBurnInResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageBurnInResp.DiscardUnknown(m)
}

var xxx_messageInfo_StorageBurnInResp proto.InternalMessageInfo

func (m *StorageBurnInResp) GetCrets() []*NvmeControllerResult {
	if m != nil {
		return m.Crets
	}
	return nil
}

func (m *StorageBurnInResp) GetMrets() []*ScmMountResult {
	if m != nil {
		return m.Mrets
	}
	return nil
}

func init() {
	proto.RegisterType((*StoragePrepareReq)(nil), "mgmt.StoragePrepareReq")
	proto.RegisterType((*StoragePrepareResp)(nil), "mgmt.StoragePrepareResp")
	proto.RegisterType((*StorageScanReq)(nil), "mgmt.StorageScanReq")
	proto.RegisterType((*StorageScanResp)(nil), "mgmt.StorageScanResp")
	proto.RegisterType((*StorageFormatReq)(nil), "mgmt.StorageFormatReq")
	proto.RegisterType((*StorageFormatResp)(nil), "mgmt.StorageFormatResp")
	proto.RegisterType((*StorageUpdateReq)(nil), "mgmt.StorageUpdateReq")
	proto.RegisterType((*StorageUpdateResp)(nil), "mgmt.StorageUpdateResp")
	proto.RegisterType((*StorageBurnInReq)(nil), "mgmt.StorageBurnInReq")
	proto.RegisterType((*StorageBurnInResp)(nil), "mgmt.StorageBurnInResp")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_storage_a5934919a5b6df32) }

var fileDescriptor_storage_a5934919a5b6df32 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5d, 0x4b, 0xf3, 0x30,
	0x18, 0x86, 0xd9, 0xd7, 0x7b, 0x90, 0xf1, 0xea, 0x1a, 0x37, 0x18, 0x3b, 0x1a, 0xd3, 0xcd, 0xa9,
	0x30, 0x64, 0xfe, 0x03, 0x05, 0xc1, 0x03, 0x45, 0x3a, 0x3c, 0x12, 0x94, 0xd8, 0x85, 0x21, 0x34,
	0x1f, 0x4b, 0xd2, 0xfd, 0x7e, 0xc9, 0x57, 0x4d, 0xea, 0xca, 0x40, 0xf0, 0x34, 0xcf, 0xfd, 0xdc,
	0x57, 0xae, 0x06, 0x0a, 0xfe, 0x4b, 0xc5, 0x04, 0xda, 0xe0, 0x05, 0x17, 0x4c, 0x31, 0xd8, 0x26,
	0x1b, 0xa2, 0x46, 0xd0, 0x1d, 0xbe, 0xd3, 0x1d, 0x71, 0x93, 0x51, 0xe2, 0xcf, 0x64, 0x46, 0xec,
	0xd1, 0x64, 0x0d, 0x92, 0x95, 0x3d, 0x7c, 0x16, 0x98, 0x23, 0x81, 0x53, 0xbc, 0x85, 0x73, 0xd0,
	0xd6, 0x5b, 0xc3, 0xc6, 0xb8, 0x31, 0xef, 0x2e, 0xfb, 0x0b, 0x5d, 0xb8, 0x70, 0xf3, 0xa7, 0x1d,
	0xd1, 0x99, 0xd4, 0x24, 0xe0, 0x14, 0xb4, 0x64, 0x46, 0x86, 0x4d, 0x13, 0x3c, 0x89, 0x82, 0xab,
	0x8c, 0xe8, 0x9c, 0x9e, 0x4f, 0x36, 0x00, 0x56, 0x29, 0x92, 0xc3, 0x8b, 0x08, 0x33, 0xd8, 0x83,
	0x91, 0xdc, 0x71, 0x66, 0x21, 0xa7, 0xff, 0x93, 0x23, 0xb9, 0x05, 0xbd, 0x82, 0x23, 0x07, 0x5a,
	0x65, 0x88, 0x6a, 0x97, 0x69, 0x04, 0x49, 0xec, 0xaa, 0x1e, 0xc6, 0x22, 0x93, 0x10, 0xd0, 0xfb,
	0x4e, 0x85, 0x16, 0x6f, 0xe0, 0x38, 0x2a, 0x97, 0x1c, 0xce, 0xa2, 0x76, 0x58, 0x6d, 0x2f, 0xef,
	0x7f, 0x1a, 0xd6, 0x27, 0x95, 0x7a, 0x7f, 0x79, 0x04, 0x7a, 0xae, 0xff, 0x9e, 0x09, 0x82, 0x94,
	0xbe, 0xfe, 0x79, 0x04, 0x70, 0x5f, 0xd8, 0x8e, 0x63, 0x81, 0xb3, 0x90, 0x00, 0xc3, 0x5c, 0xa8,
	0xb0, 0x2d, 0x9f, 0xdb, 0x23, 0x24, 0x87, 0xd7, 0xa0, 0x93, 0x09, 0xac, 0xe4, 0xb0, 0x31, 0x6e,
	0xcd, 0xbb, 0xcb, 0x91, 0x5d, 0xd6, 0xf5, 0x77, 0x8c, 0x2a, 0xc1, 0xf2, 0x1c, 0x8b, 0x14, 0xcb,
	0x22, 0x57, 0xa9, 0x0d, 0xc2, 0x4b, 0xd0, 0x21, 0x66, 0xa3, 0x69, 0x36, 0xfa, 0x5e, 0x88, 0x3c,
	0xb2, 0x82, 0x2a, 0x9f, 0x35, 0x91, 0xc0, 0xea, 0x85, 0xaf, 0x91, 0xc2, 0xb5, 0x56, 0x76, 0x7c,
	0xd8, 0xca, 0xe6, 0x42, 0x2b, 0x51, 0x5a, 0x79, 0xc4, 0xaf, 0xac, 0xae, 0x62, 0xab, 0x41, 0x60,
	0xb5, 0x2e, 0x72, 0x5c, 0xa7, 0x75, 0x5b, 0x08, 0xfa, 0x40, 0x6b, 0xb5, 0xf4, 0xf8, 0x93, 0x1e,
	0xd6, 0xb2, 0xb9, 0xfd, 0x8f, 0xe5, 0x11, 0x7f, 0xfd, 0x58, 0x1f, 0xff, 0xcc, 0x5f, 0xe1, 0xe6,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xd2, 0xdf, 0xd4, 0x53, 0x04, 0x00, 0x00,
}
