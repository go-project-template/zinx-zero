// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: proto/player.proto

package msg

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

type DBPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId      int64            `protobuf:"varint,1,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	AccountId   int64            `protobuf:"varint,3,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	Nickname    string           `protobuf:"bytes,5,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	IntAttr     map[int32]int64  `protobuf:"bytes,6,rep,name=IntAttr,proto3" json:"IntAttr,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	StrAttr     map[int32]string `protobuf:"bytes,7,rep,name=StrAttr,proto3" json:"StrAttr,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DBPlayerBag *DBPlayerBag     `protobuf:"bytes,8,opt,name=DBPlayerBag,proto3" json:"DBPlayerBag,omitempty"`
}

func (x *DBPlayer) Reset() {
	*x = DBPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBPlayer) ProtoMessage() {}

func (x *DBPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBPlayer.ProtoReflect.Descriptor instead.
func (*DBPlayer) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{0}
}

func (x *DBPlayer) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *DBPlayer) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *DBPlayer) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *DBPlayer) GetIntAttr() map[int32]int64 {
	if x != nil {
		return x.IntAttr
	}
	return nil
}

func (x *DBPlayer) GetStrAttr() map[int32]string {
	if x != nil {
		return x.StrAttr
	}
	return nil
}

func (x *DBPlayer) GetDBPlayerBag() *DBPlayerBag {
	if x != nil {
		return x.DBPlayerBag
	}
	return nil
}

type ItemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId  int64 `protobuf:"varint,1,opt,name=UniqueId,proto3" json:"UniqueId,omitempty"`
	ItemId    int32 `protobuf:"varint,2,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	ItemCount int64 `protobuf:"varint,3,opt,name=ItemCount,proto3" json:"ItemCount,omitempty"`
}

func (x *ItemInfo) Reset() {
	*x = ItemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemInfo) ProtoMessage() {}

func (x *ItemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemInfo.ProtoReflect.Descriptor instead.
func (*ItemInfo) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{1}
}

func (x *ItemInfo) GetUniqueId() int64 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

func (x *ItemInfo) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemInfo) GetItemCount() int64 {
	if x != nil {
		return x.ItemCount
	}
	return 0
}

type DBPlayerBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList []*ItemInfo `protobuf:"bytes,1,rep,name=ItemList,proto3" json:"ItemList,omitempty"` // map<int64,ItemInfo> ItemList
}

func (x *DBPlayerBag) Reset() {
	*x = DBPlayerBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBPlayerBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBPlayerBag) ProtoMessage() {}

func (x *DBPlayerBag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBPlayerBag.ProtoReflect.Descriptor instead.
func (*DBPlayerBag) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{2}
}

func (x *DBPlayerBag) GetItemList() []*ItemInfo {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_proto_player_proto protoreflect.FileDescriptor

var file_proto_player_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x02,
	0x0a, 0x08, 0x44, 0x42, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x49, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x44, 0x42, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x49, 0x6e, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x41, 0x74, 0x74, 0x72, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x44, 0x42, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x53, 0x74, 0x72, 0x41, 0x74, 0x74, 0x72, 0x12, 0x32, 0x0a, 0x0b, 0x44, 0x42, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x44, 0x42, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x67, 0x52,
	0x0b, 0x44, 0x42, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x67, 0x1a, 0x3a, 0x0a, 0x0c,
	0x49, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x38, 0x0a, 0x0b, 0x44, 0x42, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61,
	0x67, 0x12, 0x29, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_player_proto_rawDescOnce sync.Once
	file_proto_player_proto_rawDescData = file_proto_player_proto_rawDesc
)

func file_proto_player_proto_rawDescGZIP() []byte {
	file_proto_player_proto_rawDescOnce.Do(func() {
		file_proto_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_proto_rawDescData)
	})
	return file_proto_player_proto_rawDescData
}

var file_proto_player_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_player_proto_goTypes = []interface{}{
	(*DBPlayer)(nil),    // 0: msg.DBPlayer
	(*ItemInfo)(nil),    // 1: msg.ItemInfo
	(*DBPlayerBag)(nil), // 2: msg.DBPlayerBag
	nil,                 // 3: msg.DBPlayer.IntAttrEntry
	nil,                 // 4: msg.DBPlayer.StrAttrEntry
}
var file_proto_player_proto_depIdxs = []int32{
	3, // 0: msg.DBPlayer.IntAttr:type_name -> msg.DBPlayer.IntAttrEntry
	4, // 1: msg.DBPlayer.StrAttr:type_name -> msg.DBPlayer.StrAttrEntry
	2, // 2: msg.DBPlayer.DBPlayerBag:type_name -> msg.DBPlayerBag
	1, // 3: msg.DBPlayerBag.ItemList:type_name -> msg.ItemInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_player_proto_init() }
func file_proto_player_proto_init() {
	if File_proto_player_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBPlayer); i {
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
		file_proto_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemInfo); i {
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
		file_proto_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBPlayerBag); i {
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
			RawDescriptor: file_proto_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_player_proto_goTypes,
		DependencyIndexes: file_proto_player_proto_depIdxs,
		MessageInfos:      file_proto_player_proto_msgTypes,
	}.Build()
	File_proto_player_proto = out.File
	file_proto_player_proto_rawDesc = nil
	file_proto_player_proto_goTypes = nil
	file_proto_player_proto_depIdxs = nil
}
