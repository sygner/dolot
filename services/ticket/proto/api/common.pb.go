// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.3
// source: proto/protos/common.proto

package __

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

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit    int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	GetTotal bool  `protobuf:"varint,3,opt,name=get_total,json=getTotal,proto3" json:"get_total,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	mi := &file_proto_protos_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_proto_protos_common_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetGetTotal() bool {
	if x != nil {
		return x.GetTotal
	}
	return false
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	mi := &file_proto_protos_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_proto_protos_common_proto_rawDescGZIP(), []int{1}
}

func (x *UserId) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	mi := &file_proto_protos_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_proto_protos_common_proto_rawDescGZIP(), []int{2}
}

func (x *Signature) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type TicketId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TicketId) Reset() {
	*x = TicketId{}
	mi := &file_proto_protos_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketId) ProtoMessage() {}

func (x *TicketId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketId.ProtoReflect.Descriptor instead.
func (*TicketId) Descriptor() ([]byte, []int) {
	return file_proto_protos_common_proto_rawDescGZIP(), []int{3}
}

func (x *TicketId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Count struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Count) Reset() {
	*x = Count{}
	mi := &file_proto_protos_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Count) ProtoMessage() {}

func (x *Count) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Count.ProtoReflect.Descriptor instead.
func (*Count) Descriptor() ([]byte, []int) {
	return file_proto_protos_common_proto_rawDescGZIP(), []int{4}
}

func (x *Count) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_protos_common_proto protoreflect.FileDescriptor

var file_proto_protos_common_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x22, 0x57, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x21, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x29, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_protos_common_proto_rawDescOnce sync.Once
	file_proto_protos_common_proto_rawDescData = file_proto_protos_common_proto_rawDesc
)

func file_proto_protos_common_proto_rawDescGZIP() []byte {
	file_proto_protos_common_proto_rawDescOnce.Do(func() {
		file_proto_protos_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protos_common_proto_rawDescData)
	})
	return file_proto_protos_common_proto_rawDescData
}

var file_proto_protos_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_protos_common_proto_goTypes = []any{
	(*Pagination)(nil), // 0: ticket.Pagination
	(*UserId)(nil),     // 1: ticket.UserId
	(*Signature)(nil),  // 2: ticket.Signature
	(*TicketId)(nil),   // 3: ticket.TicketId
	(*Count)(nil),      // 4: ticket.Count
}
var file_proto_protos_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_protos_common_proto_init() }
func file_proto_protos_common_proto_init() {
	if File_proto_protos_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_protos_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_protos_common_proto_goTypes,
		DependencyIndexes: file_proto_protos_common_proto_depIdxs,
		MessageInfos:      file_proto_protos_common_proto_msgTypes,
	}.Build()
	File_proto_protos_common_proto = out.File
	file_proto_protos_common_proto_rawDesc = nil
	file_proto_protos_common_proto_goTypes = nil
	file_proto_protos_common_proto_depIdxs = nil
}
