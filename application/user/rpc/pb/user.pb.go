// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: user.proto

package pb

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

type DecutMoneyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Pwd    string `protobuf:"bytes,3,opt,name=Pwd,proto3" json:"Pwd,omitempty"`
}

func (x *DecutMoneyReq) Reset() {
	*x = DecutMoneyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecutMoneyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecutMoneyReq) ProtoMessage() {}

func (x *DecutMoneyReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecutMoneyReq.ProtoReflect.Descriptor instead.
func (*DecutMoneyReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *DecutMoneyReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DecutMoneyReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DecutMoneyReq) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type DecutMoneyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DecutMoneyResp) Reset() {
	*x = DecutMoneyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecutMoneyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecutMoneyResp) ProtoMessage() {}

func (x *DecutMoneyResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecutMoneyResp.ProtoReflect.Descriptor instead.
func (*DecutMoneyResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4b, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x75, 0x74, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x50, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50,
	0x77, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x32, 0x8b, 0x01, 0x0a, 0x03, 0x50, 0x61, 0x79, 0x12, 0x3d, 0x0a, 0x0a,
	0x44, 0x65, 0x63, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x63,
	0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x12, 0x44,
	0x65, 0x63, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x6f, 0x6c, 0x6c, 0x42, 0x61, 0x63,
	0x6b, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x75,
	0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_proto_goTypes = []interface{}{
	(*DecutMoneyReq)(nil),  // 0: service.DecutMoneyReq
	(*DecutMoneyResp)(nil), // 1: service.DecutMoneyResp
}
var file_user_proto_depIdxs = []int32{
	0, // 0: service.Pay.DecutMoney:input_type -> service.DecutMoneyReq
	0, // 1: service.Pay.DecutMoneyRollBack:input_type -> service.DecutMoneyReq
	1, // 2: service.Pay.DecutMoney:output_type -> service.DecutMoneyResp
	1, // 3: service.Pay.DecutMoneyRollBack:output_type -> service.DecutMoneyResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecutMoneyReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecutMoneyResp); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
