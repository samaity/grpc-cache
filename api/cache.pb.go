// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/cache.proto

package cachepb

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

type PutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *PutReq) Reset() {
	*x = PutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReq) ProtoMessage() {}

func (x *PutReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReq.ProtoReflect.Descriptor instead.
func (*PutReq) Descriptor() ([]byte, []int) {
	return file_api_cache_proto_rawDescGZIP(), []int{0}
}

func (x *PutReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutReq) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type PutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutResp) Reset() {
	*x = PutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResp) ProtoMessage() {}

func (x *PutResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResp.ProtoReflect.Descriptor instead.
func (*PutResp) Descriptor() ([]byte, []int) {
	return file_api_cache_proto_rawDescGZIP(), []int{1}
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_api_cache_proto_rawDescGZIP(), []int{2}
}

func (x *GetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_api_cache_proto_rawDescGZIP(), []int{3}
}

func (x *GetResp) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_api_cache_proto protoreflect.FileDescriptor

var file_api_cache_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x22, 0x2c, 0x0a, 0x06, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x09, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x1a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x1b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x32, 0x5f, 0x0a, 0x05,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x2a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6d, 0x61,
	0x69, 0x74, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_cache_proto_rawDescOnce sync.Once
	file_api_cache_proto_rawDescData = file_api_cache_proto_rawDesc
)

func file_api_cache_proto_rawDescGZIP() []byte {
	file_api_cache_proto_rawDescOnce.Do(func() {
		file_api_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cache_proto_rawDescData)
	})
	return file_api_cache_proto_rawDescData
}

var file_api_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_cache_proto_goTypes = []interface{}{
	(*PutReq)(nil),  // 0: cachepb.PutReq
	(*PutResp)(nil), // 1: cachepb.PutResp
	(*GetReq)(nil),  // 2: cachepb.GetReq
	(*GetResp)(nil), // 3: cachepb.GetResp
}
var file_api_cache_proto_depIdxs = []int32{
	0, // 0: cachepb.Cache.Put:input_type -> cachepb.PutReq
	2, // 1: cachepb.Cache.Get:input_type -> cachepb.GetReq
	1, // 2: cachepb.Cache.Put:output_type -> cachepb.PutResp
	3, // 3: cachepb.Cache.Get:output_type -> cachepb.GetResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_cache_proto_init() }
func file_api_cache_proto_init() {
	if File_api_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReq); i {
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
		file_api_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResp); i {
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
		file_api_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_api_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
			RawDescriptor: file_api_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cache_proto_goTypes,
		DependencyIndexes: file_api_cache_proto_depIdxs,
		MessageInfos:      file_api_cache_proto_msgTypes,
	}.Build()
	File_api_cache_proto = out.File
	file_api_cache_proto_rawDesc = nil
	file_api_cache_proto_goTypes = nil
	file_api_cache_proto_depIdxs = nil
}
