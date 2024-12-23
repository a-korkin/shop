// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: shop.proto

package common

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

type ItemId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ItemId) Reset() {
	*x = ItemId{}
	mi := &file_shop_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemId) ProtoMessage() {}

func (x *ItemId) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemId.ProtoReflect.Descriptor instead.
func (*ItemId) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{0}
}

func (x *ItemId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_shop_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ItemDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ItemDto) Reset() {
	*x = ItemDto{}
	mi := &file_shop_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDto) ProtoMessage() {}

func (x *ItemDto) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDto.ProtoReflect.Descriptor instead.
func (*ItemDto) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{2}
}

func (x *ItemDto) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ItemDto) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ItemList) Reset() {
	*x = ItemList{}
	mi := &file_shop_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemList) ProtoMessage() {}

func (x *ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemList.ProtoReflect.Descriptor instead.
func (*ItemList) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{3}
}

func (x *ItemList) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type PageParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PageParams) Reset() {
	*x = PageParams{}
	mi := &file_shop_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageParams) ProtoMessage() {}

func (x *PageParams) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageParams.ProtoReflect.Descriptor instead.
func (*PageParams) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{4}
}

func (x *PageParams) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageParams) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageParams) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_shop_proto protoreflect.FileDescriptor

var file_shop_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x35, 0x0a, 0x07, 0x49, 0x74,
	0x65, 0x6d, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x27, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4e, 0x0a, 0x0a, 0x50, 0x61,
	0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0x6b, 0x0a, 0x0b, 0x53, 0x68,
	0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x07, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x1a, 0x05, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x08, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x74, 0x6f, 0x1a, 0x05, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x0b, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x2d, 0x6b, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x2f, 0x73,
	0x68, 0x6f, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_shop_proto_rawDescOnce sync.Once
	file_shop_proto_rawDescData = file_shop_proto_rawDesc
)

func file_shop_proto_rawDescGZIP() []byte {
	file_shop_proto_rawDescOnce.Do(func() {
		file_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_proto_rawDescData)
	})
	return file_shop_proto_rawDescData
}

var file_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shop_proto_goTypes = []any{
	(*ItemId)(nil),     // 0: ItemId
	(*Item)(nil),       // 1: Item
	(*ItemDto)(nil),    // 2: ItemDto
	(*ItemList)(nil),   // 3: ItemList
	(*PageParams)(nil), // 4: PageParams
}
var file_shop_proto_depIdxs = []int32{
	1, // 0: ItemList.items:type_name -> Item
	0, // 1: ShopService.GetItem:input_type -> ItemId
	2, // 2: ShopService.CreateItem:input_type -> ItemDto
	4, // 3: ShopService.GetItems:input_type -> PageParams
	1, // 4: ShopService.GetItem:output_type -> Item
	1, // 5: ShopService.CreateItem:output_type -> Item
	3, // 6: ShopService.GetItems:output_type -> ItemList
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shop_proto_init() }
func file_shop_proto_init() {
	if File_shop_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_proto_goTypes,
		DependencyIndexes: file_shop_proto_depIdxs,
		MessageInfos:      file_shop_proto_msgTypes,
	}.Build()
	File_shop_proto = out.File
	file_shop_proto_rawDesc = nil
	file_shop_proto_goTypes = nil
	file_shop_proto_depIdxs = nil
}
