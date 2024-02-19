// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: order.proto

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

type FindOrderByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *FindOrderByIdReq) Reset() {
	*x = FindOrderByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOrderByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOrderByIdReq) ProtoMessage() {}

func (x *FindOrderByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOrderByIdReq.ProtoReflect.Descriptor instead.
func (*FindOrderByIdReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *FindOrderByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOrderByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	PayTime     string `protobuf:"bytes,2,opt,name=PayTime,proto3" json:"PayTime,omitempty"`
	PaymentType int64  `protobuf:"varint,3,opt,name=PaymentType,proto3" json:"PaymentType,omitempty"`
	Status      int64  `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	TotalFee    int64  `protobuf:"varint,5,opt,name=TotalFee,proto3" json:"TotalFee,omitempty"`
	UserId      int64  `protobuf:"varint,6,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CloseTime   string `protobuf:"bytes,7,opt,name=CloseTime,proto3" json:"CloseTime,omitempty"`
	CommentTime string `protobuf:"bytes,8,opt,name=CommentTime,proto3" json:"CommentTime,omitempty"`
	ConsignTime string `protobuf:"bytes,9,opt,name=ConsignTime,proto3" json:"ConsignTime,omitempty"`
	CreateTime  string `protobuf:"bytes,10,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	EndTime     string `protobuf:"bytes,11,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *FindOrderByIdResp) Reset() {
	*x = FindOrderByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOrderByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOrderByIdResp) ProtoMessage() {}

func (x *FindOrderByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOrderByIdResp.ProtoReflect.Descriptor instead.
func (*FindOrderByIdResp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *FindOrderByIdResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindOrderByIdResp) GetPayTime() string {
	if x != nil {
		return x.PayTime
	}
	return ""
}

func (x *FindOrderByIdResp) GetPaymentType() int64 {
	if x != nil {
		return x.PaymentType
	}
	return 0
}

func (x *FindOrderByIdResp) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindOrderByIdResp) GetTotalFee() int64 {
	if x != nil {
		return x.TotalFee
	}
	return 0
}

func (x *FindOrderByIdResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindOrderByIdResp) GetCloseTime() string {
	if x != nil {
		return x.CloseTime
	}
	return ""
}

func (x *FindOrderByIdResp) GetCommentTime() string {
	if x != nil {
		return x.CommentTime
	}
	return ""
}

func (x *FindOrderByIdResp) GetConsignTime() string {
	if x != nil {
		return x.ConsignTime
	}
	return ""
}

func (x *FindOrderByIdResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *FindOrderByIdResp) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0xc7, 0x02, 0x0a, 0x11, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x50, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x65, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x65, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73,
	0x69, 0x67, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43,
	0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0x4f, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x46, 0x0a,
	0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_order_proto_goTypes = []interface{}{
	(*FindOrderByIdReq)(nil),  // 0: service.FindOrderByIdReq
	(*FindOrderByIdResp)(nil), // 1: service.FindOrderByIdResp
}
var file_order_proto_depIdxs = []int32{
	0, // 0: service.Order.FindOrderById:input_type -> service.FindOrderByIdReq
	1, // 1: service.Order.FindOrderById:output_type -> service.FindOrderByIdResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOrderByIdReq); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOrderByIdResp); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}