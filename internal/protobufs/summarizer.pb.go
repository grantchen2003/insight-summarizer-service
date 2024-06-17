// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.3
// source: summarizer.proto

package protobufs

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

type FileComponentIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileComponentIds []int32 `protobuf:"varint,1,rep,packed,name=file_component_ids,json=fileComponentIds,proto3" json:"file_component_ids,omitempty"`
}

func (x *FileComponentIds) Reset() {
	*x = FileComponentIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summarizer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileComponentIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileComponentIds) ProtoMessage() {}

func (x *FileComponentIds) ProtoReflect() protoreflect.Message {
	mi := &file_summarizer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileComponentIds.ProtoReflect.Descriptor instead.
func (*FileComponentIds) Descriptor() ([]byte, []int) {
	return file_summarizer_proto_rawDescGZIP(), []int{0}
}

func (x *FileComponentIds) GetFileComponentIds() []int32 {
	if x != nil {
		return x.FileComponentIds
	}
	return nil
}

type FileComponentSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FileComponentId int32  `protobuf:"varint,2,opt,name=file_component_id,json=fileComponentId,proto3" json:"file_component_id,omitempty"`
	Summary         string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *FileComponentSummary) Reset() {
	*x = FileComponentSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summarizer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileComponentSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileComponentSummary) ProtoMessage() {}

func (x *FileComponentSummary) ProtoReflect() protoreflect.Message {
	mi := &file_summarizer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileComponentSummary.ProtoReflect.Descriptor instead.
func (*FileComponentSummary) Descriptor() ([]byte, []int) {
	return file_summarizer_proto_rawDescGZIP(), []int{1}
}

func (x *FileComponentSummary) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileComponentSummary) GetFileComponentId() int32 {
	if x != nil {
		return x.FileComponentId
	}
	return 0
}

func (x *FileComponentSummary) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type FileComponentSummaries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileComponentSummaries []*FileComponentSummary `protobuf:"bytes,1,rep,name=file_component_summaries,json=fileComponentSummaries,proto3" json:"file_component_summaries,omitempty"`
}

func (x *FileComponentSummaries) Reset() {
	*x = FileComponentSummaries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summarizer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileComponentSummaries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileComponentSummaries) ProtoMessage() {}

func (x *FileComponentSummaries) ProtoReflect() protoreflect.Message {
	mi := &file_summarizer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileComponentSummaries.ProtoReflect.Descriptor instead.
func (*FileComponentSummaries) Descriptor() ([]byte, []int) {
	return file_summarizer_proto_rawDescGZIP(), []int{2}
}

func (x *FileComponentSummaries) GetFileComponentSummaries() []*FileComponentSummary {
	if x != nil {
		return x.FileComponentSummaries
	}
	return nil
}

var File_summarizer_proto protoreflect.FileDescriptor

var file_summarizer_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x73, 0x22, 0x6c, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x22, 0x69, 0x0a, 0x16, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x18,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x16, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x32, 0x5f, 0x0a,
	0x11, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x11, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x73, 0x1a, 0x17, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x42, 0x0c,
	0x5a, 0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_summarizer_proto_rawDescOnce sync.Once
	file_summarizer_proto_rawDescData = file_summarizer_proto_rawDesc
)

func file_summarizer_proto_rawDescGZIP() []byte {
	file_summarizer_proto_rawDescOnce.Do(func() {
		file_summarizer_proto_rawDescData = protoimpl.X.CompressGZIP(file_summarizer_proto_rawDescData)
	})
	return file_summarizer_proto_rawDescData
}

var file_summarizer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_summarizer_proto_goTypes = []interface{}{
	(*FileComponentIds)(nil),       // 0: FileComponentIds
	(*FileComponentSummary)(nil),   // 1: FileComponentSummary
	(*FileComponentSummaries)(nil), // 2: FileComponentSummaries
}
var file_summarizer_proto_depIdxs = []int32{
	1, // 0: FileComponentSummaries.file_component_summaries:type_name -> FileComponentSummary
	0, // 1: SummarizerService.CreateFileComponentSummaries:input_type -> FileComponentIds
	2, // 2: SummarizerService.CreateFileComponentSummaries:output_type -> FileComponentSummaries
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_summarizer_proto_init() }
func file_summarizer_proto_init() {
	if File_summarizer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_summarizer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileComponentIds); i {
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
		file_summarizer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileComponentSummary); i {
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
		file_summarizer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileComponentSummaries); i {
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
			RawDescriptor: file_summarizer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_summarizer_proto_goTypes,
		DependencyIndexes: file_summarizer_proto_depIdxs,
		MessageInfos:      file_summarizer_proto_msgTypes,
	}.Build()
	File_summarizer_proto = out.File
	file_summarizer_proto_rawDesc = nil
	file_summarizer_proto_goTypes = nil
	file_summarizer_proto_depIdxs = nil
}
