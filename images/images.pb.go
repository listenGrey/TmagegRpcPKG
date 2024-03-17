// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0--rc1
// source: images/images.proto

package images

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

type TagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *TagRequest) Reset() {
	*x = TagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagRequest) ProtoMessage() {}

func (x *TagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagRequest.ProtoReflect.Descriptor instead.
func (*TagRequest) Descriptor() ([]byte, []int) {
	return file_images_images_proto_rawDescGZIP(), []int{0}
}

func (x *TagRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type FileItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Format   string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *FileItem) Reset() {
	*x = FileItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_images_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileItem) ProtoMessage() {}

func (x *FileItem) ProtoReflect() protoreflect.Message {
	mi := &file_images_images_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileItem.ProtoReflect.Descriptor instead.
func (*FileItem) Descriptor() ([]byte, []int) {
	return file_images_images_proto_rawDescGZIP(), []int{1}
}

func (x *FileItem) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileItem) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type FilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileItem `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FilesResponse) Reset() {
	*x = FilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_images_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesResponse) ProtoMessage() {}

func (x *FilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_images_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesResponse.ProtoReflect.Descriptor instead.
func (*FilesResponse) Descriptor() ([]byte, []int) {
	return file_images_images_proto_rawDescGZIP(), []int{2}
}

func (x *FilesResponse) GetFiles() []*FileItem {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_images_images_proto protoreflect.FileDescriptor

var file_images_images_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x54, 0x6d, 0x61, 0x67, 0x65, 0x67, 0x52, 0x70, 0x63,
	0x50, 0x4b, 0x47, 0x22, 0x1e, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x22, 0x3e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x54, 0x6d, 0x61, 0x67, 0x65, 0x67, 0x52, 0x70, 0x63, 0x50,
	0x4b, 0x47, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x32, 0x58, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x54, 0x61, 0x67, 0x12, 0x18, 0x2e, 0x54, 0x6d, 0x61, 0x67, 0x65, 0x67, 0x52, 0x70,
	0x63, 0x50, 0x4b, 0x47, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x54, 0x6d, 0x61, 0x67, 0x65, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_images_images_proto_rawDescOnce sync.Once
	file_images_images_proto_rawDescData = file_images_images_proto_rawDesc
)

func file_images_images_proto_rawDescGZIP() []byte {
	file_images_images_proto_rawDescOnce.Do(func() {
		file_images_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_images_images_proto_rawDescData)
	})
	return file_images_images_proto_rawDescData
}

var file_images_images_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_images_images_proto_goTypes = []interface{}{
	(*TagRequest)(nil),    // 0: TmagegRpcPKG.TagRequest
	(*FileItem)(nil),      // 1: TmagegRpcPKG.FileItem
	(*FilesResponse)(nil), // 2: TmagegRpcPKG.FilesResponse
}
var file_images_images_proto_depIdxs = []int32{
	1, // 0: TmagegRpcPKG.FilesResponse.files:type_name -> TmagegRpcPKG.FileItem
	0, // 1: TmagegRpcPKG.FileService.GetFileListByTag:input_type -> TmagegRpcPKG.TagRequest
	2, // 2: TmagegRpcPKG.FileService.GetFileListByTag:output_type -> TmagegRpcPKG.FilesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_images_images_proto_init() }
func file_images_images_proto_init() {
	if File_images_images_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_images_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagRequest); i {
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
		file_images_images_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileItem); i {
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
		file_images_images_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesResponse); i {
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
			RawDescriptor: file_images_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_images_images_proto_goTypes,
		DependencyIndexes: file_images_images_proto_depIdxs,
		MessageInfos:      file_images_images_proto_msgTypes,
	}.Build()
	File_images_images_proto = out.File
	file_images_images_proto_rawDesc = nil
	file_images_images_proto_goTypes = nil
	file_images_images_proto_depIdxs = nil
}
