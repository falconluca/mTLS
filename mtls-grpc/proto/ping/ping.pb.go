// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.20.3
// source: mtls-grpc/proto/ping.proto

package ping

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtls_grpc_proto_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtls_grpc_proto_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_mtls_grpc_proto_ping_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtls_grpc_proto_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtls_grpc_proto_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_mtls_grpc_proto_ping_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_mtls_grpc_proto_ping_proto protoreflect.FileDescriptor

var file_mtls_grpc_proto_ping_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x74, 0x6c, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0x3c, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1b, 0x5a, 0x19, 0x6d, 0x74, 0x6c, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x3b, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mtls_grpc_proto_ping_proto_rawDescOnce sync.Once
	file_mtls_grpc_proto_ping_proto_rawDescData = file_mtls_grpc_proto_ping_proto_rawDesc
)

func file_mtls_grpc_proto_ping_proto_rawDescGZIP() []byte {
	file_mtls_grpc_proto_ping_proto_rawDescOnce.Do(func() {
		file_mtls_grpc_proto_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_mtls_grpc_proto_ping_proto_rawDescData)
	})
	return file_mtls_grpc_proto_ping_proto_rawDescData
}

var file_mtls_grpc_proto_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mtls_grpc_proto_ping_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: ping.PingRequest
	(*PingResponse)(nil), // 1: ping.PingResponse
}
var file_mtls_grpc_proto_ping_proto_depIdxs = []int32{
	0, // 0: ping.PingService.Ping:input_type -> ping.PingRequest
	1, // 1: ping.PingService.Ping:output_type -> ping.PingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mtls_grpc_proto_ping_proto_init() }
func file_mtls_grpc_proto_ping_proto_init() {
	if File_mtls_grpc_proto_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mtls_grpc_proto_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_mtls_grpc_proto_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_mtls_grpc_proto_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mtls_grpc_proto_ping_proto_goTypes,
		DependencyIndexes: file_mtls_grpc_proto_ping_proto_depIdxs,
		MessageInfos:      file_mtls_grpc_proto_ping_proto_msgTypes,
	}.Build()
	File_mtls_grpc_proto_ping_proto = out.File
	file_mtls_grpc_proto_ping_proto_rawDesc = nil
	file_mtls_grpc_proto_ping_proto_goTypes = nil
	file_mtls_grpc_proto_ping_proto_depIdxs = nil
}
