// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: registry/output.proto

package registry

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_output_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_registry_output_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_registry_output_proto_rawDescGZIP(), []int{0}
}

func (x *Artifact) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_registry_output_proto protoreflect.FileDescriptor

var file_registry_output_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x22, 0x22, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x77, 0x61,
	0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registry_output_proto_rawDescOnce sync.Once
	file_registry_output_proto_rawDescData = file_registry_output_proto_rawDesc
)

func file_registry_output_proto_rawDescGZIP() []byte {
	file_registry_output_proto_rawDescOnce.Do(func() {
		file_registry_output_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_output_proto_rawDescData)
	})
	return file_registry_output_proto_rawDescData
}

var file_registry_output_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_registry_output_proto_goTypes = []interface{}{
	(*Artifact)(nil), // 0: registry.Artifact
}
var file_registry_output_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_registry_output_proto_init() }
func file_registry_output_proto_init() {
	if File_registry_output_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_output_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
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
			RawDescriptor: file_registry_output_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registry_output_proto_goTypes,
		DependencyIndexes: file_registry_output_proto_depIdxs,
		MessageInfos:      file_registry_output_proto_msgTypes,
	}.Build()
	File_registry_output_proto = out.File
	file_registry_output_proto_rawDesc = nil
	file_registry_output_proto_goTypes = nil
	file_registry_output_proto_depIdxs = nil
}
