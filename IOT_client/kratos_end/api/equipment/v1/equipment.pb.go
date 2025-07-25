// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: equipment/v1/equipment.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductsListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductsListReq) Reset() {
	*x = ProductsListReq{}
	mi := &file_equipment_v1_equipment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductsListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsListReq) ProtoMessage() {}

func (x *ProductsListReq) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsListReq.ProtoReflect.Descriptor instead.
func (*ProductsListReq) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{0}
}

type ProductsListResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greet         string                 `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductsListResp) Reset() {
	*x = ProductsListResp{}
	mi := &file_equipment_v1_equipment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductsListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsListResp) ProtoMessage() {}

func (x *ProductsListResp) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsListResp.ProtoReflect.Descriptor instead.
func (*ProductsListResp) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{1}
}

func (x *ProductsListResp) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

var File_equipment_v1_equipment_proto protoreflect.FileDescriptor

const file_equipment_v1_equipment_proto_rawDesc = "" +
	"\n" +
	"\x1cequipment/v1/equipment.proto\x12\fequipment.v1\x1a\x1cgoogle/api/annotations.proto\"\x11\n" +
	"\x0fProductsListReq\"(\n" +
	"\x10ProductsListResp\x12\x14\n" +
	"\x05greet\x18\x01 \x01(\tR\x05greet2w\n" +
	"\tEquipment\x12j\n" +
	"\fProductsList\x12\x1d.equipment.v1.ProductsListReq\x1a\x1e.equipment.v1.ProductsListResp\"\x1b\x82\xd3\xe4\x93\x02\x15:\x01*\"\x10/v1/ProductsListB4\n" +
	"\x10api.equipment.v1P\x01Z\x1ekratos_end/api/equipment/v1;v1b\x06proto3"

var (
	file_equipment_v1_equipment_proto_rawDescOnce sync.Once
	file_equipment_v1_equipment_proto_rawDescData []byte
)

func file_equipment_v1_equipment_proto_rawDescGZIP() []byte {
	file_equipment_v1_equipment_proto_rawDescOnce.Do(func() {
		file_equipment_v1_equipment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_equipment_v1_equipment_proto_rawDesc), len(file_equipment_v1_equipment_proto_rawDesc)))
	})
	return file_equipment_v1_equipment_proto_rawDescData
}

var file_equipment_v1_equipment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_equipment_v1_equipment_proto_goTypes = []any{
	(*ProductsListReq)(nil),  // 0: equipment.v1.ProductsListReq
	(*ProductsListResp)(nil), // 1: equipment.v1.ProductsListResp
}
var file_equipment_v1_equipment_proto_depIdxs = []int32{
	0, // 0: equipment.v1.Equipment.ProductsList:input_type -> equipment.v1.ProductsListReq
	1, // 1: equipment.v1.Equipment.ProductsList:output_type -> equipment.v1.ProductsListResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_equipment_v1_equipment_proto_init() }
func file_equipment_v1_equipment_proto_init() {
	if File_equipment_v1_equipment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_equipment_v1_equipment_proto_rawDesc), len(file_equipment_v1_equipment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_equipment_v1_equipment_proto_goTypes,
		DependencyIndexes: file_equipment_v1_equipment_proto_depIdxs,
		MessageInfos:      file_equipment_v1_equipment_proto_msgTypes,
	}.Build()
	File_equipment_v1_equipment_proto = out.File
	file_equipment_v1_equipment_proto_goTypes = nil
	file_equipment_v1_equipment_proto_depIdxs = nil
}
