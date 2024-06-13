// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: v1/menu.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMenuRequest) Reset() {
	*x = ListMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenuRequest) ProtoMessage() {}

func (x *ListMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenuRequest.ProtoReflect.Descriptor instead.
func (*ListMenuRequest) Descriptor() ([]byte, []int) {
	return file_v1_menu_proto_rawDescGZIP(), []int{0}
}

type ListMenuReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ListMenuReply_Menu `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListMenuReply) Reset() {
	*x = ListMenuReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMenuReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenuReply) ProtoMessage() {}

func (x *ListMenuReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenuReply.ProtoReflect.Descriptor instead.
func (*ListMenuReply) Descriptor() ([]byte, []int) {
	return file_v1_menu_proto_rawDescGZIP(), []int{1}
}

func (x *ListMenuReply) GetData() []*ListMenuReply_Menu {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListMenuReply_Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string                `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Component string                `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	Name      string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	IconCls   string                `protobuf:"bytes,4,opt,name=iconCls,proto3" json:"iconCls,omitempty"`
	Meta      bool                  `protobuf:"varint,5,opt,name=meta,proto3" json:"meta,omitempty"`
	Children  []*ListMenuReply_Menu `protobuf:"bytes,6,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ListMenuReply_Menu) Reset() {
	*x = ListMenuReply_Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMenuReply_Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenuReply_Menu) ProtoMessage() {}

func (x *ListMenuReply_Menu) ProtoReflect() protoreflect.Message {
	mi := &file_v1_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenuReply_Menu.ProtoReflect.Descriptor instead.
func (*ListMenuReply_Menu) Descriptor() ([]byte, []int) {
	return file_v1_menu_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListMenuReply_Menu) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListMenuReply_Menu) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *ListMenuReply_Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListMenuReply_Menu) GetIconCls() string {
	if x != nil {
		return x.IconCls
	}
	return ""
}

func (x *ListMenuReply_Menu) GetMeta() bool {
	if x != nil {
		return x.Meta
	}
	return false
}

func (x *ListMenuReply_Menu) GetChildren() []*ListMenuReply_Menu {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_v1_menu_proto protoreflect.FileDescriptor

var file_v1_menu_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x77, 0x65,
	0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8e, 0x02, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0xbf, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x43, 0x6c, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x43, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x32, 0x6e, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x66,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x42, 0x1e, 0x5a, 0x1c, 0x76, 0x68, 0x72, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x77, 0x65, 0x62,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_menu_proto_rawDescOnce sync.Once
	file_v1_menu_proto_rawDescData = file_v1_menu_proto_rawDesc
)

func file_v1_menu_proto_rawDescGZIP() []byte {
	file_v1_menu_proto_rawDescOnce.Do(func() {
		file_v1_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_menu_proto_rawDescData)
	})
	return file_v1_menu_proto_rawDescData
}

var file_v1_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_menu_proto_goTypes = []interface{}{
	(*ListMenuRequest)(nil),    // 0: api.frontend.web.v1.ListMenuRequest
	(*ListMenuReply)(nil),      // 1: api.frontend.web.v1.ListMenuReply
	(*ListMenuReply_Menu)(nil), // 2: api.frontend.web.v1.ListMenuReply.Menu
}
var file_v1_menu_proto_depIdxs = []int32{
	2, // 0: api.frontend.web.v1.ListMenuReply.data:type_name -> api.frontend.web.v1.ListMenuReply.Menu
	2, // 1: api.frontend.web.v1.ListMenuReply.Menu.children:type_name -> api.frontend.web.v1.ListMenuReply.Menu
	0, // 2: api.frontend.web.v1.Menu.List:input_type -> api.frontend.web.v1.ListMenuRequest
	1, // 3: api.frontend.web.v1.Menu.List:output_type -> api.frontend.web.v1.ListMenuReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_menu_proto_init() }
func file_v1_menu_proto_init() {
	if File_v1_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMenuRequest); i {
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
		file_v1_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMenuReply); i {
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
		file_v1_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMenuReply_Menu); i {
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
			RawDescriptor: file_v1_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_menu_proto_goTypes,
		DependencyIndexes: file_v1_menu_proto_depIdxs,
		MessageInfos:      file_v1_menu_proto_msgTypes,
	}.Build()
	File_v1_menu_proto = out.File
	file_v1_menu_proto_rawDesc = nil
	file_v1_menu_proto_goTypes = nil
	file_v1_menu_proto_depIdxs = nil
}
