// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: transforms.proto

package pbtransform

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CombinedFilter is a combination of "LogFilters" and "CallToFilters"
//
// It transforms the requested stream in two ways:
//   1. STRIPPING
//      The block data is stripped from all transactions that don't
//      match any of the filters.
//
//   2. SKIPPING
//      If an "block index" covers a range containing a
//      block that does NOT match any of the filters, the block will be
//      skipped altogether, UNLESS send_all_block_headers is enabled
//      In that case, the block would still be sent, but without any
//      transactionTrace
//
// The SKIPPING feature only applies to historical blocks, because
// the "block index" is always produced after the merged-blocks files
// are produced. Therefore, the "live" blocks are never filtered out.
//
type CombinedFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogFilters  []*LogFilter    `protobuf:"bytes,1,rep,name=log_filters,json=logFilters,proto3" json:"log_filters,omitempty"`
	CallFilters []*CallToFilter `protobuf:"bytes,2,rep,name=call_filters,json=callFilters,proto3" json:"call_filters,omitempty"`
	// Always send all blocks. if they don't match any log_filters or call_filters,
	// all the transactions will be filtered out, sending only the header.
	SendAllBlockHeaders bool `protobuf:"varint,3,opt,name=send_all_block_headers,json=sendAllBlockHeaders,proto3" json:"send_all_block_headers,omitempty"`
}

func (x *CombinedFilter) Reset() {
	*x = CombinedFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transforms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombinedFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombinedFilter) ProtoMessage() {}

func (x *CombinedFilter) ProtoReflect() protoreflect.Message {
	mi := &file_transforms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombinedFilter.ProtoReflect.Descriptor instead.
func (*CombinedFilter) Descriptor() ([]byte, []int) {
	return file_transforms_proto_rawDescGZIP(), []int{0}
}

func (x *CombinedFilter) GetLogFilters() []*LogFilter {
	if x != nil {
		return x.LogFilters
	}
	return nil
}

func (x *CombinedFilter) GetCallFilters() []*CallToFilter {
	if x != nil {
		return x.CallFilters
	}
	return nil
}

func (x *CombinedFilter) GetSendAllBlockHeaders() bool {
	if x != nil {
		return x.SendAllBlockHeaders
	}
	return false
}

// MultiLogFilter concatenates the results of each LogFilter (inclusive OR)
type MultiLogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogFilters []*LogFilter `protobuf:"bytes,1,rep,name=log_filters,json=logFilters,proto3" json:"log_filters,omitempty"`
}

func (x *MultiLogFilter) Reset() {
	*x = MultiLogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transforms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiLogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiLogFilter) ProtoMessage() {}

func (x *MultiLogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_transforms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiLogFilter.ProtoReflect.Descriptor instead.
func (*MultiLogFilter) Descriptor() ([]byte, []int) {
	return file_transforms_proto_rawDescGZIP(), []int{1}
}

func (x *MultiLogFilter) GetLogFilters() []*LogFilter {
	if x != nil {
		return x.LogFilters
	}
	return nil
}

// LogFilter will match calls where *BOTH*
// * the contract address that emits the log is one in the provided addresses -- OR addresses list is empty --
// * the event signature (topic.0) is one of the provided event_signatures -- OR event_signatures is empty --
//
// a LogFilter with both empty addresses and event_signatures lists is invalid and will fail.
type LogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses       [][]byte `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	EventSignatures [][]byte `protobuf:"bytes,2,rep,name=event_signatures,json=eventSignatures,proto3" json:"event_signatures,omitempty"` // corresponds to the keccak of the event signature which is stores in topic.0
}

func (x *LogFilter) Reset() {
	*x = LogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transforms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogFilter) ProtoMessage() {}

func (x *LogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_transforms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogFilter.ProtoReflect.Descriptor instead.
func (*LogFilter) Descriptor() ([]byte, []int) {
	return file_transforms_proto_rawDescGZIP(), []int{2}
}

func (x *LogFilter) GetAddresses() [][]byte {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *LogFilter) GetEventSignatures() [][]byte {
	if x != nil {
		return x.EventSignatures
	}
	return nil
}

// MultiCallToFilter concatenates the results of each CallToFilter (inclusive OR)
type MultiCallToFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallFilters []*CallToFilter `protobuf:"bytes,1,rep,name=call_filters,json=callFilters,proto3" json:"call_filters,omitempty"`
}

func (x *MultiCallToFilter) Reset() {
	*x = MultiCallToFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transforms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCallToFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCallToFilter) ProtoMessage() {}

func (x *MultiCallToFilter) ProtoReflect() protoreflect.Message {
	mi := &file_transforms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCallToFilter.ProtoReflect.Descriptor instead.
func (*MultiCallToFilter) Descriptor() ([]byte, []int) {
	return file_transforms_proto_rawDescGZIP(), []int{3}
}

func (x *MultiCallToFilter) GetCallFilters() []*CallToFilter {
	if x != nil {
		return x.CallFilters
	}
	return nil
}

// CallToFilter will match calls where *BOTH*
// * the contract address (TO) is one in the provided addresses -- OR addresses list is empty --
// * the method signature (in 4-bytes format) is one of the provided signatures -- OR signatures is empty --
//
// a CallToFilter with both empty addresses and signatures lists is invalid and will fail.
type CallToFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses  [][]byte `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Signatures [][]byte `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *CallToFilter) Reset() {
	*x = CallToFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transforms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallToFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallToFilter) ProtoMessage() {}

func (x *CallToFilter) ProtoReflect() protoreflect.Message {
	mi := &file_transforms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallToFilter.ProtoReflect.Descriptor instead.
func (*CallToFilter) Descriptor() ([]byte, []int) {
	return file_transforms_proto_rawDescGZIP(), []int{4}
}

func (x *CallToFilter) GetAddresses() [][]byte {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *CallToFilter) GetSignatures() [][]byte {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type LightBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LightBlock) Reset() {
	*x = LightBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transforms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightBlock) ProtoMessage() {}

func (x *LightBlock) ProtoReflect() protoreflect.Message {
	mi := &file_transforms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightBlock.ProtoReflect.Descriptor instead.
func (*LightBlock) Descriptor() ([]byte, []int) {
	return file_transforms_proto_rawDescGZIP(), []int{5}
}

var File_transforms_proto protoreflect.FileDescriptor

var file_transforms_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x18, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0xd6, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x44, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x49, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x66,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x33, 0x0a, 0x16, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0x56, 0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x6f,
	0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73,
	0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x54, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x61, 0x6c, 0x6c,
	0x54, 0x6f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x22, 0x4c, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42,
	0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x72, 0x65,
	0x68, 0x6f, 0x73, 0x65, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x62, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_transforms_proto_rawDescOnce sync.Once
	file_transforms_proto_rawDescData = file_transforms_proto_rawDesc
)

func file_transforms_proto_rawDescGZIP() []byte {
	file_transforms_proto_rawDescOnce.Do(func() {
		file_transforms_proto_rawDescData = protoimpl.X.CompressGZIP(file_transforms_proto_rawDescData)
	})
	return file_transforms_proto_rawDescData
}

var file_transforms_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transforms_proto_goTypes = []interface{}{
	(*CombinedFilter)(nil),    // 0: sf.ethereum.transform.v1.CombinedFilter
	(*MultiLogFilter)(nil),    // 1: sf.ethereum.transform.v1.MultiLogFilter
	(*LogFilter)(nil),         // 2: sf.ethereum.transform.v1.LogFilter
	(*MultiCallToFilter)(nil), // 3: sf.ethereum.transform.v1.MultiCallToFilter
	(*CallToFilter)(nil),      // 4: sf.ethereum.transform.v1.CallToFilter
	(*LightBlock)(nil),        // 5: sf.ethereum.transform.v1.LightBlock
}
var file_transforms_proto_depIdxs = []int32{
	2, // 0: sf.ethereum.transform.v1.CombinedFilter.log_filters:type_name -> sf.ethereum.transform.v1.LogFilter
	4, // 1: sf.ethereum.transform.v1.CombinedFilter.call_filters:type_name -> sf.ethereum.transform.v1.CallToFilter
	2, // 2: sf.ethereum.transform.v1.MultiLogFilter.log_filters:type_name -> sf.ethereum.transform.v1.LogFilter
	4, // 3: sf.ethereum.transform.v1.MultiCallToFilter.call_filters:type_name -> sf.ethereum.transform.v1.CallToFilter
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_transforms_proto_init() }
func file_transforms_proto_init() {
	if File_transforms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transforms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombinedFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transforms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiLogFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transforms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transforms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCallToFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transforms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallToFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transforms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transforms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transforms_proto_goTypes,
		DependencyIndexes: file_transforms_proto_depIdxs,
		MessageInfos:      file_transforms_proto_msgTypes,
	}.Build()
	File_transforms_proto = out.File
	file_transforms_proto_rawDesc = nil
	file_transforms_proto_goTypes = nil
	file_transforms_proto_depIdxs = nil
}
