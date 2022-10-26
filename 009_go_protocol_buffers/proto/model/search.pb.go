// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/model/search.proto

// package 명은 go package 에 지정된 경로의 제일 마지막 것으로 해야 한다.

package model

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

type Corpus int32

const (
	Corpus_CORPUS_UNSPECIFIED Corpus = 0
	Corpus_CORPUS_UNIVERSAL   Corpus = 1
	Corpus_CORPUS_WEB         Corpus = 2
	Corpus_CORPUS_IMAGES      Corpus = 3
	Corpus_CORPUS_LOCAL       Corpus = 4
	Corpus_CORPUS_NEWS        Corpus = 5
	Corpus_CORPUS_PRODUCTS    Corpus = 6
	Corpus_CORPUS_VIDEO       Corpus = 7
)

// Enum value maps for Corpus.
var (
	Corpus_name = map[int32]string{
		0: "CORPUS_UNSPECIFIED",
		1: "CORPUS_UNIVERSAL",
		2: "CORPUS_WEB",
		3: "CORPUS_IMAGES",
		4: "CORPUS_LOCAL",
		5: "CORPUS_NEWS",
		6: "CORPUS_PRODUCTS",
		7: "CORPUS_VIDEO",
	}
	Corpus_value = map[string]int32{
		"CORPUS_UNSPECIFIED": 0,
		"CORPUS_UNIVERSAL":   1,
		"CORPUS_WEB":         2,
		"CORPUS_IMAGES":      3,
		"CORPUS_LOCAL":       4,
		"CORPUS_NEWS":        5,
		"CORPUS_PRODUCTS":    6,
		"CORPUS_VIDEO":       7,
	}
)

func (x Corpus) Enum() *Corpus {
	p := new(Corpus)
	*p = x
	return p
}

func (x Corpus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Corpus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_model_search_proto_enumTypes[0].Descriptor()
}

func (Corpus) Type() protoreflect.EnumType {
	return &file_proto_model_search_proto_enumTypes[0]
}

func (x Corpus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Corpus.Descriptor instead.
func (Corpus) EnumDescriptor() ([]byte, []int) {
	return file_proto_model_search_proto_rawDescGZIP(), []int{0}
}

type EnumAllowingAlias int32

const (
	EnumAllowingAlias_EAA_UNSPECIFIED EnumAllowingAlias = 0
	EnumAllowingAlias_EAA_STARTED     EnumAllowingAlias = 1
	EnumAllowingAlias_EAA_RUNNING     EnumAllowingAlias = 1
	EnumAllowingAlias_EAA_FINISHED    EnumAllowingAlias = 2
)

// Enum value maps for EnumAllowingAlias.
var (
	EnumAllowingAlias_name = map[int32]string{
		0: "EAA_UNSPECIFIED",
		1: "EAA_STARTED",
		// Duplicate value: 1: "EAA_RUNNING",
		2: "EAA_FINISHED",
	}
	EnumAllowingAlias_value = map[string]int32{
		"EAA_UNSPECIFIED": 0,
		"EAA_STARTED":     1,
		"EAA_RUNNING":     1,
		"EAA_FINISHED":    2,
	}
)

func (x EnumAllowingAlias) Enum() *EnumAllowingAlias {
	p := new(EnumAllowingAlias)
	*p = x
	return p
}

func (x EnumAllowingAlias) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumAllowingAlias) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_model_search_proto_enumTypes[1].Descriptor()
}

func (EnumAllowingAlias) Type() protoreflect.EnumType {
	return &file_proto_model_search_proto_enumTypes[1]
}

func (x EnumAllowingAlias) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumAllowingAlias.Descriptor instead.
func (EnumAllowingAlias) EnumDescriptor() ([]byte, []int) {
	return file_proto_model_search_proto_rawDescGZIP(), []int{1}
}

type EnumNotAllowingAlias int32

const (
	EnumNotAllowingAlias_ENAA_UNSPECIFIED EnumNotAllowingAlias = 0
	EnumNotAllowingAlias_ENAA_STARTED     EnumNotAllowingAlias = 1
	// ENAA_RUNNING = 1;  // Uncommenting this line will cause a compile error inside Google and a warning message outside.
	EnumNotAllowingAlias_ENAA_FINISHED EnumNotAllowingAlias = 2
)

// Enum value maps for EnumNotAllowingAlias.
var (
	EnumNotAllowingAlias_name = map[int32]string{
		0: "ENAA_UNSPECIFIED",
		1: "ENAA_STARTED",
		2: "ENAA_FINISHED",
	}
	EnumNotAllowingAlias_value = map[string]int32{
		"ENAA_UNSPECIFIED": 0,
		"ENAA_STARTED":     1,
		"ENAA_FINISHED":    2,
	}
)

func (x EnumNotAllowingAlias) Enum() *EnumNotAllowingAlias {
	p := new(EnumNotAllowingAlias)
	*p = x
	return p
}

func (x EnumNotAllowingAlias) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumNotAllowingAlias) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_model_search_proto_enumTypes[2].Descriptor()
}

func (EnumNotAllowingAlias) Type() protoreflect.EnumType {
	return &file_proto_model_search_proto_enumTypes[2]
}

func (x EnumNotAllowingAlias) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumNotAllowingAlias.Descriptor instead.
func (EnumNotAllowingAlias) EnumDescriptor() ([]byte, []int) {
	return file_proto_model_search_proto_rawDescGZIP(), []int{2}
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int32  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_model_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchRequest) GetResultPerPage() int32 {
	if x != nil {
		return x.ResultPerPage
	}
	return 0
}

var File_proto_model_search_proto protoreflect.FileDescriptor

var file_proto_model_search_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x6e, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x2a, 0xa3, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x52, 0x50, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f,
	0x52, 0x50, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x52, 0x50, 0x55, 0x53, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x52, 0x50, 0x55, 0x53, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45,
	0x53, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x52, 0x50, 0x55, 0x53, 0x5f, 0x4c, 0x4f,
	0x43, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x52, 0x50, 0x55, 0x53, 0x5f,
	0x4e, 0x45, 0x57, 0x53, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x52, 0x50, 0x55, 0x53,
	0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x53, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x4f, 0x52, 0x50, 0x55, 0x53, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x07, 0x2a, 0x60, 0x0a,
	0x11, 0x45, 0x6e, 0x75, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x41, 0x41, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x41, 0x41, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x41, 0x41, 0x5f,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x41, 0x41,
	0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x2a,
	0x51, 0x0a, 0x14, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x6f, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x41, 0x41, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x45, 0x4e, 0x41, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x41, 0x41, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44,
	0x10, 0x02, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_model_search_proto_rawDescOnce sync.Once
	file_proto_model_search_proto_rawDescData = file_proto_model_search_proto_rawDesc
)

func file_proto_model_search_proto_rawDescGZIP() []byte {
	file_proto_model_search_proto_rawDescOnce.Do(func() {
		file_proto_model_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_model_search_proto_rawDescData)
	})
	return file_proto_model_search_proto_rawDescData
}

var file_proto_model_search_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_model_search_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_model_search_proto_goTypes = []interface{}{
	(Corpus)(0),               // 0: proto.model.Corpus
	(EnumAllowingAlias)(0),    // 1: proto.model.EnumAllowingAlias
	(EnumNotAllowingAlias)(0), // 2: proto.model.EnumNotAllowingAlias
	(*SearchRequest)(nil),     // 3: proto.model.SearchRequest
}
var file_proto_model_search_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_model_search_proto_init() }
func file_proto_model_search_proto_init() {
	if File_proto_model_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_model_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
			RawDescriptor: file_proto_model_search_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_model_search_proto_goTypes,
		DependencyIndexes: file_proto_model_search_proto_depIdxs,
		EnumInfos:         file_proto_model_search_proto_enumTypes,
		MessageInfos:      file_proto_model_search_proto_msgTypes,
	}.Build()
	File_proto_model_search_proto = out.File
	file_proto_model_search_proto_rawDesc = nil
	file_proto_model_search_proto_goTypes = nil
	file_proto_model_search_proto_depIdxs = nil
}
